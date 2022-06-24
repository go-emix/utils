package uuid

import (
	"crypto/sha1"
	"encoding/hex"
)

type UUID [16]byte

func NewString(space []byte) string {
	h := sha1.New()
	h.Reset()
	h.Write(space[:])
	s := h.Sum(nil)
	var uuid UUID
	copy(uuid[:], s)
	uuid[6] = (uuid[6] & 0x0f) | uint8((5&0xf)<<4)
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	var buf [36]byte
	encodeHex(buf[:], uuid)
	return string(buf[:])
}

func encodeHex(dst []byte, uuid UUID) {
	hex.Encode(dst, uuid[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], uuid[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], uuid[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], uuid[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], uuid[10:])
}
