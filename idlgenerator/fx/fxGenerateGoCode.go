package fx

import (
	"github.com/bhbosman/CodeGenerators/idlgenerator/AppInterfaces"
	"github.com/bhbosman/CodeGenerators/idlgenerator/scoping"
	"go.uber.org/fx"
)

func AppProvideGenerateCodeGolang() fx.Option {
	return fx.Provide(
		func() (AppInterfaces.IScopeWalker, error) {
			generateCodeGolang := scoping.NewScopeWalker()
			return generateCodeGolang, nil
		})
}
