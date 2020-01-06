// Code generated by abc. DO NOT EDIT.

package corbaFiles

//line DynamicAny.idl:21
//Exception Decl: DynamicAny_DynAny_InvalidValueException
//Usage Count: 0
type DynamicAny_DynAny_InvalidValueException struct {
}

//Constructors
func NewDynamicAny_DynAny_InvalidValueExceptionDefaultPointer() (*DynamicAny_DynAny_InvalidValueException, error) {
	return &DynamicAny_DynAny_InvalidValueException{}, nil
}

func NewDynamicAny_DynAny_InvalidValueExceptionDefaultValue() (DynamicAny_DynAny_InvalidValueException, error) {
	return DynamicAny_DynAny_InvalidValueException{}, nil
}

//line DynamicAny.idl:22
//Exception Decl: DynamicAny_DynAny_TypeMismatchException
//Usage Count: 0
type DynamicAny_DynAny_TypeMismatchException struct {
}

//Constructors
func NewDynamicAny_DynAny_TypeMismatchExceptionDefaultPointer() (*DynamicAny_DynAny_TypeMismatchException, error) {
	return &DynamicAny_DynAny_TypeMismatchException{}, nil
}

func NewDynamicAny_DynAny_TypeMismatchExceptionDefaultValue() (DynamicAny_DynAny_TypeMismatchException, error) {
	return DynamicAny_DynAny_TypeMismatchException{}, nil
}

//line DynamicAny.idl:20
//Interface Decl: DynamicAny_DynAny
//Usage Count: 15
type DynamicAny_DynAny interface {

	//line DynamicAny.idl:24
	Type() (CORBA_TypeCode, error)

	//line DynamicAny.idl:26
	Assign(dyn_any DynamicAny_DynAny) error

	//line DynamicAny.idl:28
	From_any(value CorbaAny) error

	//line DynamicAny.idl:30
	To_any() (CorbaAny, error)

	//line DynamicAny.idl:32
	Equal(dyn_any DynamicAny_DynAny) (bool, error)

	//line DynamicAny.idl:34
	Destroy() error

	//line DynamicAny.idl:35
	Copy() (DynamicAny_DynAny, error)

	//line DynamicAny.idl:37
	Insert_boolean(value bool) error

	//line DynamicAny.idl:39
	Insert_octet(value int) error

	//line DynamicAny.idl:41
	Insert_char(value AnsiChar) error

	//line DynamicAny.idl:43
	Insert_short(value int16) error

	//line DynamicAny.idl:45
	Insert_ushort(value uint16) error

	//line DynamicAny.idl:48
	Insert_long(value int32) error

	//line DynamicAny.idl:50
	Insert_ulong(value uint32) error

	//line DynamicAny.idl:53
	Insert_float(value CorbaFloat) error

	//line DynamicAny.idl:55
	Insert_double(value CorbaDouble) error

	//line DynamicAny.idl:57
	Insert_string(value AnsiString) error

	//line DynamicAny.idl:59
	Insert_reference(value CorbaObject) error

	//line DynamicAny.idl:61
	Insert_typecode(value CORBA_TypeCode) error

	//line DynamicAny.idl:64
	Insert_longlong(value int64) error

	//line DynamicAny.idl:66
	Insert_ulonglong(value uint64) error

	//line DynamicAny.idl:69
	Insert_longdouble(value CorbaLongDouble) error

	//line DynamicAny.idl:72
	Insert_wchar(value WideChar) error

	//line DynamicAny.idl:74
	Insert_wstring(value WideString) error

	//line DynamicAny.idl:76
	Insert_any(value CorbaAny) error

	//line DynamicAny.idl:78
	Insert_dyn_any(value DynamicAny_DynAny) error

	//line DynamicAny.idl:80
	Insert_val(value Valuebase) error

	//line DynamicAny.idl:83
	Get_boolean() (bool, error)

	//line DynamicAny.idl:85
	Get_octet() (int, error)

	//line DynamicAny.idl:87
	Get_char() (AnsiChar, error)

	//line DynamicAny.idl:89
	Get_short() (int16, error)

	//line DynamicAny.idl:91
	Get_ushort() (uint16, error)

	//line DynamicAny.idl:93
	Get_long() (int32, error)

	//line DynamicAny.idl:95
	Get_ulong() (uint32, error)

	//line DynamicAny.idl:97
	Get_float() (CorbaFloat, error)

	//line DynamicAny.idl:99
	Get_double() (CorbaDouble, error)

	//line DynamicAny.idl:101
	Get_string() (AnsiString, error)

	//line DynamicAny.idl:103
	Get_reference() (CorbaObject, error)

	//line DynamicAny.idl:105
	Get_typecode() (CORBA_TypeCode, error)

	//line DynamicAny.idl:107
	Get_longlong() (int64, error)

	//line DynamicAny.idl:109
	Get_ulonglong() (uint64, error)

	//line DynamicAny.idl:111
	Get_longdouble() (CorbaLongDouble, error)

	//line DynamicAny.idl:113
	Get_wchar() (WideChar, error)

	//line DynamicAny.idl:115
	Get_wstring() (WideString, error)

	//line DynamicAny.idl:117
	Get_any() (CorbaAny, error)

	//line DynamicAny.idl:119
	Get_dyn_any() (DynamicAny_DynAny, error)

	//line DynamicAny.idl:121
	Get_val() (Valuebase, error)

	//line DynamicAny.idl:124
	Seek(index int32) (bool, error)

	//line DynamicAny.idl:125
	Rewind() error

	//line DynamicAny.idl:126
	Next() (bool, error)

	//line DynamicAny.idl:127
	Component_count() (uint32, error)

	//line DynamicAny.idl:128
	Current_component() (DynamicAny_DynAny, error)

	//line DynamicAny.idl:131
	Insert_abstract(value CORBA_AbstractBase) error

	//line DynamicAny.idl:133
	Get_abstract() (CORBA_AbstractBase, error)

	//line DynamicAny.idl:136
	Insert_boolean_seq(value CORBA_BooleanSeq) error

	//line DynamicAny.idl:138
	Insert_octet_seq(value CORBA_OctetSeq) error

	//line DynamicAny.idl:140
	Insert_char_seq(value CORBA_CharSeq) error

	//line DynamicAny.idl:142
	Insert_short_seq(value CORBA_ShortSeq) error

	//line DynamicAny.idl:144
	Insert_ushort_seq(value CORBA_UShortSeq) error

	//line DynamicAny.idl:146
	Insert_long_seq(value CORBA_LongSeq) error

	//line DynamicAny.idl:148
	Insert_ulong_seq(value CORBA_ULongSeq) error

	//line DynamicAny.idl:150
	Insert_float_seq(value CORBA_FloatSeq) error

	//line DynamicAny.idl:152
	Insert_double_seq(value CORBA_DoubleSeq) error

	//line DynamicAny.idl:154
	Insert_longlong_seq(value CORBA_LongLongSeq) error

	//line DynamicAny.idl:156
	Insert_ulonglong_seq(value CORBA_ULongLongSeq) error

	//line DynamicAny.idl:158
	Insert_longdouble_seq(value CORBA_LongDoubleSeq) error

	//line DynamicAny.idl:160
	Insert_wchar_seq(value CORBA_WCharSeq) error

	//line DynamicAny.idl:162
	Get_boolean_seq() (CORBA_BooleanSeq, error)

	//line DynamicAny.idl:164
	Get_octet_seq() (CORBA_OctetSeq, error)

	//line DynamicAny.idl:166
	Get_char_seq() (CORBA_CharSeq, error)

	//line DynamicAny.idl:168
	Get_short_seq() (CORBA_ShortSeq, error)

	//line DynamicAny.idl:170
	Get_ushort_seq() (CORBA_UShortSeq, error)

	//line DynamicAny.idl:172
	Get_long_seq() (CORBA_LongSeq, error)

	//line DynamicAny.idl:174
	Get_ulong_seq() (CORBA_ULongSeq, error)

	//line DynamicAny.idl:176
	Get_float_seq() (CORBA_FloatSeq, error)

	//line DynamicAny.idl:178
	Get_double_seq() (CORBA_DoubleSeq, error)

	//line DynamicAny.idl:180
	Get_longlong_seq() (CORBA_LongLongSeq, error)

	//line DynamicAny.idl:182
	Get_ulonglong_seq() (CORBA_ULongLongSeq, error)

	//line DynamicAny.idl:184
	Get_longdouble_seq() (CORBA_LongDoubleSeq, error)

	//line DynamicAny.idl:186
	Get_wchar_seq() (CORBA_WCharSeq, error)
}

//line DynamicAny.idl:190
//Interface Decl: DynamicAny_DynFixed
//Usage Count: 0
type DynamicAny_DynFixed interface {

	//line DynamicAny.idl:191
	Get_value() (AnsiString, error)

	//line DynamicAny.idl:192
	Set_value(val AnsiString) (bool, error)
}

//line DynamicAny.idl:196
//Interface Decl: DynamicAny_DynEnum
//Usage Count: 0
type DynamicAny_DynEnum interface {

	//line DynamicAny.idl:197
	Get_as_string() (AnsiString, error)

	//line DynamicAny.idl:198
	Set_as_string(value AnsiString) error

	//line DynamicAny.idl:200
	Get_as_ulong() (uint32, error)

	//line DynamicAny.idl:201
	Set_as_ulong(value uint32) error
}

//line DynamicAny.idl:204
type DynamicAny_FieldName AnsiString

//line DynamicAny.idl:205
//Struct Decl: DynamicAny_NameValuePair
//Usage Count: 0
type DynamicAny_NameValuePair struct {
	Id    DynamicAny_FieldName
	Value CorbaAny
}

//Constructors
func NewDynamicAny_NameValuePairDefaultPointer() (*DynamicAny_NameValuePair, error) {
	return &DynamicAny_NameValuePair{}, nil
}

func NewDynamicAny_NameValuePairDefaultValue() (DynamicAny_NameValuePair, error) {
	return DynamicAny_NameValuePair{}, nil
}

func NewDynamicAny_NameValuePairValue(
	id DynamicAny_FieldName,
	value CorbaAny) (DynamicAny_NameValuePair, error) {
	return DynamicAny_NameValuePair{
		Id:    id,
		Value: value,
	}, nil
}

//line DynamicAny.idl:210
//Sequence Decl: Sequence_NameValuePair
type Sequence_NameValuePair interface{}

//line DynamicAny.idl:210
//Typedef Decl: DynamicAny_NameValuePairSeq
type DynamicAny_NameValuePairSeq Sequence_NameValuePair

//line DynamicAny.idl:211
//Struct Decl: DynamicAny_NameDynAnyPair
//Usage Count: 0
type DynamicAny_NameDynAnyPair struct {
	Id    DynamicAny_FieldName
	Value DynamicAny_DynAny
}

//Constructors
func NewDynamicAny_NameDynAnyPairDefaultPointer() (*DynamicAny_NameDynAnyPair, error) {
	return &DynamicAny_NameDynAnyPair{}, nil
}

func NewDynamicAny_NameDynAnyPairDefaultValue() (DynamicAny_NameDynAnyPair, error) {
	return DynamicAny_NameDynAnyPair{}, nil
}

func NewDynamicAny_NameDynAnyPairValue(
	id DynamicAny_FieldName,
	value DynamicAny_DynAny) (DynamicAny_NameDynAnyPair, error) {
	return DynamicAny_NameDynAnyPair{
		Id:    id,
		Value: value,
	}, nil
}

//line DynamicAny.idl:216
//Sequence Decl: Sequence_NameDynAnyPair
type Sequence_NameDynAnyPair interface{}

//line DynamicAny.idl:216
//Typedef Decl: DynamicAny_NameDynAnyPairSeq
type DynamicAny_NameDynAnyPairSeq Sequence_NameDynAnyPair

//line DynamicAny.idl:217
//Interface Decl: DynamicAny_DynStruct
//Usage Count: 0
type DynamicAny_DynStruct interface {

	//line DynamicAny.idl:218
	Current_member_name() (DynamicAny_FieldName, error)

	//line DynamicAny.idl:220
	Current_member_kind() (CORBA_TCKind, error)

	//line DynamicAny.idl:222
	Get_members() (DynamicAny_NameValuePairSeq, error)

	//line DynamicAny.idl:223
	Set_members(value DynamicAny_NameValuePairSeq) error

	//line DynamicAny.idl:225
	Get_members_as_dyn_any() (DynamicAny_NameDynAnyPairSeq, error)

	//line DynamicAny.idl:226
	Set_members_as_dyn_any(value DynamicAny_NameDynAnyPairSeq) error
}

//line DynamicAny.idl:230
//Interface Decl: DynamicAny_DynUnion
//Usage Count: 0
type DynamicAny_DynUnion interface {

	//line DynamicAny.idl:231
	Get_discriminator() (DynamicAny_DynAny, error)

	//line DynamicAny.idl:232
	Set_discriminator(d DynamicAny_DynAny) error

	//line DynamicAny.idl:234
	Set_to_default_member() error

	//line DynamicAny.idl:236
	Set_to_no_active_member() error

	//line DynamicAny.idl:238
	Has_no_active_member() (bool, error)

	//line DynamicAny.idl:239
	Discriminator_kind() (CORBA_TCKind, error)

	//line DynamicAny.idl:240
	Member() (DynamicAny_DynAny, error)

	//line DynamicAny.idl:242
	Member_name() (DynamicAny_FieldName, error)

	//line DynamicAny.idl:244
	Member_kind() (CORBA_TCKind, error)
}

//line DynamicAny.idl:248
//Sequence Decl: Sequence_CorbaAny
type Sequence_CorbaAny interface{}

//line DynamicAny.idl:248
//Typedef Decl: DynamicAny_AnySeq
type DynamicAny_AnySeq Sequence_CorbaAny

//line DynamicAny.idl:249
//Sequence Decl: Sequence_DynAny
type Sequence_DynAny interface{}

//line DynamicAny.idl:249
//Typedef Decl: DynamicAny_DynAnySeq
type DynamicAny_DynAnySeq Sequence_DynAny

//line DynamicAny.idl:250
//Interface Decl: DynamicAny_DynSequence
//Usage Count: 0
type DynamicAny_DynSequence interface {

	//line DynamicAny.idl:251
	Get_length() (uint32, error)

	//line DynamicAny.idl:252
	Set_length(len uint32) error

	//line DynamicAny.idl:254
	Get_elements() (DynamicAny_AnySeq, error)

	//line DynamicAny.idl:255
	Set_elements(value DynamicAny_AnySeq) error

	//line DynamicAny.idl:257
	Get_elements_as_dyn_any() (DynamicAny_DynAnySeq, error)

	//line DynamicAny.idl:258
	Set_elements_as_dyn_any(value DynamicAny_DynAnySeq) error
}

//line DynamicAny.idl:262
//Interface Decl: DynamicAny_DynArray
//Usage Count: 0
type DynamicAny_DynArray interface {

	//line DynamicAny.idl:263
	Get_elements() (DynamicAny_AnySeq, error)

	//line DynamicAny.idl:264
	Set_elements(value DynamicAny_AnySeq) error

	//line DynamicAny.idl:266
	Get_elements_as_dyn_any() (DynamicAny_DynAnySeq, error)

	//line DynamicAny.idl:267
	Set_elements_as_dyn_any(value DynamicAny_DynAnySeq) error
}

//line DynamicAny.idl:271
//Interface Decl: DynamicAny_DynValueCommon
//Usage Count: 0
type DynamicAny_DynValueCommon interface {

	//line DynamicAny.idl:272
	Is_null() (bool, error)

	//line DynamicAny.idl:273
	Set_to_null() error

	//line DynamicAny.idl:274
	Set_to_value() error
}

//line DynamicAny.idl:277
//Interface Decl: DynamicAny_DynValue
//Usage Count: 0
type DynamicAny_DynValue interface {

	//line DynamicAny.idl:278
	Current_member_name() (DynamicAny_FieldName, error)

	//line DynamicAny.idl:280
	Current_member_kind() (CORBA_TCKind, error)

	//line DynamicAny.idl:282
	Get_members() (DynamicAny_NameValuePairSeq, error)

	//line DynamicAny.idl:284
	Set_members(value DynamicAny_NameValuePairSeq) error

	//line DynamicAny.idl:286
	Get_members_as_dyn_any() (DynamicAny_NameDynAnyPairSeq, error)

	//line DynamicAny.idl:288
	Set_members_as_dyn_any(value DynamicAny_NameDynAnyPairSeq) error
}

//line DynamicAny.idl:292
//Interface Decl: DynamicAny_DynValueBox
//Usage Count: 0
type DynamicAny_DynValueBox interface {

	//line DynamicAny.idl:293
	Get_boxed_value() (CorbaAny, error)

	//line DynamicAny.idl:295
	Set_boxed_value(boxed CorbaAny) error

	//line DynamicAny.idl:297
	Get_boxed_value_as_dyn_any() (DynamicAny_DynAny, error)

	//line DynamicAny.idl:299
	Set_boxed_value_as_dyn_any(boxed DynamicAny_DynAny) error
}

//line DynamicAny.idl:303
//Exception Decl: DynamicAny_MustTruncateException
//Usage Count: 0
type DynamicAny_MustTruncateException struct {
}

//Constructors
func NewDynamicAny_MustTruncateExceptionDefaultPointer() (*DynamicAny_MustTruncateException, error) {
	return &DynamicAny_MustTruncateException{}, nil
}

func NewDynamicAny_MustTruncateExceptionDefaultValue() (DynamicAny_MustTruncateException, error) {
	return DynamicAny_MustTruncateException{}, nil
}

//line DynamicAny.idl:306
//Exception Decl: DynamicAny_DynAnyFactory_InconsistentTypeCodeException
//Usage Count: 0
type DynamicAny_DynAnyFactory_InconsistentTypeCodeException struct {
}

//Constructors
func NewDynamicAny_DynAnyFactory_InconsistentTypeCodeExceptionDefaultPointer() (*DynamicAny_DynAnyFactory_InconsistentTypeCodeException, error) {
	return &DynamicAny_DynAnyFactory_InconsistentTypeCodeException{}, nil
}

func NewDynamicAny_DynAnyFactory_InconsistentTypeCodeExceptionDefaultValue() (DynamicAny_DynAnyFactory_InconsistentTypeCodeException, error) {
	return DynamicAny_DynAnyFactory_InconsistentTypeCodeException{}, nil
}

//line DynamicAny.idl:305
//Interface Decl: DynamicAny_DynAnyFactory
//Usage Count: 0
type DynamicAny_DynAnyFactory interface {

	//line DynamicAny.idl:307
	Create_dyn_any(value CorbaAny) (DynamicAny_DynAny, error)

	//line DynamicAny.idl:309
	Create_dyn_any_from_type_code(type_ CORBA_TypeCode) (DynamicAny_DynAny, error)

	//line DynamicAny.idl:311
	Create_dyn_any_without_truncation(value CorbaAny) (DynamicAny_DynAny, error)

	//line DynamicAny.idl:313
	Create_multiple_dyn_anys(values DynamicAny_AnySeq, allow_truncate bool) (DynamicAny_DynAnySeq, error)

	//line DynamicAny.idl:317
	Create_multiple_anys(values DynamicAny_DynAnySeq) (DynamicAny_AnySeq, error)
}
