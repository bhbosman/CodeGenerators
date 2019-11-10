package scopedObjects

import (
	"fmt"
	"path"
)

type ReservedWordData struct {
	FileInformationBase
	identifier string
}

func (self *ReservedWordData) String() string {
	_, fileName := path.Split(self.FileName)
	return fmt.Sprintf("ReservedWordData: %v (%v: %v, %v)", self.identifier, fileName, self.Row, self.Col)
}

func NewReservedWordData(identifier string, fileName string, row int, col int) *ReservedWordData {
	return &ReservedWordData{
		FileInformationBase: NewFileInformationBase01(fileName, row, col),
		identifier:          identifier,
	}
}
