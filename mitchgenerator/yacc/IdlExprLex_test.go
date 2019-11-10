package yacc

import (
	"bufio"
	"fmt"
	ctx2 "github.com/bhbosman/CodeGenerators/ctx"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestIdlExprLex(t *testing.T) {

	t.Run("One Enum Value assigned int", func(t *testing.T) {
		data := `enum ABC {AA='a'};`
		reader := bufio.NewReader(strings.NewReader(data))
		ctx := ctx2.NewIdlExprContext()
		idlExprLex, _ := NewIdlExprLex(
			reader, ctx)
		if !assert.Equal(t, 0, IdlExprParse(idlExprLex)) {
			return
		}
		DeclaredTypes := ctx.GetSpecification()
		if !assert.NotNil(t, DeclaredTypes) {
			return
		}
		assert.Len(t, DeclaredTypes, 1)
	})

	t.Run("", func(t *testing.T) {
		a := byte(65)

		fmt.Println(string(a))
	})

}
