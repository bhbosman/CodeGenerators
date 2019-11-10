package yacc

//go:generate goyacc -o idl.go -p "IdlExpr" idl.y
//go:generate golex idl.l

import (
	"github.com/bhbosman/CodeGenerators/ctx"
	"io"
	"log"
	"strconv"
)

type IdlExprLex struct {
	InputStream    io.Reader
	idlExprContext *ctx.IdlExprContext
}

func (x *IdlExprLex) Lex(value *IdlExprSymType) int {
	for {
		yyin = x.InputStream
		token := yylex()
		lexValue := yytext
		switch token {
		case Character_literal:
			value.StringValue = lexValue[1:2]

		case Identifier:
			value.Identifier = lexValue
			break
		case Integer_literal:
			integerValue, _ := strconv.ParseInt(lexValue, 10, 64)
			value.IntegerValue = integerValue
			break
		case Hex_literal:
			integerValue, _ := strconv.ParseInt(lexValue, 0, 64)
			value.IntegerValue = integerValue
			token = Integer_literal
			break
		}
		return token
	}
}

func (x *IdlExprLex) Error(s string) {
	log.Printf("parse error: %s at (%d,%d).", s, row, column)
}

func NewIdlExprLex(
	inputStream io.Reader,
	IdlExprContext *ctx.IdlExprContext) (*IdlExprLex, error) {
	IdlExprErrorVerbose = true

	return &IdlExprLex{
		InputStream:    inputStream,
		idlExprContext: IdlExprContext,
	}, nil
}
