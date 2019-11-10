package fx

import (
	"context"
	"flag"
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/AppInterfaces"
	"go.uber.org/fx"
	"os"
	"path"
)

type IDLGeneratorFlags struct {
	files []string
}

func NewIDLGeneratorFlags() *IDLGeneratorFlags {
	return &IDLGeneratorFlags{}
}

func (self *IDLGeneratorFlags) SetFiles(files []string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fullFileNames := make([]string, 0)
	for _, fileName := range files {
		absoluteFileName := path.Join(dir, fileName)
		info, err := os.Stat(absoluteFileName)
		if err != nil {
			return fmt.Errorf("could not find file %v. error %v", absoluteFileName, err.Error())
		}
		if info.IsDir() {
			return fmt.Errorf("a file must be specfified, not a folder. %v", absoluteFileName)
		}
		fullFileNames = append(fullFileNames, absoluteFileName)
	}
	self.files = fullFileNames
	if len(self.files) == 0 {
		return fmt.Errorf("no files in argument")
	}
	return nil
}

func (self *IDLGeneratorFlags) Files() []string {
	return self.files
}

func AppProvideContext() fx.Option {
	return fx.Provide(func() (AppInterfaces.IIdlGeneratorFlags, AppInterfaces.ISetIdlGeneratorFlags, error) {
		ctx := NewIDLGeneratorFlags()
		return ctx, ctx, nil
	})
}

func AppInvokeContext() fx.Option {
	return fx.Invoke(func(lifecycle fx.Lifecycle, ctx AppInterfaces.ISetIdlGeneratorFlags) error {
		lifecycle.Append(fx.Hook{
			OnStart: func(contextStart context.Context) error {
				flag.Parse()
				return ctx.SetFiles(flag.Args())
			},
			OnStop: nil,
		})
		return nil
	})
}
