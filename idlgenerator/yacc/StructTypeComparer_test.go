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

func TestStructTypeDecl(t *testing.T) {
	factory := NewCompleteIdlLexerFactoryImpl(
		log.New(os.Stdout, "test", 0),
		scopedObjects.NewNextNumber(),
		scoping.NewScopingContext(scoping.NewDefaultTypeService(), nil))
	t.Run("", func(t *testing.T) {
		stream := fmt.Sprintf(`
			struct abc;
		`)
		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))
		actual, _ := factory.Create("(string test)", reader)

		mock := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
		result := CompleteIdlParse(mock)
		assert.Equal(t, 0, result, mock.LastError())
		spec, _ := actual.GetSpec()
		assert.NotNil(t, spec)
		interfaceDcl, ok := spec.(ScopingInterfaces.IStructType)

		assert.True(t, ok)
		assert.True(t, interfaceDcl.Forward())
	})

	t.Run("", func(t *testing.T) {
		stream := fmt.Sprintf(`
			struct abc
			{
			};
		`)
		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))
		actual, _ := factory.Create("(string test)", reader)

		mock := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
		result := CompleteIdlParse(mock)
		assert.Equal(t, 0, result, mock.LastError())
		spec, _ := actual.GetSpec()
		assert.NotNil(t, spec)
		interfaceDcl, ok := spec.(ScopingInterfaces.IStructType)

		assert.True(t, ok)
		assert.False(t, interfaceDcl.Forward())
	})

	t.Run("", func(t *testing.T) {
		stream := fmt.Sprintf(`
			struct abc;
			struct abc
			{
			};
		`)
		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))
		actual, _ := factory.Create("(string test)", reader)

		mock := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
		result := CompleteIdlParse(mock)
		if !assert.Equal(t, 0, result, mock.LastError()) {
			return
		}
		//1
		spec, err := actual.GetSpec()
		assert.NoError(t, err)
		assert.NotNil(t, spec)
		structType, ok := spec.(ScopingInterfaces.IStructType)
		assert.True(t, ok)
		assert.True(t, structType.Forward())

		// 2
		spec, err = structType.GetNextTypeSpec()
		assert.NoError(t, err)
		assert.NotNil(t, spec)
		structType, ok = spec.(ScopingInterfaces.IStructType)
		assert.True(t, ok)
		assert.False(t, structType.Forward())

	})

}
