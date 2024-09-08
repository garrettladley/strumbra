package stumbra

import (
	"bytes"
	"errors"
	"unsafe"
)

const (
	inlinedLength = 12
	prefixLength  = 4
	suffixLength  = 8
)

// ErrTooLong is returned when trying to create a UmbraString from a string that's too long (> 2^32-1 bytes)
var ErrTooLong = errors.New("string is too long")

// UmbraString is a string implementation that allows for the very important “short string optimization”
// this data structure is described in this paper: https://www.cidrdb.org/cidr2020/papers/p29-neumann-cidr20.pdf
type UmbraString struct {
	len      int32
	prefix   [prefixLength]byte
	trailing unsafe.Pointer
}

func New(s string) (UmbraString, error) {
	if len(s) > 1<<32-1 {
		return UmbraString{}, ErrTooLong
	}

	us := UmbraString{
		len: int32(len(s)),
	}

	switch {
	case us.len <= prefixLength:
		copy(us.prefix[:], s)
	case us.len <= inlinedLength:
		copy(us.prefix[:], s[:prefixLength])
		buf := new([suffixLength]byte)
		copy(buf[:], s[prefixLength:])
		us.trailing = unsafe.Pointer(buf)
	default:
		copy(us.prefix[:], s[:prefixLength])
		us.trailing = unsafe.Pointer(&s)
	}

	return us, nil
}

func (us *UmbraString) Len() int {
	return int(us.len)
}

func (us *UmbraString) IsEmpty() bool {
	return us.len == 0
}

func (us *UmbraString) String() string {
	return string(us.Bytes())
}

func (us *UmbraString) Bytes() []byte {
	if us.len <= inlinedLength {
		return append(us.prefix[:], (*[suffixLength]byte)(us.trailing)[:us.len-prefixLength]...)
	}
	return append(us.prefix[:], us.suffix()...)
}

func (us *UmbraString) Equal(other UmbraString) bool {
	// get the first 8 bytes, this includes the length and the prefix
	lhs := *(*[8]byte)(unsafe.Pointer(&us))
	rhs := *(*[8]byte)(unsafe.Pointer(&other))
	if lhs != rhs {
		return false
	}

	if us.len <= inlinedLength {
		if (*[suffixLength]byte)(us.trailing) == nil && (*[suffixLength]byte)(other.trailing) == nil {
			return true
		}
		return bytes.Equal((*[suffixLength]byte)(us.trailing)[:], (*[suffixLength]byte)(us.trailing)[:])
	}

	return bytes.Equal(us.suffix(), other.suffix())

}

func (us *UmbraString) Compare(other UmbraString) int {
	if prefixCompare := bytes.Compare(us.prefix[:], other.prefix[:]); prefixCompare != 0 {
		return prefixCompare
	}

	if us.len <= prefixLength && other.len <= prefixLength {
		return i32Compare(us.len, other.len)
	}

	if us.len <= inlinedLength && other.len <= inlinedLength {
		if trailing := bytes.Compare((*[suffixLength]byte)(us.trailing)[:], (*[suffixLength]byte)(other.trailing)[:]); trailing != 0 {
			return trailing
		}
		return i32Compare(us.len, other.len)
	}

	return bytes.Compare(us.suffix(), other.suffix())
}

func (us *UmbraString) suffix() []byte {
	if us.len <= inlinedLength {
		return (*[suffixLength]byte)(us.trailing)[:us.len-prefixLength]
	}
	return unsafe.Slice((*byte)(us.trailing), us.len)[prefixLength:]
}

func i32Compare(a, b int32) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
