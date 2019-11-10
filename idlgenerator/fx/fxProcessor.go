package fx

import (
	"bufio"
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/AppInterfaces"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	"github.com/bhbosman/CodeGenerators/idlgenerator/scoping"
	"github.com/bhbosman/CodeGenerators/idlgenerator/yacc"

	"go.uber.org/fx"
	"log"
)

type processor struct {
	readerClosers            AppInterfaces.IIoReaders
	context                  AppInterfaces.IIdlGeneratorFlags
	logger                   *log.Logger
	definitionContextFactory AppInterfaces.IDefinitionContextFactory
	nextNumber               ScopingInterfaces.INextNumber
	generateCode             AppInterfaces.IScopeWalker
}

func (self *processor) Process() error {
	for _, fileInformation := range self.readerClosers.GetFileInformation() {
		typeSpec, err := func() (ScopingInterfaces.ITypeSpec, error) {
			self.logger.Printf("Processing %v\n", fileInformation.GetArg())

			lex, err := yacc.NewCompleteIdlLexImpl(
				fileInformation.GetArg(),
				bufio.NewReader(fileInformation.GetReader()),
				self.logger,
				self.definitionContextFactory.Create(),
				self.nextNumber,
				scoping.NewScopingContext(scoping.NewDefaultTypeService(), nil))
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
		err = self.generateCode.Generate(
			scoping.NewScopingContext(scoping.NewDefaultTypeService(), nil),
			0,
			typeSpec,
			fileInformation.GetFileName())
		if err != nil {
			continue
		}
	}
	return nil
}

func AppProvideProcessor() fx.Option {
	return fx.Provide(
		func(
			context AppInterfaces.IIdlGeneratorFlags,
			logger *log.Logger,
			ioReaders AppInterfaces.IIoReaders,
			definitionContextFactory AppInterfaces.IDefinitionContextFactory,
			nextNumber ScopingInterfaces.INextNumber,
			generateCode AppInterfaces.IScopeWalker) (AppInterfaces.IProcessor, error) {

			processor := &processor{
				readerClosers:            ioReaders,
				context:                  context,
				logger:                   logger,
				definitionContextFactory: definitionContextFactory,
				nextNumber:               nextNumber,
				generateCode:             generateCode,
			}
			return processor, nil
		})
}

func AppProvideDefaultTypeService() fx.Option {
	return fx.Provide(
		func() (ScopingInterfaces.IDefaultTypeService, error) {
			return scoping.NewDefaultTypeService(), nil
		})
}
