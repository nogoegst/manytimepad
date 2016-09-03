// manytimepad.go - use to have fun, not for real things. seriously.
//
// To the extent possible under law, Ivan Markin waived all copyright
// and related or neighboring rights to this module of manytimepad, using the creative
// commons "cc0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package manytimepad

import (
	"crypto/cipher"
)

type manyTimePadCipher struct {
	key []byte
}

func (c manyTimePadCipher) XORKeyStream(dst, src []byte) {
	for i:=0; i < len(src); i++ {
		dst[i] = src[i] ^ c.key[i%len(c.key)]
	}
}

func NewCipher(key []byte) (cipher.Stream, error) {
	c := manyTimePadCipher{key: make([]byte, len(key))}
	copy(c.key, key)
	return c, nil
}

