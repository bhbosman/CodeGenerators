package Common

type IErrorList interface {
	Add(err error)
	Error() error
}
