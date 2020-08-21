// 2FA双因素认证 Golang 代码实现 TOTP
package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"time"
)

func main() {
	key := []byte("MOJOTV_CN_IS_AWESOME_AND_AWESOME_SECRET_KEY")
	number := totp(key, time.Now(), 6)
	fmt.Printf("2FA code: %06d\n", number)
}

func hotp(key []byte, counter uint64, digits int) int {
	//RFC 6238
	h := hmac.New(sha1.New, key)
	binary.Write(h, binary.BigEndian, counter)
	sum := h.Sum(nil)
	//取sha1的最后4byte
	//0x7FFFFFFF 是long int的最大值
	//math.MaxUint32 == 2^32-1
	//& 0x7FFFFFFF == 2^31  Set the first bit of truncatedHash to zero  //remove the most significant bit
	// len(sum)-1]&0x0F 最后 像登陆 (bytes.len-4)
	//取sha1 bytes的最后4byte 转换成 uint32
	v := binary.BigEndian.Uint32(sum[sum[len(sum)-1]&0x0F:]) & 0x7FFFFFFF
	d := uint32(1)

	//取十进制的余数
	for i := 0; i < digits && i < 8; i++ {
		d *= 10
	}
	return int(v % d)
}

func totp(key []byte, t time.Time, digits int) int {
	return hotp(key, uint64(t.Unix())/30, digits)
	//return hotp(key, uint64(t.UnixNano())/30e9, digits)
}
