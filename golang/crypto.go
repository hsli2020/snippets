// 怎么使用 AES

// 加密

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)
func Encrypt(data []byte, key [32]byte) ([]byte, error) {
	//	初始化 block cipher
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}
	//	设置 block cipher mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	//	生成随机 nonce
	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, err
	}
	//	封装、返回
	return gcm.Seal(nonce, nonce, data, nil), nil
}

// 解密

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)
func Decrypt(ciphertext []byte, key [32]byte) (plaintext []byte, err error) {
	//	初始化 block cipher
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}
	//	设置 block cipher mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	//	返回解开的包，注意这里的 nonce 是直接取的。
	return gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
}

// 我们应该使用 HMAC，而不要直接用 Hash

// 实现 Hash

import (
	"crypto/hmac"
	"crypto/sha512"
)
func Hash(tag string, data []byte) []byte {
	h := hmac.New(sha512.New512_256, []byte(tag))
	h.Write(data)
	return h.Sum(nil)
}
func ExampleHash() error {
	tag := "hashing file for storage key"
	contents, err := ioutil.ReadFile("testfile")
	if err != nil {
		return error
	}
	digest := Hash(tag, contents)
	fmt.Println(hex.EncodeToString(digest))
}
//	Output:
//	9f4c795d8ae5e207f19184ccebee6a606c1fdfe509c793614006d613580f03e1

// 使用 bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)
func HashPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, 14)
}
func CheckPasswordHash(hash, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}
func Example() {
	myPassword := []byte("password")
	hashed, err := HashPassword(myPassword)
	if err != nil {
		return
	}
	fmt.Println(string(hashed))
}

// 签名

// 生成密钥

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)
func NewSigningKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

// 签名数据

func Sign(data []byte, priv *ecdsa.PrivateKey) ([]byte, error) {
	digest := sha256.Sum256(data)
	r, s, err := ecdsa.Sign(rand.Reader, priv, digest[:])
	if err != nil {
		return nil, err
	}
	//	encode the signature {R, S}
	params := priv.Curve.Params()
	curveByteSize := params.P.BitLen() / 8
	rBytes, sBytes := r.Bytes(), s.Bytes()
	signature := make([]byte, curveByteSize * 2)
	copy(signature[curveByteSize - len(rBytes):], rBytes)
	copy(signature[curveByteSize*2 - len(sBytes):], sBytes)
	return signature, nil
}

// 验证签名

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"math/big"
)
//	验证成功返回 true，否则 false
func Verify(data, sig []byte, pub *ecdsa.PublicKey) bool {
	digest := sha256.Sum256(data)
	curveByteSize := pub.Curve.Params().P.BitLen() / 8
	r, s := new(big.Int), new(big.Int)
	r.SetBytes(signature[:curveByteSize])
	s.SetBytes(signature[curveByteSize:])
	return ecdsa.Verify(pub, digest[:], r, s)
}
