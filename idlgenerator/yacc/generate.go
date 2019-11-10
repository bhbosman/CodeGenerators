package yacc

//go:generate mockgen -source CompleteIdlLexer.go  -package yacc -destination CompleteIdlLexerMock.go
///////go:generate golex completeIdl.l
//go:generate goyacc -o completeIdl.go  -p "CompleteIdl"  completeIdl.y
