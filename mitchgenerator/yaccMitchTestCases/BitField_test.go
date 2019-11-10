package yaccIdlTests

import (
	"bufio"
	"github.com/bhbosman/CodeGenerators/ctx"
	"github.com/bhbosman/CodeGenerators/mitchgenerator/yacc"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestBitField(t *testing.T) {

	createContext := func() *ctx.IdlExprContext {
		return ctx.NewIdlExprContext()
	}

	t.Run("No Decl", func(t *testing.T) {
		data := `typedef MitchBitField<a0, a1, a2, a3, a4, a5, a6, a7> newBitType;`

		reader := bufio.NewReader(strings.NewReader(data))
		idlExprLex, _ := yacc.NewIdlExprLex(
			reader,
			createContext())
		assert.Equal(t, 0, yacc.IdlExprParse(idlExprLex))
	})

	t.Run("No Decl", func(t *testing.T) {
		data := `typedef bitfield<a0, a1, a2, a3, a4, a5, a6, a7> newBitType;`

		reader := bufio.NewReader(strings.NewReader(data))
		idlExprLex, _ := yacc.NewIdlExprLex(
			reader,
			createContext())
		assert.Equal(t, yacc.DefNotFound, yacc.IdlExprParse(idlExprLex))
	})

}
