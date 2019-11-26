package yacc

//go:generate golex completeIdl.l

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	"github.com/bhbosman/CodeGenerators/idlgenerator/scopedObjects"
	"io"
	"log"
	"strconv"
)

type CompleteIdlLexImpl struct {
	tokenizer      *Tokenizer
	logger         *log.Logger
	lastError      string
	typeSpec       ScopingInterfaces.ITypeSpec
	nextNumber     ScopingInterfaces.INextNumber
	scopingContext ScopingInterfaces.IScopingContext
}

func (self *CompleteIdlLexImpl) FindPrimitive(fileInformation ScopingInterfaces.IFileInformation, s string) (ScopingInterfaces.IBaseDeclaredType, error) {
	return self.scopingContext.FindTypeSpec(fileInformation, s)
}

func NewCompleteIdlLexImpl(
	fileName string,
	inputStream io.Reader,
	logger *log.Logger,
	definitionContext IDefinitionContext,
	nextNumber ScopingInterfaces.INextNumber,
	scopingContext ScopingInterfaces.IScopingContext) (*CompleteIdlLexImpl, error) {

	return &CompleteIdlLexImpl{
		tokenizer:      NewTokenizer(fileName, inputStream, NewAdditionalInformation(definitionContext)),
		logger:         logger,
		lastError:      "",
		typeSpec:       nil,
		nextNumber:     nextNumber,
		scopingContext: scopingContext,
	}, nil
}

func (self *CompleteIdlLexImpl) NewTypeDeclarator(simpleTypeSpec ScopingInterfaces.IBaseDeclaredType, declarator ScopingInterfaces.IDeclarator) (ScopingInterfaces.ITypeDeclarator, error) {

	return scopedObjects.NewTypeDeclarator(simpleTypeSpec, declarator)
}

func (self *CompleteIdlLexImpl) NewIdlConstDcl(fileInformation ScopingInterfaces.IFileInformation, identifier string, value int) (ScopingInterfaces.IIdlConstDcl, error) {
	return scopedObjects.NewIdlConstDcl(fileInformation, identifier, value), nil
}

func (self *CompleteIdlLexImpl) NewMember(typeSpec ScopingInterfaces.IBaseDeclaredType, declarator ScopingInterfaces.IDeclarator) (ScopingInterfaces.IStructMember, error) {
	return scopedObjects.NewStructMember(typeSpec, declarator)
}

func (self *CompleteIdlLexImpl) CreateInterfaceKind(fileInformation ScopingInterfaces.IFileInformation, local bool, abstract bool) (ScopingInterfaces.IInterfaceKind, error) {
	return scopedObjects.NewInterfaceKindImpl(fileInformation, local, abstract), nil
}

func (self *CompleteIdlLexImpl) NewDeclarator(fileInformation ScopingInterfaces.IFileInformation, identifier string) (ScopingInterfaces.IDeclarator, error) {
	return scopedObjects.NewDeclarator(fileInformation, identifier), nil
}

func (self *CompleteIdlLexImpl) GetSpec() (ScopingInterfaces.ITypeSpec, error) {
	return self.typeSpec, nil
}

func (self *CompleteIdlLexImpl) GetFileName() string {
	return self.tokenizer.currentStream.fileName
}

func (self *CompleteIdlLexImpl) GetRow() int {
	return self.tokenizer.currentStream.row
}

func (self *CompleteIdlLexImpl) GetCol() int {
	return self.tokenizer.currentStream.column
}

func (self *CompleteIdlLexImpl) AssignSpec(typeSpec ScopingInterfaces.ITypeSpec) (ScopingInterfaces.ITypeSpec, error) {
	self.typeSpec = typeSpec
	return typeSpec, nil
}

func (self *CompleteIdlLexImpl) CreateModuleDcl(fileInformation ScopingInterfaces.IFileInformation, identifier string, typeSpec ScopingInterfaces.ITypeSpec) (ScopingInterfaces.IIdlModuleDcl, error) {
	return scopedObjects.NewModuleDcl(fileInformation, identifier, typeSpec), nil
}

func (self *CompleteIdlLexImpl) CreateInterfaceDcl(fileInformation ScopingInterfaces.IFileInformation, identifier string, forward, abstract, local bool, body ScopingInterfaces.ITypeSpec) (ScopingInterfaces.ITypeSpec, error) {
	return scopedObjects.NewInterfaceDcl(fileInformation, identifier, forward, abstract, local, body)
}

func (self *CompleteIdlLexImpl) CreateTypePrefixDcl(fileInformation ScopingInterfaces.IFileInformation, scopedName, stringLiteral string) (ScopingInterfaces.ITypeSpec, error) {
	return scopedObjects.NewCreateTypePrefixDcl(fileInformation, scopedName, stringLiteral)
}

func (self *CompleteIdlLexImpl) TransformString(value string) (ScopingInterfaces.IPrimaryExpression, error) {
	return scopedObjects.NewPrimaryExpression(value, ScopingInterfaces.PetString), nil
}

func (self *CompleteIdlLexImpl) TransformInteger(value int) (ScopingInterfaces.IPrimaryExpression, error) {
	return scopedObjects.NewPrimaryExpression(value, ScopingInterfaces.PetInteger), nil
}

func (self *CompleteIdlLexImpl) TransformValue(value interface{}, valuetype ScopingInterfaces.IPrimaryExpressionType) (ScopingInterfaces.IPrimaryExpression, error) {
	return scopedObjects.NewPrimaryExpression(value, valuetype), nil
}

func (self *CompleteIdlLexImpl) NewStructType(fileInformation ScopingInterfaces.IFileInformation, identifier string, members ScopingInterfaces.IStructMember, forward bool) (ScopingInterfaces.IStructType, error) {
	if identifier == "" {
		return nil, fmt.Errorf("invalid identifier")
	}
	structType := scopedObjects.NewStructType(fileInformation, identifier, members, forward)
	return structType, nil
}

func (self *CompleteIdlLexImpl) AddExpr(a int) (int, error) {
	return a, nil
}

func (self *CompleteIdlLexImpl) MultiExpr(a int) (int, error) {
	return a, nil
}

func (self *CompleteIdlLexImpl) AddOperator(a, b int) (int, error) {
	return a + b, nil
}

func (self *CompleteIdlLexImpl) MinusOperator(a, b int) (int, error) {
	return a - b, nil
}

func (self *CompleteIdlLexImpl) DivideOperator(a, b int) (int, error) {
	return a / b, nil
}

func (self *CompleteIdlLexImpl) MultiplyOperator(a, b int) (int, error) {
	return a * b, nil
}

func (self *CompleteIdlLexImpl) LastError() string {
	return self.lastError
}

func (self *CompleteIdlLexImpl) InfoAt(info string, params ...interface{}) {
	//if info == "" {
	//	return
	//}
	//s := ""
	//for _, p := range params {
	//
	//	if s == "" {
	//		s = fmt.Sprintf("(%v)", p)
	//	} else {
	//		s = fmt.Sprintf("%v, (%v)", s, p)
	//	}
	//
	//}
	//if self.tokenizer.currentStream != nil {
	//	self.logger.Printf(" %s [%v].", info, s)
	//} else {
	//	self.logger.Printf(" >>>%s [%v].<<<", info, s)
	//}
}

func (self *CompleteIdlLexImpl) Lex(value *CompleteIdlSymType) int {
	CompleteIdlErrorVerbose = true

	for {
		token := self.tokenizer.yylex()
		lexValue := self.tokenizer.yytext
		if token == ErrorFileNotFound {
			self.Error("File not found")
			return token
		}

		switch token {
		case RWvoid, RWin, RWinout, RWout, RWreadonly, RWattribute, RWsequence, RWsupports, RWinterface, RWconst, RWexception,
			RWvaluetype, RWcustom, RWstring, RWwstring, RWtypeprefix, RWlocal, RWabstract, RWstruct, RWenum, RWmodule,
			RWfloat, RWdouble, RWlong, RWshort, RWint16, RWint32, RWint64, RWunsigned, RWuint16, RWuint32, RWuint64,
			RWchar, RWwchar, RWboolean, RWoctet, RWany, RWObject, RWint8, RWuint8, RWfixed, RWValueBase:
			value.ReservedWord = scopedObjects.NewReservedWordData(lexValue, self.tokenizer.currentStream.fileName, self.tokenizer.currentStream.row, self.tokenizer.currentStream.column)
			return token
		case string_literal:
			value.StringValue = lexValue
			return token
		case identifier:
			value.Identifier = scopedObjects.NewIdlIdentifier(
				lexValue,
				self.tokenizer.currentStream.fileName,
				self.tokenizer.currentStream.row,
				self.tokenizer.currentStream.column)
			return token
		case integer_literal:
			integerValue, _ := strconv.ParseInt(lexValue, 10, 64)
			value.IntegerValue = int(integerValue)
			return token
		case Hex_literal:
			integerValue, _ := strconv.ParseInt(lexValue, 0, 64)
			value.IntegerValue = int(integerValue)
			return integer_literal
		default:
			return token
		}
	}
}

func (self *CompleteIdlLexImpl) Error(s string) {
	if self.tokenizer.currentStream != nil {
		self.lastError = fmt.Sprintf("parse error: %s at %v(%d,%d).", s, self.tokenizer.currentStream.fileName, self.tokenizer.currentStream.row, self.tokenizer.currentStream.column)
	} else {
		self.lastError = fmt.Sprintf("parse error: %s.", s)
	}

	self.logger.Printf(self.lastError)
}

type yyStack struct {
	yyStart    yystartcondition
	stackState StackState
	next       *yyStack
}

type AdditionalInformation struct {
	reservedWords     map[string]int
	stackState        StackState
	DefinitionContext IDefinitionContext
	YY_STACK          *yyStack
}

type StackState struct {
	expressionDirection bool
	expressionValue     bool
	ifdefBlock          bool
	defineFlag          bool
}

func (self *AdditionalInformation) push(abc *Tokenizer) {
	item := &yyStack{
		yyStart:    abc.YY_START,
		next:       self.YY_STACK,
		stackState: self.stackState,
	}
	self.YY_STACK = item
}
func (self *AdditionalInformation) pop(abc *Tokenizer) error {
	if self.YY_STACK == nil {
		return fmt.Errorf("stack is empty")
	}
	abc.YY_START = self.YY_STACK.yyStart
	self.stackState = self.YY_STACK.stackState
	self.YY_STACK = self.YY_STACK.next
	return nil
}

func NewAdditionalInformation(definitionContext IDefinitionContext) *AdditionalInformation {
	reservedWords := make(map[string]int)
	reservedWords["abstract"] = RWabstract
	reservedWords["any"] = RWany
	reservedWords["alias"] = RWalias
	reservedWords["attribute"] = RWattribute
	reservedWords["bitfield"] = RWbitfield
	reservedWords["bitmask"] = RWbitmask
	reservedWords["bitset"] = RWbitset
	reservedWords["boolean"] = RWboolean
	reservedWords["case"] = RWcase
	reservedWords["char"] = RWchar
	reservedWords["component"] = RWcomponent
	reservedWords["connector"] = RWconnector
	reservedWords["const"] = RWconst
	reservedWords["consumes"] = RWconsumes
	reservedWords["userContext"] = RWcontext
	reservedWords["custom"] = RWcustom
	reservedWords["default"] = RWdefault
	reservedWords["double"] = RWdouble
	reservedWords["exception"] = RWexception
	reservedWords["emits"] = RWemits
	reservedWords["enum"] = RWenum
	reservedWords["eventtype"] = RWeventtype
	reservedWords["factory"] = RWfactory
	reservedWords["FALSE"] = RWFALSE
	reservedWords["finder"] = RWfinder
	reservedWords["fixed"] = RWfixed
	reservedWords["float"] = RWfloat
	reservedWords["getraises"] = RWgetraises
	reservedWords["home"] = RWhome
	reservedWords["import"] = RWimport
	reservedWords["in"] = RWin
	reservedWords["inout"] = RWinout
	reservedWords["interface"] = RWinterface
	reservedWords["local"] = RWlocal
	reservedWords["long"] = RWlong
	reservedWords["manages"] = RWmanages
	reservedWords["map"] = RWmap
	reservedWords["mirrorport"] = RWmirrorport
	reservedWords["module"] = RWmodule
	reservedWords["multiple"] = RWmultiple
	reservedWords["native"] = RWnative
	reservedWords["Object"] = RWObject
	reservedWords["octet"] = RWoctet
	reservedWords["oneway"] = RWoneway
	reservedWords["out"] = RWout
	reservedWords["primarykey"] = RWprimarykey
	reservedWords["private"] = RWprivate
	reservedWords["port"] = RWport
	reservedWords["porttype"] = RWporttype
	reservedWords["provides"] = RWprovides
	reservedWords["public"] = RWpublic
	reservedWords["publishes"] = RWpublishes
	reservedWords["raises"] = RWraises
	reservedWords["readonly"] = RWreadonly
	reservedWords["setraises"] = RWsetraises
	reservedWords["sequence"] = RWsequence
	reservedWords["short"] = RWshort
	reservedWords["string"] = RWstring
	reservedWords["struct"] = RWstruct
	reservedWords["supports"] = RWsupports
	reservedWords["switch"] = RWswitch
	reservedWords["TRUE"] = RWTRUE
	reservedWords["truncatable"] = RWtruncatable
	reservedWords["typedef"] = RWtypedef
	reservedWords["typeid"] = RWtypeid
	reservedWords["typename"] = RWtypename
	reservedWords["typeprefix"] = RWtypeprefix
	reservedWords["unsigned"] = RWunsigned
	reservedWords["union"] = RWunion
	reservedWords["uses"] = RWuses
	reservedWords["ValueBase"] = RWValueBase
	reservedWords["valuetype"] = RWvaluetype
	reservedWords["void"] = RWvoid
	reservedWords["wchar"] = RWwchar
	reservedWords["wstring"] = RWwstring
	reservedWords["int8"] = RWint8
	reservedWords["uint8"] = RWuint8
	reservedWords["int16"] = RWint16
	reservedWords["int32"] = RWint32
	reservedWords["int64"] = RWint64
	reservedWords["uint16"] = RWuint16
	reservedWords["uint32"] = RWuint32
	reservedWords["uint64"] = RWuint64
	return &AdditionalInformation{
		reservedWords: reservedWords,
		stackState: StackState{
			expressionDirection: false,
			expressionValue:     true,
			ifdefBlock:          false,
			defineFlag:          false,
		},
		DefinitionContext: definitionContext,
	}
}

func (self *AdditionalInformation) StartIfDefBlock(anc *Tokenizer, isIfDef bool) {
	self.push(anc)
	self.stackState.expressionDirection = isIfDef
	anc.yyBEGIN(ifdefExpression)
}
