package yaccIdlTests

import (
	"github.com/bhbosman/CodeGenerators/mitchgenerator/MitchDefinedTypes"
	"github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMitchTypeInformation(t *testing.T) {
	typeInformation := MitchDefinedTypes.NewMitchTypeInformation()
	t.Run("CheckCreateType", func(t *testing.T) {

		data, err := typeInformation.CreateType(interfaces.Struct, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.Int64, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.Char, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.Enum, nil)
		assert.Error(t, err)
		assert.Nil(t, data)

		data, err = typeInformation.CreateType(interfaces.MitchAlpha, int64(1))
		assert.NoError(t, err)
		assert.NotNil(t, data)

		data, err = typeInformation.CreateType(interfaces.MitchBitField, []string{"1", "2", "3", "4", "5", "6", "7", "8"})
		assert.NoError(t, err)
		assert.NotNil(t, data)

		data, err = typeInformation.CreateType(interfaces.MitchByte, nil)
		assert.NoError(t, err)
		assert.NotNil(t, data)

		data, err = typeInformation.CreateType(interfaces.MitchDate, nil)
		assert.NoError(t, err)
		assert.NotNil(t, data)

		data, err = typeInformation.CreateType(interfaces.MitchTime, nil)
		assert.NoError(t, err)
		assert.NotNil(t, data)

		data, err = typeInformation.CreateType(interfaces.MitchPrice04, nil)
		assert.NoError(t, err)
		assert.NotNil(t, data)

		data, err = typeInformation.CreateType(interfaces.MitchPrice08, nil)
		assert.NoError(t, err)
		assert.NotNil(t, data)

		data, err = typeInformation.CreateType(interfaces.MitchUInt08, nil)
		assert.NoError(t, err)
		assert.NotNil(t, data)

		data, err = typeInformation.CreateType(interfaces.MitchUInt16, nil)
		assert.NoError(t, err)
		assert.NotNil(t, data)

		data, err = typeInformation.CreateType(interfaces.MitchUInt32, nil)
		assert.NoError(t, err)
		assert.NotNil(t, data)

		data, err = typeInformation.CreateType(interfaces.MitchUInt64, nil)
		assert.NoError(t, err)
		assert.NotNil(t, data)
	})

}
