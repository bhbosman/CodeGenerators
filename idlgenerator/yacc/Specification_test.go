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

func TestSpecification(t *testing.T) {
	factory := NewCompleteIdlLexerFactoryImpl(
		log.New(os.Stdout, "test", 0),
		scopedObjects.NewNextNumber(),
		scoping.NewScopingContext(scoping.NewDefaultTypeService(), nil))

	t.Run("0001", func(t *testing.T) {
		stream := fmt.Sprintf(`
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
		spec, _ := actual.GetSpec()
		if !assert.NotNil(t, spec, "GetSpec does not have a value") {
			return
		}
		interfaceDcl, ok := spec.(ScopingInterfaces.IInterfaceDcl)
		if !assert.True(t, ok, "IInterfaceDcl not found") {
			return
		}
		assert.False(t, interfaceDcl.Forward())
	})

}
