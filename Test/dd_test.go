package Test

import (
	"bufio"
	"bytes"
	"io"

	"fmt"

	"regexp"
	"testing"
)

func TestName(t *testing.T) {

	t.Run("", func(t *testing.T) {
		var buffer ITokenBuffer = NewReader(
			bytes.NewBufferString("Brendan Bosman"))

		reg := regexp.MustCompile("[A-Za-z]*")
		s := reg.FindReaderIndex(buffer)
		fmt.Println(s)
		buffer.Next(s[1])
		s = reg.FindReaderIndex(buffer)
		fmt.Println(s)

	})
	t.Run("", func(t *testing.T) {
		var buffer io.RuneReader = bufio.NewReader(
			bytes.NewBufferString("Brendan"))

		reg := regexp.MustCompile("[A-Za-z]*")
		s := reg.FindReaderIndex(buffer)
		fmt.Println(s)
		//buffer.SetNextTypeSpec()
		s = reg.FindReaderIndex(buffer)
		fmt.Println(s)

	})

}
