package magic8ball

import (
	"crypto/rand"
	"encoding/binary"
	mrand "math/rand"
)

// CryptoRandSource is a `crypto/rand` backed source that satisfies
// the `math/rand.Source` interface definition. It can be used as
// `r := rand.New(NewCryptoRandSource())`
// See: https://stackoverflow.com/a/35208651/1908967
type CryptoRandSource struct{}

func NewCryptoRandSource() CryptoRandSource {
	return CryptoRandSource{}
}

func (CryptoRandSource) Int63() int64 {
	var b [8]byte
	_, err := rand.Read(b[:])
	if err != nil {
		panic(err) // `math/rand.Source` interface only returns `int64`
	}
	// mask off sign bit to ensure positive number
	return int64(binary.LittleEndian.Uint64(b[:]) & (1<<63 - 1))
}

func (CryptoRandSource) Seed(int64) {}

// Intn returns a random number backed by `crypto/rand`.
func Intn(n uint) int {
	return mrand.New(NewCryptoRandSource()).Intn(int(n)) // #nosec G404
}
