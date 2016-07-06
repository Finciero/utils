package uuid

import (
	"crypto/rand"
	"fmt"
	"io"
)

// UUID type definition
type UUID string

func (uuid *UUID) String() string {
	return string(*uuid)
}

// NewFromReader creates a new UUID from the reader r
func NewFromReader(r io.Reader) (*UUID, error) {
	base := make([]byte, 16)
	n, err := io.ReadFull(r, base)
	if n != len(base) || err != nil {
		return new(UUID), err
	}
	base[8] = base[8]&^0xc0 | 0x80
	base[6] = base[6]&^0xf0 | 0x40

	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", base[0:4], base[4:6], base[6:8], base[8:10], base[10:])
	return (*UUID)(&uuid), nil
}

// New returns a new random uuid
func New() (*UUID, error) {
	return NewFromReader(rand.Reader)
}

// RandomUUIDString returns a random uuid string, if any errors accours retunrs an empty string.
func RandomUUIDString() string {
	uuid, err := New()
	if err != nil {
		return ""
	}
	return uuid.String()
}
