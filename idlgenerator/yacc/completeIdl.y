%{
package yacc

import "github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
import "github.com/bhbosman/CodeGenerators/idlgenerator/scopedObjects"

//go:generate goyacc -o completeIdl.go  -p "CompleteIdl"  completeIdl.y

%}

%token ScopeOp // "::"
%token ShlOp // "<<"
%token ShrOp // ">>"
%token Annotation //"@annotation"
%token string_literal
%token identifier
%token integer_literal
%token floating_pt_literal
%token fixed_pt_literal
%token character_literal
%token wide_character_literal
%token wide_string_literal
%token Hex_literal
%token RWabstract
%token RWany
%token RWalias
%token RWattribute
%token RWbitfield
%token RWbitmask
%token RWbitset
%token RWboolean
%token RWcase
%token RWchar
%token RWcomponent
%token RWconnector
%token RWconst
%token RWconsumes
%token RWcontext
%token RWcustom
%token RWdefault
%token RWdouble
%token RWexception
%token RWemits
%token RWenum
%token RWeventtype
%token RWfactory
%token RWFALSE
%token RWfinder
%token RWfixed
%token RWfloat
%token RWgetraises
%token RWhome
%token RWimport
%token RWin
%token RWinout
%token RWinterface
%token RWlocal
%token RWlong
%token RWmanages
%token RWmap
%token RWmirrorport
%token RWmodule
%token RWmultiple
%token RWnative
%token RWObject
%token RWoctet
%token RWoneway
%token RWout
%token RWprimarykey
%token RWprivate
%token RWport
%token RWporttype
%token RWprovides
%token RWpublic
%token RWpublishes
%token RWraises
%token RWreadonly
%token RWsetraises
%token RWsequence
%token RWshort
%token RWstring
%token RWstruct
%token RWsupports
%token RWswitch
%token RWTRUE
%token RWtruncatable
%token RWtypedef
%token RWtypeid
%token RWtypename
%token RWtypeprefix
%token RWunsigned
%token RWunion
%token RWuses
%token RWValueBase
%token RWvaluetype
%token RWvoid
%token RWwchar
%token RWwstring
%token RWint8
%token RWuint8
%token RWint16
%token RWint32
%token RWint64
%token RWuint16
%token RWuint32
%token RWuint64

%union{
	ReservedWord             ScopingInterfaces.IReservedWordData
	InterfaceKind            ScopingInterfaces.IInterfaceKind
	Identifier               ScopingInterfaces.IIdlIdentifier
	IntegerValue             int
	StringValue              string
	PrimaryExpression        ScopingInterfaces.IPrimaryExpression
	ModuleDcl                ScopingInterfaces.IIdlModuleDcl
	Definition               ScopingInterfaces.IIdlDefinition
	Member                   ScopingInterfaces.IStructMember
	TypeSpec                 ScopingInterfaces.ITypeSpec
	Declarator               ScopingInterfaces.IDeclarator
	Enumerator               ScopingInterfaces.IEnumerator
	TypedefDcl               ScopingInterfaces.ITypedefDcl
	TypeDeclarator           ScopingInterfaces.ITypeDeclarator
	TypeDcl                  ScopingInterfaces.ITypeDcl
	constr_type_dcl          ScopingInterfaces.Iconstr_type_dcl
	native_dcl               ScopingInterfaces.Inative_dcl
	union_dcl                ScopingInterfaces.Iunion_dcl
	enum_dcl                 ScopingInterfaces.Ienum_dcl
	bitset_dcl               ScopingInterfaces.Ibitset_dcl
	bitmask_dcl              ScopingInterfaces.Ibitmask_dcl
	struct_def               ScopingInterfaces.IStructType
	union_def                ScopingInterfaces.Iunion_def
	union_forward_dcl        ScopingInterfaces.Iunion_forward_dcl
	parameter_dcls           ScopingInterfaces.IParameterDeclarations
	value_kind               ScopingInterfaces.IIdlValueKind
	InterfaceHeader          ScopingInterfaces.IInterfaceHeader
	InterfaceInheritanceSpec ScopingInterfaces.IInterfaceInheritanceSpec
	InterfaceNamePlus        ScopingInterfaces.IInterfaceNamePlus
	AttrDeclarator           ScopingInterfaces.IAttrDeclarator
	ValueInheritanceSpec     ScopingInterfaces.IValueInheritanceSpec
	IdlValueHeader           ScopingInterfaces.IIdlValueHeader
	DeclaredType             ScopingInterfaces.IDeclaredType
	ScopedName               ScopingInterfaces.IScopedName
	ParamAttribute           ScopingInterfaces.IParamAttribute
}

%type <ParamAttribute> param_attribute
%type <IdlValueHeader> value_header
%type <ValueInheritanceSpec> value_inheritance_spec
%type <AttrDeclarator> readonly_attr_declarator attr_declarator
%type <Identifier>  value_namePlus
%type <InterfaceNamePlus> interface_namePlus
%type <InterfaceInheritanceSpec> interface_inheritance_spec
%type <value_kind> value_kind
%type <parameter_dcls> parameter_dcls param_dclPlus param_dcl
%type <InterfaceHeader> interface_header
%type <ReservedWord> RWmodule RWcustom RWsupports
%type <ReservedWord>  RWfixed RWvoid
%type <ReservedWord> RWstruct
%type <ReservedWord> RWfloat
%type <ReservedWord> RWdouble
%type <ReservedWord> RWlong
%type <ReservedWord> RWshort
%type <ReservedWord> RWint16
%type <ReservedWord> RWint32
%type <ReservedWord> RWint64
%type <ReservedWord> RWunsigned
%type <ReservedWord> RWuint16
%type <ReservedWord> RWuint32
%type <ReservedWord> RWuint64
%type <ReservedWord> RWstring
%type <ReservedWord> RWwstring
%type <ReservedWord> RWchar
%type <ReservedWord> RWwchar
%type <ReservedWord> RWboolean
%type <ReservedWord> RWoctet
%type <ReservedWord> RWany
%type <ReservedWord> RWObject
%type <ReservedWord> RWValueBase
%type <ReservedWord> RWint8
%type <ReservedWord> RWuint8
%type <ReservedWord> RWinterface
%type <ReservedWord> RWlocal
%type <ReservedWord> RWabstract
%type <ReservedWord> RWenum
%type <ReservedWord> RWin
%type <ReservedWord> RWinout
%type <ReservedWord> RWout
%type <ReservedWord> RWconst
%type <ReservedWord> RWexception
%type <ReservedWord> RWtypedef
%type <ReservedWord> RWvaluetype
%type <ReservedWord> RWtypeprefix
%type <ReservedWord> RWreadonly
%type <ReservedWord> RWattribute
%type <ReservedWord> RWtruncatable
%type <ReservedWord> RWsequence
%type <TypeSpec> op_dcl attr_dcl
%type <TypeSpec> op_oneway_dcl
%type <TypeSpec> op_with_context
%type <TypeSpec> readonly_attr_spec attr_spec
%type	<InterfaceKind>   interface_kind
%type	<PrimaryExpression> literal primary_expr
%type	<IntegerValue> unary_expr mult_expr add_expr shift_expr
%type	<Declarator>	fixed_array_sizePlus  array_declarator
%type	<Identifier>	identifier
%type	<Identifier>    raises_expr interface_name value_name
%type	<StringValue> 	string_literal
%type	<Member>    	member memberPlus memberStar
%type	<DeclaredType>    	type_spec
%type	<DeclaredType>    	op_type_spec
%type	<TypeSpec>    	value_abs_def
%type	<TypeSpec>    	value_def
%type	<DeclaredType>    	simple_type_spec
%type	<TypeSpec>    	value_forward_dcl
%type	<TypeSpec>    	value_box_def
%type	<DeclaredType>    	base_type_spec
%type	<DeclaredType>    	template_type_spec
%type	<DeclaredType>  floating_pt_type
%type	<DeclaredType>  integer_type const_type fixed_pt_const_type
%type	<DeclaredType>    	signed_int
%type	<DeclaredType>    	unsigned_int
%type	<DeclaredType>    	signed_short_int
%type	<DeclaredType>    	signed_long_int
%type	<DeclaredType>    	signed_longlong_int signed_tiny_int signed_tiny_int
%type	<DeclaredType>    	char_type
%type	<DeclaredType>    	wide_char_type
%type	<DeclaredType>    	boolean_type
%type	<DeclaredType>    	octet_type
%type	<DeclaredType>    	any_type
%type	<DeclaredType>    	object_type
%type	<DeclaredType>    	value_base_type
%type	<TypeSpec>    	sequence_type
%type	<DeclaredType>    	string_type
%type	<DeclaredType>    	wide_string_type
%type	<TypeSpec>    	interface_def interface_forward_dcl
%type	<TypeSpec>    	fixed_pt_type map_type
%type	<DeclaredType>    	unsigned_short_int
%type	<DeclaredType>    	unsigned_long_int
%type	<DeclaredType>    	unsigned_longlong_int unsigned_tiny_int


%type   <TypeSpec>	const_dcl
%type   <TypeSpec>	module_dcl
%type 	<Declarator> 	declarators declarator declaratorPlus
%type <Declarator> any_declarators
%type <Declarator> any_declaratorsPlus
%type <Declarator> any_declarator
%type	<native_dcl>    	 native_dcl
%type	<struct_def>    	 struct_dcl
%type	<union_dcl>    	 union_dcl
%type	<enum_dcl>    	 enum_dcl
%type	<bitset_dcl>    	 bitset_dcl
%type	<bitmask_dcl>    	 bitmask_dcl
%type	<union_def>    	 union_def
%type	<union_forward_dcl>    	 union_forward_dcl
%type	<TypeDeclarator>    	type_declarator
%type	<TypeSpec>   constr_type_dcl
%type	<TypeSpec>	type_dcl
%type	<TypeDeclarator>    typedef_dcl
%type	<struct_def>    	struct_def struct_forward_dcl
%type	<TypeSpec>    	except_dcl
%type	<TypeSpec>    	interface_dcl
%type	<TypeSpec>    	value_dcl
%type	<TypeSpec>    	type_id_dcl
%type	<TypeSpec>    	type_prefix_dcl
%type	<TypeSpec>    	import_dcl
%type	<TypeSpec>    	component_dcl
%type	<TypeSpec>    	home_dcl
%type	<TypeSpec>    	event_dcl
%type	<TypeSpec>    	porttype_dcl
%type	<TypeSpec>    	connector_dcl
%type	<TypeSpec>    	template_module_dcl
%type	<TypeSpec>    	template_module_inst
%type	<TypeSpec>    	annotation_dcl
%type 	<Enumerator> 	enumerator enumerators
%type 	<IntegerValue>  positive_int_const fixed_array_size
%type 	<IntegerValue>  unary_operator  const_expr or_expr xor_expr and_expr
%type 	<IntegerValue>  integer_literal
%type 	<IntegerValue>  floating_pt_literal
%type 	<IntegerValue>  fixed_pt_literal
%type 	<IntegerValue>  character_literal
%type 	<IntegerValue>  wide_character_literal
%type 	<IntegerValue>  boolean_literal
%type 	<IntegerValue>  wide_string_literal
%type 	<IntegerValue>  RWTRUE
%type 	<IntegerValue>  RWFALSE
%type 	<ScopedName>  scoped_name
%type 	<ScopedName>  scoped_namePlus
%type 	<ScopedName>  simple_declarator
%type 	<ScopedName>  simple_declaratorPlus


%type	<TypeSpec>    	definition
%type	<TypeSpec>    	definitionPlus
%type	<TypeSpec>    	specification


%type <TypeSpec> export
%type <TypeSpec> exportStar
%type <TypeSpec> exportPlus
%type	<TypeSpec>    	interface_body
%%

//(1)
specification:
	definitionPlus
	{
		CompleteIdllex.InfoAt("specification/definitionPlus", $1)
		var err error
		_, err = CompleteIdllex.AssignSpec($1)
		$$ = $1
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}

//(2)(71)//(98)//(111)//(133)//(144)//(153)//(171)//(184)//(218)
definitionPlus:
	definition
	{
		CompleteIdllex.InfoAt("definitionPlus/definition", $1)
		var err error
		$$, err = $1, nil
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| definitionPlus definition {
		CompleteIdllex.InfoAt("definitionPlus/definitionPlus definition", $1, $2)
		$1.SetNextTypeSpec($2)
		$$ = $1
	}

definition:
	module_dcl ';'
	{
		CompleteIdllex.InfoAt("definition/module_dcl", $1)
		$$ = $1

	}
	| const_dcl ';'
	{
		CompleteIdllex.InfoAt("definition/const_dcl", $1)
		$$= $1
	}
	| type_dcl ';'
	{
		CompleteIdllex.InfoAt("definition/type_dcl", $1)
		$$= $1
	}
	| except_dcl ';'
	{
		CompleteIdllex.InfoAt("definition/except_dcl", $1)
		$$= $1
	}
	| interface_dcl ';'
	{
		CompleteIdllex.InfoAt("definition/interface_dcl", $1)
		$$= $1
	}
	| value_dcl ';'
	{
		CompleteIdllex.InfoAt("definition/value_dcl", $1)
		$$= $1
	}
	| type_id_dcl ';'
	{
		CompleteIdllex.InfoAt("definition/type_id_dcl", $1)
		$$= $1
	}
	| type_prefix_dcl ';'
	{
		CompleteIdllex.InfoAt("definition/type_prefix_dcl", $1)
		$$= $1
	}
	| import_dcl ';'
	{
		CompleteIdllex.InfoAt("definition/import_dcl", $1)
		$$= $1
	}
	| component_dcl ';'
	{
		CompleteIdllex.InfoAt("definition/component_dcl", $1)
		$$= $1
	}
	| home_dcl ';'
	{
		CompleteIdllex.InfoAt("definition/home_dcl", $1)
		$$= $1
	}
	| event_dcl ';'
	{
		CompleteIdllex.InfoAt("definition/event_dcl", $1)
		$$= $1
	}
	| porttype_dcl ';'
	{
		CompleteIdllex.InfoAt("definition/porttype_dcl", $1)
		$$= $1
	}
	| connector_dcl ';'
	{
		CompleteIdllex.InfoAt("definition/connector_dcl", $1)
		$$= $1
	}
	| template_module_dcl ';'
	{
		CompleteIdllex.InfoAt("definition/template_module_dcl", $1)
		$$= $1
	}
	| template_module_inst ';'
	{
		CompleteIdllex.InfoAt("definition/template_module_inst", $1)
		$$= $1
	}
	| annotation_dcl ';'
	{
		CompleteIdllex.InfoAt("definition/annotation_dcl", $1)
	}

//(3)
module_dcl:
	RWmodule identifier '{' definitionPlus '}'
	{
		CompleteIdllex.InfoAt("RWmodule identifier '{' definitionPlus '}'", $2, $4)
		var err error
		$$, err = CompleteIdllex.CreateModuleDcl($1, $2.Identifier(), $4)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| RWmodule identifier '{'  '}'
	{
		CompleteIdllex.InfoAt("RWmodule identifier '{'  '}'", $2)
		var err error
		$$, err = CompleteIdllex.CreateModuleDcl($1, $2.Identifier(), nil)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}


//(4)
scoped_namePlus:
	scoped_name
	{
		CompleteIdllex.InfoAt("scoped_namePlus/scoped_name", $1)
		$$ = scopedObjects.NewScopedName02($1, $1.Identifier())
	}
	| scoped_namePlus ',' scoped_name
	{
		CompleteIdllex.InfoAt("scoped_namePlus/scoped_namePlus ',' scoped_name", $1, $3)
		$$ = $1
		err := $1.NextScopedName($3)
		if err != nil{
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
scoped_name:
	identifier
	{
		CompleteIdllex.InfoAt("scoped_name/identifier", $1)
		$$ = scopedObjects.NewScopedName02($1, $1.Identifier())
	}
	| ScopeOp identifier
	{
		CompleteIdllex.InfoAt("scoped_name/ScopeOp identifier", $2)
		$$ = scopedObjects.NewScopedName02($2, "::" + $2.Identifier())
	}
	| scoped_name ScopeOp identifier
	{
		CompleteIdllex.InfoAt("scoped_name/scoped_name ScopeOp identifier", $1, $3)
		$$ = scopedObjects.NewScopedName02($1, $1.Identifier() +  "::" + $3.Identifier())
	}
//(5)
const_dcl:
	RWconst const_type identifier '=' const_expr
	{
		CompleteIdllex.InfoAt("RWconst const_type identifier '=' const_expr", $1, $2, $3 , $5)

		var err error
		$$, err = CompleteIdllex.NewIdlConstDcl($1, $3.Identifier(), $5)
		if err != nil{
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(6)
const_type:
	integer_type
	{
		CompleteIdllex.InfoAt("const_type/integer_type", $1)
		$$ = $1
	}
	| floating_pt_type
	{
		CompleteIdllex.InfoAt("const_type/floating_pt_type", $1)
		$$ = $1
	}
	| fixed_pt_const_type
	{
		CompleteIdllex.InfoAt("const_type/fixed_pt_const_type", $1)
		$$ = $1
	}
	| char_type
	{
		CompleteIdllex.InfoAt("const_type/char_type", $1)
		$$ = $1
	}
	| wide_char_type
	{
		CompleteIdllex.InfoAt("const_type/wide_char_type", $1)
		$$ = $1
	}
	| boolean_type
	{
		CompleteIdllex.InfoAt("const_type/boolean_type", $1)
		$$ = $1
	}
	| octet_type
	{
		CompleteIdllex.InfoAt("const_type/octet_type", $1)
		$$ = $1
	}
	| string_type
	{
		CompleteIdllex.InfoAt("const_type/string_type", $1)
		$$ = $1
	}
	| wide_string_type
	{
		CompleteIdllex.InfoAt("const_type/wide_string_type", $1)
		$$ = $1
	}
	| scoped_name
	{
		var err error
		$$, err = scopedObjects.NewEmptyIdlDefinition($1, $1.Identifier()), nil
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(7)
const_expr:
	or_expr
	{
		CompleteIdllex.InfoAt("const_expr/or_expr", $1)
		$$ = $1
	}
//(8)
or_expr:
	xor_expr
	{
		CompleteIdllex.InfoAt("or_expr/xor_expr", $1)
		$$ = $1
	}
	| or_expr '|' xor_expr
	{
		CompleteIdllex.InfoAt("or_expr/or_expr '|' xor_expr", $1, $3)
		$$ = $1 | $3
	}
//(9)
xor_expr:
	and_expr
	{
		CompleteIdllex.InfoAt("xor_expr/and_expr", $1)
		$$ = $1
	}
	| xor_expr '^' and_expr
	{
		CompleteIdllex.InfoAt("xor_expr/xor_expr '^' and_expr", $1, $3)
		$$ = $1 ^ $3
	}

//(10)
and_expr:
	shift_expr
	{
		CompleteIdllex.InfoAt("and_expr/shift_expr", $1)
		$$ = $1
	}
	| and_expr '&' shift_expr
	{
		CompleteIdllex.InfoAt("and_expr/and_expr '&' shift_expr", $1, $3)
		$$ = $1 & $3
	}

//(11)
shift_expr:
	add_expr
	{
		CompleteIdllex.InfoAt("shift_expr/add_expr", $1)
		var err error
		$$, err = CompleteIdllex.AddExpr($1)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| shift_expr ShrOp add_expr
	{
		CompleteIdllex.InfoAt("shift_expr/shift_expr ShrOp add_expr", $1, $3)
		if $3 < 0 {
			CompleteIdllex.Error("shift count must be positive")
		}
		$$ = $1 >> uint64($3)
	}
	| shift_expr ShlOp add_expr
	{
		CompleteIdllex.InfoAt("shift_expr/shift_expr ShlOp add_expr", $1, $3)
		if $3 < 0 {
			CompleteIdllex.Error("shift count must be positive")
		}

		$$ = $1 << uint64($3)
	}
//(12)
add_expr:
	mult_expr
	{
		CompleteIdllex.InfoAt("add_expr/mult_expr", $1)
		var err error
		$$, err = CompleteIdllex.MultiExpr($1)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| add_expr '+' mult_expr
	{
		CompleteIdllex.InfoAt("add_expr/add_expr '+' mult_expr", $1, $3)
		var err error
		$$, err = CompleteIdllex.AddOperator($1, $3)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| add_expr '-' mult_expr
	{
		CompleteIdllex.InfoAt("add_expr/add_expr '-' mult_expr", $1, $3)
		var err error
		$$, err = CompleteIdllex.MinusOperator($1, $3)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(13)
mult_expr:
	unary_expr
	{
		CompleteIdllex.InfoAt("mult_expr/unary_expr", $1)
		$$ = $1
	}
	| mult_expr '*' unary_expr
	{
		CompleteIdllex.InfoAt("mult_expr/mult_expr '*' unary_expr", $1, $3)
		$$ = $1 * $3
	}
	| mult_expr '/' unary_expr
	{
		CompleteIdllex.InfoAt("mult_expr/mult_expr '/' unary_expr", $1, $3)
		$$ = $1 / $3
	}
	| mult_expr '%' unary_expr
	{
		CompleteIdllex.InfoAt("mult_expr/mult_expr '%' unary_expr", $1, $3)
		$$ = $1 % $3
	}
//(14)
unary_expr:
	unary_operator primary_expr
	{
		CompleteIdllex.InfoAt("unary_expr/unary_operator primary_expr", $1, $2)
		if $2.Type() == ScopingInterfaces.PetInteger{
			$$ = $1 * $2.Value().(int)
		} else {
			CompleteIdllex.Error("expression not an integer")
			return 1
		}
	}
	| primary_expr
	{
		CompleteIdllex.InfoAt("unary_expr/primary_expr", $1)
		if $1.Type() == ScopingInterfaces.PetInteger{
			$$ = $1.Value().(int)
		} else {
			CompleteIdllex.Error("expression not an integer")
			return 1
		}


	}
//(15)
unary_operator:
	'-'
	{
		$$ = -1
	}
	| '+'
	{
		$$ = +1
	}
	| '~'
	{

	}
//(16)
primary_expr:
	scoped_name
	{
		CompleteIdllex.InfoAt("primary_expr/scoped_name", $1)
//		v, err := strconv.Atoi($1)
//		if err != nil {
//			CompleteIdllex.Error(err.Error())
//			return 1
//		}
		CompleteIdllex.Error("implement me......... primary_expr/scoped_name")
		//$$ = int(v)
	}
	| literal
	{
		CompleteIdllex.InfoAt("primary_expr/literal", $1)
		$$ = $1
	}
	| '(' const_expr ')'
	{
		CompleteIdllex.InfoAt("primary_expr/'(' const_expr ')'", $2)
		var err error
		$$, err = CompleteIdllex.TransformInteger($2)
		if err != nil{
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(17)
literal:
	integer_literal
	{
		CompleteIdllex.InfoAt("literal/integer_literal", $1)
		var err error
		$$, err = CompleteIdllex.TransformInteger($1)
		if err != nil{
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| floating_pt_literal
	{
		CompleteIdllex.InfoAt("literal/floating_pt_literal", $1)
		var err error
		$$, err = CompleteIdllex.TransformValue($1, ScopingInterfaces.PetFloatingPoint)
		if err != nil{
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| fixed_pt_literal
	{
		CompleteIdllex.InfoAt("literal/fixed_pt_literal", $1)
		var err error
		$$, err = CompleteIdllex.TransformValue($1, ScopingInterfaces.PetFixedPoint)
		if err != nil{
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| character_literal
	{
		CompleteIdllex.InfoAt("literal/character_literal", $1)
		var err error
		$$, err = CompleteIdllex.TransformValue($1, ScopingInterfaces.PetCharacter)
		if err != nil{
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| wide_character_literal
	{
		CompleteIdllex.InfoAt("literal/wide_character_literal", $1)
		var err error
		$$, err = CompleteIdllex.TransformValue($1, ScopingInterfaces.PetWideCharacter)
		if err != nil{
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| boolean_literal
	{
		CompleteIdllex.InfoAt("literal/boolean_literal", $1)
		var err error
		$$, err = CompleteIdllex.TransformValue($1, ScopingInterfaces.PetBoolean)
		if err != nil{
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| string_literal
	{
		CompleteIdllex.InfoAt("literal/string_literal", $1)
		var err error
		$$, err = CompleteIdllex.TransformString($1)
		if err != nil{
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| wide_string_literal
	{
		CompleteIdllex.InfoAt("literal/wide_string_literal", $1)
		var err error
		$$, err = CompleteIdllex.TransformValue($1, ScopingInterfaces.PetWideString)
		if err != nil{
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(18)
boolean_literal:
	RWTRUE
	{
		$$ = $1
	}
	| RWFALSE
	{
		$$ = $1
	}
//(19)
positive_int_const:
	const_expr
	{
		CompleteIdllex.InfoAt("const_expr", $1)
		if $1 < 0 {
			CompleteIdllex.Error("expression must be positive")
			return 1
		}
		$$ = $1
	}
//(20)
type_dcl:
	constr_type_dcl
	{
		$$ = $1
	}
	| native_dcl
	{
		$$ = $1
	}
	| typedef_dcl
	{
		$$ = $1
	}
//(21)
type_spec:
	simple_type_spec
	{
		CompleteIdllex.InfoAt("type_spec/simple_type_spec", $1)
		$$ = $1
	}
	|template_type_spec
	{
		CompleteIdllex.InfoAt("type_spec/template_type_spec")
		$$ = $1
	}
//(22)
// ISimpleTypeSpec
simple_type_spec:
	base_type_spec
	{
		CompleteIdllex.InfoAt("simple_type_spec/base_type_spec", $1)
		$$ =$1
	}
	| scoped_name
	{
		CompleteIdllex.InfoAt("simple_type_spec/scoped_name", $1)
		var err error
		$$, err = scopedObjects.NewEmptyIdlDefinition($1, $1.Identifier()), nil
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}

	}
//(23)(69)//(117)//(131)
base_type_spec:
	integer_type
	{
		CompleteIdllex.InfoAt("base_type_spec/integer_type", $1)
        	$$ = $1
        }
	| floating_pt_type
	{
		CompleteIdllex.InfoAt("base_type_spec/floating_pt_type", $1)
		$$ = $1
	}
	| char_type
	{
		CompleteIdllex.InfoAt("base_type_spec/char_type", $1)
		$$ = $1
	}
	| wide_char_type
	{
		CompleteIdllex.InfoAt("base_type_spec/wide_char_type", $1)
		$$ = $1
	}
	| boolean_type
	{
		CompleteIdllex.InfoAt("base_type_spec/boolean_type", $1)
		$$ = $1
	}
	| octet_type
	{
		CompleteIdllex.InfoAt("base_type_spec/octet_type", $1)
		$$ =$1
	}
	| any_type
	{
		CompleteIdllex.InfoAt("base_type_spec/any_type", $1)
		$$ =$1
	}
	| object_type
	{
		CompleteIdllex.InfoAt("base_type_spec/object_type", $1)
		$$ =$1
	}
	| value_base_type
	{
		CompleteIdllex.InfoAt("base_type_spec/value_base_type", $1)
		$$ =$1
	}

//(24)
floating_pt_type:
	RWfloat {
		CompleteIdllex.InfoAt("RWfloat")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "float")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| RWdouble {
		CompleteIdllex.InfoAt("RWdouble")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "double")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| RWlong RWdouble {
		CompleteIdllex.InfoAt("RWlong RWdouble")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "long double")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(25)
integer_type:
	signed_int
	{
		$$ = $1
	}
	| unsigned_int
	{
		$$ = $1
	}

//(26)//(206)
signed_int:
	signed_short_int {
		$$ =$1
	}
	| signed_long_int {
		$$ =$1
	}
	| signed_longlong_int {
		$$ =$1
	}
	| signed_tiny_int {
		$$ =$1
	}
//(27)//(210)
signed_short_int:
	RWshort
	{
		CompleteIdllex.InfoAt("RWshort")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "short")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| RWint16
	{
		CompleteIdllex.InfoAt("RWint16")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "int16")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}

	}
//(28)//(211)
signed_long_int:
	RWlong
	{
		CompleteIdllex.InfoAt("RWlong")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "long")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}

	}
	| RWint32
	{
		CompleteIdllex.InfoAt("RWint32")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "int32")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}

	}

//(29)//(212)
signed_longlong_int:
	RWlong RWlong
	{
		CompleteIdllex.InfoAt("RWlong RWlong")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "long long")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	|RWint64
	{
		CompleteIdllex.InfoAt("RWint64")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "int64")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}

//(30)//(207)
unsigned_int:
	unsigned_short_int
	{
		CompleteIdllex.InfoAt("unsigned_short_int", $1)
		$$ = $1
	}
	| unsigned_long_int {
		CompleteIdllex.InfoAt("unsigned_long_int", $1)
		$$ = $1
	}
	| unsigned_longlong_int {
		CompleteIdllex.InfoAt("unsigned_longlong_int", $1)
		$$ = $1
	}
	|unsigned_tiny_int {
		$$ = $1
	}

//(31)//(213)
unsigned_short_int:
	RWunsigned RWshort
	{
		CompleteIdllex.InfoAt("RWunsigned RWshort")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "unsigned short")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	|RWuint16
	{
		CompleteIdllex.InfoAt("RWuint16")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "uint16")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}

//(32)//(214)
unsigned_long_int:
	RWunsigned RWlong
	{
		CompleteIdllex.InfoAt("RWunsigned RWlong")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "unsigned long")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| RWuint32
	{
		CompleteIdllex.InfoAt("RWuint32")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "uint32")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}

//(33)//(215)
unsigned_longlong_int:
	RWunsigned RWlong RWlong
	{
		CompleteIdllex.InfoAt("RWunsigned RWlong RWlong")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "unsigned long long")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| RWuint64
	{
		CompleteIdllex.InfoAt("RWuint64")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "uint64")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}

//(34)
char_type:
	RWchar
	{
		CompleteIdllex.InfoAt("RWchar")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "char")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(35)
wide_char_type:
	RWwchar {
		CompleteIdllex.InfoAt("RWwchar")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "wchar")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(36)
boolean_type:
	RWboolean
	{
		CompleteIdllex.InfoAt("RWboolean")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "boolean")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(37)
octet_type:
	RWoctet
	{
		CompleteIdllex.InfoAt("RWoctet")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "octet")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(38)//(197)
template_type_spec:
	sequence_type
	{
		CompleteIdllex.InfoAt("sequence_type")
		$$ =$1
	}
	| string_type
	{
		CompleteIdllex.InfoAt("string_type")
		$$ =$1
	}
	| wide_string_type
	{
		CompleteIdllex.InfoAt("wide_string_type")
		$$ =$1
	}
	| fixed_pt_type
	{
		CompleteIdllex.InfoAt("fixed_pt_type")
		$$ =$1
	}
	| map_type
	{
		CompleteIdllex.InfoAt("map_type")
		$$ =$1
	}

//(39)
sequence_type:
	RWsequence '<' type_spec ',' positive_int_const '>'
	{
		CompleteIdllex.InfoAt("RWsequence '<' type_spec ',' positive_int_const '>'", $3, $5)
		if $5 <= 0 {
			CompleteIdllex.Error("need a positive int")
			return 1
		}
		var err error
		$$, err = scopedObjects.NewSequenceType(
			$1,
			$3, $5)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| RWsequence '<' type_spec '>'
	{
		CompleteIdllex.InfoAt("RWsequence '<' type_spec '>'", $3)
		var err error
		$$, err = scopedObjects.NewSequenceType(
			$1,
			$3, 0)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(40)
string_type:
	RWstring '<' positive_int_const '>'
	{
		CompleteIdllex.InfoAt("string_type/RWstring '<' positive_int_const '>'", $3)
		if $3 <= 0 {
			CompleteIdllex.Error("need a positive int")
			return 1
		}
		var err error
		$$, err = scopedObjects.NewStringType($1, $3)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| RWstring
	{
		CompleteIdllex.InfoAt("string_type/| RWstring")
		var err error
		$$, err = scopedObjects.NewStringType($1, 0)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(41)
wide_string_type:
	RWwstring '<' positive_int_const '>'
	{
		CompleteIdllex.InfoAt("wide_string_type/RWwstring '<' positive_int_const '>'", $3)
		if $3 <= 0 {
			CompleteIdllex.Error("need a positive int")
			return 1
		}
		var err error
		$$, err = scopedObjects.NewWideStringType($1, $3)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| RWwstring
	{
		CompleteIdllex.InfoAt("wide_string_type/RWwstring")
		var err error
		$$, err = scopedObjects.NewWideStringType($1, 0)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(42)
fixed_pt_type:
	RWfixed '<' positive_int_const ',' positive_int_const '>'
	{
		CompleteIdllex.InfoAt("fixed_pt_type/RWfixed '<' positive_int_const ',' positive_int_const '>'", $3)
		if $3 <= 0 {
			CompleteIdllex.Error("need a positive int")
			return 1
		}
		if $5 <= 0 {
			CompleteIdllex.Error("need a positive int")
			return 1
		}
		var err error
		$$, err = scopedObjects.NewFixedPointType($1, $3, $5)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(43)
fixed_pt_const_type:
	RWfixed
	{
		CompleteIdllex.Error("fix this")
		return 1
	}
//(44)//(198)
constr_type_dcl:
	struct_dcl
	{
		$$ = $1
	}
	| union_dcl
	{
		$$ = $1
	}
	| enum_dcl
	{
		$$ = $1
	}
	| bitset_dcl
	{
		$$ = $1
	}
	| bitmask_dcl
	{
		$$ = $1
	}

//(45)
struct_dcl:
	struct_def
	{
		$$ = $1
	}
	| struct_forward_dcl
	{
	 	$$ = $1
	}
//(46)//(195)
struct_def:
	RWstruct identifier '{' memberPlus '}'{
		CompleteIdllex.InfoAt("RWstruct identifier '{' memberPlus '}'", $1, $2, $4)
		var err error
		$$, err = CompleteIdllex.NewStructType($1, $2.Identifier(), $4, false)
		if err != nil{
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| RWstruct identifier ':' scoped_name '{' memberStar '}'
	{
		var err error
		$$, err =CompleteIdllex.NewStructType($1, $2.Identifier(), $6, false)
		if err != nil{
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| RWstruct identifier '{' '}'
	{
		var err error
		$$, err = CompleteIdllex.NewStructType($1, $2.Identifier(), nil, false)
		if err != nil{
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}

//(47)
memberPlus:
	member{
		CompleteIdllex.InfoAt("memberPlus/member")
		$$ = $1
	}
	|memberPlus member{
		CompleteIdllex.InfoAt("memberPlus/memberPlus member")
		$$ = $1.NextStructMember($2)
	}
memberStar:
	{
		CompleteIdllex.InfoAt("memberStar")
		$$ = nil
	}
	|memberPlus

member:
	type_spec declarators ';'
	{
		CompleteIdllex.InfoAt("member/type_spec declarators", $1, $2)
		var err error
		$$, err = CompleteIdllex.NewMember($1, $2)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(48)
struct_forward_dcl:
	RWstruct identifier
	{
		CompleteIdllex.InfoAt("struct_forward_dcl/RWstruct identifier")
		var err error
		$$, err = CompleteIdllex.NewStructType($1, $2.Identifier(), nil, true)
		if err != nil{
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(49)
union_dcl:
	union_def
	{
		$$ = $1
	}
	| union_forward_dcl
	{
		$$ = $1
	}
//(50)
union_def:
	RWunion identifier RWswitch '(' switch_type_spec ')' '{' switch_body '}'
	{	CompleteIdllex.Error("implement union_def")
		return 1
	}
//(51)//(196)
switch_type_spec:
	integer_type
	| char_type
	| boolean_type
	| scoped_name
	| wide_char_type
	| octet_type

//(52)
switch_body: casePlus
//(53)
casePlus:
	case
	|casePlus case
case: case_labelPlus element_spec ';'
//(54)
case_labelPlus:
case_label: RWcase const_expr ':' | RWdefault ':'
//(55)
element_spec:
	type_spec declarator
	{

	}
//(56)
union_forward_dcl:
	RWunion identifier
	{
		CompleteIdllex.Error("implement union_forward_dcl")
		return 1
	}
//(57)
enum_dcl:
	RWenum identifier '{' enumerators '}'
	{
		var err error
		$$, err = scopedObjects.NewEnumType($1, $2.Identifier(), $4)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
enumerators:
	enumerator
	{
		$$ = $1
		CompleteIdllex.InfoAt("enumerators/enumerator| enumerators ',' enumerator", $1)
	}
	| enumerators ',' enumerator
	{
		$$ = $1.Last($3)
		CompleteIdllex.InfoAt("enumerators/enumerator| enumerators ',' enumerator", $1, $3)

	}
//(58)
enumerator:
	identifier {
		CompleteIdllex.InfoAt("enumerator/identifier", $1)

		var err error
		$$, err = scopedObjects.NewEnumerator($1, $1.Identifier())
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(59)
array_declarator:
	simple_declarator fixed_array_sizePlus
	{
		CompleteIdllex.InfoAt("array_declarator/identifier fixed_array_sizePlus", $1, $2)
		var err error
		$$, err = CompleteIdllex.NewDeclarator($1, $1.Identifier())
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(60)
fixed_array_sizePlus:
	fixed_array_size
	{
		CompleteIdllex.InfoAt("fixed_array_size", $1)
		$$ = nil
	}
	|fixed_array_sizePlus fixed_array_size
	{
		CompleteIdllex.InfoAt("fixed_array_sizePlus fixed_array_size", $1, $2)
		$$ = nil
	}
fixed_array_size:
	'[' positive_int_const ']'
	{
		CompleteIdllex.InfoAt("'[' positive_int_const ']'", $2)
		$$ = $2

	}
//(61)
native_dcl:
	RWnative simple_declarator
	{
		CompleteIdllex.Error("implement native_dcl")
		return 1
	}
//(62)
simple_declaratorPlus:
	simple_declarator
	{
		CompleteIdllex.InfoAt("simple_declaratorPlus/simple_declarator")
		$$ = scopedObjects.NewScopedName02($1, $1.Identifier())
	}|
	simple_declaratorPlus ',' simple_declarator
	{
		CompleteIdllex.InfoAt("simple_declaratorPlus/simple_declaratorPlus ',' simple_declarator")
		$1.NextScopedName($3)
		$$ = $1

	}

simple_declarator:
	identifier
	{
		$$ = scopedObjects.NewScopedName02($1, $1.Identifier())
	}
//(63)
typedef_dcl:
	RWtypedef type_declarator
	{
		CompleteIdllex.InfoAt("typedef_dcl/RWtypedef type_declarator", $2)
		$$ = $2
	}
//(64)
// ITypeDeclarator interface
type_declarator:
	simple_type_spec any_declarators
	{
		CompleteIdllex.InfoAt("type_declarator/simple_type_spec any_declarators", $1)
		var err error = nil
		$$, err = CompleteIdllex.NewTypeDeclarator($1, $2)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| template_type_spec any_declarators
	{
		CompleteIdllex.InfoAt("type_declarator/template_type_spec any_declarators", $1)
		var err error = nil
		$$, err = CompleteIdllex.NewTypeDeclarator($1, $2)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| constr_type_dcl any_declarators
	{
		CompleteIdllex.InfoAt("type_declarator/constr_type_dcl any_declarators", $1)
		var err error = nil
		$$, err = CompleteIdllex.NewTypeDeclarator($1, $2)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(65)
any_declarators:
	any_declaratorsPlus
	{
		CompleteIdllex.InfoAt("any_declarators/any_declaratorsPlus", $1)
		$$ = $1
	}
any_declaratorsPlus:
	any_declarator
	{
		CompleteIdllex.InfoAt("any_declaratorsPlus/any_declarator", $1)
		$$ = $1

	}
	| any_declaratorsPlus ',' any_declarator
	{
		CompleteIdllex.InfoAt("any_declaratorsPlus/any_declaratorsPlus ',' any_declarator", $1, $3)
		$$ = $1.Next($3)

	}

//(66)
any_declarator:
	simple_declarator
	{
		CompleteIdllex.InfoAt("declarator/simple_declarator", $1)
		var err error
		$$, err = CompleteIdllex.NewDeclarator($1, $1.Identifier())
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| array_declarator
	{

	}
//(67)
declarators:
	declaratorPlus
	{
		CompleteIdllex.InfoAt("declaratorPlus", $1)
		$$ = $1
	}
//(68)//(217)
declaratorPlus:
	declarator
	{
		CompleteIdllex.InfoAt("declarator", $1)
		$$ = $1
	}
	| declaratorPlus ',' declarator
	{
		CompleteIdllex.InfoAt("declaratorPlus ',' declarator", $1, $3)
		$$ = $1.Next($3)
	}
declarator:
	simple_declarator
	{
		CompleteIdllex.InfoAt("declarator/simple_declarator", $1)
		var err error
		$$, err = CompleteIdllex.NewDeclarator($1, $1.Identifier())
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	|array_declarator
	{
         	CompleteIdllex.InfoAt("array_declarator", $1)
        }


//(70)
any_type:
	RWany
	{
		CompleteIdllex.InfoAt("")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "any")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(72)
except_dcl:
	RWexception identifier '{' memberStar '}'
	{
		CompleteIdllex.InfoAt("RWexception identifier '{' memberStar '}'", $1, $2)

		var err error
		$$, err = scopedObjects.NewIdlRwException($1, $2.Identifier(), $4)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(73)
interface_dcl:
	interface_def
	{
		CompleteIdllex.InfoAt("interface_dcl/interface_def", $1)
		$$ = $1
	}
	|
	interface_forward_dcl
	{
		CompleteIdllex.InfoAt("interface_dcl/interface_forward_dcl", $1)
		$$ = $1
	}
//(74)
interface_def:
	interface_header '{' interface_body '}'
	{
		CompleteIdllex.InfoAt("interface_def/interface_header '{' interface_body '}'", $1, $3)
		var err error
		$$, err = CompleteIdllex.CreateInterfaceDcl(
			$1,
			$1.Identifier(),
			false,
			$1.Abstract(),
			$1.Local(),
			$3)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(75)
interface_forward_dcl:
	interface_kind identifier
	{
		CompleteIdllex.InfoAt("interface_forward_dcl/interface_kind identifier", $2)
		var err error
		$$, err = CompleteIdllex.CreateInterfaceDcl(
			$1,
			$2.Identifier(),
			true,
			$1.Abstract(),
			$1.Local(),
			nil)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(76)
interface_header:
	interface_kind identifier
	{
		CompleteIdllex.InfoAt("interface_header/interface_kind identifier", $1, $2)
		var err error
		$$, err = scopedObjects.NewInterfaceHeader(
			$1,
			$2.Identifier(),
			$1.Local(),
			$1.Abstract(),
			nil)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}

	|
	interface_kind identifier interface_inheritance_spec
	{
		CompleteIdllex.InfoAt("interface_header/interface_kind identifier interface_inheritance_spec", $1, $2)
		CompleteIdllex.InfoAt("interface_header/interface_kind identifier", $1, $2, $3)
		var err error
		$$, err = scopedObjects.NewInterfaceHeader(
			$1,
			$2.Identifier(),
			$1.Local(),
			$1.Abstract(),
			$3)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}

	}
//(77)//(119)//(129)
interface_kind:
	RWinterface
	{
		CompleteIdllex.InfoAt("interface_kind/RWinterface", $1)
		var err error
		$$, err = CompleteIdllex.CreateInterfaceKind($1, false, false)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| RWlocal RWinterface
	{
		CompleteIdllex.InfoAt("interface_kind/RWlocal RWinterface", $1)
		var err error
		$$, err = CompleteIdllex.CreateInterfaceKind($1, true, false)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	|RWabstract RWinterface
	{
		CompleteIdllex.InfoAt("interface_kind/RWabstract RWinterface", $1)
		var err error
		$$, err = CompleteIdllex.CreateInterfaceKind($1, false, true)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}


//(78)
interface_inheritance_spec:
	':' interface_namePlus
	{
		CompleteIdllex.InfoAt("interface_inheritance_spec/':' interface_namePlus")
		var err error
		$$, err = scopedObjects.NewInterfaceInheritanceSpec()
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(79)
interface_namePlus:
	interface_name
	{
		CompleteIdllex.InfoAt("interface_namePlus/interface_name")
		var err error
		$$, err = scopedObjects.NewInterfaceNamePlus($1.Identifier())
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| interface_namePlus ',' interface_name
	{
		CompleteIdllex.InfoAt("interface_namePlus/interface_namePlus ',' interface_name")
		var err error
		$$ = $1.Next($3.Identifier())
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
interface_name:
	scoped_name
	{
		CompleteIdllex.InfoAt("interface_name/scoped_name")
		$$ = $1
	}
//(80)
interface_body:
	exportStar
	{
		CompleteIdllex.InfoAt("interface_body/exportStar")
		$$ =$1
	}
	|{

	}
//(81)//(97)//(112)
exportStar:
	{
		CompleteIdllex.InfoAt("exportStar")
		$$ = nil
	}
	| exportPlus
	{
		$$ = $1
	}

exportPlus:
	exportPlus export {
		CompleteIdllex.InfoAt("exportPlus/exportPlus export", $1, $2)
		var err error
		$$, err = $1, $1.SetNextTypeSpec($2)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| export {
		CompleteIdllex.InfoAt("exportPlus/export", $1)
		var err error
		$$, err = $1, nil
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}

export:
	op_dcl ';'
	{
		CompleteIdllex.InfoAt("export/op_dcl ';'", $1)
		$$ = $1
	}
	| attr_dcl ';'
	{
		CompleteIdllex.InfoAt("export/attr_dcl ';'", $1)
		$$ = $1
	}
	| type_dcl ';'
	{
		CompleteIdllex.InfoAt("export/type_dcl ';'", $1)
		$$ = $1
	}
	| const_dcl ';'
	{
		CompleteIdllex.InfoAt("export/const_dcl ';'", $1)
		$$ = $1
	}
	| except_dcl ';'
	{
		CompleteIdllex.InfoAt("export/except_dcl ';'", $1)
		$$ = $1
	}
	| type_id_dcl ';'
	{
		CompleteIdllex.InfoAt("export/type_id_dcl ';'", $1)
		$$ = $1
	}
	| type_prefix_dcl ';'
	{
		CompleteIdllex.InfoAt("export/type_prefix_dcl ';'", $1)
		$$ = $1
	}
	| import_dcl ';'
	{
		CompleteIdllex.InfoAt("export/import_dcl ';'", $1)
		$$ = $1
	}
	| op_oneway_dcl ';'
	{
		CompleteIdllex.InfoAt("export/op_oneway_dcl ';'", $1)
		$$ = $1
	}
	| op_with_context ';'
	{
		CompleteIdllex.InfoAt("export/op_with_context ';'", $1)
		$$ = $1
	}

//(82)
op_dcl:
	op_type_spec identifier '('  ')'
	{
		CompleteIdllex.InfoAt("op_dcl/op_type_spec identifier '('  ')'", $1, $2)
		var err error
		$$, err = scopedObjects.NewOperationDeclarations($2, $2.Identifier(), $1, nil, nil), nil
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	|op_type_spec identifier '('  ')'  raises_expr
	{
		CompleteIdllex.InfoAt("op_dcl/op_type_spec identifier '('  ')'  raises_expr", $1, $2, $5)
		var err error
		$$, err = scopedObjects.NewOperationDeclarations($2, $2.Identifier(), $1, nil, $5), nil
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	|op_type_spec identifier '('  parameter_dcls ')'
	{
		CompleteIdllex.InfoAt("op_dcl/op_type_spec identifier '('  parameter_dcls ')'", $1, $2, $4)
		var err error
		$$, err = scopedObjects.NewOperationDeclarations($2, $2.Identifier(), $1, $4, nil), nil
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	|op_type_spec identifier '('  parameter_dcls ')'  raises_expr
	{
		CompleteIdllex.InfoAt("op_dcl/op_type_spec identifier '('  parameter_dcls ')'  raises_expr", $1, $2, $4, $6)
		var err error
		$$, err = scopedObjects.NewOperationDeclarations($2, $2.Identifier(), $1, $4, $6), nil
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(83)
op_type_spec:
	type_spec
	{
		$$ = $1
	}| RWvoid
	{
		CompleteIdllex.InfoAt("op_type_spec/RWvoid")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "void")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(84)
parameter_dcls:
	param_dclPlus
	{
		$$ = $1
	}
//(85)
param_dclPlus:
	param_dcl {
		CompleteIdllex.InfoAt("param_dclPlus/param_dcl")
		$$ = $1
	}
	| param_dclPlus ',' param_dcl
	{
		CompleteIdllex.InfoAt("param_dclPlus/param_dclPlus ',' param_dcl")
		$$ = $1.NextParameterDeclarations($3)
	}
param_dcl:
	param_attribute type_spec simple_declarator
	{
		CompleteIdllex.InfoAt("param_dcl/param_attribute type_spec simple_declarator")
		var err error
		$$, err = scopedObjects.NewParameterDeclarations($3, $1.In(), $1.Out(), $3.Identifier(), $2), nil
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(86)
param_attribute:
	RWin
	{
		CompleteIdllex.InfoAt("param_attribute/RWin")
		$$ = scopedObjects.NewParamAttribute($1, true, false)
	}
	| RWout
	{
		CompleteIdllex.InfoAt("param_attribute/RWout")
		$$ = scopedObjects.NewParamAttribute($1, false, true)
	}
	| RWinout
	{
		CompleteIdllex.InfoAt("param_attribute/RWinout")
		$$ = scopedObjects.NewParamAttribute($1, true, true)
	}
//(87)
raises_expr:
	RWraises '(' scoped_namePlus ')'
	{
		CompleteIdllex.InfoAt("raises_expr/RWraises '(' scoped_namePlus ')'", $3)
	}
//(88)
attr_dcl:
	readonly_attr_spec
	{
		CompleteIdllex.InfoAt("attr_dcl/readonly_attr_spec")
		$$ = $1
	}
	| attr_spec
	{
		CompleteIdllex.InfoAt("attr_dcl/attr_spec")
		$$ = $1
	}
//(89)
readonly_attr_spec:
	RWreadonly RWattribute type_spec readonly_attr_declarator
	{
		CompleteIdllex.InfoAt("readonly_attr_spec/RWreadonly RWattribute type_spec readonly_attr_declarator")
		var err error
		$$, err = scopedObjects.NewAttributeDcl($1, $4, $3, true)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(90)
readonly_attr_declarator:
	simple_declarator raises_expr
	{
		CompleteIdllex.InfoAt("readonly_attr_declarator/simple_declarator raises_expr")
		var err error
		$$, err = scopedObjects.NewAttrDeclarator($1)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| simple_declaratorPlus
	{
		CompleteIdllex.InfoAt("readonly_attr_declarator/simple_declaratorPlus")
		var err error
		$$, err = scopedObjects.NewAttrDeclarator($1)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(91)
attr_spec:
	RWattribute type_spec attr_declarator
	{
		CompleteIdllex.InfoAt("attr_spec/RWattribute type_spec attr_declarator")
		var err error
		$$, err = scopedObjects.NewAttributeDcl($1, $3, $2, false)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(92)
attr_declarator:
	simple_declarator attr_raises_expr
	{
		CompleteIdllex.InfoAt("attr_declarator/simple_declarator attr_raises_expr")
		var err error
		$$, err = scopedObjects.NewAttrDeclarator($1)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| simple_declaratorPlus
	{
		CompleteIdllex.InfoAt("attr_declarator/simple_declarator attr_raises_expr")
		var err error
		$$, err = scopedObjects.NewAttrDeclarator($1)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(93)
attr_raises_expr:
	get_excep_expr  set_excep_expr
	{
		CompleteIdllex.InfoAt("")
	}
	| get_excep_expr
	{
		CompleteIdllex.InfoAt("")
	}
	| set_excep_expr
	{
		CompleteIdllex.InfoAt("")
	}
//(94)
get_excep_expr:
	RWgetraises exception_list
	{
		CompleteIdllex.InfoAt("")
	}
//(95)
set_excep_expr:
	RWsetraises exception_list
	{
		CompleteIdllex.InfoAt("")
	}
//(96)
exception_list:
 	'(' scoped_namePlus ')'
	{
		CompleteIdllex.InfoAt("")
	}
//(99)//(125)
value_dcl:
	value_def
	{
		CompleteIdllex.InfoAt("value_dcl/value_def")
		$$ = $1
	}
	| value_forward_dcl
	{
		CompleteIdllex.InfoAt("value_dcl/value_forward_dcl")
		$$ = $1
	}
	| value_box_def
	{
		CompleteIdllex.InfoAt("value_dcl/value_box_def")
		$$ = $1
	}
	| value_abs_def
	{
		CompleteIdllex.InfoAt("value_dcl/value_abs_def")
		$$ = $1
	}
//(100)
value_def:
	value_header '{' value_elementPlus '}'
	{
		CompleteIdllex.InfoAt("value_header '{' value_elementPlus '}'")
	}
//(101)
value_header:
	value_kind identifier value_inheritance_spec
	{
		CompleteIdllex.InfoAt("value_header/value_kind identifier value_inheritance_spec")
		var err error
		$$, err = scopedObjects.NewIdlValueHeader($1, $1, $2, $3)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| value_kind identifier
	{
		CompleteIdllex.InfoAt("value_header/value_kind identifier")
		var err error
		$$, err = scopedObjects.NewIdlValueHeader($1, $1, $2, nil)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(102) //(128)
value_kind:
	RWvaluetype
	{
		CompleteIdllex.InfoAt("value_kind/RWvaluetype")
		var err error
		$$, err = scopedObjects.NewIdlValueKind($1, false)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	|RWcustom RWvaluetype
	{
		CompleteIdllex.InfoAt("value_kind/RWcustom RWvaluetype")
		var err error
		$$, err = scopedObjects.NewIdlValueKind($1, true)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}

//(103)//(130)
value_inheritance_spec:
	{
		CompleteIdllex.InfoAt("value_inheritance_spec")
		var err error
		$$, err = nil, nil
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	|   RWsupports interface_name
	{
		CompleteIdllex.InfoAt("value_inheritance_spec/RWsupports interface_name")
		var err error
		$$, err = scopedObjects.NewValueInheritanceSpec($1)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	|  ':' value_name
	{
		CompleteIdllex.InfoAt("value_inheritance_spec/':' value_name")
		var err error
		$$, err = scopedObjects.NewValueInheritanceSpec($2)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	|  ':' value_name  RWsupports interface_name
	{
		CompleteIdllex.InfoAt("value_inheritance_spec/':' value_name  RWsupports interface_name")
		var err error
		$$, err = scopedObjects.NewValueInheritanceSpec($2)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	|  ':' value_namePlus
	{
		CompleteIdllex.InfoAt("value_inheritance_spec/':' value_namePlus")
		var err error
		$$, err = scopedObjects.NewValueInheritanceSpec($2)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	|  ':' value_namePlus RWsupports interface_namePlus
	{
		CompleteIdllex.InfoAt("value_inheritance_spec/':' value_namePlus RWsupports interface_namePlus")
		var err error
		$$, err = scopedObjects.NewValueInheritanceSpec($2)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	|  ':' RWtruncatable value_namePlus
	{
		CompleteIdllex.InfoAt("value_inheritance_spec/':' RWtruncatable value_namePlus")
		var err error
		$$, err = scopedObjects.NewValueInheritanceSpec($2)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	|  ':' RWtruncatable value_namePlus  RWsupports interface_namePlus
	{
		CompleteIdllex.InfoAt("value_inheritance_spec/':' RWtruncatable value_namePlus  RWsupports interface_namePlus")
		var err error
		$$, err = scopedObjects.NewValueInheritanceSpec($2)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}

//(104)
value_namePlus:
	value_name
	{
		CompleteIdllex.InfoAt("value_namePlus/value_name")
		$$ = $1
	}
	| value_namePlus ',' value_name
	{
		CompleteIdllex.InfoAt("value_namePlus/value_namePlus ',' value_name")
		$$ = $1
	}
value_name:
	scoped_name
	{
		CompleteIdllex.InfoAt("value_name/scoped_name")
		$$ = $1
	}
//(105)
value_elementPlus:
	{
		CompleteIdllex.InfoAt("")
	}
value_elementStar:
	{
		CompleteIdllex.InfoAt("")
	}
value_element:
	export
	{
		CompleteIdllex.InfoAt("")
	}
	| state_member
	{
		CompleteIdllex.InfoAt("")
	}
	| init_dcl
	{
		CompleteIdllex.InfoAt("")
	}
//(106)
state_member:
	RWpublic type_spec declarators ';'
	{
		CompleteIdllex.InfoAt("")
	}
	| RWprivate  type_spec declarators ';'
	{
		CompleteIdllex.InfoAt("")
	}
//(107)
init_dcl:
	RWfactory identifier '('  ')'  ';'
	{
		CompleteIdllex.InfoAt("")
	}
	| RWfactory identifier '('  ')'  raises_expr  ';'
	{
		CompleteIdllex.InfoAt("")
	}
	| RWfactory identifier '('  init_param_dcls  ')'  ';'
	{
		CompleteIdllex.InfoAt("")
	}
	| RWfactory identifier '('  init_param_dcls  ')'  raises_expr  ';'
	{
		CompleteIdllex.InfoAt("")
	}
//(108)
init_param_dcls:
	init_param_dclPlus
	{
		CompleteIdllex.InfoAt("")
	}
//(109)
init_param_dclPlus:
	init_param_dcl
	{
		CompleteIdllex.InfoAt("")
	}
	|init_param_dclPlus ',' init_param_dcl
	{
		CompleteIdllex.InfoAt("")
	}
init_param_dcl:
	RWin type_spec simple_declarator
	{
		CompleteIdllex.InfoAt("")
	}
//(110)
value_forward_dcl:
	value_kind identifier
	{
		CompleteIdllex.InfoAt("value_forward_dcl/value_kind identifier")

		var err error
		$$, err = scopedObjects.NewIdlvalue_forward_dcl($1, $2.Identifier(), $1.GetFileName(), $1.GetRow(), $1.GetCol())
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}


	}
//(113)
type_id_dcl: RWtypeid scoped_name string_literal {}
//(114)
type_prefix_dcl:
	RWtypeprefix scoped_name string_literal
	{
		CompleteIdllex.InfoAt("type_prefix_dcl/RWtypeprefix scoped_name string_literal", $2, $3)
		var err error
		$$, err = CompleteIdllex.CreateTypePrefixDcl($1, $2.Identifier(), $3)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(115)
import_dcl: RWimport imported_scope {}
//(116)
imported_scope: scoped_name {} | string_literal {}
//(118)
object_type:
	RWObject
	{
		CompleteIdllex.InfoAt("")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "Object")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(120)
op_oneway_dcl: RWoneway RWvoid identifier '('  in_parameter_dclsStar  ')' {}
//(121)
in_parameter_dclsStar:
	 {
		CompleteIdllex.InfoAt("")
	 }
	| in_parameter_dcls
	{
		CompleteIdllex.InfoAt("")
	}
	| in_parameter_dclsStar in_parameter_dcls
	{
		CompleteIdllex.InfoAt("")
	}
in_parameter_dcls:
	in_param_dclPlus
	{
		CompleteIdllex.InfoAt("")
	}
//(122)
in_param_dclPlus:
	in_param_dcl
	{
		CompleteIdllex.InfoAt("")
	}
	| in_param_dclPlus ',' in_param_dcl
	{
		CompleteIdllex.InfoAt("")
	}
in_param_dcl:
	RWin type_spec simple_declarator
	{
		CompleteIdllex.InfoAt("")
	}
//(123)
op_with_context:
 	 op_oneway_dcl context_expr
 	 {
		CompleteIdllex.InfoAt("")
 	 }
 	| op_dcl  context_expr
 	{
		CompleteIdllex.InfoAt("")
 	}
//(124)
string_literalPlus:
	string_literal
	{
		CompleteIdllex.InfoAt("")
	}
	| string_literalPlus ',' string_literal
	{
		CompleteIdllex.InfoAt("")
	}
context_expr:
	RWcontext '(' string_literalPlus  ')'
	{
		CompleteIdllex.InfoAt("")
	}
//(126)
value_box_def:
	RWvaluetype identifier type_spec
	{
		CompleteIdllex.InfoAt("")
	}
//(127)
value_abs_def:
	RWabstract RWvaluetype identifier
	{
		CompleteIdllex.InfoAt("value_abs_def/RWabstract   identifier", $3, true)

		var err error
		$$, err = scopedObjects.NewIdlValueAbsoluteDefinition($1, $3.Identifier(), nil, scopedObjects.VADAbstract | scopedObjects.VADForward)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| RWabstract RWvaluetype identifier  value_inheritance_spec  '{' exportStar '}'
	{
		CompleteIdllex.InfoAt("value_abs_def/RWabstract RWvaluetype identifier  value_inheritance_spec  '{' exportStar '}'")

		var err error
		$$, err = scopedObjects.NewIdlValueAbsoluteDefinition($1, $3.Identifier(), $6, scopedObjects.VADAbstract)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
	| RWabstract RWvaluetype identifier '{' exportStar '}'
	{
		CompleteIdllex.InfoAt("value_abs_def/RWabstract RWvaluetype identifier '{' exportStar '}'")

		var err error
		$$, err = scopedObjects.NewIdlValueAbsoluteDefinition($1, $3.Identifier(), $5, scopedObjects.VADAbstract)
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(132)
value_base_type:
	RWValueBase
	{
		CompleteIdllex.InfoAt("value_base_type/RWValueBase")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "valuebase")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(134)
component_dcl:
	component_def
	{
		CompleteIdllex.InfoAt("")
	}
	| component_forward_dcl
	{
		CompleteIdllex.InfoAt("")
	}
//(135)
component_forward_dcl:
	RWcomponent identifier
	{
		CompleteIdllex.InfoAt("")
	}
//(136)
component_def:
	component_header '{' component_body '}'
	{
		CompleteIdllex.InfoAt("")
	}

//(137) //(154)
component_header:
	RWcomponent identifier
	{
		CompleteIdllex.InfoAt("")
	}
	| RWcomponent identifier  component_inheritance_spec
	{
		CompleteIdllex.InfoAt("")
	}
	| RWcomponent identifier  supported_interface_spec
	{
		CompleteIdllex.InfoAt("")
	}
	| RWcomponent identifier  component_inheritance_spec  supported_interface_spec
	{
		CompleteIdllex.InfoAt("")
	}

//(138)
component_inheritance_spec:
	':' scoped_name
	{
		CompleteIdllex.InfoAt("")
	}
//(139)
component_body:
	component_exportStar
	{
		CompleteIdllex.InfoAt("")
	}
//(140)//(156)//(179)
component_exportStar:
	{
		CompleteIdllex.InfoAt("")
	}
	| component_export
	{
		CompleteIdllex.InfoAt("")
	}
	| component_exportStar component_export
	{
		CompleteIdllex.InfoAt("")
	}
component_export:
	provides_dcl ';'
	{
		CompleteIdllex.InfoAt("")
	}
	| uses_dcl ';'
	{
		CompleteIdllex.InfoAt("")
	}
	| attr_dcl ';'
	{
		CompleteIdllex.InfoAt("")
	}
	| emits_dcl ';'
	{
		CompleteIdllex.InfoAt("")
	}
	| publishes_dcl ';'
	{
		CompleteIdllex.InfoAt("")
	}
	| consumes_dcl ';'
	{
		CompleteIdllex.InfoAt("")
	}
	| port_dcl ';'
	{
		CompleteIdllex.InfoAt("")
	}

//(141)
provides_dcl:
	RWprovides interface_type identifier
	{
		CompleteIdllex.InfoAt("")
	}
//(142)//(157)
interface_type:
	scoped_name
	{
		CompleteIdllex.InfoAt("")
	}
	| RWObject
	{
		CompleteIdllex.InfoAt("")
	}

//(143)//(158)
uses_dcl:
	RWuses interface_type identifier
	{
		CompleteIdllex.InfoAt("")
	}
	|RWuses RWmultiple interface_type identifier
	{
		CompleteIdllex.InfoAt("")
	}

//(145)
home_dcl:
	home_header '{' home_body '}'
	{
		CompleteIdllex.InfoAt("")
	}
//(146)//(162)
home_header:
	RWhome identifier  RWmanages scoped_name
	{
		CompleteIdllex.InfoAt("")
	}
	| RWhome identifier  home_inheritance_spec  RWmanages scoped_name
	{
		CompleteIdllex.InfoAt("")
	}
	| RWhome identifier  RWmanages scoped_name
	{
		CompleteIdllex.InfoAt("")
	}
	| RWhome identifier RWmanages scoped_name  primary_key_spec
	{
		CompleteIdllex.InfoAt("")
	}
	| RWhome identifier   supported_interface_spec  RWmanages scoped_name
	{
		CompleteIdllex.InfoAt("")
	}
	| RWhome identifier   supported_interface_spec  RWmanages scoped_name  primary_key_spec
	{
		CompleteIdllex.InfoAt("")
	}
	| RWhome identifier  home_inheritance_spec   RWmanages scoped_name
	{
		CompleteIdllex.InfoAt("")
	}
	| RWhome identifier  home_inheritance_spec   RWmanages scoped_name  primary_key_spec
	{
		CompleteIdllex.InfoAt("")
	}
	| RWhome identifier  home_inheritance_spec   supported_interface_spec  RWmanages scoped_name
	{
		CompleteIdllex.InfoAt("")
	}
	| RWhome identifier  home_inheritance_spec   supported_interface_spec  RWmanages scoped_name  primary_key_spec
	{
		CompleteIdllex.InfoAt("")
	}


//(147)
home_inheritance_spec:
	':' scoped_name
	{
		CompleteIdllex.InfoAt("")
	}
//(148)
home_body:
	home_exportStar
	{
		CompleteIdllex.InfoAt("")
	}
//(149)//(164)
home_exportStar:
	{
		CompleteIdllex.InfoAt("")
	}
	| home_export
	{
		CompleteIdllex.InfoAt("")
	}
	| home_exportStar home_export
	{
		CompleteIdllex.InfoAt("")
	}
home_export:
	export
	{
		CompleteIdllex.InfoAt("")
	}
	| factory_dcl ';'
	{
		CompleteIdllex.InfoAt("")
	}
	| finder_dcl ';'
	{
		CompleteIdllex.InfoAt("")
	}

//(150)
factory_dcl:
	RWfactory identifier '('  ')'
	{
		CompleteIdllex.InfoAt("")
	}
	|RWfactory identifier '('  ')'  raises_expr
	{
		CompleteIdllex.InfoAt("")
	}
	|RWfactory identifier '('  factory_param_dcls  ')'
	{
		CompleteIdllex.InfoAt("")
	}
	|RWfactory identifier '('  factory_param_dcls  ')'  raises_expr
	{
		CompleteIdllex.InfoAt("")
	}

//(151)
factory_param_dcls:
	factory_param_dclPlus
	{
		CompleteIdllex.InfoAt("")
	}
//(152)
factory_param_dclPlus:
	factory_param_dcl
	{
		CompleteIdllex.InfoAt("")
	}
	| factory_param_dclPlus ',' factory_param_dcl
	{
		CompleteIdllex.InfoAt("")
	}
factory_param_dcl:
	RWin type_spec simple_declarator
	{
		CompleteIdllex.InfoAt("")
	}
//(155)
supported_interface_spec:
	RWsupports scoped_namePlus
	{
		CompleteIdllex.InfoAt("")
	}
//(159)
emits_dcl:
	RWemits scoped_name identifier
	{
		CompleteIdllex.InfoAt("")
	}
//(160)
publishes_dcl:
	RWpublishes scoped_name identifier
	{
		CompleteIdllex.InfoAt("")
	}
//(161)
consumes_dcl:
	RWconsumes scoped_name identifier
	{
		CompleteIdllex.InfoAt("")
	}
//(163)
primary_key_spec:
	RWprimarykey scoped_name
	{
		CompleteIdllex.InfoAt("")
	}
//(165)
finder_dcl:
	RWfinder identifier '('  ')'
	{
		CompleteIdllex.InfoAt("")
	}
	|RWfinder identifier '('  ')'  raises_expr
	{
		CompleteIdllex.InfoAt("")
	}
	|RWfinder identifier '('  init_param_dcls  ')'
	{
		CompleteIdllex.InfoAt("")
	}
	|RWfinder identifier '('  init_param_dcls  ')'  raises_expr
	{
		CompleteIdllex.InfoAt("")
	}
//(166)
event_dcl:
	{
		CompleteIdllex.InfoAt("")
	}
	| event_def
	{
		CompleteIdllex.InfoAt("")
	}
	| event_abs_def
	{
		CompleteIdllex.InfoAt("")
	}
	| event_forward_dcl
	{
		CompleteIdllex.InfoAt("")
	}
//(167)
event_forward_dcl:
	RWeventtype identifier
	{
		CompleteIdllex.InfoAt("")
	}
	|  RWabstract  RWeventtype identifier
	{
		CompleteIdllex.InfoAt("")
	}
//(168)
event_abs_def:
	RWabstract RWeventtype identifier  '{' exportStar '}'
	{
		CompleteIdllex.InfoAt("")
	}
	| RWabstract RWeventtype identifier  value_inheritance_spec  '{' exportStar '}'
	{
		CompleteIdllex.InfoAt("")
	}
//(169)
event_def:
	event_header '{' value_elementStar '}'
	{
		CompleteIdllex.InfoAt("")
	}
//(170)
event_header:
	RWeventtype identifier
	{
		CompleteIdllex.InfoAt("")
	}
	|  RWeventtype identifier  value_inheritance_spec
	{
		CompleteIdllex.InfoAt("")
	}
	|  RWcustom  RWeventtype identifier
	{
		CompleteIdllex.InfoAt("")
	}
	|  RWcustom  RWeventtype identifier  value_inheritance_spec
	{
		CompleteIdllex.InfoAt("")
	}
//(172)
porttype_dcl:
	porttype_def
	{
		CompleteIdllex.InfoAt("")
	}
	| porttype_forward_dcl
	{
		CompleteIdllex.InfoAt("")
	}
//(173)
porttype_forward_dcl:
	RWporttype identifier
	{
		CompleteIdllex.InfoAt("")
	}
//(174)
porttype_def:
	RWporttype identifier '{' port_body '}'
	{
		CompleteIdllex.InfoAt("")
	}
//(175)
port_body: port_ref port_exportStar{}
//(176)
port_ref:
	provides_dcl ';'{}
	| uses_dcl ';' {}
	| port_dcl ';'{}
//(177)
port_exportStar:
	{
		CompleteIdllex.InfoAt("")
	}
port_export:
	port_ref
	{
		CompleteIdllex.InfoAt("")
	}
	| attr_dcl ';'
	{
		CompleteIdllex.InfoAt("")
	}
//(178)
port_dcl:
 	RWmirrorport scoped_name identifier
 	{
		CompleteIdllex.InfoAt("")
 	}
 	| RWport  scoped_name identifier
 	{
		CompleteIdllex.InfoAt("")
 	}
//(180)
connector_dcl: connector_header '{' connector_exportPlus '}'{}
//(181)
connector_header:
	RWconnector identifier{}
	|RWconnector identifier  connector_inherit_spec{}
//(182)
connector_inherit_spec: ':' scoped_name{}
//(183)
connector_exportPlus:{}
connector_export: port_ref {}| attr_dcl ';'{}
//(185)
template_module_dcl: RWmodule identifier '<' formal_parameters '>' '{' tpl_definitionPlus'}'{}
//(186)
formal_parameters: formal_parameterPlus{}
//(187)
formal_parameterPlus:
	formal_parameter {}
	| formal_parameterPlus ',' formal_parameter{}
formal_parameter: formal_parameter_type identifier
//(188)
formal_parameter_type: RWtypename | RWinterface | RWvaluetype | RWeventtype | RWstruct | RWunion | RWexception | RWenum | RWsequence | RWconst const_type | sequence_type
//(189)
tpl_definitionPlus:
tpl_definition: definition | template_module_ref ';'
//(190)
template_module_inst:
	RWmodule scoped_name '<' actual_parameters '>' identifier{}
//(191)
actual_parameters: actual_parameterPlus{}
//(192)
actual_parameterPlus: actual_parameter{}| actual_parameterPlus ',' actual_parameter{}
actual_parameter: type_spec {}| const_expr{}
//(193)
template_module_ref: RWalias scoped_name '<' formal_parameter_names '>' identifier{}
//(194)
identifierPlus: identifier{}| identifierPlus ',' identifier{}
formal_parameter_names: identifierPlus{}
//(199)
map_type:
 	RWmap '<' type_spec ',' type_spec ',' positive_int_const '>'
 	{

 	}
 	| RWmap '<' type_spec ',' type_spec '>'{

 	}
//(200)
bitset_dcl:
	RWbitset identifier  '{' bitfieldStar '}'
	{
		CompleteIdllex.Error("implement bitset_dcl1")
		return 1
	}
	RWbitset identifier ':' scoped_name '{' bitfieldStar '}'
	{
		CompleteIdllex.Error("implement bitset_dcl2")
		return 1
	}
//(201)
identifierStar:
bitfieldStar:
bitfield: bitfield_spec identifierStar ';'
//(202)
bitfield_spec: RWbitfield '<' positive_int_const '>' | RWbitfield '<' positive_int_const ',' destination_type '>'
//(203)
destination_type: boolean_type | octet_type | integer_type
//(204)
bitmask_dcl:
	RWbitmask identifier '{' bit_valueStar '}'
	{
		CompleteIdllex.Error("implement bitmask_dcl")
		return 1
	}
//(205)
bit_valueStar:
bit_value: identifier
//(208)
signed_tiny_int: RWint8
	{
		CompleteIdllex.InfoAt("")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "int8")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(209)
unsigned_tiny_int: RWuint8
	{
		CompleteIdllex.InfoAt("")
		var err error
		$$, err = CompleteIdllex.FindPrimitive($1, "uint8")
		if err != nil {
			CompleteIdllex.Error(err.Error())
			return 1
		}
	}
//(216)

//(219)
annotation_dcl: annotation_header '{' annotation_body '}'{}
//(220)
annotation_header: Annotation identifier{}
//(221)
annotation_bodyEntry:
	annotation_member{}
	| enum_dcl ';'{}
	| const_dcl ';'{}
	| typedef_dcl ';'{}

annotation_bodyEntryStar:{}
annotation_body: annotation_bodyEntryStar{}
//(222)
annotation_member:
	annotation_member_type simple_declarator  ';'{}
	| annotation_member_type simple_declarator  RWdefault const_expr  ';'{}
//(223)
annotation_member_type: const_type {}| any_const_type {}| scoped_name{}
//(224)
any_const_type:
	RWany
	{
	}
//(225)
annotation_appl:
	'@' scoped_name{}
	| '@' scoped_name  '(' annotation_appl_params ')'{}
//(226)
annotation_appl_params:
	const_expr{}
	| annotation_appl_paramPlus{}
//(227)
annotation_appl_paramPlus:
	annotation_appl_param{}
	| annotation_appl_paramPlus ',' annotation_appl_param{}
annotation_appl_param: identifier '=' const_expr{}
%%