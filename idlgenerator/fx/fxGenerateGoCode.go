package fx

import (
	"github.com/bhbosman/CodeGenerators/idlgenerator/AppInterfaces"
	"github.com/bhbosman/CodeGenerators/idlgenerator/scoping"
	"go.uber.org/fx"
	"log"
)

func AppProvideGenerateCodeGolang() fx.Option {
	return fx.Provide(
		func(logger *log.Logger) (AppInterfaces.IScopeWalker, error) {
			generateCodeGolang := scoping.NewScopeWalker(logger)
			return generateCodeGolang, nil
		})
}
