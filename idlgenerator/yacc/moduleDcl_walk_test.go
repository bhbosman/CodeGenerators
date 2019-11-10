package yacc

import (
	"bufio"
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	"github.com/bhbosman/CodeGenerators/idlgenerator/scopedObjects"
	"github.com/bhbosman/CodeGenerators/idlgenerator/scoping"
	"github.com/bhbosman/gomock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strings"

	"testing"
)

func TestModuleDcl_Walk(t *testing.T) {
	factory := NewCompleteIdlLexerFactoryImpl(
		log.New(os.Stdout, "test", 0),
		scopedObjects.NewNextNumber(),
		scoping.NewScopingContext(scoping.NewDefaultTypeService(), nil))
	run := func(controller *gomock.Controller, stream string) (ScopingInterfaces.IIdlModuleDcl, error) {
		reader := bufio.NewReader(strings.NewReader(stream))
		actual, _ := factory.Create("(string test)", reader)
		result := CompleteIdlParse(actual)
		if !assert.Equal(t, 0, result, actual.LastError()) {
			return nil, fmt.Errorf(actual.LastError())
		}
		spec, err := actual.GetSpec()
		if !assert.NoError(t, err) {
			return nil, err
		}
		moduleDcl, ok := spec.(ScopingInterfaces.IIdlModuleDcl)
		if !assert.True(t, ok) {
			return nil, fmt.Errorf("no moduule")
		}
		return moduleDcl, nil
	}

	t.Run("Nested modules 01", func(t *testing.T) {
		stream := fmt.Sprintf(`
			module ABC 
			{
				module DEF
				{
					struct  GHI
					{
						long aa;
					};
					struct  JKL
					{
						long aa;
					};
				};
				struct MNO 
				{
					DEF::GHI a;
				};
			};
		`)
		controller := gomock.NewController(t)
		defer controller.Finish()
		moduleDcl, err := run(controller, stream)
		if !assert.NoError(t, err) {
			return
		}
		assert.NotNil(t, moduleDcl)

	})

	t.Run("Nested modules 01", func(t *testing.T) {
		stream := fmt.Sprintf(`
			module ABC 
			{
				struct  GHI
				{
					long aa;
				};
				struct  JKL
				{
					GHI aa;
				};

			};
		`)
		controller := gomock.NewController(t)
		defer controller.Finish()
		moduleDcl, err := run(controller, stream)
		if !assert.NoError(t, err) {
			return
		}
		assert.NotNil(t, moduleDcl)

	})

}
