package yacc

import (
	"bufio"
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/scopedObjects"
	"github.com/bhbosman/CodeGenerators/idlgenerator/scoping"
	"github.com/bhbosman/gomock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strings"
	"testing"
)

func TestScopeName(t *testing.T) {
	factory := NewCompleteIdlLexerFactoryImpl(
		log.New(os.Stdout, "test", 0),
		scopedObjects.NewNextNumber(),
		scoping.NewScopingContext("", scoping.NewDefaultTypeService(), nil))
	t.Run("Struct scoping", func(t *testing.T) {
		stream := fmt.Sprintf(`
			struct A 
			{
				long B;
				double C;
			};
			struct D 
			{
				long E;
				double F;
				A 		G;
				A		H;
			};
		`)
		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))
		actual, _ := factory.Create("(string test)", reader)
		mock := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
		mock.EXPECT().NewStructType(gomock.Any(), "A", gomock.Any(), gomock.Any()).Times(1)
		mock.EXPECT().NewStructType(gomock.Any(), "D", gomock.Any(), gomock.Any()).Times(1)

		result := CompleteIdlParse(mock)
		assert.Equal(t, 0, result, mock.LastError())
	})

	t.Run("Struct scoping in module", func(t *testing.T) {
		stream := fmt.Sprintf(`
			module M 
			{
				struct A 
				{
					long B;
					double C;
				};
				struct D 
				{
					long E;
					double F;
					A 		G;
					A		H;
				};
			};
		`)
		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))
		actual, _ := factory.Create("(string test)", reader)
		mock := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
		mock.EXPECT().NewStructType(gomock.Any(), "A", gomock.Any(), gomock.Any()).Times(1)
		mock.EXPECT().NewStructType(gomock.Any(), "D", gomock.Any(), gomock.Any()).Times(1)

		result := CompleteIdlParse(mock)
		assert.Equal(t, 0, result, mock.LastError())
	})

	t.Run("Struct scoping declared inside module, used outside module", func(t *testing.T) {
		stream := fmt.Sprintf(`
			module M 
			{
				struct A 
				{
					long B;
					double C;
				};
				struct D 
				{
					long E;
					double F;
					A 		G;
					A		H;
				};
			};
			struct I 
			{
				double J;
				M::A K;
			};
		`)
		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))
		actual, _ := factory.Create("(string test)", reader)
		mock := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
		mock.EXPECT().NewStructType(gomock.Any(), "A", gomock.Any(), gomock.Any()).Times(1)
		mock.EXPECT().NewStructType(gomock.Any(), "D", gomock.Any(), gomock.Any()).Times(1)
		mock.EXPECT().CreateModuleDcl(gomock.Any(), "M", gomock.Any()).Times(1)
		mock.EXPECT().NewStructType(gomock.Any(), "I", gomock.Any(), gomock.Any()).Times(1)
		result := CompleteIdlParse(mock)
		if !assert.Equal(t, 0, result, mock.LastError()) {
			return
		}

		typeSpec, e := mock.GetSpec()
		assert.NoError(t, e)
		assert.NotNil(t, typeSpec)
	})
}
