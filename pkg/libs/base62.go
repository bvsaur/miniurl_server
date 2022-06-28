package libs

import (
	"errors"
	"math"
	"strings"
)

const (
	base    = 62
	charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

func EncodeBase62(id uint) string {
	b := make([]byte, 0)

	for id > 0 {
		r := math.Mod(float64(id), float64(base))
		id /= base
		b = append([]byte{charset[int(r)]}, b...)
	}

	return string(b)
}

func DecodeBase62(s string) (uint, error) {
	var r, pow int
	for i, v := range s {
		pow = len(s) - (i + 1)
		pos := strings.IndexRune(charset, v)

		if pos == -1 {
			return 0, errors.New("invalid character: " + string(v))
		}

		r += pos * int(math.Pow(float64(base), float64(pow)))
	}

	return uint(r), nil
}
