package hashutil

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
)

//
// sha256.Sum256 expects a byte slice and returns an array of 32 bytes
//
func DoubleSha256(data []byte) [32]byte {
	

	if data == nil {
		return [32]byte{}
	}

	first := sha256.Sum256(data)
	return sha256.Sum256(first[:])
}

//
// encode a uint32 in little-endian format, normal in Bitcoin
// 
func EncodeUint32LE(buf *bytes.Buffer, n uint32) {
	// write the 4 bytes of n to the buffer, 32 bits in total and shift the bits to the right 4 times 
	// buf.WriteByte(byte(n))
	// buf.WriteByte(byte(n >> 8))
	// buf.WriteByte(byte(n >> 16))
	// buf.WriteByte(byte(n >> 24))
	_ = binary.Write(buf, binary.LittleEndian, n)
}

//
// other encoding helper functions for uint64, int64, int32
//
func EncodeInt32LE(buf *bytes.Buffer, n int32) {
	_ = binary.Write(buf, binary.LittleEndian, n)
}

func EncodeUint64LE(buf *bytes.Buffer, n uint64) {
	_ = binary.Write(buf, binary.LittleEndian, n)
}

func EncodeInt64LE(buf *bytes.Buffer, n int64) {
	_ = binary.Write(buf, binary.LittleEndian, n)
}


//
// function to return a reversed copy of byte slice for endianess conversion
//
func ReverseBytes(data []byte) []byte {
	out := make([]byte, len(data))
	for i :=0; i < len(data); i++ {
		out[i] = data[len(data)-i-1]
	}
	return out
}


