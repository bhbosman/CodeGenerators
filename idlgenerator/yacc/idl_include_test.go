package yacc

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/CodeGeneration"
	"github.com/bhbosman/CodeGenerators/idlgenerator/scopedObjects"
	"github.com/bhbosman/CodeGenerators/idlgenerator/scoping"
	"github.com/bhbosman/gomock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"path"
	"testing"
)

func TestName(t *testing.T) {
	factory := NewCompleteIdlLexerFactoryImpl(
		log.New(os.Stdout, "test", 0),
		scopedObjects.NewNextNumber(),
		scoping.NewScopingContext(scoping.NewDefaultTypeService(), nil))
	t.Run("Basic struct in a definition block", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()
		mock := NewMockIDefinitionContext(controller)
		wd, err := os.Getwd()
		if !assert.NoError(t, err) {
			return
		}
		filePath := path.Join(wd, "./testFiles/first.idl")
		f, err := os.Open(filePath)
		if !assert.NoError(t, err) {
			return
		}
		defer func() {
			_ = f.Close()
		}()

		stats, err := f.Stat()
		if !assert.NoError(t, err) {
			return
		}
		fmt.Println(stats.Name())

		actual, _ := factory.CreateWithDefinitionContext(filePath, f, mock)
		mockCompleteIdlLexer := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
		mockCompleteIdlLexer.EXPECT().NewStructType(gomock.Any(), "ABC", gomock.Any(), gomock.Any()).Times(1)
		result := CompleteIdlParse(mockCompleteIdlLexer)
		assert.Equal(t, 0, result, mockCompleteIdlLexer.LastError())
	})

	t.Run("Basic struct in a definition block", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()
		mock := NewMockIDefinitionContext(controller)
		wd, err := os.Getwd()
		if !assert.NoError(t, err) {
			return
		}
		filePath := path.Join(wd, "./testFiles/second.idl")
		f, err := os.Open(filePath)
		if !assert.NoError(t, err) {
			return
		}
		defer func() {
			_ = f.Close()
		}()

		stats, err := f.Stat()
		if !assert.NoError(t, err) {
			return
		}
		fmt.Println(stats.Name())

		actual, _ := factory.CreateWithDefinitionContext(filePath, f, mock)
		mockCompleteIdlLexer := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
		mockCompleteIdlLexer.EXPECT().NewStructType(gomock.Any(), "B", gomock.Any(), gomock.Any()).Times(1)
		result := CompleteIdlParse(mockCompleteIdlLexer)
		assert.Equal(t, 0, result, mockCompleteIdlLexer.LastError())

		goLang := CodeGeneration.NewGenerateCodeGolang()
		typeSpec, _ := actual.GetSpec()
		goLang.Generate(typeSpec)

	})

	t.Run("Basic struct in a definition block", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()
		mock := NewMockIDefinitionContext(controller)
		wd, err := os.Getwd()
		if !assert.NoError(t, err) {
			return
		}
		filePath := path.Join(wd, "./testFiles/third.idl")
		f, err := os.Open(filePath)
		if !assert.NoError(t, err) {
			return
		}
		defer func() {
			_ = f.Close()
		}()

		stats, err := f.Stat()
		if !assert.NoError(t, err) {
			return
		}
		fmt.Println(stats.Name())

		actual, _ := factory.CreateWithDefinitionContext(filePath, f, mock)

		result := CompleteIdlParse(actual)
		assert.Equal(t, 0, result, actual.LastError())

		goLang := CodeGeneration.NewGenerateCodeGolang()
		typeSpec, _ := actual.GetSpec()
		goLang.Generate(typeSpec)

	})

}
