package ScopingInterfaces

type IIdlKind interface {
	GetKind() IDlSupportedTypes
}

type IIdlName interface {
	GetName() string
	SetName(string)
}

type IFileInformation interface {
	GetFileName() string
	GetRow() int
	GetCol() int
}

type IIdlDefinition interface {
	IFileInformation
	IIdlName
	IIdlKind
}

type IDeclaredType interface {
	IIdlDefinition
	IsDefined() bool
	IsPrimitive() bool
}

type ITypeSpec interface {
	IDeclaredType
	Forward() bool
	Abstract() bool
	SetNextTypeSpec(next ITypeSpec) error
	GetNextTypeSpec() (ITypeSpec, error)
}

type IValueAbsoluteDefinition interface {
	IParentModule
}

type IDeclarator interface {
	IFileInformation
	GetIdentifier() string
	GetNext() IDeclarator
	SetNext(next IDeclarator)
	Next(next IDeclarator) IDeclarator
}

type IStructMemberInformation interface {
	GetId() string
	GetTypeSpec() IDeclaredType
}

type IStructMember interface {
	NextStructMember(next IStructMember) IStructMember
	Count() int
	GetMembers() []IStructMemberInformation
	DeclaredType() IDeclaredType
	GetDeclarator() IDeclarator
	GetNext() IStructMember
}

type IEnumerator interface {
	Last(next IEnumerator) IEnumerator
	SetLast(next IEnumerator)
	Next() IEnumerator
	Id() string
}

type IParentModule interface {
	ITypeSpec
	Iterate(cb func(typeSpec ITypeSpec) error) error
}

type IIdlModuleDcl interface {
	IParentModule
	GetModuleExports() ITypeSpec
	SetModuleExports(moduleExports ITypeSpec)
}

type IStructType interface {
	ITypeSpec
	Members() IStructMember
	FindMemberType(memberIdentifier string) IDeclaredType
}

type IInterfaceKind interface {
	IFileInformation
	Local() bool
	Abstract() bool
}

type ILocalDeclaration interface {
	Local() bool
}

type IInterfaceDcl interface {
	IParentModule
	ILocalDeclaration
	GetBody() ITypeSpec
	SetBody(ITypeSpec)
	BodyCount() int
	BodyArray() []IIdlDefinition
}

type ITypeDcl interface {
	ITypeSpec
}

type Iconstr_type_dcl interface {
	ITypeDcl
	//Iconstr_type_dcl() Iconstr_type_dcl
}

type Iunion_dcl interface {
	Iconstr_type_dcl
}

type Iunion_def interface {
	Iunion_dcl
}

type Iunion_forward_dcl interface {
	Iunion_dcl
}

type Ienum_dcl interface {
	Iconstr_type_dcl
}
type Ibitset_dcl interface {
	Iconstr_type_dcl
}
type Ibitmask_dcl interface {
	Iconstr_type_dcl
}

type Inative_dcl interface {
	ITypeDcl
}

type ITypedefDcl interface {
	ITypeDcl
}

type ITypeDeclarator interface {
	ITypedefDcl
	TypeSpec() IDeclaredType
	//Declarator() IDeclarator
}

type IReservedWordData interface {
	IFileInformation
}

type IIdlIdentifier interface {
	IFileInformation
	IIdlName
	Identifier() string
}

type IIdlValueKind interface {
	IFileInformation
}

type IInterfaceHeader interface {
	IFileInformation
	Identifier() string
	Local() bool
	Abstract() bool
}

type IInterfaceInheritanceSpec interface {
}

type IInterfaceNamePlus interface {
	Next(identifier string) IInterfaceNamePlus
}

type IAttrDeclarator interface {
	IFileInformation
	IIdlName
	Names() []string
}

type IValueInheritanceSpec interface {
	IFileInformation
}

type IIdlValueHeader interface {
	IFileInformation
}

type IScopedName interface {
	IIdlIdentifier
	NextScopedName(next IScopedName) error
	GetNextScopedName() (IScopedName, error)
}

type IParameterDeclarations interface {
	NextParameterDeclarations(IParameterDeclarations) IParameterDeclarations
	GetNextParameterDeclarations() IParameterDeclarations
	GetParamIn() bool
	GetParamOut() bool
	GetParamName() string
	GetParamDeclarationType() IDeclaredType
}

type IPrimaryExpressionType int

const (
	PetInteger IPrimaryExpressionType = iota
	PetString
	PetFloatingPoint
	PetWideString
	PetCharacter
	PetBoolean
	PetWideCharacter
	PetFixedPoint
)

type IPrimaryExpression interface {
	Type() IPrimaryExpressionType
	Value() interface{}
}

type IIdlComparer interface {
	Compare(x, y IIdlDefinition) (IIdlDefinition, error)
}

type IIdlCompareFactory interface {
	Create() IIdlComparer
}

type IIdlPlaceHolder interface {
	IdlName() string
	IdlType() IIdlDefinition
	DoCompare(x, iIdlDefinition IIdlDefinition) (IIdlDefinition, error)
}

type IEnumType interface {
	Ienum_dcl
	Enumerator() IEnumerator
}

type IIdlConstDcl interface {
	ITypeSpec
}

//noinspection GoSnakeCaseUsage
type IIdlExceptDcl interface {
	IIdlDefinition
	IIdlExceptDcl() IIdlExceptDcl
}

//noinspection GoSnakeCaseUsage
type IIdlEventDcl interface {
	IIdlDefinition
	Idlevent_dcl() IIdlEventDcl
}

//noinspection GoSnakeCaseUsage
type IIdlPortTypeDcl interface {
	IIdlDefinition
	IdlPortTypeDcl() IIdlPortTypeDcl
}

//noinspection GoSnakeCaseUsage
type IIdlhome_dcl interface {
	IIdlDefinition
	Idlhome_dcl() IIdlhome_dcl
}

//noinspection GoSnakeCaseUsage
type IIdlAnnotationDcl interface {
	IIdlDefinition
	IdlAnnotationDcl() IIdlAnnotationDcl
}

//noinspection GoSnakeCaseUsage
type IIdlComponentDcl interface {
	IIdlDefinition
	IdlComponentDcl() IIdlComponentDcl
}

//noinspection GoSnakeCaseUsage
type IIdlConnectorDcl interface {
	IIdlDefinition
	IdlConnectorDcl() IIdlConnectorDcl
}

type ITypePrefixDefinition interface {
	ITypeSpec
	StringLiteral() string
}

//noinspection GoSnakeCaseUsage
type IIdlTypeDcl interface {
	IIdlDefinition
	IdlTypeDcl() IIdlTypeDcl
}

//noinspection GoSnakeCaseUsage
type IIdlvalue_dcl interface {
	IIdlDefinition
	Idlvalue_dcl() IIdlvalue_dcl
}

//noinspection GoSnakeCaseUsage
type IIdltype_prefix_dcl interface {
	IIdlDefinition
	Idltype_prefix_dcl() IIdltype_prefix_dcl
}

//noinspection GoSnakeCaseUsage
type IIdltype_id_dcl interface {
	IIdlDefinition
	Idltype_id_dcl() IIdltype_id_dcl
}

//noinspection GoSnakeCaseUsage
type IIdlImportDcl interface {
	IIdlDefinition
	IdlImportDcl() IIdlImportDcl
}

type IInterfaceMember interface {
	Next(next IInterfaceMember) (IInterfaceMember, error)
}

type IIdlException interface {
	ITypeSpec
	GetMembers() IStructMember
}

type ISequenceType interface {
	TypeSpec() IDeclaredType
	Count() int
}

type IParamAttribute interface {
	In() bool
	Out() bool
	IFileInformation
}

type IAttributeDcl interface {
	ReadOnly() bool
	AttrDeclarator() IAttrDeclarator
	DeclaredType() IDeclaredType
}

type INextNumber interface {
	NextNumber() int
}

//noinspection GoSnakeCaseUsage
type IIdltemplate_module_dcl interface {
	IIdlDefinition
	Idltemplate_module_dcl() IIdltemplate_module_dcl
}

type IOperationDeclarations interface {
	ITypeSpec
	GetOperationDeclaratorType() IDeclaredType
	GetOperationName() string
	GetParams() IParameterDeclarations
	GetExceptionList() interface{}
}
