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

func TestInterfaceOperationsDcl(t *testing.T) {
	factory := NewCompleteIdlLexerFactoryImpl(
		log.New(os.Stdout, "test", 0),
		scopedObjects.NewNextNumber(),
		scoping.NewScopingContext(scoping.NewDefaultTypeService(), nil))
	t.Run("Interface with void operation with no params", func(t *testing.T) {
		stream := fmt.Sprintf(`
			interface IABC
			{
				void VoidOperation();
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
		assert.False(t, interfaceDcl.Forward())
		assert.Equal(t, "IABC", interfaceDcl.GetName())
		body := interfaceDcl.GetBody()
		declarations, ok := body.(ScopingInterfaces.IOperationDeclarations)
		assert.True(t, ok)
		assert.NotNil(t, body)
		assert.Equal(t, "VoidOperation", declarations.GetOperationName())
		assert.Equal(t, "void", declarations.GetOperationDeclaratorType().GetName())
	})
	t.Run("Interface with four void operation with no params", func(t *testing.T) {
		stream := fmt.Sprintf(`
			interface IABC
			{
				void VoidOperation01();
				int32 VoidOperation02();
				double VoidOperation03();
				string VoidOperation04();
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
		assert.False(t, interfaceDcl.Forward())
		assert.Equal(t, "IABC", interfaceDcl.GetName())

		// method 1
		body := interfaceDcl.GetBody()
		declarations, ok := body.(ScopingInterfaces.IOperationDeclarations)
		assert.True(t, ok)
		assert.NotNil(t, body)
		assert.Equal(t, "VoidOperation01", declarations.GetOperationName())
		assert.Equal(t, "void", declarations.GetOperationDeclaratorType().GetName())

		var err error
		// method 2
		body, err = body.GetNextTypeSpec()
		assert.NoError(t, err)
		assert.NotNil(t, body)
		declarations, ok = body.(ScopingInterfaces.IOperationDeclarations)
		assert.True(t, ok)
		assert.NotNil(t, body)
		assert.Equal(t, "VoidOperation02", declarations.GetOperationName())
		assert.Equal(t, "int32", declarations.GetOperationDeclaratorType().GetName())

		// method 3
		body, err = body.GetNextTypeSpec()
		assert.NoError(t, err)
		assert.NotNil(t, body)
		declarations, ok = body.(ScopingInterfaces.IOperationDeclarations)
		assert.True(t, ok)
		assert.NotNil(t, body)
		assert.Equal(t, "VoidOperation03", declarations.GetOperationName())
		assert.Equal(t, "double", declarations.GetOperationDeclaratorType().GetName())

		// method 4
		body, err = body.GetNextTypeSpec()
		assert.NoError(t, err)
		assert.NotNil(t, body)
		declarations, ok = body.(ScopingInterfaces.IOperationDeclarations)
		assert.True(t, ok)
		assert.NotNil(t, body)
		assert.Equal(t, "VoidOperation04", declarations.GetOperationName())
		assert.Equal(t, "string", declarations.GetOperationDeclaratorType().GetName())

		//typeSpec, err := actual.GetSpec()
		//golang := fx.GenerateCodeGolang{}
		//golang.InternalGenerate(0, typeSpec, "")
	})
	t.Run("Interface with return value operation with params", func(t *testing.T) {
		stream := fmt.Sprintf(`
			interface IABC
			{
				long Operation01(in long a, in long b);
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
		assert.False(t, interfaceDcl.Forward())
		assert.Equal(t, "IABC", interfaceDcl.GetName())

		//method Operation01
		body := interfaceDcl.GetBody()
		assert.NotNil(t, body)
		operationDeclaration, ok := body.(ScopingInterfaces.IOperationDeclarations)
		assert.True(t, ok)
		assert.NotNil(t, body)
		assert.Equal(t, "Operation01", operationDeclaration.GetOperationName())
		assert.Equal(t, "long", operationDeclaration.GetOperationDeclaratorType().GetName())

		//params
		params := operationDeclaration.GetParams()
		assert.NotNil(t, params)
		//params a
		paramDecl, ok := params.(ScopingInterfaces.IParameterDeclarations)
		assert.True(t, ok)
		assert.NotNil(t, paramDecl)
		assert.Equal(t, "long", paramDecl.GetParamDeclarationType().GetName())
		assert.Equal(t, "a", paramDecl.GetParamName())
		//params b
		paramDecl = paramDecl.GetNextParameterDeclarations()
		assert.NotNil(t, paramDecl)
		assert.Equal(t, "long", paramDecl.GetParamDeclarationType().GetName())
		assert.Equal(t, "b", paramDecl.GetParamName())
	})
	t.Run("Interface with return value operation using interface", func(t *testing.T) {
		stream := fmt.Sprintf(`
			interface IABC
			{
				IABC Operation01(in IABC a, in IABC b);
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
		assert.False(t, interfaceDcl.Forward())
		assert.Equal(t, "IABC", interfaceDcl.GetName())

		//method Operation01
		body := interfaceDcl.GetBody()
		assert.NotNil(t, body)
		operationDeclaration, ok := body.(ScopingInterfaces.IOperationDeclarations)
		assert.True(t, ok)
		assert.NotNil(t, body)
		assert.Equal(t, "Operation01", operationDeclaration.GetOperationName())
		assert.Equal(t, "IABC", operationDeclaration.GetOperationDeclaratorType().GetName())

		//params
		params := operationDeclaration.GetParams()
		assert.NotNil(t, params)
		//params a
		paramDecl, ok := params.(ScopingInterfaces.IParameterDeclarations)
		assert.True(t, ok)
		assert.NotNil(t, paramDecl)
		assert.Equal(t, "IABC", paramDecl.GetParamDeclarationType().GetName())
		assert.Equal(t, "a", paramDecl.GetParamName())
		//params b
		paramDecl = paramDecl.GetNextParameterDeclarations()
		assert.NotNil(t, paramDecl)
		assert.Equal(t, "IABC", paramDecl.GetParamDeclarationType().GetName())
		assert.Equal(t, "b", paramDecl.GetParamName())
	})
}
