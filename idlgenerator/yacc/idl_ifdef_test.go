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
	"strconv"
	"strings"
	"testing"
)

func TestIfDef(t *testing.T) {
	factory := NewCompleteIdlLexerFactoryImpl(
		log.New(os.Stdout, "test", 0),
		scopedObjects.NewNextNumber(),
		scoping.NewScopingContext(scoping.NewDefaultTypeService(), nil))
	t.Run("Basic struct in a definition block", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()
		mock := NewMockIDefinitionContext(controller)
		mock.EXPECT().ParseExpression(gomock.Any()).MinTimes(1).DoAndReturn(func(s string) bool {
			b, _ := strconv.ParseBool(strings.TrimSpace(s))
			return b
		})

		stream := fmt.Sprintf(`
			#ifdef TRUE 
				struct a{long b;};
			#endif`)
		reader := bufio.NewReader(strings.NewReader(stream))
		actual, _ := factory.CreateWithDefinitionContext("(string test)", reader, mock)
		mockCompleteIdlLexer := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
		mockCompleteIdlLexer.EXPECT().NewStructType(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(1)
		result := CompleteIdlParse(mockCompleteIdlLexer)
		assert.Equal(t, 0, result, mockCompleteIdlLexer.LastError())
	})

	t.Run("Basic struct in a definition block", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()
		mock := NewMockIDefinitionContext(controller)
		mock.EXPECT().ParseExpression(gomock.Any()).MinTimes(1).DoAndReturn(func(s string) bool {
			b, _ := strconv.ParseBool(strings.TrimSpace(s))
			return b
		})

		stream := fmt.Sprintf(`
			#ifdef FALSE 
				struct a{long b;};
			#endif
			;`)
		reader := bufio.NewReader(strings.NewReader(stream))
		actual, _ := factory.CreateWithDefinitionContext("(string test)", reader, mock)
		mockCompleteIdlLexer := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
		mockCompleteIdlLexer.EXPECT().NewStructType(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
		result := CompleteIdlParse(mockCompleteIdlLexer)
		assert.Equal(t, 0, result, mockCompleteIdlLexer.LastError())
	})

	t.Run("Basic struct in a definition block, with 2 others", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()
		mock := NewMockIDefinitionContext(controller)
		mock.EXPECT().ParseExpression(gomock.Any()).MinTimes(1).DoAndReturn(func(s string) bool {
			b, _ := strconv.ParseBool(strings.TrimSpace(s))
			return b
		})

		stream := fmt.Sprintf(`
			struct a{long b;};
			#ifdef TRUE 
				struct b{long b;};
			#endif
			struct c{long b;};`)
		reader := bufio.NewReader(strings.NewReader(stream))
		actual, _ := factory.CreateWithDefinitionContext("(string test)", reader, mock)
		mockCompleteIdlLexer := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
		mockCompleteIdlLexer.EXPECT().NewStructType(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(3)
		result := CompleteIdlParse(mockCompleteIdlLexer)
		assert.Equal(t, 0, result, mockCompleteIdlLexer.LastError())
	})

	t.Run("Basic struct in a definition block, with 2 others", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()
		mock := NewMockIDefinitionContext(controller)
		mock.EXPECT().ParseExpression(gomock.Any()).MinTimes(1).DoAndReturn(func(s string) bool {
			b, _ := strconv.ParseBool(strings.TrimSpace(s))
			return b
		})

		stream := fmt.Sprintf(`
			struct a{long b;};
			#ifndef TRUE 
				struct b{long b;};
			#endif
			struct c{long b;};`)
		reader := bufio.NewReader(strings.NewReader(stream))
		actual, _ := factory.CreateWithDefinitionContext("(string test)", reader, mock)
		mockCompleteIdlLexer := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
		mockCompleteIdlLexer.EXPECT().NewStructType(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(2)
		result := CompleteIdlParse(mockCompleteIdlLexer)
		assert.Equal(t, 0, result, mockCompleteIdlLexer.LastError())
	})
}
