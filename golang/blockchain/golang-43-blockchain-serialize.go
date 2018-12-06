// 本文链接： https://dreamerjonson.com/2018/12/12/golang-43-blockchain-serialize/ 
// golang[43]-区块链-真实比特币序列化 2018-12-12

真实比特币序列化

package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"fmt"
	"encoding/hex"
	"crypto/sha256"
)

//将类型转化为了字节数组
func IntToHex(num int32) []byte{
		buff := new(bytes.Buffer)
//binary.LittleEndian 小端模式
		err:= binary.Write(buff,binary.LittleEndian,num)

		if err!=nil{
			log.Panic(err)
		}

		return buff.Bytes()
}

//字节反转
func ReverseBytes4(data []byte){
	for i,j :=0,len(data) - 1;i<j;i,j = i+1,j - 1{
		data[i],data[j] = data[j],data[i]
	}
}

func main(){

	//版本号

	var version int32 = 2

	fmt.Printf("%x\n",IntToHex(version))

	//前一个区块的hash

	prev,_ := hex.DecodeString("000000000000000016145aa12fa7e81a304c38aec3d7c5208f1d33b587f966a6")
	ReverseBytes4(prev)
	fmt.Printf("%x\n",prev)
//默克尔根
	merkleroot,_ := hex.DecodeString("3a4f410269fcc4c7885770bc8841ce6781f15dd304ae5d2770fc93a21dbd70d7")
	ReverseBytes4(merkleroot)
	fmt.Printf("%x\n",merkleroot)
//时间
	var time int32 = 1418755780
	fmt.Printf("%x\n",IntToHex(time))

//难度
	var bits int32 = 404454260
	fmt.Printf("%x\n",IntToHex(bits))
//随机数
	var nonce int32 = 1865996595
	fmt.Printf("%x\n",IntToHex(nonce))


//拼接
	result := bytes.Join([][]byte{IntToHex(version),prev,merkleroot,IntToHex(time),IntToHex(bits),IntToHex(nonce)},[]byte{})

	fmt.Printf("%x\n",result)

	//double hash256
	firsthash := sha256.Sum256(result)
	resulthash:= sha256.Sum256(firsthash[:])

	ReverseBytes4(resulthash[:])
	fmt.Printf("%x",resulthash)
}

// 参考资料：
// https://www.blockchain.com/btc/block/00000000000000000a1f57cd656e5522b7bac263aa33fc98c583ad68de309603
