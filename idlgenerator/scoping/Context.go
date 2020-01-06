package scoping

import (
	"fmt"
	si "github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	"go.uber.org/multierr"
)

type Context struct {
	name            string
	declaredTypes   map[string]si.IBaseDeclaredType
	unresolvedTypes map[string][]si.IFileInformation
	prevContext     si.IScopingContext
}

func (self *Context) Replace(name string, dcl si.IBaseDeclaredType) error {
	var err error
	if incomingDeclaredType, ok22 := dcl.(si.IDeclaredType); ok22 {
		if existing, ok := self.declaredTypes[name]; ok {
			if existingDeclaredType, ok2 := existing.(si.IDeclaredType); ok2 {
				for _, item := range existingDeclaredType.GetLinkedItems() {
					err = multierr.Append(
						err,
						incomingDeclaredType.Link(item.(si.IDeclaredTypePlaceHolder)))
				}
			}
		}
	}

	self.declaredTypes[name] = dcl
	return nil
}

func (self *Context) IterateUnresolved(cb func(name string, information []si.IFileInformation) error) error {
	var err, errInstance error
	if cb != nil {
		for k, v := range self.unresolvedTypes {
			errInstance = cb(k, v)
			err = multierr.Append(err, errInstance)
		}
	}
	return err
}

func (self *Context) Previous() si.IScopingContext {
	return self.prevContext
}

func (self *Context) FindTypeSpec(fileInformation si.IFileInformation, s string) (si.IBaseDeclaredType, error) {
	b, declaredType := self.Find(s, true)
	if b {
		return declaredType, nil
	}
	return nil, fmt.Errorf("can not find \"%v\"", s)
}

func NewScopingContext(name string, defaultTypeService si.IDefaultTypeService, prevContext si.IScopingContext) *Context {
	declaredTypes := make(map[string]si.IBaseDeclaredType)
	if defaultTypeService != nil {
		defaultTypeService.Iterate(func(key string, declaredType si.IBaseDeclaredType) {
			declaredTypes[key] = declaredType
		})
	}
	return &Context{
		name:            name,
		declaredTypes:   declaredTypes,
		unresolvedTypes: make(map[string][]si.IFileInformation),
		prevContext:     prevContext,
	}
}

func (self *Context) AddUnresolved(name string, information si.IFileInformation) error {
	l, ok := self.unresolvedTypes[name]
	if !ok {
		l = make([]si.IFileInformation, 0, 4)
	}
	l = append(l, information)
	self.unresolvedTypes[name] = l
	return nil
}

func (self *Context) Iterate(cb func(key string, value si.IBaseDeclaredType) error) error {
	var err error
	if cb != nil {
		for key, value := range self.declaredTypes {
			errInstance := cb(key, value)
			err = multierr.Append(err, errInstance)
		}
	}
	return err
}

func (self *Context) Find(name string, allContext bool) (bool, si.IBaseDeclaredType) {
	result, ok := self.declaredTypes[name]
	if !ok && self.prevContext != nil && allContext {
		return self.prevContext.Find(name, allContext)
	}
	return ok, result
}

func (self *Context) Add(name string, structType si.IBaseDeclaredType) error {
	_, ok := self.declaredTypes[name]
	if ok {
		return fmt.Errorf("%v already added", name)
	}
	self.declaredTypes[name] = structType
	return nil
}
