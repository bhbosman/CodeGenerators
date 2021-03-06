// Code generated by abc. DO NOT EDIT.

package corbaFiles

//line CORBA_Pollable.idl:8
//Interface Decl: CORBA_Pollable
//Usage Count: 3
type CORBA_Pollable interface {

	//line CORBA_Pollable.idl:9
	Is_ready(timeout uint32) (bool, error)

	//line CORBA_Pollable.idl:13
	Create_pollable_set() (CORBA_PollableSet, error)
}

//line CORBA_Pollable.idl:16
//Interface Decl: CORBA_DIIPollable
//Usage Count: 1
type CORBA_DIIPollable interface {
}

//line CORBA_Pollable.idl:20
//Exception Decl: CORBA_PollableSet_NoPossiblePollableException
//Usage Count: 0
type CORBA_PollableSet_NoPossiblePollableException struct {
}

//Constructors
func NewCORBA_PollableSet_NoPossiblePollableExceptionDefaultPointer() (*CORBA_PollableSet_NoPossiblePollableException, error) {
	return &CORBA_PollableSet_NoPossiblePollableException{}, nil
}

func NewCORBA_PollableSet_NoPossiblePollableExceptionDefaultValue() (CORBA_PollableSet_NoPossiblePollableException, error) {
	return CORBA_PollableSet_NoPossiblePollableException{}, nil
}

//line CORBA_Pollable.idl:21
//Exception Decl: CORBA_PollableSet_UnknownPollableException
//Usage Count: 0
type CORBA_PollableSet_UnknownPollableException struct {
}

//Constructors
func NewCORBA_PollableSet_UnknownPollableExceptionDefaultPointer() (*CORBA_PollableSet_UnknownPollableException, error) {
	return &CORBA_PollableSet_UnknownPollableException{}, nil
}

func NewCORBA_PollableSet_UnknownPollableExceptionDefaultValue() (CORBA_PollableSet_UnknownPollableException, error) {
	return CORBA_PollableSet_UnknownPollableException{}, nil
}

//line CORBA_Pollable.idl:18
//Interface Decl: CORBA_PollableSet
//Usage Count: 1
type CORBA_PollableSet interface {

	//line CORBA_Pollable.idl:23
	Create_dii_pollable() (CORBA_DIIPollable, error)

	//line CORBA_Pollable.idl:25
	Add_pollable(potential CORBA_Pollable) error

	//line CORBA_Pollable.idl:29
	Get_ready_pollable(timeout uint32) (CORBA_Pollable, error)

	//line CORBA_Pollable.idl:33
	Remove(potential CORBA_Pollable) error

	//line CORBA_Pollable.idl:37
	Number_left() (uint16, error)
}
