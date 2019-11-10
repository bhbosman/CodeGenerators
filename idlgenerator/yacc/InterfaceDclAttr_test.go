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

func TestInterfaceDclWithAttr(t *testing.T) {
	factory := NewCompleteIdlLexerFactoryImpl(
		log.New(os.Stdout, "test", 0),
		scopedObjects.NewNextNumber(),
		scoping.NewScopingContext(scoping.NewDefaultTypeService(), nil))
	t.Run("0001", func(t *testing.T) {
		stream := fmt.Sprintf(`
			interface IABC
			{
        		attribute long id;
				readonly attribute long idReadOnly;
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
		assert.Equal(t, 2, interfaceDcl.BodyCount())
		bodyArray := interfaceDcl.BodyArray()
		assert.NotNil(t, bodyArray)
		assert.Len(t, bodyArray, 2)
		attrDcl, ok := bodyArray[0].(ScopingInterfaces.IAttributeDcl)
		assert.True(t, ok)
		assert.False(t, attrDcl.ReadOnly())
		assert.Equal(t, attrDcl.DeclaredType().GetKind(), ScopingInterfaces.RWlongIdlType)
		assert.Equal(t, "id", attrDcl.AttrDeclarator().GetName())

		attrDcl, ok = bodyArray[1].(ScopingInterfaces.IAttributeDcl)
		assert.True(t, ok)
		assert.True(t, attrDcl.ReadOnly())
		assert.Equal(t, attrDcl.DeclaredType().GetKind(), ScopingInterfaces.RWlongIdlType)
		assert.Equal(t, "idReadOnly", attrDcl.AttrDeclarator().GetName())
	})

	t.Run("0002", func(t *testing.T) {
		stream := fmt.Sprintf(`
			interface IABC
			{
        		attribute long id01;
				attribute double id02, id03;
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
		assert.Equal(t, 2, interfaceDcl.BodyCount())
		bodyArray := interfaceDcl.BodyArray()
		assert.NotNil(t, bodyArray)
		assert.Len(t, bodyArray, 2)

		attrDcl, ok := bodyArray[0].(ScopingInterfaces.IAttributeDcl)
		assert.True(t, ok)
		assert.False(t, attrDcl.ReadOnly())
		assert.Equal(t, attrDcl.DeclaredType().GetKind(), ScopingInterfaces.RWlongIdlType)
		assert.Equal(t, "id01", attrDcl.AttrDeclarator().GetName())

		attrDcl, ok = bodyArray[1].(ScopingInterfaces.IAttributeDcl)
		assert.True(t, ok)
		assert.False(t, attrDcl.ReadOnly())
		assert.Equal(t, attrDcl.DeclaredType().GetKind(), ScopingInterfaces.RWdoubleIdlType)
		assert.Equal(t, "id02,id03", attrDcl.AttrDeclarator().GetName())
		array := attrDcl.AttrDeclarator().Names()
		assert.Len(t, array, 2)
		assert.Equal(t, array[0], "id02")
		assert.Equal(t, array[1], "id03")
	})

	t.Run("0003", func(t *testing.T) {
		stream := fmt.Sprintf(`
			interface IABC
			{
        		attribute long id;
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
		attrDcl, ok := bodyArray[0].(ScopingInterfaces.IAttributeDcl)
		assert.True(t, ok)
		assert.False(t, attrDcl.ReadOnly())
		assert.Equal(t, "id", attrDcl.AttrDeclarator().GetName())
	})

}
