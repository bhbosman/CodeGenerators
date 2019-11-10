package fx

import (
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	"github.com/bhbosman/CodeGenerators/idlgenerator/scopedObjects"
	"go.uber.org/fx"
)

func AppNextNumber() fx.Option {
	return fx.Provide(func() (ScopingInterfaces.INextNumber, error) {
		nextNumber := &scopedObjects.NextNumber{
			Number: 0,
		}
		return nextNumber, nil
	})
}
