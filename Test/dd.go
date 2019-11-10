package Test

import (
	"errors"
	"io"
	"unicode/utf8"
)

type ITokenBuffer interface {
	io.RuneReader
	Next(n int)
	Reset()
}

type Reader struct {
	buf  []byte
	rd   io.Reader // reader provided by the client
	r, w int       // buf read and write positions
}

func (b *Reader) ReadRune() (r rune, size int, err error) {
	if b.r == b.w {
		return 0, 0, io.EOF
	}
	r, size = rune(b.buf[b.r]), 1
	if r >= utf8.RuneSelf {
		r, size = utf8.DecodeRune(b.buf[b.r:b.w])
	}
	b.r += size
	return r, size, nil
}

func (b *Reader) fill() error {
	if b.r > 0 {
		copy(b.buf, b.buf[b.r:b.w])
		b.w -= b.r
		b.r = 0
	}

	if b.w >= len(b.buf) {
		panic("bufio: tried to fill full buffer")
	}

	for i := maxConsecutiveEmptyReads; i > 0; i-- {
		n, err := b.rd.Read(b.buf[b.w:])
		if n < 0 {
			panic(errNegativeRead)
		}
		b.w += n
		if err != nil {

			return err
		}
		if n > 0 {
			return nil
		}
	}
	return io.ErrNoProgress
}

func (b *Reader) Next(n int) {
	b.r = n
	_ = b.fill()
}

func (b *Reader) Reset() {
	b.r = 0
}

const maxConsecutiveEmptyReads = 100

func NewReader(rd io.Reader) *Reader {
	return NewReaderSize(rd, 1024)
}

func NewReaderSize(rd io.Reader, size int) *Reader {
	r := new(Reader)
	r.init(make([]byte, size), rd)
	return r
}
func (b *Reader) init(buf []byte, r io.Reader) {
	*b = Reader{
		buf: buf,
		rd:  r,
	}
	_ = b.fill()
}

var errNegativeRead = errors.New("bufio: reader returned negative count from Read")
