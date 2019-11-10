%{
package yacc
	//go:generate goyacc -o idl.go -p "IdlExpr" idl.y
	import (
		"github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"
		"github.com/bhbosman/CodeGenerators/mitchgenerator/MitchDefinedTypes"
		)
%}

%token Identifier
%token Integer_literal
%token Hex_literal
%token Character_literal
%token Scope
%token Rwenum
%token Rwstruct
%token Rwtypedef
%token RwMitchAlpha
%token RwMitchBitField
%token RwMitchByte
%token RwMitchDate
%token RwMitchTime
%token RwMitchPrice04
%token RwMitchPrice08
%token RwMitchUInt08
%token RwMitchUInt16
%token RwMitchUInt32
%token RwMitchUInt64
//%token RwMitchMessageNumberType
//%token RwMitchMessageLengthType


%union{
	Identifier 	string
	IntegerValue    int64
	StringValue string
	FloatValue float64
	ConstValue interfaces.IConstantValue
	BoolValue bool

	Member 		*Member
	Declarator 	interfaces.IDeclarator
	DefinedType 	interfaces.IDefinedType
	DefinitionDeclaration interfaces.IDefinitionDeclaration
	Specification []interfaces.IDefinitionDeclaration
}


%type	<StringValue>	Character_literal
%type	<Identifier>	Identifier simple_declarator
%type	<ConstValue>	literal primary_expr
%type	<IntegerValue>	positive_int_const
%type	<IntegerValue>	const_expr
%type	<IntegerValue>	Integer_literal
%type	<IntegerValue>	unary_expr
%type	<IntegerValue>	unary_operator
%type	<Member>	member
%type	<DefinedType>	type_spec
%type	<DefinedType>	simple_type_spec
%type	<DefinedType>	mitch_type_spec
%type	<DefinedType>	scoped_name
%type	<DefinedType>	mitch_bit_field
%type	<DefinedType>	template_type_spec
%type	<Specification>	specification
%type	<DefinitionDeclaration> definition
%type	<DefinitionDeclaration> definitions
%type	<DefinitionDeclaration> type_dcl
%type	<DefinitionDeclaration> struct_dcl
%type	<DefinitionDeclaration>	struct_def constr_type_dcl enum_dcl
%type	<DefinitionDeclaration> type_declarator typedef_dcl
%type	<Declarator>	declarator any_declarators enumerator



%%
specification :
	definitions{
		$$ = AddDefinitions($1)
		context, _ := GetIdlExprContext(IdlExprlex)
		context.Specification = $$
	}

definitions:
	definition{
		$$ = $1
	}
	|definitions definition{
		$$ = $1
		GetLast($$).SetNext($2)
	}

definition :
	type_dcl ';'{
		err := AddTypeDclToContext(IdlExprlex, $1)
		if err != nil {
			SendError(IdlExprlex, "AddTypedefDcl error")
			return ErrorOnAddTypedefDcl
		}

		$$ = $1
	}

scoped_name :
	Identifier{
		lex, err := GetIdlExprContext(IdlExprlex)
		if err == nil {
			definitionDeclaration := lex.FindScopeName($1)
			if definitionDeclaration == nil{
				IdlExprlex.Error(__yyfmt__.Sprintf("Value %v is not declared", $1))
				return DefNotFound
			}else{
				$$ = definitionDeclaration
			}
		}else{
			IdlExprlex.Error(__yyfmt__.Sprintf("GetIdlExprLex failure. %v", $1))
			return NoLex
		}
	}


const_expr :
	unary_expr{
		$$ = $1
	}


unary_expr :
	unary_operator primary_expr {
		value, ok := $2.Value().(int64)
		if ok{
			$$ = value
		}else{
			SendError(IdlExprlex, "Value must be an integer (int64)")
			return ErrorMustbeAnInt
		}
	}
	| primary_expr{
		value, ok := $1.Value().(int64)
		if ok{
			$$ = value
		}else{
			SendError(IdlExprlex, "Value must be an integer (int64)")
			return ErrorMustbeAnInt
		}
	}

unary_operator :
	'-' {
		$$ = -1
	}
	| '+'{
		$$ = 0
	}
	| '~'
	{
		$$ = 99
	}

primary_expr :
	scoped_name{
		lex, err := GetIdlExprContext(IdlExprlex)
		if err == nil {
			data := lex.FindScopeName($1.GetName())
			if data == nil{
				IdlExprlex.Error(__yyfmt__.Sprintf("Could not find defined value %v", $1.GetName()))
				return 10003
			}
		} else {
			IdlExprlex.Error("Could not find lex")
			return NoLex
		}
	}
	| literal{
		$$ = $1
	}
	| '(' const_expr ')'{
		$$ = newConstantValueWithNoLength($2, interfaces.Int64)
	}

literal :
	Integer_literal{
		$$ = newConstantValueWithNoLength($1, interfaces.Int64)
	}
	| Character_literal{
		$$ = newConstantValue([]byte($1)[0], interfaces.Char,1)
        }


positive_int_const : const_expr

type_dcl :
	constr_type_dcl{
		$$ = $1
	}
	| typedef_dcl{
		if $1 == nil{
			SendError(IdlExprlex, "AddTypedefDcl error")
			return ErrorOnAddTypedefDcl

		}
		if typeDecl, ok := $1.(interfaces.ITypeDeclaration); ok{
			err := AddTypedefDcl(IdlExprlex, typeDecl)
			if err != nil {
				SendError(IdlExprlex, "AddTypedefDcl error")
				return ErrorOnAddTypedefDcl
			}
		}else{
			SendError(IdlExprlex, "AddTypedefDcl error")
			return ErrorOnAddTypedefDcl
		}

	}

type_spec : simple_type_spec{
	$$ = $1
}

mitch_type_spec:
	RwMitchAlpha '<' positive_int_const '>' {$$ =&MitchDefinedTypes.MitchAlpha{Length:$3}}
	|RwMitchByte{$$ =&MitchDefinedTypes.MitchByte{}}
	|RwMitchDate{$$ =&MitchDefinedTypes.MitchDate{}}
	|RwMitchTime{$$ =&MitchDefinedTypes.MitchTime{}}
	|RwMitchPrice04{$$ =&MitchDefinedTypes.MitchPrice04{}}
	|RwMitchPrice08{$$ =&MitchDefinedTypes.MitchPrice08{}}
	|RwMitchUInt08{$$ =&MitchDefinedTypes.MitchUInt08{}}
	|RwMitchUInt16{$$ =&MitchDefinedTypes.MitchUInt16{}}
	|RwMitchUInt32{$$ =&MitchDefinedTypes.MitchUInt32{}}
	|RwMitchUInt64 {$$ =&MitchDefinedTypes.MitchUInt64{}}

simple_type_spec :
	mitch_type_spec{
		$$ = $1
	}
	| scoped_name{
		$$ = $1
	}


template_type_spec :
	mitch_bit_field{
		$$ = $1
	}

mitch_bit_field:
	RwMitchBitField '<' simple_declarator ',' simple_declarator ',' simple_declarator ',' simple_declarator ',' simple_declarator ',' simple_declarator ',' simple_declarator ',' simple_declarator '>'{
		$$ = MitchDefinedTypes.NewMitchBitField($3, $5, $7, $9, $11, $13, $15, $17)
	}

constr_type_dcl :
	struct_dcl{
		$$ = $1
	}
	| enum_dcl{
		$$ = $1
	}

struct_dcl :
	struct_def{
		$$ = $1

    	}


struct_def :
	Rwstruct Identifier '{' member '}'{
		def := NewMitchMessageDefinition($2, 0, 0)
		member := $4
		for member != nil {
			decl := member.Declarator
			for decl != nil  {
				def.AddMember(member.DefinedType, decl)
			decl = decl.Next()
			}
			member = member.Next
		}
		$$ = def
	}
	| Rwstruct Identifier '{'  '}'{
		def := NewMitchMessageDefinition($2, 0, 0)
		$$ = def
	}
	|
	Rwstruct '<' positive_int_const ',' positive_int_const '>' Identifier '{' member '}'{
		def := NewMitchMessageDefinition($7, $3, $5)
		member := $9
		for member != nil {
			decl := member.Declarator
			for decl != nil  {
				def.AddMember(member.DefinedType, decl)
			decl = decl.Next()
			}
			member = member.Next
		}
		$$ = def
	}
	| Rwstruct  '<' positive_int_const ',' positive_int_const  '>' Identifier '{'  '}'{
		def := NewMitchMessageDefinition($7, $3, $5)
		$$ = def
	}

member :
	type_spec declarator ';'{
 		$$ = NewMember($1, $2, nil)
	}
	|member member{
		$$ = NewMember($1.DefinedType, $1.Declarator, $2)
	}



enum_dcl :
	Rwenum Identifier '{' enumerator  '}'{
		def := NewEnumDcl($2)
		decl := $4
		for decl != nil  {
			def.AddMember(decl)
			decl = decl.Next()
		}
		$$ = def
	}
enumerator :
	declarator{}
	| enumerator ',' enumerator{
		$1.SetNext($3)
	}


array_declarator :
	Identifier fixed_array_sizes{

	}

fixed_array_sizes:
	fixed_array_size{

	}
	| fixed_array_sizes fixed_array_size{

	}

fixed_array_size :
	'[' positive_int_const ']'{
	}


simple_declarator :
	Identifier{
		$$ =$1
	}

typedef_dcl :
	Rwtypedef type_declarator{
		$$ = $2
	}
type_declarator :
	simple_type_spec any_declarators{
		$$ = Newtypedef_dcl($1, $2)
	}
	| template_type_spec any_declarators{
		$$ = Newtypedef_dcl($1, $2)
	}
	| constr_type_dcl  any_declarators{
		$$ = Newtypedef_dcl($1, $2)
	}
any_declarators:
	simple_declarator{
		$$ = NewDeclarator($1, nil)
	}
	|array_declarator{

	}
	|any_declarators ',' any_declarators{
		$1.SetNext($3)
		$$ = $1
	}
declarator :
	simple_declarator '=' literal {
		$$ = NewDeclarator($1, $3)
	}
	|simple_declarator{
		$$ = NewDeclarator($1, nil)
	}
	|declarator ',' declarator{
		$1.SetNext($3)
		$$ = $1
	}
%%


