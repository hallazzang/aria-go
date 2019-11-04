package aria

import "encoding/binary"

func cryptBlock(xk []uint32, dst, src []byte) {
	// TODO: implement
}

func expandKey(key []byte, enc, dec []uint32) {
	k := len(key)
	n := 8 + k/4

	var kl, kr [16]byte

	for i := 0; i < k; i++ {
		if i < 16 {
			kl[i] = key[i]
		} else {
			kr[i-16] = key[i]
		}
	}

	var ck1, ck2, ck3 [16]byte

	switch k {
	case 128 / 8:
		ck1 = c1
		ck2 = c2
		ck3 = c3
	case 192 / 8:
		ck1 = c2
		ck2 = c3
		ck3 = c1
	case 256 / 8:
		ck1 = c3
		ck2 = c1
		ck3 = c2
	}

	var w0, w1, w2, w3 [16]byte

	w0 = kl // TODO: use kl instead of w0
	w1 = xor(roundOdd(w0, ck1), kr)
	w2 = xor(roundEven(w1, ck2), w0)
	w3 = xor(roundOdd(w2, ck3), w1)

	setEncKey(enc, xor(w0, rrot(w1, 19)))
	setEncKey(enc[4:], xor(w1, rrot(w2, 19)))
	setEncKey(enc[8:], xor(w2, rrot(w3, 19)))
	setEncKey(enc[12:], xor(w3, rrot(w0, 19)))
	setEncKey(enc[16:], xor(w0, rrot(w1, 31)))
	setEncKey(enc[20:], xor(w1, rrot(w2, 31)))
	setEncKey(enc[24:], xor(w2, rrot(w3, 31)))
	setEncKey(enc[28:], xor(w3, rrot(w0, 31)))
	setEncKey(enc[32:], xor(w0, lrot(w1, 61)))
	setEncKey(enc[36:], xor(w1, lrot(w2, 61)))
	setEncKey(enc[40:], xor(w2, lrot(w3, 61)))
	setEncKey(enc[44:], xor(w3, lrot(w0, 61)))
	setEncKey(enc[48:], xor(w0, lrot(w1, 31)))
	if n > 12 {
		setEncKey(enc[52:], xor(w1, lrot(w2, 31)))
		setEncKey(enc[56:], xor(w2, lrot(w3, 31)))
	}
	if n > 14 {
		setEncKey(enc[60:], xor(w3, lrot(w0, 31)))
		setEncKey(enc[64:], xor(w0, lrot(w1, 19)))
	}

	copy(dec, enc[n*4:(n+1)*4])
	copy(dec[4:], enc[(n-1)*4:n*4])
	copy(dec[8:], enc[(n-2)*4:(n-1)*4])
	copy(dec[12:], enc[(n-3)*4:(n-2)*4])
	copy(dec[16:], enc[(n-4)*4:(n-3)*4])
	copy(dec[20:], enc[(n-5)*4:(n-4)*4])
	copy(dec[24:], enc[(n-6)*4:(n-5)*4])
	copy(dec[28:], enc[(n-7)*4:(n-6)*4])
	copy(dec[32:], enc[(n-8)*4:(n-7)*4])
	copy(dec[36:], enc[(n-9)*4:(n-8)*4])
	copy(dec[40:], enc[(n-10)*4:(n-9)*4])
	copy(dec[44:], enc[(n-11)*4:(n-10)*4])
	copy(dec[48:], enc[(n-12)*4:(n-11)*4])
	if n > 12 {
		copy(dec[52:], enc[(n-13)*4:(n-12)*4])
		copy(dec[56:], enc[(n-14)*4:(n-13)*4])
	}
	if n > 14 {
		copy(dec[60:], enc[(n-15)*4:(n-14)*4])
		copy(dec[64:], enc[(n-16)*4:(n-15)*4])
	}
}

func setEncKey(enc []uint32, x [16]byte) {
	enc[0] = binary.BigEndian.Uint32(x[:])
	enc[1] = binary.BigEndian.Uint32(x[4:])
	enc[2] = binary.BigEndian.Uint32(x[8:])
	enc[3] = binary.BigEndian.Uint32(x[12:])
}

// Round Function Fo
func roundOdd(d, rk [16]byte) [16]byte {
	return diffuse(substitute1(xor(d, rk)))
}

// Round Function Fe
func roundEven(d, rk [16]byte) [16]byte {
	return diffuse(substitute2(xor(d, rk)))
}

// Substitution Layer SL1
func substitute1(x [16]byte) (y [16]byte) {
	y[0] = sb1[x[0]]
	y[1] = sb2[x[1]]
	y[2] = sb3[x[2]]
	y[3] = sb4[x[3]]
	y[4] = sb1[x[4]]
	y[5] = sb2[x[5]]
	y[6] = sb3[x[6]]
	y[7] = sb4[x[7]]
	y[8] = sb1[x[8]]
	y[9] = sb2[x[9]]
	y[10] = sb3[x[10]]
	y[11] = sb4[x[11]]
	y[12] = sb1[x[12]]
	y[13] = sb2[x[13]]
	y[14] = sb3[x[14]]
	y[15] = sb4[x[15]]
	return
}

// Substitution Layer SL2
func substitute2(x [16]byte) (y [16]byte) {
	y[0] = sb3[x[0]]
	y[1] = sb4[x[1]]
	y[2] = sb1[x[2]]
	y[3] = sb2[x[3]]
	y[4] = sb3[x[4]]
	y[5] = sb4[x[5]]
	y[6] = sb1[x[6]]
	y[7] = sb2[x[7]]
	y[8] = sb3[x[8]]
	y[9] = sb4[x[9]]
	y[10] = sb1[x[10]]
	y[11] = sb2[x[11]]
	y[12] = sb3[x[12]]
	y[13] = sb4[x[13]]
	y[14] = sb1[x[14]]
	y[15] = sb2[x[15]]
	return
}

// Diffuse Layer A
func diffuse(x [16]byte) (y [16]byte) {
	y[0] = x[3] ^ x[4] ^ x[6] ^ x[8] ^ x[9] ^ x[13] ^ x[14]
	y[1] = x[2] ^ x[5] ^ x[7] ^ x[8] ^ x[9] ^ x[12] ^ x[15]
	y[2] = x[1] ^ x[4] ^ x[6] ^ x[10] ^ x[11] ^ x[12] ^ x[15]
	y[3] = x[0] ^ x[5] ^ x[7] ^ x[10] ^ x[11] ^ x[13] ^ x[14]
	y[4] = x[0] ^ x[2] ^ x[5] ^ x[8] ^ x[11] ^ x[14] ^ x[15]
	y[5] = x[1] ^ x[3] ^ x[4] ^ x[9] ^ x[10] ^ x[14] ^ x[15]
	y[6] = x[0] ^ x[2] ^ x[7] ^ x[9] ^ x[10] ^ x[12] ^ x[13]
	y[7] = x[1] ^ x[3] ^ x[6] ^ x[8] ^ x[11] ^ x[12] ^ x[13]
	y[8] = x[0] ^ x[1] ^ x[4] ^ x[7] ^ x[10] ^ x[13] ^ x[15]
	y[9] = x[0] ^ x[1] ^ x[5] ^ x[6] ^ x[11] ^ x[12] ^ x[14]
	y[10] = x[2] ^ x[3] ^ x[5] ^ x[6] ^ x[8] ^ x[13] ^ x[15]
	y[11] = x[2] ^ x[3] ^ x[4] ^ x[7] ^ x[9] ^ x[12] ^ x[14]
	y[12] = x[1] ^ x[2] ^ x[6] ^ x[7] ^ x[9] ^ x[11] ^ x[12]
	y[13] = x[0] ^ x[3] ^ x[6] ^ x[7] ^ x[8] ^ x[10] ^ x[13]
	y[14] = x[0] ^ x[3] ^ x[4] ^ x[5] ^ x[9] ^ x[11] ^ x[14]
	y[15] = x[1] ^ x[2] ^ x[4] ^ x[5] ^ x[8] ^ x[10] ^ x[15]
	return
}

func xor(a, b [16]byte) (r [16]byte) {
	r[0] = a[0] ^ b[0]
	r[1] = a[1] ^ b[1]
	r[2] = a[2] ^ b[2]
	r[3] = a[3] ^ b[3]
	r[4] = a[4] ^ b[4]
	r[5] = a[5] ^ b[5]
	r[6] = a[6] ^ b[6]
	r[7] = a[7] ^ b[7]
	r[8] = a[8] ^ b[8]
	r[9] = a[9] ^ b[9]
	r[10] = a[10] ^ b[10]
	r[11] = a[11] ^ b[11]
	r[12] = a[12] ^ b[12]
	r[13] = a[13] ^ b[13]
	r[14] = a[14] ^ b[14]
	r[15] = a[15] ^ b[15]
	return
}

func lrot(x [16]byte, n uint) (y [16]byte) {
	q, r := n/8, n%8
	s := 8 - r
	y[0] = x[q%16]<<r | x[(q+1)%16]>>s
	y[1] = x[(q+1)%16]<<r | x[(q+2)%16]>>s
	y[2] = x[(q+2)%16]<<r | x[(q+3)%16]>>s
	y[3] = x[(q+3)%16]<<r | x[(q+4)%16]>>s
	y[4] = x[(q+4)%16]<<r | x[(q+5)%16]>>s
	y[5] = x[(q+5)%16]<<r | x[(q+6)%16]>>s
	y[6] = x[(q+6)%16]<<r | x[(q+7)%16]>>s
	y[7] = x[(q+7)%16]<<r | x[(q+8)%16]>>s
	y[8] = x[(q+8)%16]<<r | x[(q+9)%16]>>s
	y[9] = x[(q+9)%16]<<r | x[(q+10)%16]>>s
	y[10] = x[(q+10)%16]<<r | x[(q+11)%16]>>s
	y[11] = x[(q+11)%16]<<r | x[(q+12)%16]>>s
	y[12] = x[(q+12)%16]<<r | x[(q+13)%16]>>s
	y[13] = x[(q+13)%16]<<r | x[(q+14)%16]>>s
	y[14] = x[(q+14)%16]<<r | x[(q+15)%16]>>s
	y[15] = x[(q+15)%16]<<r | x[q%16]>>s
	return
}

func rrot(x [16]byte, n uint) (y [16]byte) {
	q, r := n/8%16, n%8
	s := 8 - r
	y[0] = x[(16-q)%16]>>r | x[(15-q)%16]<<s
	y[1] = x[(17-q)%16]>>r | x[(16-q)%16]<<s
	y[2] = x[(18-q)%16]>>r | x[(17-q)%16]<<s
	y[3] = x[(19-q)%16]>>r | x[(18-q)%16]<<s
	y[4] = x[(20-q)%16]>>r | x[(19-q)%16]<<s
	y[5] = x[(21-q)%16]>>r | x[(20-q)%16]<<s
	y[6] = x[(22-q)%16]>>r | x[(21-q)%16]<<s
	y[7] = x[(23-q)%16]>>r | x[(22-q)%16]<<s
	y[8] = x[(24-q)%16]>>r | x[(23-q)%16]<<s
	y[9] = x[(25-q)%16]>>r | x[(24-q)%16]<<s
	y[10] = x[(26-q)%16]>>r | x[(25-q)%16]<<s
	y[11] = x[(27-q)%16]>>r | x[(26-q)%16]<<s
	y[12] = x[(28-q)%16]>>r | x[(27-q)%16]<<s
	y[13] = x[(29-q)%16]>>r | x[(28-q)%16]<<s
	y[14] = x[(30-q)%16]>>r | x[(29-q)%16]<<s
	y[15] = x[(31-q)%16]>>r | x[(30-q)%16]<<s
	return
}
