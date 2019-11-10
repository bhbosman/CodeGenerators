package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	"path"
)

type FileInformationBase struct {
	FileName string
	Row      int
	Col      int
}

func (self *FileInformationBase) String() string {
	_, fileName := path.Split(self.FileName)
	return fmt.Sprintf("(%v: %v, %v)", fileName, self.Row, self.Col)

}

func NewFileInformationBase01(fileName string, row int, col int) FileInformationBase {
	return FileInformationBase{
		FileName: fileName,
		Row:      row,
		Col:      col,
	}
}

func NewFileInformationBase02(fileInformation ScopingInterfaces.IFileInformation) FileInformationBase {
	if fileInformation == nil {
		return NewFileInformationBase01("", 0, 0)
	}
	return NewFileInformationBase01(
		fileInformation.GetFileName(),
		fileInformation.GetRow(),
		fileInformation.GetCol())
}

func (self *FileInformationBase) GetFileName() string {
	return self.FileName
}

func (self *FileInformationBase) GetRow() int {
	return self.Row
}

func (self *FileInformationBase) GetCol() int {
	return self.Col
}
