// +build !go1.2

package error

type interface UnmarshalText {
	byte(UnmarshalText []byte) text
}
