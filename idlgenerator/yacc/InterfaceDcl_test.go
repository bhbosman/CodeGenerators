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

func TestInterfaceDcl(t *testing.T) {
	factory := NewCompleteIdlLexerFactoryImpl(
		log.New(os.Stdout, "test", 0),
		scopedObjects.NewNextNumber(),
		scoping.NewScopingContext("", scoping.NewDefaultTypeService(), nil))
	t.Run("0001", func(t *testing.T) {
		stream := fmt.Sprintf(`
			interface IABC;
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
		interfaceDcl, ok := spec.(ScopingInterfaces.IInterfaceDcl)
		assert.True(t, ok)
		assert.True(t, interfaceDcl.Forward())
	})

	t.Run("0002", func(t *testing.T) {
		stream := fmt.Sprintf(`
			abstract interface IABC;
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
		interfaceDcl, ok := spec.(ScopingInterfaces.IInterfaceDcl)
		assert.True(t, ok)
		assert.True(t, interfaceDcl.Abstract())
	})

	t.Run("0003", func(t *testing.T) {
		stream := fmt.Sprintf(`
			local interface IABC;
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
		interfaceDcl, ok := spec.(ScopingInterfaces.IInterfaceDcl)
		assert.True(t, ok)
		assert.True(t, interfaceDcl.Local())
	})

	t.Run("0004", func(t *testing.T) {
		stream := fmt.Sprintf(`
			interface IABC;

			interface IABC
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

		f := func(a func() (ScopingInterfaces.ITypeSpec, error)) ScopingInterfaces.IInterfaceDcl {
			spec, _ := a()
			assert.NotNil(t, spec)
			interfaceDcl, ok := spec.(ScopingInterfaces.IInterfaceDcl)
			assert.True(t, ok)
			return interfaceDcl

		}
		interfaceDcl := f(actual.GetSpec)
		assert.True(t, interfaceDcl.Forward())

		interfaceDcl = f(interfaceDcl.GetNextTypeSpec)
		assert.False(t, interfaceDcl.Forward())
	})

	t.Run("0005", func(t *testing.T) {
		stream := fmt.Sprintf(`
			interface IABC;
			interface IABC;
			interface IABC
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

		f := func(a func() (ScopingInterfaces.ITypeSpec, error)) ScopingInterfaces.IInterfaceDcl {
			spec, _ := a()
			assert.NotNil(t, spec)
			interfaceDcl, ok := spec.(ScopingInterfaces.IInterfaceDcl)
			assert.True(t, ok)
			return interfaceDcl

		}
		interfaceDcl := f(actual.GetSpec)
		assert.True(t, interfaceDcl.Forward())

		interfaceDcl = f(interfaceDcl.GetNextTypeSpec)
		assert.True(t, interfaceDcl.Forward())

		interfaceDcl = f(interfaceDcl.GetNextTypeSpec)
		assert.False(t, interfaceDcl.Forward())
	})

}
