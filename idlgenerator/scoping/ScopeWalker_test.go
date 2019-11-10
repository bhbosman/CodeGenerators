package scoping

import (
	"bufio"
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/CodeGeneration"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	"github.com/bhbosman/CodeGenerators/idlgenerator/scopedObjects"
	"github.com/bhbosman/CodeGenerators/idlgenerator/yacc"
	"github.com/bhbosman/gomock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strings"
	"testing"
)

func TestGenerateCodeGolang_Generate(t *testing.T) {
	var scopingContext ScopingInterfaces.IScopingContext
	scopingContext = NewScopingContext(NewDefaultTypeService(), nil)
	factory := yacc.NewCompleteIdlLexerFactoryImpl(
		log.New(os.Stdout, "test", 0),
		scopedObjects.NewNextNumber(),
		scopingContext)

	t.Run("Interface with return value operation with params", func(t *testing.T) {
		stream := fmt.Sprintf(`
			module A
			{
				struct A
				{
					int8 adsada,aadasdas,adsadasa,a;
				};
				struct B
				{
				};
				
				interface IABC
				{
					B aa();
					A aa();
				};
			};
			struct A
			{
				A::A a;
			};

		`)
		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))
		actual, _ := factory.Create("(string test)", reader)

		mock := yacc.NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
		result := yacc.CompleteIdlParse(mock)
		if !assert.Equal(t, 0, result, mock.LastError()) {
			return
		}

		interfaceDcl, err := actual.GetSpec()
		assert.NoError(t, err)
		assert.NotNil(t, interfaceDcl)

		dd := NewScopeWalker(log.New(os.Stdout, "", 0))

		newScopingContext := NewScopingContext(nil, scopingContext)
		err = dd.Generate(newScopingContext, 0, interfaceDcl, "")

		_ = scopingContext.Iterate(func(key string, value ScopingInterfaces.IDeclaredType) error {
			if !value.IsPrimitive() {
				fmt.Println(key)
			}
			return nil
		})
		generateCodeGolang := CodeGeneration.NewGenerateCodeGolang(log.New(os.Stdout, "", 0))
		_ = generateCodeGolang.Generate(interfaceDcl)

		assert.NoError(t, err)
	})
}