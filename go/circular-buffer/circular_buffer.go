package circular

import "errors"

const testVersion = 4

type Buffer struct {
	used, pos int
	buff      []byte
}

func NewBuffer(size int) *Buffer {
	var buffer = Buffer{0, 0, make([]byte, size)}
	return &buffer
}

func (buffer *Buffer) ReadByte() (byte, error) {
	var e error
	if buffer.used == 0 {
		return 0, errors.New("Buffer is emty")
	}
	var id int = buffer.pos
	for i := 0; i < buffer.used; i++ {
		if id == 0 {
			id = len(buffer.buff) - 1
		} else {
			id--
		}
	}
	buffer.used--
	return buffer.buff[id], e
}

func (buffer *Buffer) WriteByte(c byte) error {
	var e error
	if len(buffer.buff) == buffer.used {
		return errors.New("Buffer is full. Use 'Overwrite'.")
	}
	buffer.buff[buffer.pos] = c
	buffer.used++
	if buffer.pos+1 == len(buffer.buff) {
		buffer.pos = 0
	} else {
		buffer.pos++
	}
	return e
}

func (buffer *Buffer) Overwrite(c byte) {
	if len(buffer.buff) != buffer.used {
		buffer.WriteByte(c)
		return
	} else {
		buffer.buff[buffer.pos] = c
		if buffer.pos+1 == len(buffer.buff) {
			buffer.pos = 0
		} else {
			buffer.pos++
		}
	}
	return
}

func (buffer *Buffer) Reset() {
	buffer.pos = 0
	buffer.used = 0
}
