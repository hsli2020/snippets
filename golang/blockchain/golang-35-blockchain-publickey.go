// 本文链接： https://dreamerjonson.com/2018/12/07/golang-35-blockchain-publickey/ 
// golang[35]-区块链-私钥公钥生成 2018-12-07

//生成私钥和公钥
func newKeyPair() (ecdsa.PrivateKey,[]byte){

	//生成椭圆曲线,  secp256r1 曲线。    比特币当中的曲线是secp256k1
	curve :=elliptic.P256()

	private,err :=ecdsa.GenerateKey(curve,rand.Reader)

	if err !=nil{

		fmt.Println("error")
	}
	pubkey :=append(private.PublicKey.X.Bytes(),private.PublicKey.Y.Bytes()...)
	return *private,pubkey
}

func main(){

    //调用函数生成公钥
    privatekey,public :=newKeyPair()

    //打印私钥  曲线上的x点
    fmt.Printf("%x\n",privatekey.D.Bytes())

    //打印公钥， 曲线上的x点和y点
    fmt.Printf("%x",public)
}
