package yacc

//go:generate goyacc -o idl.go -p "IdlExpr" idl.y

type ScopeName struct {
	name string
}

func (self *ScopeName) GetData() interface{} {
	return self
}

func (self *ScopeName) GetName() string {
	return self.name
}

func NewScopeName(name string) *ScopeName {
	return &ScopeName{name: name}
}
