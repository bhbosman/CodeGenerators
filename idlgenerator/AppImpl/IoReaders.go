package AppImpl

import (
	"errors"
	"github.com/bhbosman/CodeGenerators/idlgenerator/AppInterfaces"
	"go.uber.org/multierr"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type IoReaders struct {
	FileInformation []AppInterfaces.IFileInformation
	Context         AppInterfaces.IIdlGeneratorFlags
	Logger          *log.Logger
}

func (self *IoReaders) GetFileInformation() []AppInterfaces.IFileInformation {
	return self.FileInformation
}

func (self *IoReaders) Close() error {
	var resultError error = nil
	for _, fileInformation := range self.FileInformation {
		self.Logger.Println("closing reader...")
		resultError = multierr.Append(resultError, fileInformation.Close())
	}
	return resultError
}

type fileInformation struct {
	readCloser io.ReadCloser
	arg        string
	fileName   string
	actualFile bool
}

func newFileInformation(actualFile bool, readCloser io.ReadCloser, arg string, fileName string) *fileInformation {
	return &fileInformation{
		readCloser: readCloser,
		arg:        arg,
		fileName:   fileName,
		actualFile: actualFile,
	}
}

func (self *fileInformation) GetReader() io.Reader {
	return self.readCloser
}

func (self *fileInformation) GetFileName() string {
	return self.fileName
}

func (self *fileInformation) GetArg() string {
	return self.arg
}

func (self *fileInformation) Close() error {
	return self.readCloser.Close()
}

func (self *IoReaders) Start() error {
	self.Logger.Println("Read files...")
	if len(self.Context.Files()) > 0 {
		self.FileInformation = make([]AppInterfaces.IFileInformation, 0, len(self.Context.Files()))
		for _, item := range self.Context.Files() {
			self.Logger.Printf("Opening item %v", item)
			if self.Exists(item) {
				self.Logger.Printf("Item %v exist", item)
				file, err := os.Open(item)
				if err != nil {
					continue
				}
				stat, _ := os.Stat(item)
				self.FileInformation = append(self.FileInformation, newFileInformation(true, file, item, stat.Name()))
			} else {
				self.Logger.Printf("Create wrapper for text %v", item)
				readerCloser := ioutil.NopCloser(strings.NewReader(item))
				self.FileInformation = append(self.FileInformation, newFileInformation(false, readerCloser, "", item))
			}
		}
		return nil
	}
	self.Logger.Println("No files processed")
	return errors.New("no input files or strings")
}

func (self *IoReaders) Stop() error {
	return self.Close()
}

func (self *IoReaders) Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
