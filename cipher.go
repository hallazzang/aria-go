package aria

import (
	"crypto/cipher"
	"fmt"
)

// BlockSize is the ARIA block size in bytes.
const BlockSize = 16

type ariaCipher struct {
	enc []uint32
	dec []uint32
}

type KeySizeError int

func (k KeySizeError) Error() string {
	return fmt.Sprintf("aria: invalid key size %d", int(k))
}

// NewCipher creates and returns a new cipher.Block.
// The key argument should be the ARIA key,
// either 16, 24, or 32 bytes to select
// ARIA-128, ARIA-192, or AES-256.
func NewCipher(key []byte) (cipher.Block, error) {
	k := len(key)
	switch k {
	default:
		return nil, KeySizeError(k)
	case 128 / 8, 192 / 8, 256 / 8:
		break
	}
	c := ariaCipher{enc: make([]uint32, k), dec: make([]uint32, k)}
	expandKey(key, c.enc, c.dec)
	return &c, nil
}

func (c *ariaCipher) BlockSize() int { return BlockSize }

func (c *ariaCipher) Encrypt(dst, src []byte) {
	if len(src) < BlockSize {
		panic("aria: input not full block")
	}
	if len(dst) < BlockSize {
		panic("aria: output not full block")
	}
	if inexactOverlap(dst[:BlockSize], src[:BlockSize]) {
		panic("aria: invalid buffer overlap")
	}
	cryptBlock(c.enc, dst, src)
}

func (c *ariaCipher) Decrypt(dst, src []byte) {
	if len(src) < BlockSize {
		panic("aria: input not full block")
	}
	if len(dst) < BlockSize {
		panic("aria: output not full block")
	}
	if inexactOverlap(dst[:BlockSize], src[:BlockSize]) {
		panic("aria: invalid buffer overlap")
	}
	cryptBlock(c.dec, dst, src)
}
