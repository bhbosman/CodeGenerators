package main

import (
	"context"
	"flag"
	"github.com/bhbosman/CodeGenerators/idlgenerator/AppInterfaces"
	"github.com/bhbosman/CodeGenerators/idlgenerator/CodeGeneration"
	fx2 "github.com/bhbosman/CodeGenerators/idlgenerator/fx"
	"go.uber.org/fx"
	"io"
	"log"
	"os"
)

func main() {
	verbose := flag.Bool("verbose", false, "Logging")
	flag.Parse()

	out := getLogger(*verbose)




	applicationLogger := log.New(out, "Idl Generator: ", log.Ldate|log.Ltime|log.Lmicroseconds)
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

func getLogger(verbose bool) io.Writer  {
	if verbose {
		return os.Stdout
	}
	return  &CodeGeneration.NullWriter{}
}
