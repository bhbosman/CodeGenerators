package scoping

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	"go.uber.org/multierr"
)

type Context struct {
	declaredTypes   map[string]ScopingInterfaces.IBaseDeclaredType
	unresolvedTypes map[string][]ScopingInterfaces.IFileInformation
	prevContext     ScopingInterfaces.IScopingContext
}

func (self *Context) Replace(name string, structType ScopingInterfaces.IBaseDeclaredType) error {
	_, _ = self.declaredTypes[name]
	return nil
}

func (self *Context) IterateUnresolved(cb func(name string, information []ScopingInterfaces.IFileInformation) error) error {
	var err, errInstance error
	if cb != nil {
		for k, v := range self.unresolvedTypes {
			errInstance = cb(k, v)
			err = multierr.Append(err, errInstance)
		}
	}
	return err
}

func (self *Context) Previous() ScopingInterfaces.IScopingContext {
	return self.prevContext
}

func (self *Context) FindTypeSpec(fileInformation ScopingInterfaces.IFileInformation, s string) (ScopingInterfaces.IBaseDeclaredType, error) {
	b, declaredType := self.Find(s)
	if b {
		return declaredType, nil
	}
	return nil, fmt.Errorf("can not find \"%v\"", s)
}

func NewScopingContext(defaultTypeService ScopingInterfaces.IDefaultTypeService, prevContext ScopingInterfaces.IScopingContext) *Context {
	declaredTypes := make(map[string]ScopingInterfaces.IBaseDeclaredType)
	if defaultTypeService != nil {
		defaultTypeService.Iterate(func(key string, declaredType ScopingInterfaces.IBaseDeclaredType) {
			declaredTypes[key] = declaredType
		})
	}
	return &Context{
		declaredTypes:   declaredTypes,
		unresolvedTypes: make(map[string][]ScopingInterfaces.IFileInformation),
		prevContext:     prevContext,
	}
}

func (self *Context) AddUnresolved(name string, information ScopingInterfaces.IFileInformation) error {
	l, ok := self.unresolvedTypes[name]
	if !ok {
		l = make([]ScopingInterfaces.IFileInformation, 0, 4)
	}
	l = append(l, information)
	self.unresolvedTypes[name] = l
	return nil
}

func (self *Context) Iterate(cb func(key string, value ScopingInterfaces.IBaseDeclaredType) error) error {
	var err error
	if cb != nil {
		for key, value := range self.declaredTypes {
			errInstance := cb(key, value)
			err = multierr.Append(err, errInstance)
		}
	}
	return err
}

func (self *Context) Find(name string) (bool, ScopingInterfaces.IBaseDeclaredType) {
	result, ok := self.declaredTypes[name]
	if !ok && self.prevContext != nil {
		return self.prevContext.Find(name)
	}
	return ok, result
}

func (self *Context) Add(name string, structType ScopingInterfaces.IBaseDeclaredType) error {
	_, ok := self.declaredTypes[name]
	if ok {
		return fmt.Errorf("%v already added", name)
	}
	self.declaredTypes[name] = structType
	return nil
}
