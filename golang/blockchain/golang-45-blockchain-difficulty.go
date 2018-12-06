// 本文链接： https://dreamerjonson.com/2018/12/12/golang-45-blockchain-difficulty/ 
// golang[45]-区块链-挖矿困难度 2018-12-12

// ##比特币挖矿困难度

// 比特币的挖矿困难度 = 目标hash / 创世hash
// 比特币挖矿的计算

/*
 * 计算挖矿difficulty
 */
func CalculateDifficulty(strTargetHash string) string {
	strGeniusBlockHash := "00000000ffff0000000000000000000000000000000000000000000000000000" // 创世块编号

	var biGeniusHash big.Int
	var biTargetHash big.Int
	biGeniusHash.SetString(strGeniusBlockHash, 16)
	biTargetHash.SetString(strTargetHash, 16)

	difficulty := big.NewInt(0)
	difficulty.Div(&biGeniusHash, &biTargetHash)
	//fmt.Printf("%T \n" , difficulty)
	return fmt.Sprintf("%s", difficulty)
}

