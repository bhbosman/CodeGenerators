package fx

import (
	"github.com/bhbosman/CodeGenerators/idlgenerator/AppInterfaces"
	"github.com/bhbosman/CodeGenerators/idlgenerator/yacc"
	"go.uber.org/fx"
	"strconv"
)

type DefinitionContextFactory struct {
}

func (d DefinitionContextFactory) Create() yacc.IDefinitionContext {
	return &DefinitionContext{
		values: make(map[string]int),
	}
}

type DefinitionContext struct {
	values map[string]int
}

func (d DefinitionContext) ParseExpression(s string) bool {
	b, e := strconv.ParseBool(s)
	if e == nil {
		return b
	}
	_, ok := d.values[s]
	return ok
}

func (d DefinitionContext) ParseDefinition(s string, b bool) {
	if b {
		d.values[s] = 1
	} else {
		delete(d.values, s)
	}
}

func (d DefinitionContext) ParsePragma(string) {

}

func FxAppProvideDefinitionFactoryContext() fx.Option {
	return fx.Provide(func() (AppInterfaces.IDefinitionContextFactory, error) {
		definitionContextFactory := &DefinitionContextFactory{}
		return definitionContextFactory, nil
	})
}
