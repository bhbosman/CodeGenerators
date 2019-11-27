package scoping

import (
	"fmt"
	si "github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	"go.uber.org/multierr"
	"reflect"
)

type IFileWriter interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

type ScopeWalker struct {
	Logger IFileWriter
}

func NewScopeWalker(logger IFileWriter) *ScopeWalker {
	return &ScopeWalker{
		Logger: logger,
	}
}

func (self ScopeWalker) Scope(scopingContext si.IScopingContext, indent int, dcl si.ITypeSpec, fileName string) error {
	var err error
	for typeSpec := dcl; typeSpec != nil; typeSpec, _ = typeSpec.GetNextTypeSpec() {
		err = multierr.Append(
			err,
			self.InternalGenerate(scopingContext, 0, typeSpec))
	}
	return err
}

func (self ScopeWalker) InternalGenerate(scopingContext si.IScopingContext, indent int, dcl si.ITypeSpec) error {
	switch dcl.GetKind() {

	case si.Attr_specIdlType:
		attributeDecl, ok := dcl.(si.IAttributeDcl)
		if ok {
			return self.ScopeAttributeDcl(indent+1, attributeDecl)
		}

	case si.ExceptionIdlType:
		exceptionDecl, ok := dcl.(si.IIdlException)
		if ok {
			return self.ScopeExceptionDcl(scopingContext, indent+1, exceptionDecl)
		}

	case si.ConstDclType:
		constantDecl, ok := dcl.(si.IIdlConstDcl)
		if ok {
			return self.ScopeConstantDcl(indent+1, constantDecl)
		}
	case si.ModuleIdlType:
		moduleDcl, ok := dcl.(si.IIdlModuleDcl)
		if ok {
			return self.ScopeModuleDcl(scopingContext, indent+1, moduleDcl)
		}
	case si.StructIdlType:
		structType, ok := dcl.(si.IStructType)
		if ok {
			return self.ScopeStructDcl(scopingContext, indent+1, structType)
		}
	case si.RWEnumIdlType:
		enumType, ok := dcl.(si.IEnumType)
		if ok {
			return self.ScopeEnumDcl(scopingContext, indent+1, enumType)
		}

	case si.InterfaceIdlType:
		interfaceDcl, ok := dcl.(si.IInterfaceDcl)
		if ok {
			return self.ScopeInterfaceDcl(scopingContext, indent+1, interfaceDcl)
		}
	case si.Op_dclIdlType:
		operation, ok := dcl.(si.IOperationDeclarations)
		if ok {
			return self.ScopeOperationDcl(scopingContext, indent+1, operation)
		}
	case si.IdlValue_Abs_DefType:
		value, ok := dcl.(si.IValueAbsoluteDefinition)
		if ok {
			return self.ScopeValueAbsolute(scopingContext, indent+1, value)
		}

	case si.TypeDeclaratorIdlType:
		typeDecl, ok := dcl.(si.ITypeDeclarator)
		if ok {
			return self.ScopeTypeDcl(scopingContext, indent+1, typeDecl)
		}
	default:
		self.Logger.Println(fmt.Sprintf(">>>>>>>>>>>>>%v<<<<<<<<<", dcl))
		return nil
	}
	return nil
}

func (self ScopeWalker) ScopeModuleDcl(scopingContext si.IScopingContext, indent int, dcl si.IIdlModuleDcl) error {
	self.Logger.Println(fmt.Sprintf("(%d): %v", indent, dcl))
	var err error
	scopeName := dcl.GetName()
	newScopingContext := NewScopingContext(scopeName, nil, scopingContext)
	defer func() {
		_ = newScopingContext.Iterate(
			func(key string, value si.IBaseDeclaredType) error {
				scopeName := self.buildDeclarationName(dcl.GetName(), key)
				value.SetName(scopeName)
				return scopingContext.Add(scopeName, value)
			})
	}()
	err = multierr.Append(
		err,
		dcl.Iterate(func(typeSpec si.ITypeSpec) error {
			return self.InternalGenerate(newScopingContext, indent, typeSpec)
		}))

	return err

}

func (self ScopeWalker) ScopeStructDcl(scopingContext si.IScopingContext, indent int, dcl si.IStructType) error {
	structScope := func(scopingContext si.IScopingContext, indent int, dcl si.IStructType) error {
		var err error
		if !dcl.Forward() {
			indent++
			members := dcl.Members()
			if members != nil {
				for _, memberInformation := range members.GetMembers() {
					self.Logger.Println(fmt.Sprintf("(%d): %v", indent, memberInformation))
					if placeHolder, isPlaceHolder := memberInformation.GetTypeSpec().(si.IDeclaredTypePlaceHolder); isPlaceHolder {
						if placeHolder.GetKind() == si.DeclareTypePlaceHolderType {
							if found, foundDeclareType := scopingContext.Find(placeHolder.GetName(), true); found {
								if declaredType, ok := foundDeclareType.(si.IDeclaredType); ok {
									err = multierr.Append(
										err,
										declaredType.Link(placeHolder))
								}
							}
						}
					}
				}
			}
		}
		return err
	}

	var err error
	self.Logger.Println(fmt.Sprintf("(%d): %v", indent, dcl))
	if found, declType := scopingContext.Find(dcl.GetName(), false); found {
		if existingDecl, ok := declType.(si.IInterfaceDcl); ok {
			if dcl.Forward() && existingDecl.Forward() {
				err = multierr.Append(
					err,
					fmt.Errorf("multiple forward declarations of interface %v", dcl.GetName()))
			} else if dcl.Forward() && !existingDecl.Forward() {
				err = multierr.Append(
					err,
					fmt.Errorf("forward declaration after declaraton interface %v", dcl.GetName()))
			} else if !dcl.Forward() && existingDecl.Forward() {
				err = multierr.Append(
					err,
					scopingContext.Replace(dcl.GetName(), dcl))
				err = multierr.Append(
					err,
					structScope(scopingContext, indent, dcl))

			} else if !dcl.Forward() && !existingDecl.Forward() {
				err = multierr.Append(
					err,
					fmt.Errorf("two declarations of structs %v", dcl.GetName()))
			}
		} else {
			err = multierr.Append(
				err,
				fmt.Errorf("two declarations of %v. An interface and a  %v", dcl.GetName(), reflect.TypeOf(declType)))
		}
	} else {
		err = multierr.Append(
			err,
			scopingContext.Add(dcl.GetName(), dcl))
		err = multierr.Append(
			err,
			structScope(scopingContext, indent, dcl))

	}
	return err
}

func (self ScopeWalker) ScopeEnumDcl(scopingContext si.IScopingContext, indent int, enumType si.IEnumType) error {
	self.Logger.Println(fmt.Sprintf("(%d): %v", indent, enumType))
	indent++

	err := scopingContext.Add(enumType.GetName(), enumType)
	for m := enumType.Enumerator(); m != nil; m = m.Next() {
		self.Logger.Println(fmt.Sprintf("(%d): %v", indent, m))
	}
	return err
}

func (self ScopeWalker) buildDeclarationName(scope, name string) string {
	if scope == "" {
		return name
	}
	return fmt.Sprintf("%v::%v", scope, name)

}

func (self ScopeWalker) ScopeInterfaceDcl(scopingContext si.IScopingContext, indent int, dcl si.IInterfaceDcl) error {
	self.Logger.Println(fmt.Sprintf("(%d): %v", indent, dcl))
	if dcl.GetName() == "Container" {
		fmt.Println("Container")
	}

	interfaceScope := func(scopingContext si.IScopingContext, indent int, dcl si.IInterfaceDcl) error {
		var err error
		if !dcl.Forward() {
			scopeName := dcl.GetName()
			newScopingContext := NewScopingContext(scopeName, nil, scopingContext)
			defer func() {
				_ = newScopingContext.Iterate(
					func(key string, value si.IBaseDeclaredType) error {
						scopeName := self.buildDeclarationName(dcl.GetName(), key)
						return scopingContext.Add(scopeName, value)
					})
			}()
			err = multierr.Append(
				err,
				dcl.Iterate(func(typeSpec si.ITypeSpec) error {
					return self.InternalGenerate(newScopingContext, indent, typeSpec)
				}))
		}
		return err
	}
	var err error
	if found, declType := scopingContext.Find(dcl.GetName(), false); found {
		if existingInterfaceDecl, ok := declType.(si.IInterfaceDcl); ok {
			if dcl.Forward() && existingInterfaceDecl.Forward() {
				err = multierr.Append(
					err,
					fmt.Errorf("multiple forward declarations of interface %v", dcl.GetName()))
			} else if dcl.Forward() && !existingInterfaceDecl.Forward() {
				err = multierr.Append(
					err,
					fmt.Errorf("forward declaration after declaraton interface %v", dcl.GetName()))
			} else if !dcl.Forward() && existingInterfaceDecl.Forward() {
				err = multierr.Append(
					err,
					scopingContext.Replace(dcl.GetName(), dcl))

				err = multierr.Append(
					err,
					interfaceScope(scopingContext, indent, dcl))

			} else if !dcl.Forward() && !existingInterfaceDecl.Forward() {
				err = multierr.Append(
					err,
					fmt.Errorf("two declarations of interface %v", dcl.GetName()))
			}
		} else {
			err = multierr.Append(
				err,
				fmt.Errorf("two declarations of %v. An struct and a  %v", dcl.GetName(), reflect.TypeOf(declType)))
		}
	} else {
		err = multierr.Append(
			err,
			scopingContext.Add(dcl.GetName(), dcl))
		err = multierr.Append(
			err,
			interfaceScope(scopingContext, indent, dcl))
	}
	return err
}

//func (self ScopeWalker) findType(scopingContext si.IScopingContext, dcl si.IBaseDeclaredType) error {
//	b, _ := scopingContext.Find(dcl.GetName())
//	if b {
//		return nil
//	}
//	return fmt.Errorf("type not found %v", dcl.GetName())
//}

func (self ScopeWalker) ScopeOperationDcl(scopingContext si.IScopingContext, indent int, dcl si.IOperationDeclarations) error {
	var err error
	self.Logger.Println(fmt.Sprintf("(%d): %v", indent, dcl))
	indent++
	if found, foundDeclareType := scopingContext.Find(dcl.GetOperationDeclaratorType().GetName(), true); found {
		if declaredType, ok := foundDeclareType.(si.IDeclaredType); ok {
			if placeHolder, isPlaceHolder := dcl.GetOperationDeclaratorType().(si.IDeclaredTypePlaceHolder); isPlaceHolder {
				err = multierr.Append(
					err,
					declaredType.Link(placeHolder))
			}
		}
	}
	for param := dcl.GetParams(); param != nil; param = param.GetNextParameterDeclarations() {
		if found, foundDeclareType := scopingContext.Find(param.GetParamDeclarationType().GetName(), true); found {
			if declaredType, ok := foundDeclareType.(si.IDeclaredType); ok {
				if placeHolder, isPlaceHolder := param.GetParamDeclarationType().(si.IDeclaredTypePlaceHolder); isPlaceHolder {
					err = multierr.Append(
						err,
						declaredType.Link(placeHolder))
				}
			}
		}
	}
	return err
}

func (self ScopeWalker) ScopeTypeDcl(scopingContext si.IScopingContext, indent int, dcl si.ITypeDeclarator) error {
	self.Logger.Println(fmt.Sprintf("(%d): %v", indent, dcl))
	return scopingContext.Add(dcl.GetName(), dcl)
}

func (self ScopeWalker) ScopeValueAbsolute(scopingContext si.IScopingContext, indent int, dcl si.IValueAbsoluteDefinition) error {
	self.Logger.Println(fmt.Sprintf("(%d): %v", indent, dcl))
	var err error
	err = multierr.Append(
		err,
		scopingContext.Add(dcl.GetName(), dcl))

	err = multierr.Append(
		err,
		dcl.Iterate(
			func(typeSpec si.ITypeSpec) error {
				return self.InternalGenerate(scopingContext, indent, typeSpec)
			}))

	return err
}

func (self ScopeWalker) ScopeConstantDcl(indent int, constDcl si.IIdlConstDcl) error {
	self.Logger.Println(fmt.Sprintf("(%d): %v", indent, constDcl))
	return nil
}

func (self ScopeWalker) ScopeExceptionDcl(scopingContext si.IScopingContext, indent int, dcl si.IIdlException) error {
	self.Logger.Println(fmt.Sprintf("(%d): %v", indent, dcl))
	indent++
	err := scopingContext.Add(dcl.GetName(), dcl)
	members := dcl.Members()
	if members != nil {
		for _, memberInformation := range members.GetMembers() {
			self.Logger.Println(fmt.Sprintf("(%d): %v", indent, memberInformation))
			memberInformationTypeSpec := memberInformation.GetTypeSpec()
			switch memberInformationTypeSpec.GetKind() {
			case si.DeclareTypePlaceHolderType:
				if found, foundDeclareType := scopingContext.Find(memberInformationTypeSpec.GetName(), true); found {
					if declaredType, ok := foundDeclareType.(si.IDeclaredType); ok {
						if placeHolder, isPlaceHolder := memberInformationTypeSpec.(si.IDeclaredTypePlaceHolder); isPlaceHolder {
							err = multierr.Append(
								err,
								declaredType.Link(placeHolder))
						}
					}
				}
				break
			case si.SequenceIdlType:
				//if found, foundDeclareType := scopingContext.Find(memberInformationTypeSpec.GetName()); found {
				//	if declaredType, ok := foundDeclareType.(si.IDeclaredType); ok {
				//		declaredType.SequenceRequired(true)
				//	}
				//}
				break
			}
		}
	}
	return err
}

func (self ScopeWalker) ScopeAttributeDcl(indent int, attributeDcl si.IAttributeDcl) error {
	self.Logger.Println(fmt.Sprintf("(%d): %v", indent, attributeDcl))
	return nil
}
