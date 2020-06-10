package main

// https://darjun.github.io/2020/05/21/godailylib/rpcx/

// 使用 Go 语言我们能很方便地生成一个证书和私钥：
// 下面代码生成了一个证书和私钥，有效期为 1 年。运行程序，得到两个文件server.pem和server.key。

import (
  "crypto/rand"
  "crypto/rsa"
  "crypto/x509"
  "crypto/x509/pkix"
  "encoding/pem"
  "math/big"
  "net"
  "os"
  "time"
)

func main() {
  max := new(big.Int).Lsh(big.NewInt(1), 128)
  serialNumber, _ := rand.Int(rand.Reader, max)
  subject := pkix.Name{
    Organization:       []string{"Go Daily Lib"},
    OrganizationalUnit: []string{"TechBlog"},
    CommonName:         "go daily lib",
  }

  template := x509.Certificate{
    SerialNumber: serialNumber,
    Subject:      subject,
    NotBefore:    time.Now(),
    NotAfter:     time.Now().Add(365 * 24 * time.Hour),
    KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
    ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
    IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
  }

  pk, _ := rsa.GenerateKey(rand.Reader, 2048)

  derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
  certOut, _ := os.Create("server.pem")
  pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
  certOut.Close()

  keyOut, _ := os.Create("server.key")
  pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
  keyOut.Close()
}