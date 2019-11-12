package main

import (
	"context"
	"github.com/bhbosman/CodeGenerators/idlgenerator/AppInterfaces"
	fx2 "github.com/bhbosman/CodeGenerators/idlgenerator/fx"
	"go.uber.org/fx"
	"log"
	"os"
)

func main() {
	var applicationLogger = log.New(os.Stdout, "Idl Generator: ", log.Ldate|log.Ltime|log.Lmicroseconds)
	var processor AppInterfaces.IProcessor
	app := fx.New(
		fx.StartTimeout(fx.DefaultTimeout),
		fx.StopTimeout(fx.DefaultTimeout),
		FxAppProvideFxAppOverrideLogger(applicationLogger),
		FxAppProvideApplicationLogger(applicationLogger),
		fx2.AppProvideCodeGenerator(),
		fx2.AppProvideDefaultTypeService(),
		fx2.AppProvideContext(),
		fx2.AppNextNumber(),
		fx2.AppInvokeContext(),
		fx2.FxAppProvideIoReaders(),
		fx2.FxAppInvokeIoReaders(),
		fx2.AppProvideProcessor(),
		fx2.FxAppProvideDefinitionFactoryContext(),
		fx2.AppProvideGenerateCodeGolang(),
		fx.Populate(&processor))

	startError := app.Start(context.TODO())
	if startError != nil {
		os.Exit(1)
	}
	defer func() {
		_ = app.Stop(context.TODO())
	}()

	err := processor.Process()
	if err == nil {

	}
}
