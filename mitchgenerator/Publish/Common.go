package Publish

import (
	"errors"
	"github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"
	"io"
)

type OutputType int

const (
	Unknown OutputType = iota
	Json
	Go
)

func ToOutputType(s string) OutputType {
	switch s {
	case "json":
		return Json
	case "go":
		return Go
	default:
		return Unknown
	}
}

type ExportParams struct {
	OutputStream  io.Writer
	PackageName   string
	DeclaredTypes []interfaces.IDefinitionDeclaration
}

type IPublish interface {
	Export(params ExportParams) error
}

var registrations map[OutputType]IPublish

func init() {
	registrations = make(map[OutputType]IPublish)
}

func Register(outputType OutputType, publish IPublish) {
	registrations[outputType] = publish
}

func HasOutputType(outputType OutputType) (IPublish, error) {
	result, ok := registrations[outputType]
	if ok {
		return result, nil
	}
	return nil, errors.New("Could not find publisher")
}

func PublishOutputType(
	outputType OutputType,
	writer io.Writer,
	packageName string,
	declaredTypes []interfaces.IDefinitionDeclaration) error {
	result, ok := registrations[outputType]
	if ok {
		err := result.Export(
			ExportParams{
				OutputStream:  writer,
				PackageName:   packageName,
				DeclaredTypes: declaredTypes})
		if err != nil {
			return err
		}
	}
	return nil
}
