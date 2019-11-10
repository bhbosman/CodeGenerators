package fx

import (
	"context"
	"github.com/bhbosman/CodeGenerators/idlgenerator/AppImpl"
	"github.com/bhbosman/CodeGenerators/idlgenerator/AppInterfaces"
	"go.uber.org/fx"
	"log"
)

func FxAppProvideIoReaders() fx.Option {
	return fx.Provide(func(context AppInterfaces.IIdlGeneratorFlags, logger *log.Logger) (AppInterfaces.IIoReaders, error) {
		ioReaders := &AppImpl.IoReaders{
			FileInformation: nil,
			Context:         context,
			Logger:          logger,
		}
		return ioReaders, nil
	})
}

func FxAppInvokeIoReaders() fx.Option {
	return fx.Invoke(func(lifecycle fx.Lifecycle, ioReaders AppInterfaces.IIoReaders) error {
		lifecycle.Append(fx.Hook{
			OnStart: func(contextStart context.Context) error {
				return ioReaders.Start()
			},
			OnStop: func(contextStop context.Context) error {
				return ioReaders.Stop()
			},
		})
		return nil
	})
}