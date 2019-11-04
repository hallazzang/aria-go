package aria

import (
	"bytes"
	"crypto/aes"
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

func BenchmarkAES128(b *testing.B) { benchmarkAES(b, 16) }
func BenchmarkAES192(b *testing.B) { benchmarkAES(b, 24) }
func BenchmarkAES256(b *testing.B) { benchmarkAES(b, 32) }

func BenchmarkARIA128(b *testing.B) { benchmarkARIA(b, 16) }
func BenchmarkARIA192(b *testing.B) { benchmarkARIA(b, 24) }
func BenchmarkARIA256(b *testing.B) { benchmarkARIA(b, 32) }

func benchmarkAES(b *testing.B, k int) {
	b.StopTimer()
	key := make([]byte, k)
	input := make([]byte, 16)
	block, err := aes.NewCipher(key)
	if err != nil {
		b.Fatal(err)
	}
	cipher := make([]byte, 16)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		block.Encrypt(cipher, input)
	}
}

func benchmarkARIA(b *testing.B, k int) {
	b.StopTimer()
	key := make([]byte, k)
	input := make([]byte, 16)
	block, err := NewCipher(key)
	if err != nil {
		b.Fatal(err)
	}
	cipher := make([]byte, 16)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		block.Encrypt(cipher, input)
	}
}
