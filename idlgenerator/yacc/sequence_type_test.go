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

func TestSequenceType(t *testing.T) {
	factory := NewCompleteIdlLexerFactoryImpl(
		log.New(os.Stdout, "test", 0),
		scopedObjects.NewNextNumber(),
		scoping.NewScopingContext("", scoping.NewDefaultTypeService(), nil))
	t.Run("001", func(t *testing.T) {
		stream := fmt.Sprintf(`
			typedef sequence<unsigned long> type1;
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
		spec, _ := actual.GetSpec()
		assert.NotNil(t, spec)

		typeDeclarator, ok := spec.(ScopingInterfaces.ITypeDeclarator)
		if !assert.True(t, ok) {
			return
		}
		assert.NotNil(t, typeDeclarator)
		assert.Equal(t, "type1", typeDeclarator.GetName())
		assert.Equal(t, "sequence_unsigned long_0", typeDeclarator.TypeSpec().GetName())
		spec, _ = spec.GetNextTypeSpec()
		assert.Nil(t, spec)
	})

	t.Run("001", func(t *testing.T) {
		stream := fmt.Sprintf(`
			typedef sequence<unsigned long, 1024> type1;
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
		spec, _ := actual.GetSpec()
		assert.NotNil(t, spec)

	})

	t.Run("001 array", func(t *testing.T) {
		stream := fmt.Sprintf(`
			typedef sequence<unsigned long, -1> type1;
		`)
		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))
		actual, _ := factory.Create("(string test)", reader)

		mock := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
		result := CompleteIdlParse(mock)
		if !assert.Equal(t, 1, result, mock.LastError()) {
			return
		}

	})

	t.Run("002", func(t *testing.T) {
		stream := fmt.Sprintf(`
			typedef sequence<unsigned long> type1, type2; 
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

		typeDeclarator, ok := spec.(ScopingInterfaces.ITypeDeclarator)
		if !assert.True(t, ok) {
			return
		}
		assert.NotNil(t, typeDeclarator)
		assert.Equal(t, "type1,type2", typeDeclarator.GetName())
		assert.Equal(t, "sequence_unsigned long_0", typeDeclarator.TypeSpec().GetName())

	})

	t.Run("001 in module", func(t *testing.T) {
		stream := fmt.Sprintf(`
			module ABC
			{
				typedef sequence<unsigned long> type1;
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
		spec, _ := actual.GetSpec()
		assert.NotNil(t, spec)

		_, ok := spec.(ScopingInterfaces.IIdlModuleDcl)
		if !assert.True(t, ok) {
			return
		}

	})

	t.Run("002 in module", func(t *testing.T) {
		stream := fmt.Sprintf(`
			module ABC
			{
				typedef sequence<unsigned long> type1, type2;
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

	})

	t.Run("002 in two module", func(t *testing.T) {
		stream := fmt.Sprintf(`
			module ABC
			{
				typedef sequence<unsigned long> type1, type2;
			};
			module DEF
			{
				typedef sequence<unsigned long> type3, type4;
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

	})

}
