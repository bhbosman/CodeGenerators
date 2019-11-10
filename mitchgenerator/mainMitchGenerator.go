package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/bhbosman/CodeGenerators/ctx"
	"github.com/bhbosman/CodeGenerators/mitchgenerator/Publish"
	_ "github.com/bhbosman/CodeGenerators/mitchgenerator/Publish/json"
	_ "github.com/bhbosman/CodeGenerators/mitchgenerator/Publish/publishGo"
	"github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"
	"github.com/bhbosman/CodeGenerators/mitchgenerator/yacc"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

//go:generate goyacc -o yacc/idl.go -p "IdlExpr" yacc/idl.y
//go:generate go clean
//go:generate go build
//go:generate go install
//go:generate mitchgenerator -outType go -packageName gogo test.idl

func main() {
	outputFile := flag.String("out", "", "outputFile")
	outputType := flag.String("outType", "json", "outputType")
	verbose := flag.Bool("verbose", false, "verbose")
	packageName := flag.String("packageName", "default", "packageName")
	defaultTypes := flag.Bool("defaultTypes", false, "InternalGenerate the built in types for the type in use")
	typesToUseAsString := flag.String("typesToUse", "IdlNative", `values "IdlNative" or "Mitch"`)
	showHelp := flag.Bool("help", false, "Show this")

	flag.Parse()
	if *showHelp {
		flag.PrintDefaults()

		fmt.Println(">>>>>>>>>>>>>")
		if *verbose {
			flag.VisitAll(func(i *flag.Flag) {
				fmt.Println(i.Name, i.Value)
			})
		}
		fmt.Println("<<<<<<<<<<<<<")

		os.Exit(1)
		return
	}

	tempFileName, err := func() (string, error) {
		outTempStream, err := GetTempOutput(*outputFile)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Output file could not be created.\n")
			os.Exit(2)
			return "", err
		}

		defer func() {
			_ = outTempStream.Close()
		}()

		if err != nil {
			_, _ = os.Stderr.WriteString(fmt.Sprintf("invalid type in use (%v).", typesToUseAsString))
			return "", fmt.Errorf("error 7")
		}

		var definitionDeclarations []interfaces.IDefinitionDeclaration = nil
		if !*defaultTypes {
			definitionDeclarations = ReadFromSource()

		}
		publisher, err := Publish.HasOutputType(Publish.ToOutputType(*outputType))
		if err != nil {
			_, _ = os.Stderr.Write([]byte(fmt.Sprintf("Error: %v\n", err.Error())))
			return "", fmt.Errorf("error 7")
		}

		err = publisher.Export(
			Publish.ExportParams{
				OutputStream:  outTempStream,
				PackageName:   *packageName,
				DeclaredTypes: definitionDeclarations})
		if err != nil {
			_, _ = os.Stderr.Write([]byte(fmt.Sprintf("Error: %v\n", err.Error())))
			return "", fmt.Errorf("error 7")
		}

		if f, ok := outTempStream.(*os.File); ok {
			return f.Name(), nil
		} else {
			return "", fmt.Errorf("no os.File as temp")
		}
	}()

	if err != nil {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("%s\n", err))
		return
	}
	_ = os.Rename(tempFileName, *outputFile)
}

func ReadFromSource() []interfaces.IDefinitionDeclaration {
	inStream, err := GetInput(flag.Args())
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "No input file. See help:\n")
		os.Exit(3)
		return nil
	}

	defer func() {
		for _, closer := range inStream {
			_ = closer.Close()
		}
	}()

	_, definitionDeclarations, err := func(inStream []io.ReadCloser) ([]int, []interfaces.IDefinitionDeclaration, error) {
		result := make([]int, len(inStream), len(inStream))

		definitionDeclarations := make([]interfaces.IDefinitionDeclaration, 0, 16)

		idlExprContext := ctx.NewIdlExprContext()
		for i, reader := range inStream {
			lex, err := yacc.NewIdlExprLex(
				bufio.NewReader(reader),
				idlExprContext)
			if err != nil {
				continue
			}
			resultInstance := yacc.IdlExprParse(lex)
			if resultInstance == 0 {
				for _, definitionDeclaration := range idlExprContext.GetSpecification() {
					definitionDeclarations = append(definitionDeclarations, definitionDeclaration)
				}
			}
			result[i] = resultInstance
		}

		return result, definitionDeclarations, nil

	}(inStream)
	if err != nil {
		os.Exit(4)
	}

	return definitionDeclarations
}

type WriterNoCloser struct {
	Writer io.Writer
}

func (self *WriterNoCloser) Close() error {
	return nil
}

func (self *WriterNoCloser) Write(p []byte) (n int, err error) {
	return self.Writer.Write(p)
}

func GetTempOutput(fileName string) (io.WriteCloser, error) {
	if fileName == "" {
		wc := &WriterNoCloser{
			Writer: os.Stdout,
		}
		return wc, nil
	}
	return ioutil.TempFile("", "tempfile")
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func GetInput(args []string) ([]io.ReadCloser, error) {
	if len(args) > 0 {
		list := make([]io.ReadCloser, 0, len(args))
		for _, item := range args {
			if Exists(item) {
				file, err := os.Open(item)
				if err != nil {
					return nil, err
				}
				list = append(list, file)
				continue
			} else {
				readerCloser := ioutil.NopCloser(strings.NewReader(item))
				list = append(list, readerCloser)
			}
		}
		return list, nil
	}
	return nil, errors.New("no input files or strings")
}
