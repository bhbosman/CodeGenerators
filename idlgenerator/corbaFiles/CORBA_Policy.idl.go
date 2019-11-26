package corbaFiles

//Code generated by  DO NOT EDIT.
//line CORBA_Policy.idl:7
//Interface Decl: Policy
//Usage Count: 0
type Policy interface {
	//line CORBA_Policy.idl:8
	//line CORBA_Policy.idl:9
	Copy() (Policy, error)
	//line CORBA_Policy.idl:10
	Destroy() (void, error)
}
//line CORBA_Policy.idl:17
//Exception Decl: CORBA_PolicyError
//Usage Count: 0
type CORBA_PolicyError struct {
	Reason PolicyErrorCode
}

//Constructors
func NewCORBA_PolicyErrorDefaultPointer() (*CORBA_PolicyError, error) {
	return &CORBA_PolicyError {
		Reason: place holder,
	}, nil
}

func NewCORBA_PolicyErrorDefaultValue() (CORBA_PolicyError, error) {
	return CORBA_PolicyError {
		Reason: place holder,
	}, nil
}

func NewCORBA_PolicyErrorValue(
	reason PolicyErrorCode) (CORBA_PolicyError, error) {
	return CORBA_PolicyError {
		Reason: reason,
	}, nil
}
//line CORBA_Policy.idl:25
//Usage Count: 0
type CORBA_SetOverrideType int
//noinspection ALL
const (
	CORBA_SetOverrideType_SET_OVERRIDE = iota
	CORBA_SetOverrideType_ADD_OVERRIDE
)
//line CORBA_Policy.idl:27
//Exception Decl: CORBA_InvalidPolicies
//Usage Count: 0
type CORBA_InvalidPolicies struct {
	Indicies Sequence_unsigned short_0
}

//Constructors
func NewCORBA_InvalidPoliciesDefaultPointer() (*CORBA_InvalidPolicies, error) {
	return &CORBA_InvalidPolicies {
		Indicies: need default value for ''
,
	}, nil
}

func NewCORBA_InvalidPoliciesDefaultValue() (CORBA_InvalidPolicies, error) {
	return CORBA_InvalidPolicies {
		Indicies: need default value for ''
,
	}, nil
}

func NewCORBA_InvalidPoliciesValue(
	indicies Sequence_unsigned short_0) (CORBA_InvalidPolicies, error) {
	return CORBA_InvalidPolicies {
		Indicies: indicies,
	}, nil
}
//line CORBA_Policy.idl:31
//Interface Decl: CORBA_PolicyManager
//Usage Count: 0
type CORBA_PolicyManager interface {
	//line CORBA_Policy.idl:32
	Get_policy_overrides(ts PolicyTypeSeq) (PolicyList, error)
	//line CORBA_Policy.idl:34
	Set_policy_overrides(policies PolicyList,set_add SetOverrideType) (void, error)
}
//line CORBA_Policy.idl:39
//Interface Decl: CORBA_PolicyCurrent
//Usage Count: 0
type CORBA_PolicyCurrent interface {
}
