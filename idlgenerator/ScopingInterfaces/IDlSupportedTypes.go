package ScopingInterfaces

import "strconv"

type IDlSupportedTypes byte

const (
	Invalid IDlSupportedTypes = iota
	Unknown
	DeclareTypePlaceHolderType
	TypeDeclaratorIdlType
	Idlvalue_forward_dclType
	TypePrefixDefinitionIdlType
	ModuleIdlType
	FixedPointTypeIdlType
	StructIdlType
	InterfaceIdlType
	ConstDclType
	SequenceIdlType

	Op_dclIdlType
	Attr_specIdlType
	IdlValue_Abs_DefType

	ExceptionIdlType
	// Begin of primitive types
	RWfloatIdlType
	StringIdlType
	WideStringIdlType
	RWVoidType
	RWvaluebaseIdlType
	RWdoubleIdlType
	RWlongRWdoubleIdlType
	RWshortIdlType
	RWlongIdlType
	RWint8IdlType
	RWlongRWlongIdlType
	RWuint8IdlType
	RWunsignedRWshortIdlType
	RWunsignedRWlongIdlType
	RWunsignedRWlongRWlongIdlType
	RWcharIdlType
	RWwcharIdlType
	RWbooleanIdlType
	RWoctetIdlType
	RWanyIdlType
	RWObjectIdlType
	RWEnumIdlType
	RWint16IdlType
	RWint32IdlType
	RWint64IdlType
	RWuint16IdlType
	RWuint32IdlType
	RWuint64IdlType
	// End of primitive types
	PrimitiveTypesBegin = RWfloatIdlType
	PrimitiveTypesEnd   = RWuint64IdlType
)

var IDlSupportedTypesNames = []string{
	Invalid:                       "Invalid",
	StructIdlType:                 "Struct",
	RWfloatIdlType:                "RWFloat",
	RWdoubleIdlType:               "double",
	RWlongRWdoubleIdlType:         "RWLongRWDouble",
	RWshortIdlType:                "RWShort",
	RWlongIdlType:                 "long",
	RWint8IdlType:                 "RWInt8",
	RWlongRWlongIdlType:           "RWLongRWLong",
	RWuint8IdlType:                "RWUint8",
	RWunsignedRWshortIdlType:      "RWUnsignedRWShort",
	RWunsignedRWlongIdlType:       "RWUnsignedRWLong",
	RWunsignedRWlongRWlongIdlType: "RWUnsignedRWLongRWLong",
	RWcharIdlType:                 "RWChar",
	RWwcharIdlType:                "RWWchar",
	RWbooleanIdlType:              "RWBoolean",
	RWoctetIdlType:                "RWOctet",
	RWanyIdlType:                  "RWAny",
	RWObjectIdlType:               "Object",
	RWEnumIdlType:                 "RWEnum",
	RWint16IdlType:                "int16",
	RWint32IdlType:                "int32",
	RWint64IdlType:                "int64",
	RWuint16IdlType:               "uint16",
	RWuint32IdlType:               "uint32",
	RWuint64IdlType:               "uint64",
	RWVoidType:                    "void",
	StringIdlType:                 "string",
	WideStringIdlType:             "wstring",
}

func (k IDlSupportedTypes) String() string {
	if int(k) < len(IDlSupportedTypesNames) {
		return IDlSupportedTypesNames[k]
	}
	return "kind" + strconv.Itoa(int(k))
}

var IDlTokens = []string{
	RWfloatIdlType:                "float",
	RWdoubleIdlType:               "double",
	RWlongRWdoubleIdlType:         "long double",
	RWshortIdlType:                "short",
	RWlongIdlType:                 "long",
	RWint8IdlType:                 "int8",
	RWlongRWlongIdlType:           "long long",
	RWuint8IdlType:                "uint8",
	RWunsignedRWshortIdlType:      "unsigned short",
	RWunsignedRWlongIdlType:       "unsigned long",
	RWunsignedRWlongRWlongIdlType: "unsigned long long",
	RWcharIdlType:                 "char",
	RWwcharIdlType:                "wchar",
	RWbooleanIdlType:              "boolean",
	RWoctetIdlType:                "octet",
	RWanyIdlType:                  "any",
	RWObjectIdlType:               "object",
	RWint16IdlType:                "int16",
	RWint32IdlType:                "int32",
	RWint64IdlType:                "int64",
	RWuint16IdlType:               "uint16",
	RWuint32IdlType:               "uint32",
	RWuint64IdlType:               "uint64",
	RWVoidType:                    "void",
	RWvaluebaseIdlType:            "valuebase",
}

func (k IDlSupportedTypes) IDLToken() string {
	if int(k) < len(IDlTokens) {
		return IDlTokens[k]
	}
	return "kind" + strconv.Itoa(int(k))
}
