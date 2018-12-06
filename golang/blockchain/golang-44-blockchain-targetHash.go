// 本文链接： https://dreamerjonson.com/2018/12/12/golang-44-blockchain-targetHash/ 
// 
// golang[44]-区块链-目标值
// 2018-12-12 go go 10 评论 字数统计: 165(字) 阅读时长: 1(分)
// 比特币目标hash计算过程
// 
// 以之前的bits：181B7B74为例子

func main(){

	bits,_:= hex.DecodeString("181B7B74")


	fmt.Printf("%x",CalculateTargetFast(bits))

}

//18   1B7B74
func CalculateTargetFast(bits []byte) []byte{

		var result []byte
		//第一个字节  计算指数
		exponent := bits[:1]
		fmt.Printf("%x\n",exponent)

		//计算后面3个字节 系数
		coeffient:= bits[1:]
	fmt.Printf("%x\n",coeffient)


		//将字节，他的16进制为"18"  转化为了string "18"
		str:= hex.EncodeToString(exponent)  //"18"
		fmt.Printf("str=%s\n",str)
	   //将字符串18转化为了10进制int64 24
		exp,_:=strconv.ParseInt(str,16,8)

			fmt.Printf("exp=%d\n",exp)
		//拼接，计算出目标hash
		result  = append(bytes.Repeat([]byte{0x00},32-int(exp)),coeffient...)
		result  =  append(result,bytes.Repeat([]byte{0x00},32-len(result))...)


	return result
}
