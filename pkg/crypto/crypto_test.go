package crypto

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

// obsolete test struct implementing Encoder
// ------------------------------------------------
type testStruct struct {
	x string
	y string
}

func (struc *testStruct) Encode() []byte {
	return []byte(struc.x + struc.y)
}

// Sanity checks of SHA3_256
// ------------------------------------------------
func TestSha3_256(t *testing.T) {
	fmt.Println("testing Sha3_256:")
	input := []byte("test")
	expected, _ := hex.DecodeString("36f028580bb02cc8272a9a020f4200e346e276ae664e45ee80745574e2f5ab80")

	alg := InitHashAlgo(SHA3_256)
	hash := alg.ComputeBytesHash(input).ToBytes()
	checkBytes(t, input, expected, hash)

	hash = alg.ComputeStructHash(&testStruct{"te", "st"}).ToBytes()
	checkBytes(t, input, expected, hash)

	alg.Reset()
	alg.AddBytes([]byte("te"))
	alg.AddBytes([]byte("s"))
	alg.AddBytes([]byte("t"))
	hash = alg.SumHash().ToBytes()
	checkBytes(t, input, expected, hash)
}

func checkBytes(t *testing.T, input, expected, result []byte) {
	if !bytes.Equal(expected, result) {
		t.Errorf("hash mismatch: expect: %x have: %x, input is %x", expected, result, input)
	} else {
		t.Logf("hash test ok: expect: %x, input: %x", expected, input)
	}
}

// SHA3_256 bench
// ------------------------------------------------
func BenchmarkSha3_256(b *testing.B) {
	a := []byte("Bench me!")
	alg := InitHashAlgo(SHA3_256)
	for i := 0; i < b.N; i++ {
		alg.ComputeBytesHash(a)
	}
	return
}

// BLS tests
// ------------------------------------------------
func TestBLS_BLS12381(t *testing.T) {
	fmt.Println("testing BLS on bls12_381:")
	input := []byte("test")
	halg := InitHashAlgo(SHA3_256)
	h := halg.ComputeBytesHash(input)

	salg := InitSignatureAlgo(BLS_BLS12381)
	sk := salg.GeneratePrKey()
	pk := sk.GetPubkey()

	s := salg.SignHash(sk, h)
	result := salg.VerifyHash(pk, s, h)

	if result == false {
		t.Errorf("Verification failed: signature is %x", s)
	} else {
		t.Logf("Verification passed: signature is %x", s)
	}

	message := &testStruct{"te", "st"}
	s = salg.SignStruct(sk, message, halg)
	result = salg.VerifyStruct(pk, s, message, halg)

	if result == false {
		t.Errorf("Verification failed: signature is %x", s)
	} else {
		t.Logf("Verification passed: signature is %x", s)
	}

}
