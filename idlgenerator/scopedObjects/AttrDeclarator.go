package scopedObjects

import (
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	"strings"
)

//90 92

type AttrDeclarator struct {
	identifier ScopingInterfaces.IScopedName
}

func NewAttrDeclarator(identifier ScopingInterfaces.IScopedName) (*AttrDeclarator, error) {
	return &AttrDeclarator{
		identifier: identifier,
	}, nil
}

func (self *AttrDeclarator) Names() []string {
	result := make([]string, 0, 4)
	item := self.identifier
	for item != nil {
		result = append(result, item.Identifier())
		next, _ := item.GetNextScopedName()
		if next == nil {
			break
		}
		var ok bool
		item, ok = next.(ScopingInterfaces.IScopedName)
		if !ok {
			break
		}
	}
	return result
}

func (self *AttrDeclarator) GetFileName() string {
	return self.identifier.GetFileName()
}

func (self *AttrDeclarator) GetRow() int {
	return self.identifier.GetRow()
}

func (self *AttrDeclarator) GetCol() int {
	return self.identifier.GetCol()
}

func (self *AttrDeclarator) GetName() string {
	return strings.Join(self.Names(), ",")
}


func (self *AttrDeclarator) SetName(string) {
	panic("implement me")
}
