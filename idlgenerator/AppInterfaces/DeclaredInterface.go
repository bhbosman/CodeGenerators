package AppInterfaces

import (
	si "github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	"github.com/bhbosman/CodeGenerators/idlgenerator/yacc"
	"io"
)

type IProcessInformation interface {
	GetFileName() string
	GetArg() string
}

type IFileInformation interface {
	io.Closer
	IProcessInformation
	GetReader() io.Reader
}

type IIoReaders interface {
	Start() error
	Stop() error
	GetFileInformation() []IFileInformation
}

type IDefinitionContextFactory interface {
	Create() yacc.IDefinitionContext
}

type IProcessor interface {
	Process() error
}

type IScopeWalker interface {
	Scope(scopingContext si.IScopingContext, indent int, declaredType si.ITypeSpec, fileName string) error
}

type IIdlGeneratorFlags interface {
	Files() []string
}

type ISetIdlGeneratorFlags interface {
	SetFiles(files []string) error
}

type ICodeGenerator interface {
	io.Closer
	Generate(dcl si.ITypeSpec) error
}
