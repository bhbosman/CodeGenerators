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

func TestModuleDecl(t *testing.T) {
	factory := NewCompleteIdlLexerFactoryImpl(
		log.New(os.Stdout, "test", 0),
		scopedObjects.NewNextNumber(),
		scoping.NewScopingContext("", scoping.NewDefaultTypeService(), nil))
	t.Run("Empty Module", func(t *testing.T) {
		stream := fmt.Sprintf(`
			module ABC 
			{
			};`)
		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))
		actual, _ := factory.Create("(string test)", reader)

		mock := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
		mock.EXPECT().CreateModuleDcl(gomock.Any(), "ABC", nil).Times(1)
		result := CompleteIdlParse(mock)
		assert.Equal(t, 0, result, mock.LastError())
	})

	t.Run("Module with forward decls", func(t *testing.T) {
		stream := fmt.Sprintf(`
			module ABC 
			{
				interface IDEF;
				interface IGHI;
			};
		`)
		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))
		actual, _ := factory.Create("(string test)", reader)

		mock := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
		mock.EXPECT().CreateModuleDcl(gomock.Any(), "ABC", gomock.Any()).Times(1)
		mock.EXPECT().CreateInterfaceDcl(gomock.Any(), "IDEF", gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(1)
		mock.EXPECT().CreateInterfaceDcl(gomock.Any(), "IGHI", gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(1)
		result := CompleteIdlParse(mock)
		assert.Equal(t, 0, result, mock.LastError())
		spec, _ := actual.GetSpec()
		assert.NotNil(t, spec)

		moduleDcl := spec.(ScopingInterfaces.IIdlModuleDcl)
		assert.NotNil(t, moduleDcl)
		assert.Equal(t, "ABC", moduleDcl.GetName())

		spec = moduleDcl.GetModuleExports()
		assert.NotNil(t, spec)
		assert.Equal(t, "IDEF", spec.GetName())

		spec, _ = spec.GetNextTypeSpec()
		assert.NotNil(t, spec)
		assert.Equal(t, "IGHI", spec.GetName())

		spec, _ = spec.GetNextTypeSpec()
		assert.Nil(t, spec)
	})

	t.Run("Module with forward decls", func(t *testing.T) {
		stream := fmt.Sprintf(`
			module ABC 
			{
				module DEF 
				{
					interface IGHI;
					interface IJKL;
				};
				struct MNO {
					DEF::IGHI a;
				};
			};
		`)
		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))
		actual, _ := factory.Create("(string test)", reader)

		mock := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
		mock.EXPECT().CreateModuleDcl(gomock.Any(), "ABC", gomock.Any()).Times(1)
		mock.EXPECT().CreateModuleDcl(gomock.Any(), "DEF", gomock.Any()).Times(1)
		mock.EXPECT().CreateInterfaceDcl(gomock.Any(), "IGHI", gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(1)
		mock.EXPECT().CreateInterfaceDcl(gomock.Any(), "IJKL", gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(1)
		mock.EXPECT().NewStructType(gomock.Any(), "MNO", gomock.Any(), gomock.Any()).Times(1)
		result := CompleteIdlParse(mock)
		if !assert.Equal(t, 0, result, mock.LastError()) {
			return
		}
		spec, _ := actual.GetSpec()
		assert.NotNil(t, spec)

		moduleDcl := spec.(ScopingInterfaces.IIdlModuleDcl)
		assert.NotNil(t, moduleDcl)
		assert.Equal(t, "ABC", moduleDcl.GetName())

		spec = moduleDcl.GetModuleExports()
		assert.NotNil(t, spec)

		moduleDcl = spec.(ScopingInterfaces.IIdlModuleDcl)
		assert.NotNil(t, moduleDcl)

		assert.NotNil(t, spec)
		assert.Equal(t, "DEF", spec.GetName())

		spec = moduleDcl.GetModuleExports()
		assert.NotNil(t, spec)
		assert.Equal(t, "IGHI", spec.GetName())

		spec, _ = spec.GetNextTypeSpec()
		assert.NotNil(t, spec)
		assert.Equal(t, "IJKL", spec.GetName())

		spec, _ = spec.GetNextTypeSpec()
		assert.Nil(t, spec)
	})

}
