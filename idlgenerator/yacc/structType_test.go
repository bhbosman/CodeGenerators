package yacc

import (
	"bufio"
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	"github.com/bhbosman/CodeGenerators/idlgenerator/scopedObjects"
	"github.com/bhbosman/CodeGenerators/idlgenerator/scoping"

	"github.com/bhbosman/gomock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strings"
	"testing"
)

func TestIfPrimitivesTypesPassParsing(t *testing.T) {
	m := make(map[ScopingInterfaces.IDlSupportedTypes]string)
	m[ScopingInterfaces.RWfloatIdlType] = "float"
	m[ScopingInterfaces.RWdoubleIdlType] = "double"
	m[ScopingInterfaces.RWlongRWdoubleIdlType] = "long double"
	m[ScopingInterfaces.RWshortIdlType] = "short"
	m[ScopingInterfaces.RWlongIdlType] = "long"
	m[ScopingInterfaces.RWint8IdlType] = "int8"
	m[ScopingInterfaces.RWlongRWlongIdlType] = "long long"
	m[ScopingInterfaces.RWuint8IdlType] = "uint8"
	m[ScopingInterfaces.RWbooleanIdlType] = "boolean"
	m[ScopingInterfaces.RWunsignedRWshortIdlType] = "unsigned short"
	m[ScopingInterfaces.RWunsignedRWlongIdlType] = "unsigned long"
	m[ScopingInterfaces.RWunsignedRWlongRWlongIdlType] = "unsigned long long"
	m[ScopingInterfaces.RWcharIdlType] = "char"
	m[ScopingInterfaces.RWwcharIdlType] = "wchar"
	m[ScopingInterfaces.RWbooleanIdlType] = "boolean"
	m[ScopingInterfaces.RWoctetIdlType] = "octet"
	m[ScopingInterfaces.RWanyIdlType] = "any"
	m[ScopingInterfaces.RWObjectIdlType] = "Object"
	m[ScopingInterfaces.RWint16IdlType] = "int16"
	m[ScopingInterfaces.RWint32IdlType] = "int32"
	m[ScopingInterfaces.RWint64IdlType] = "int64"
	m[ScopingInterfaces.RWuint16IdlType] = "uint16"
	m[ScopingInterfaces.RWuint32IdlType] = "uint32"
	m[ScopingInterfaces.RWuint64IdlType] = "uint64"

	factory := NewCompleteIdlLexerFactoryImpl(
		log.New(os.Stdout, "test", 0),
		scopedObjects.NewNextNumber(),
		scoping.NewScopingContext(scoping.NewDefaultTypeService(), nil))
	t.Run("Simple TypeStruct", func(t *testing.T) {
		for k, v := range m {
			t.Run(fmt.Sprintf("%v", k), func(t *testing.T) {
				stream := fmt.Sprintf(`
					struct demo
					{
						%s c, d;
					};`,
					v)
				controller := gomock.NewController(t)
				defer controller.Finish()
				reader := bufio.NewReader(strings.NewReader(stream))

				actual, _ := factory.Create("(string test)", reader)
				mock := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
				mock.EXPECT().NewStructType(gomock.Any(), "demo", gomock.Any(), gomock.Any()).Times(1)
				result := CompleteIdlParse(mock)
				assert.Equal(t, 0, result, v)
				if !assert.Equal(t, 0, result, mock.LastError()) {
					return
				}
				typeSpec, _ := mock.GetSpec()
				assert.Equal(t, ScopingInterfaces.StructIdlType, typeSpec.GetKind())
			})
		}
	})
	t.Run("Single array struct decl", func(t *testing.T) {
		for k, v := range m {

			t.Run(fmt.Sprintf("%v", k), func(t *testing.T) {
				stream := fmt.Sprintf(`
					struct demo
					{
						%s c[12], d[2];
					};`,
					v)

				controller := gomock.NewController(t)
				defer controller.Finish()
				reader := bufio.NewReader(strings.NewReader(stream))

				actual, _ := factory.Create("(string test)", reader)
				mock := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
				mock.EXPECT().NewStructType(gomock.Any(), "demo", gomock.Any(), gomock.Any()).Times(1)
				mock.EXPECT().NewDeclarator(gomock.Any(), "c").Times(1)
				mock.EXPECT().NewDeclarator(gomock.Any(), "d").Times(1)
				result := CompleteIdlParse(mock)
				assert.Equal(t, 0, result, mock.LastError())
			})
		}
	})

	t.Run("Double array struct decl", func(t *testing.T) {
		for k, v := range m {
			t.Run(fmt.Sprintf("%v", k), func(t *testing.T) {
				stream := fmt.Sprintf(`
					struct demo
					{
						%s c[12][333], d[2][333];
					};`,
					v)

				controller := gomock.NewController(t)
				defer controller.Finish()
				reader := bufio.NewReader(strings.NewReader(stream))

				actual, _ := factory.Create("(string test)", reader)
				mock := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
				mock.EXPECT().NewStructType(gomock.Any(), "demo", gomock.Any(), gomock.Any()).Times(1)
				mock.EXPECT().NewDeclarator(gomock.Any(), "c").Times(1)
				mock.EXPECT().NewDeclarator(gomock.Any(), "d").Times(1)
				result := CompleteIdlParse(mock)
				assert.Equal(t, 0, result, mock.LastError())
			})
		}
	})

	t.Run("Mixture array struct decl", func(t *testing.T) {
		for k, v := range m {
			t.Run(fmt.Sprintf("%v", k), func(t *testing.T) {
				stream := fmt.Sprintf(`
					struct demo
					{
						%s c[12], d[2][333], e;
					};`,
					v)

				controller := gomock.NewController(t)
				defer controller.Finish()
				reader := bufio.NewReader(strings.NewReader(stream))
				actual, _ := factory.Create("(string test)", reader)
				mock := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
				mock.EXPECT().NewStructType(gomock.Any(), "demo", gomock.Any(), gomock.Any()).Times(1)
				mock.EXPECT().NewDeclarator(gomock.Any(), "c").Times(1)
				mock.EXPECT().NewDeclarator(gomock.Any(), "d").Times(1)
				mock.EXPECT().NewDeclarator(gomock.Any(), "e").Times(1)
				result := CompleteIdlParse(mock)
				assert.Equal(t, 0, result, actual.LastError())
			})
		}
	})
}
