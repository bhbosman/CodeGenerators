package interfaces

import "strconv"

type Kind byte

const (
	Invalid Kind = iota
	Struct
	Int64
	Char
	Enum
	//TypeDeclarator
	MitchAlpha
	MitchBitField
	MitchByte
	MitchDate
	MitchTime
	MitchPrice04
	MitchPrice08
	MitchUInt08
	MitchUInt16
	MitchUInt32
	MitchUInt64
	//MitchMessageNumber
	//MitchMessageLength
)

// String returns the name of k.
func (k Kind) String() string {
	if int(k) < len(kindNames) {
		return kindNames[k]
	}
	return "kind" + strconv.Itoa(int(k))
}

var kindNames = []string{
	Invalid:       "invalid",
	Struct:        "Struct",
	Int64:         "Int64",
	Char:          "Char",
	Enum:          "Enum",
	MitchAlpha:    "string",
	MitchBitField: "MitchBitField",
	MitchByte:     "byte",
	MitchDate:     "Time",
	MitchTime:     "Time",
	MitchPrice04:  "float64",
	MitchPrice08:  "float64",
	MitchUInt08:   "uint8",
	MitchUInt16:   "uint16",
	MitchUInt32:   "uint32",
	MitchUInt64:   "uint64",
	//MitchMessageNumber: "byte",
	//MitchMessageLength: "uint16",
}
