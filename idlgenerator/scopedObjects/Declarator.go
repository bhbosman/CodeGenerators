package scopedObjects

import "github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"

type Declarator struct {
	FileInformationBase
	Identifier     string
	nextDeclarator ScopingInterfaces.IDeclarator
}

func (self *Declarator) SetNext(next ScopingInterfaces.IDeclarator) {
	self.nextDeclarator = next
}

func (self *Declarator) String() string {
	return self.GetIdentifier()
}

func NewDeclarator(fileInformation ScopingInterfaces.IFileInformation, identifier string) *Declarator {
	return &Declarator{
		FileInformationBase: NewFileInformationBase02(fileInformation),
		Identifier:          identifier,
		nextDeclarator:      nil,
	}
}

//func NewDeclaratorWithIdentifier(identifier string) *Declarator {
//	return NewDeclarator(identifier, nil)
//}

func (self *Declarator) GetIdentifier() string {
	s := self.Identifier
	return s
}

func (self *Declarator) GetNext() ScopingInterfaces.IDeclarator {
	return self.nextDeclarator
}

func (self *Declarator) Next(next ScopingInterfaces.IDeclarator) ScopingInterfaces.IDeclarator {
	var item ScopingInterfaces.IDeclarator = self
	for item.GetNext() != nil {
		item = item.GetNext()
	}
	item.SetNext(next)
	return self
}
