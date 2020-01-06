package fx

import (
	"bufio"
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/AppImpl"
	ai "github.com/bhbosman/CodeGenerators/idlgenerator/AppInterfaces"
	"github.com/bhbosman/CodeGenerators/idlgenerator/CodeGeneration"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	"github.com/bhbosman/CodeGenerators/idlgenerator/scoping"
	"github.com/bhbosman/CodeGenerators/idlgenerator/yacc"
	"go.uber.org/fx"
	"go.uber.org/multierr"
	"log"
)

type processor struct {
	readerClosers            ai.IIoReaders
	context                  ai.IIdlGeneratorFlags
	logger                   *log.Logger
	definitionContextFactory ai.IDefinitionContextFactory
	nextNumber               ScopingInterfaces.INextNumber
	scopeWalker              ai.IScopeWalker
	codeGenerator            ai.ICodeGenerator
}

func (self *processor) Process() error {
	var err error
	for _, fileInformation := range self.readerClosers.GetFileInformation() {
		var typeSpec ScopingInterfaces.ITypeSpec
		typeSpec, err = func() (ScopingInterfaces.ITypeSpec, error) {
			self.logger.Printf("Processing %v\n", fileInformation.GetArg())

			lex, err := yacc.NewCompleteIdlLexImpl(
				fileInformation.GetArg(),
				bufio.NewReader(fileInformation.GetReader()),
				self.logger,
				self.definitionContextFactory.Create(),
				self.nextNumber,
				scoping.NewScopingContext("", scoping.NewDefaultTypeService(), nil))
			if err != nil {
				return nil, err
			}
			resultInstance := yacc.CompleteIdlParse(lex)
			if resultInstance != 0 {
				return nil, fmt.Errorf(lex.LastError())
			}
			return lex.GetSpec()
		}()
		if err != nil {
			continue
		}
		err = multierr.Append(
			err,
			self.scopeWalker.Scope(
				scoping.NewScopingContext("", scoping.NewDefaultTypeService(), nil),
				0,
				typeSpec,
				fileInformation.GetFileName()))

		err = multierr.Append(
			err,
			self.codeGenerator.Generate(typeSpec))
		if err != nil {
			continue
		}
	}
	return nil
}

func AppProvideProcessor() fx.Option {
	return fx.Provide(
		func(
			context ai.IIdlGeneratorFlags,
			logFactory AppImpl.ILogFactory,
			ioReaders ai.IIoReaders,
			definitionContextFactory ai.IDefinitionContextFactory,
			nextNumber ScopingInterfaces.INextNumber,
			scopeWalker ai.IScopeWalker,
			codeGenerator ai.ICodeGenerator) (ai.IProcessor, error) {

			var processor = &processor{
				readerClosers:            ioReaders,
				context:                  context,
				logger:                   logFactory.Create(),
				definitionContextFactory: definitionContextFactory,
				nextNumber:               nextNumber,
				scopeWalker:              scopeWalker,
				codeGenerator:            codeGenerator,
			}
			return processor, nil
		})
}

func AppProvideCodeGenerator() fx.Option {
	return fx.Provide(
		func() (ai.ICodeGenerator, error) {
			return CodeGeneration.NewGenerateCodeGolang(), nil
		})
}

func AppProvideDefaultTypeService() fx.Option {
	return fx.Provide(
		func() (ScopingInterfaces.IDefaultTypeService, error) {
			return scoping.NewDefaultTypeService(), nil
		})
}
