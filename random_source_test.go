package faker

import (
	cryptorand "crypto/rand"
	"io"
	mathrand "pgregory.net/rand"
	"testing"
)

func TestSetRandomSource(t *testing.T) {
	SetRandomSource(mathrand.New())

	_ = rand.Int31n(100)
}

func TestSetCryptoSource(t *testing.T) {
	SetCryptoSource(cryptorand.Reader)

	buf := make([]byte, 10)
	_, err := io.ReadFull(crypto, buf)
	if err != nil {
		t.Error("Expected Not Error, But Got: ", err)
	}
}
