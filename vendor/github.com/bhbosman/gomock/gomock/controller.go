// Copyright 2010 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package gomock is a mock framework for Go.
//
// Standard usage:
//   (1) Define an interface that you wish to mock.
//         type MyInterface interface {
//           SomeMethod(x int64, y string)
//         }
//   (2) Use mockgen to generate a mock from the interface.
//   (3) Use the mock in a test:
//         func TestMyThing(t *testing.T) {
//           mockCtrl := gomock.NewController(t)
//           defer mockCtrl.Finish()
//
//           mockObj := something.NewMockMyInterface(mockCtrl)
//           mockObj.EXPECT().SomeMethod(4, "blah")
//           // pass mockObj to a real object and play with it.
//         }
//
// By default, expected calls are not enforced to run in any particular order.
// Call order dependency can be enforced by use of InOrder and/or Call.After.
// Call.After can create more varied call order dependencies, but InOrder is
// often more convenient.
//
// The following examples create equivalent call order dependencies.
//
// Example of using Call.After to chain expected call order:
//
//     firstCall := mockObj.EXPECT().SomeMethod(1, "first")
//     secondCall := mockObj.EXPECT().SomeMethod(2, "second").After(firstCall)
//     mockObj.EXPECT().SomeMethod(3, "third").After(secondCall)
//
// Example of using InOrder to declare expected call order:
//
//     gomock.InOrder(
//         mockObj.EXPECT().SomeMethod(1, "first"),
//         mockObj.EXPECT().SomeMethod(2, "second"),
//         mockObj.EXPECT().SomeMethod(3, "third"),
//     )
//
// TODO:
//	- Handle different argument/return types (e.g. ..., chan, map, interface).
package gomock

import (
	"context"
	"fmt"
	"reflect"
	"runtime"
	"sync"
)

// A TestReporter is something that can be used to report test failures.  It
// is satisfied by the standard library's *testing.T.
type TestReporter interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

// TestHelper is a TestReporter that has the Helper method.  It is satisfied
// by the standard library's *testing.T.
type TestHelper interface {
	TestReporter
	Helper()
}

// A Controller represents the top-level control of a mock ecosystem.  It
// defines the scope and lifetime of mock objects, as well as their
// expectations.  It is safe to call Controller's methods from multiple
// goroutines. Each test should create a new Controller and invoke Finish via
// defer.
//
//   func TestFoo(t *testing.T) {
//     ctrl := gomock.NewController(t)
//     defer ctrl.Finish()
//     // ..
//   }
//
//   func TestBar(t *testing.T) {
//     t.Run("Sub-Test-1", st) {
//       ctrl := gomock.NewController(st)
//       defer ctrl.Finish()
//       // ..
//     })
//     t.Run("Sub-Test-2", st) {
//       ctrl := gomock.NewController(st)
//       defer ctrl.Finish()
//       // ..
//     })
//   })
type Controller struct {
	// T should only be called within a generated mock. It is not intended to
	// be used in user code and may be changed in future versions. T is the
	// TestReporter passed in when creating the Controller via NewController.
	// If the TestReporter does not implement a TestHelper it will be wrapped
	// with a nopTestHelper.
	T             TestHelper
	mu            sync.Mutex
	expectedCalls *callSet
	finished      bool
}

// NewController returns a new Controller. It is the preferred way to create a
// Controller.
func NewController(t TestReporter) *Controller {
	h, ok := t.(TestHelper)
	if !ok {
		h = nopTestHelper{t}
	}

	return &Controller{
		T:             h,
		expectedCalls: newCallSet(),
	}
}

type cancelReporter struct {
	TestHelper
	cancel func()
}

func (r *cancelReporter) Errorf(format string, args ...interface{}) {
	r.TestHelper.Errorf(format, args...)
}
func (r *cancelReporter) Fatalf(format string, args ...interface{}) {
	defer r.cancel()
	r.TestHelper.Fatalf(format, args...)
}

// WithContext returns a new Controller and a Context, which is cancelled on any
// fatal failure.
func WithContext(ctx context.Context, t TestReporter) (*Controller, context.Context) {
	h, ok := t.(TestHelper)
	if !ok {
		h = nopTestHelper{t}
	}

	ctx, cancel := context.WithCancel(ctx)
	return NewController(&cancelReporter{h, cancel}), ctx
}

type nopTestHelper struct {
	TestReporter
}

func (h nopTestHelper) Helper() {}

// RecordCall is called by a mock. It should not be called by user code.
func (ctrl *Controller) RecordCall(receiverInfo ReceiverInfo, args ...interface{}) *Call {
	ctrl.T.Helper()

	recv := reflect.ValueOf(receiverInfo.MockReceiver)
	for i := 0; i < recv.Type().NumMethod(); i++ {
		if recv.Type().Method(i).Name == receiverInfo.Method {
			return ctrl.RecordCallWithMethodType(
				ReceiverInfo{
					MockReceiver: receiverInfo.MockReceiver,
					Method:       receiverInfo.Method,
					MethodType:   recv.Method(i).Type(),
				},
				args...)
		}
	}
	ctrl.T.Fatalf("gomock: failed finding method %s on %T", receiverInfo.Method, receiverInfo.MockReceiver)
	panic("unreachable")
}

// RecordCallWithMethodType is called by a mock. It should not be called by user code.
func (ctrl *Controller) RecordCallWithMethodType(receiverInfo ReceiverInfo, args ...interface{}) *Call {
	ctrl.T.Helper()

	call := newCall(ctrl.T, receiverInfo, args...)

	ctrl.mu.Lock()
	defer ctrl.mu.Unlock()
	ctrl.expectedCalls.Add(call)

	return call
}

// Call is called by a mock. It should not be called by user code.
func (ctrl *Controller) Call(receiver interface{}, instance interface{}, method string, args ...interface{}) []interface{} {
	ctrl.T.Helper()

	actions := func() []func([]interface{}) []interface{} {
		ctrl.T.Helper()
		ctrl.mu.Lock()
		defer ctrl.mu.Unlock()

		expected, err := ctrl.expectedCalls.FindMatch(receiver, method, args)
		if err != nil && instance == nil {
			origin := callerInfo(2)
			ctrl.T.Fatalf("Unexpected call to %T.%v(%v) at %s because: %s", receiver, method, args, origin, err)
		}

		// Two things happen here:
		// * the matching call no longer needs to check prerequite calls,
		// * and the prerequite calls are no longer expected, so remove them.
		if expected != nil {
			preReqCalls := expected.dropPrereqs()
			for _, preReqCall := range preReqCalls {
				ctrl.expectedCalls.Remove(preReqCall)
			}
		}
		var actions []func([]interface{}) []interface{}
		if expected != nil {
			actions = expected.call(args)
			if expected.exhausted() && instance == nil {
				ctrl.expectedCalls.Remove(expected)
			}
		}


		b := expected == nil
		b = b || (expected != nil && !expected.additionalActions)


		if b && instance != nil {
			reflectType := reflect.TypeOf(instance)
			if methodByName, b := reflectType.MethodByName(method); b {
				createInstanceMethodCall := func() func([]interface{}) []interface{} {
					if instance != nil {
						reflectType := reflect.TypeOf(instance)
						if methodByName, b := reflectType.MethodByName(method); b {
							return func(args []interface{}) []interface{} {
								ft := methodByName.Func.Type()
								vargs := make([]reflect.Value, len(args)+1)
								vargs[0] = reflect.ValueOf(instance)
								for i := 0; i < len(args); i++ {
									if args[i] != nil {
										vargs[i+1] = reflect.ValueOf(args[i])
									} else {
										vargs[i+1] = reflect.Zero(
											ft.In(
												func() int {
													if i+1 > ft.NumIn()-1 {
														return ft.NumIn() - 1
													}
													return i + 1
												}()))
									}
								}
								vrets := methodByName.Func.Call(vargs)
								returns := make([]interface{}, len(vrets))
								for i, ret := range vrets {
									returns[i] = ret.Interface()
								}
								return returns
							}
						}
					}
					return nil
				}

				return []func([]interface{}) []interface{}{
					func([]interface{}) []interface{} {
						// Synthesize the zero value for each of the return args' types.
						rets := make([]interface{}, methodByName.Type.NumOut())
						for i := 0; i < methodByName.Type.NumOut(); i++ {
							rets[i] = reflect.Zero(methodByName.Type.Out(i)).Interface()
						}
						return rets
					},
					createInstanceMethodCall(),
				}
			}
			return nil
		}

		return actions
	}()

	var rets []interface{}
	for _, action := range actions {
		if action != nil {
			if r := action(args); r != nil {
				rets = r
			}
		}
	}

	return rets
}

// Finish checks to see if all the methods that were expected to be called
// were called. It should be invoked for each Controller. It is not idempotent
// and therefore can only be invoked once.
func (ctrl *Controller) Finish() {
	ctrl.T.Helper()

	ctrl.mu.Lock()
	defer ctrl.mu.Unlock()

	if ctrl.finished {
		ctrl.T.Fatalf("Controller.Finish was called more than once. It has to be called exactly once.")
	}
	ctrl.finished = true

	// If we're currently panicking, probably because this is a deferred call,
	// pass through the panic.
	if err := recover(); err != nil {
		panic(err)
	}

	// Check that all remaining expected calls are satisfied.
	failures := ctrl.expectedCalls.Failures()
	for _, call := range failures {
		ctrl.T.Errorf("missing call(s) (min(%v) <= actual(%v) <= max(%v)) to %v", call.minCalls, call.numCalls, call.maxCalls, call)
	}
	if len(failures) != 0 {
		ctrl.T.Fatalf("aborting test due to missing call(s)")
	}
}

func callerInfo(skip int) string {
	if _, file, line, ok := runtime.Caller(skip + 1); ok {
		return fmt.Sprintf("%s:%d", file, line)
	}
	return "unknown file"
}
