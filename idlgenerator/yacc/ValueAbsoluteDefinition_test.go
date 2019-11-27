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

func TestValueType(t *testing.T) {
	factory := NewCompleteIdlLexerFactoryImpl(
		log.New(os.Stdout, "test", 0),
		scopedObjects.NewNextNumber(),
		scoping.NewScopingContext("", scoping.NewDefaultTypeService(), nil))
	t.Run("abstract valuetype forward declaration", func(t *testing.T) {
		stream := fmt.Sprintf(`
			abstract valuetype CustomMarshal;
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
		interfaceDcl, ok := spec.(ScopingInterfaces.IValueAbsoluteDefinition)

		assert.True(t, ok)
		assert.True(t, interfaceDcl.Abstract())
		assert.True(t, interfaceDcl.Forward())
	})

	t.Run("abstract valuetype  declaration", func(t *testing.T) {
		stream := fmt.Sprintf(`
			abstract valuetype CustomMarshal
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
		interfaceDcl, ok := spec.(ScopingInterfaces.IValueAbsoluteDefinition)
		assert.True(t, ok)
		assert.True(t, interfaceDcl.Abstract())
		assert.False(t, interfaceDcl.Forward())
	})

	t.Run("abstract valuetype  forward and declaration", func(t *testing.T) {
		stream := fmt.Sprintf(`
			abstract valuetype CustomMarshal;
			abstract valuetype CustomMarshal
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

		f := func(a func() (ScopingInterfaces.ITypeSpec, error)) ScopingInterfaces.IValueAbsoluteDefinition {
			spec, _ := a()
			assert.NotNil(t, spec)
			v, ok := spec.(ScopingInterfaces.IValueAbsoluteDefinition)
			assert.True(t, ok)
			return v
		}
		valueDefinition := f(actual.GetSpec)
		assert.True(t, valueDefinition.Abstract())
		assert.True(t, valueDefinition.Forward())

		valueDefinition = f(valueDefinition.GetNextTypeSpec)
		assert.True(t, valueDefinition.Abstract())
		assert.False(t, valueDefinition.Forward())
	})
}
