// 本文链接： https://dreamerjonson.com/2018/12/05/golang-32-base58/ 
// golang[32]-区块链-base58 2018-12-05
// base58
// 
// Base58是用于Bitcoin中使用的一种独特的编码方式，主要用于产生Bitcoin的钱包地址。相比Base64，
// Base58不使用数字"0"，字母大写"O"，字母大写"I"，和字母小写"l"，以及"+“和”/"符号。
// 
// 设计Base58主要的目的是：
// 避免混淆。在某些字体下，数字0和字母大写O，以及字母大写I和字母小写l会非常相似。
// 不使用"+“和”/"的原因是非字母或数字的字符串作为帐号较难被接受。
// 没有标点符号，通常不会被从中间分行。
// 大部分的软件支持双击选择整个字符串。
// base58编码

package main

import (
	"math/big"
	"fmt"
)

//切片存储base58字母
var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")


func Base58Encode(input []byte) []byte{
//定义一个字节切片，返回值
	var result []byte

//把字节数组input转化为了大整数big.Int
	x:= big.NewInt(0).SetBytes(input)

//长度58的大整数
	base := big.NewInt(int64(len(b58Alphabet)))
  //0的大整数
	zero := big.NewInt(0)
//大整数的指针
	mod := &big.Int{}

  //循环，不停地对x取余数,大小为58
	for x.Cmp(zero) != 0 {
		x.DivMod(x,base,mod)  // 对x取余数

    //讲余数添加到数组当中
		result =  append(result, b58Alphabet[mod.Int64()])
	}


//反转字节数组
	ReverseBytes(result)

//如果这个字节数组的前面为字节0，会把它替换为1.
for _,b:=range input{

		if b ==0x00{
			result =  append([]byte{b58Alphabet[0]},result...)
		}else{
			break
		}
	}


	return result

}

//反转字节数组
func ReverseBytes(data []byte){
	for i,j :=0,len(data) - 1;i<j;i,j = i+1,j - 1{
		data[i],data[j] = data[j],data[i]
	}
}

//测试 反转操作
func main(){
	org := []byte("qwerty")
	fmt.Println(string(org))

	ReverseBytes(org)

	fmt.Println(string(org))
//测试编码
  fmt.Printf("%s",string( Base58Encode([]byte("hello jonson"))))
}

//解码

func Base58Decode(input []byte) []byte{
	result :=  big.NewInt(0)
	zeroBytes :=0
	for _,b :=range input{
		if b=='1'{
			zeroBytes++
		}else{
			break
		}
	}

	payload:= input[zeroBytes:]

	for _,b := range payload{
		charIndex := bytes.IndexByte(b58Alphabet,b)  //反推出余数

		result.Mul(result,big.NewInt(58))   //之前的结果乘以58

		result.Add(result,big.NewInt(int64(charIndex)))  //加上这个余数

	}

	decoded :=result.Bytes()


	decoded =  append(bytes.Repeat([]byte{0x00},zeroBytes),decoded...)
	return decoded
}

// 完整代码

package main

import (
	"math/big"
	"fmt"
	"bytes"
)

var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")


func Base58Encode(input []byte) []byte{
	var result []byte

	x:= big.NewInt(0).SetBytes(input)

	base := big.NewInt(int64(len(b58Alphabet)))
	zero := big.NewInt(0)

	mod := &big.Int{}
	for x.Cmp(zero) != 0 {
		x.DivMod(x,base,mod)  // 对x取余数
		result =  append(result, b58Alphabet[mod.Int64()])
	}

	ReverseBytes(result)

	for _,b:=range input{

		if b ==0x00{
			result =  append([]byte{b58Alphabet[0]},result...)
		}else{
			break
		}
	}


	return result
}


func Base58Decode(input []byte) []byte{
	result :=  big.NewInt(0)
	zeroBytes :=0
	for _,b :=range input{
		if b=='1'{
			zeroBytes++
		}else{
			break
		}
	}

	payload:= input[zeroBytes:]

	for _,b := range payload{
		charIndex := bytes.IndexByte(b58Alphabet,b)  //反推出余数

		result.Mul(result,big.NewInt(58))   //之前的结果乘以58

		result.Add(result,big.NewInt(int64(charIndex)))  //加上这个余数

	}

	decoded :=result.Bytes()


	decoded =  append(bytes.Repeat([]byte{0x00},zeroBytes),decoded...)
	return decoded
}

func ReverseBytes(data []byte){
	for i,j :=0,len(data) - 1;i<j;i,j = i+1,j - 1{
		data[i],data[j] = data[j],data[i]
	}
}

func main(){
	org := []byte("qwerty")
	fmt.Println(string(org))

	ReverseBytes(org)

	fmt.Println(string(org))


	fmt.Printf("%s\n",string( Base58Encode([]byte("hello jonson"))))

	fmt.Printf("%s",string(Base58Decode([]byte("2yGEbwRFyav6CimZ7"))))
}

// 参考资料
// 
// (比特币wiki-base58编码)[https://en.bitcoin.it/wiki/Base58Check_encoding#Version_bytes]
// (维基百科-base58编码)[https://zh.wikipedia.org/wiki/Base58]
