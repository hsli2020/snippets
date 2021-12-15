package main

type ICipher interface {
	Encrypt([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
}

// 然后根据这个接口，分别实现 AESCipher 和 DESCipher 两个加密类。
// AESCipher:

type AESCipher struct {
}

func NewAESCipher() *AESCipher {
	return &AESCipher{}
}

func (c AESCipher) Encrypt(data []byte) ([]byte, error) {
	return nil, nil
}

func (c AESCipher) Decrypt(data []byte) ([]byte, error) {
	return nil, nil
}

// DESCipher:
type DESCipher struct {
}

func NewDesCipher() *DESCipher {
	return &DESCipher{}
}

func (c DESCipher) Encrypt(data []byte) ([]byte, error) {
	return nil, nil
}

func (c DESCipher) Decrypt(data []byte) ([]byte, error) {
	return nil, nil
}

// 最后是一个工厂角色，根据传入的参数返回对应的加密类，
// Java 需要实现一个工厂类，这里我们用一个函数来做加密类工厂：

func CipherFactory(cType string) ICipher {
	switch cType {
	case "AES":
		return NewAESCipher()
	case "DES":
		return NewDesCipher()
	default:
		return nil
	}
}

// 这样，通过调用 CipherFactory 传入所需的加密类型，就可以得到所需要的加密类实例了。
func main() {
	c := CipherFactory("RSA")
	if c != nil {
		log.Fatalf("unsupport RSA")
	}

	c = CipherFactory("AES")
	if reflect.TypeOf(c) != reflect.TypeOf(&AESCipher{}) {
		log.Fatalf("cipher type should be AES")
	}

	c = CipherFactory("DES")
	if reflect.TypeOf(c) != reflect.TypeOf(&DESCipher{}) {
		log.Fatalf("cipher type should be DES")
	}
}
