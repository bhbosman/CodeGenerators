package yacc

import (
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	"io"
	"log"
)

type ICompleteIdlLexerFactory interface {
	Create(fileName string, inputStream io.Reader) (*CompleteIdlLexImpl, error)
	CreateWithDefinitionContext(fileName string, inputStream io.Reader, definitionContext IDefinitionContext) (*CompleteIdlLexImpl, error)
}

type CompleteIdlLexerFactoryImpl struct {
	logger         *log.Logger
	nextNumber     ScopingInterfaces.INextNumber
	scopingContext ScopingInterfaces.IScopingContext
}

func NewCompleteIdlLexerFactoryImpl(
	logger *log.Logger,
	nextNumber ScopingInterfaces.INextNumber,
	scopingContext ScopingInterfaces.IScopingContext) *CompleteIdlLexerFactoryImpl {
	return &CompleteIdlLexerFactoryImpl{
		logger:         logger,
		nextNumber:     nextNumber,
		scopingContext: scopingContext,
	}
}

func (self *CompleteIdlLexerFactoryImpl) Create(fileName string, inputStream io.Reader) (*CompleteIdlLexImpl, error) {
	return NewCompleteIdlLexImpl(
		fileName,
		inputStream,
		self.logger,
		nil,
		self.nextNumber,
		self.scopingContext)
}

func (self *CompleteIdlLexerFactoryImpl) CreateWithDefinitionContext(fileName string, inputStream io.Reader, definitionContext IDefinitionContext) (*CompleteIdlLexImpl, error) {
	return NewCompleteIdlLexImpl(
		fileName,
		inputStream,
		self.logger,
		definitionContext,
		self.nextNumber,
		self.scopingContext)
}
