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

func TestMathOperandsOnArraySize(t *testing.T) {
	factory := NewCompleteIdlLexerFactoryImpl(
		log.New(os.Stdout, "test", 0),
		scopedObjects.NewNextNumber(),
		scoping.NewScopingContext(scoping.NewDefaultTypeService(), nil))
	t.Run("SHL", func(t *testing.T) {
		stream := fmt.Sprintf("struct demo{double c[1024 << 2];};")

		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))

		idlExprLex, _ := factory.Create("(string test)", reader)
		result := CompleteIdlParse(idlExprLex)
		assert.Equal(t, 0, result, idlExprLex.LastError())
	})

	t.Run("SHR", func(t *testing.T) {
		stream := fmt.Sprintf("struct demo{double c[1024 >> 2];};")

		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))

		idlExprLex, _ := factory.Create("(string test)", reader)
		result := CompleteIdlParse(idlExprLex)
		assert.Equal(t, 0, result, idlExprLex.LastError())
	})

	t.Run("Plus", func(t *testing.T) {
		stream := fmt.Sprintf("struct demo{double c[1024 + 2];};")

		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))
		idlExprLex, _ := factory.Create("(string test)", reader)
		result := CompleteIdlParse(idlExprLex)
		assert.Equal(t, 0, result, idlExprLex.LastError())
	})

	t.Run("Minus", func(t *testing.T) {
		stream := fmt.Sprintf("struct demo{double c[1024 - 2];};")

		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))
		idlExprLex, _ := factory.Create("(string test)", reader)
		result := CompleteIdlParse(idlExprLex)
		assert.Equal(t, 0, result, idlExprLex.LastError())
	})

	t.Run("Multiply", func(t *testing.T) {
		stream := fmt.Sprintf("struct demo{double c[1024 * 2];};")

		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))
		idlExprLex, _ := factory.Create("(string test)", reader)
		result := CompleteIdlParse(idlExprLex)
		assert.Equal(t, 0, result, idlExprLex.LastError())
	})

	t.Run("Divide", func(t *testing.T) {
		stream := fmt.Sprintf("struct demo{double c[1024 / 2];};")

		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))

		idlExprLex, _ := factory.Create("(string test)", reader)
		result := CompleteIdlParse(idlExprLex)
		assert.Equal(t, 0, result, idlExprLex.LastError())
	})

	t.Run("Divide", func(t *testing.T) {
		stream := fmt.Sprintf("struct demo{double c[1024- -2];};")

		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))

		idlExprLex, _ := factory.Create("(string test)", reader)
		result := CompleteIdlParse(idlExprLex)
		assert.Equal(t, 0, result, idlExprLex.LastError())
	})

	t.Run("%", func(t *testing.T) {
		stream := fmt.Sprintf("struct demo{double c[1024%%233];};")

		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))

		idlExprLex, _ := factory.Create("(string test)", reader)
		result := CompleteIdlParse(idlExprLex)
		assert.Equal(t, 0, result, idlExprLex.LastError())
	})

	t.Run("&", func(t *testing.T) {
		stream := fmt.Sprintf("struct demo{double c[1024&233];};")

		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))

		idlExprLex, _ := factory.Create("(string test)", reader)
		result := CompleteIdlParse(idlExprLex)
		assert.Equal(t, 0, result, idlExprLex.LastError())
	})

	t.Run("|", func(t *testing.T) {
		stream := fmt.Sprintf("struct demo{double c[1024|233];};")

		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))

		idlExprLex, _ := factory.Create("(string test)", reader)
		result := CompleteIdlParse(idlExprLex)
		assert.Equal(t, 0, result, idlExprLex.LastError())
	})

	t.Run("+", func(t *testing.T) {
		stream := fmt.Sprintf("struct demo{double c[1+2];};")

		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))
		actual, _ := factory.Create("(string test)", reader)
		mockCompleteIdlLexer := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
		mockCompleteIdlLexer.EXPECT().AddOperator(1, 2).Times(1)
		mockCompleteIdlLexer.EXPECT().AddExpr(3).Times(1)
		result := CompleteIdlParse(mockCompleteIdlLexer)
		assert.Equal(t, 0, result, mockCompleteIdlLexer.LastError())
	})

	t.Run("Empty File", func(t *testing.T) {
		stream := fmt.Sprintf(";")
		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))
		actual, _ := factory.Create("(string test)", reader)
		mockCompleteIdlLexer := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
		result := CompleteIdlParse(mockCompleteIdlLexer)
		assert.Equal(t, 0, result, mockCompleteIdlLexer.LastError())
	})

	t.Run("Invalid array size 01", func(t *testing.T) {
		stream := fmt.Sprintf("struct demo{double c[-1];};")

		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))
		actual, _ := factory.Create("(string test)", reader)

		mockCompleteIdlLexer := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
		result := CompleteIdlParse(mockCompleteIdlLexer)
		assert.Equal(t, 1, result, mockCompleteIdlLexer.LastError())
		assert.NotEqual(t, 0, strings.Contains(mockCompleteIdlLexer.LastError(), "expression must be positive at"))
	})

	t.Run("Invalid array size 02", func(t *testing.T) {
		stream := fmt.Sprintf("struct demo{double c[100-100001];};")

		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))
		actual, _ := factory.Create("(string test)", reader)

		mockCompleteIdlLexer := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
		result := CompleteIdlParse(mockCompleteIdlLexer)
		assert.Equal(t, 1, result, mockCompleteIdlLexer.LastError())
		assert.NotEqual(t, 0, strings.Contains(mockCompleteIdlLexer.LastError(), "expression must be positive at"))
	})
}
