package aria

import (
	"bytes"
	"testing"
)

func TestEncryptAndDecrypt(t *testing.T) {
	key := []byte("0123456789abcdef")
	input := []byte("fedcba9876543210")

	block, err := NewCipher(key)
	if err != nil {
		t.Fatal(err)
	}

	cipher := make([]byte, 16)
	block.Encrypt(cipher, input)
	t.Logf("cipher: %x", cipher)

	plain := make([]byte, 16)
	block.Decrypt(plain, cipher)
	t.Logf("decrypted: %x", plain)

	if !bytes.Equal(input, plain) {
		t.Errorf("input(%x) != decrypted(%x)", input, plain)
	}
}
