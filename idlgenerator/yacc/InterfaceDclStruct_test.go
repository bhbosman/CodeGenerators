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

func TestInterfaceDclStruct(t *testing.T) {
	factory := NewCompleteIdlLexerFactoryImpl(
		log.New(os.Stdout, "test", 0),
		scopedObjects.NewNextNumber(),
		scoping.NewScopingContext("", scoping.NewDefaultTypeService(), nil))
	t.Run("0001", func(t *testing.T) {
		stream := fmt.Sprintf(`
			interface IABC
			{
        		struct SS 
				{
				};
				SS ABC();
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
		interfaceDcl, ok := spec.(ScopingInterfaces.IInterfaceDcl)
		assert.True(t, ok)
		assert.False(t, interfaceDcl.Forward())
		assert.NotNil(t, interfaceDcl.GetBody())
		assert.Equal(t, 2, interfaceDcl.BodyCount())
		bodyArray := interfaceDcl.BodyArray()
		assert.NotNil(t, bodyArray)
		assert.Len(t, bodyArray, 2)
		structDcl, ok := bodyArray[0].(ScopingInterfaces.IStructType)
		assert.True(t, ok)
		assert.NotNil(t, structDcl)

		operation, ok := bodyArray[1].(ScopingInterfaces.IOperationDeclarations)
		assert.True(t, ok)
		assert.NotNil(t, operation)
	})

	t.Run("0002", func(t *testing.T) {
		stream := fmt.Sprintf(`
			interface IABC
			{
        		struct SS 
				{
				};
        		struct TT 
				{
				};
				SS ABC(in TT abc);
				TT DEF(in SS abc);
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
		interfaceDcl, ok := spec.(ScopingInterfaces.IInterfaceDcl)
		assert.True(t, ok)
		assert.False(t, interfaceDcl.Forward())
		assert.NotNil(t, interfaceDcl.GetBody())
		assert.Equal(t, 4, interfaceDcl.BodyCount())
		bodyArray := interfaceDcl.BodyArray()
		assert.NotNil(t, bodyArray)
		assert.Len(t, bodyArray, 4)

		structDcl, ok := bodyArray[0].(ScopingInterfaces.IStructType)
		assert.True(t, ok)
		assert.NotNil(t, structDcl)

		structDcl, ok = bodyArray[1].(ScopingInterfaces.IStructType)
		assert.True(t, ok)
		assert.NotNil(t, structDcl)

	})

	t.Run("0001", func(t *testing.T) {
		stream := fmt.Sprintf(`
			interface IABC
			{
        		struct SS 
				{
				};
				
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
		interfaceDcl, ok := spec.(ScopingInterfaces.IInterfaceDcl)
		assert.True(t, ok)
		assert.False(t, interfaceDcl.Forward())
		assert.NotNil(t, interfaceDcl.GetBody())
		assert.Equal(t, 1, interfaceDcl.BodyCount())
		bodyArray := interfaceDcl.BodyArray()
		assert.NotNil(t, bodyArray)
		assert.Len(t, bodyArray, 1)
		structDcl, ok := bodyArray[0].(ScopingInterfaces.IStructType)
		assert.True(t, ok)
		assert.NotNil(t, structDcl)
	})
}
