# 如何安全创建 1万个以太坊钱包和私钥? 撸空投必备, 批量操作第一步! 工具源码解析
# https://www.youtube.com/watch?v=h2CMvlqma-w
# https://eth.antcave.club/1000#heading-5

from eth_account import Account
from web3 import Web3
import csv 

# 安装： pip3 install eth_account
#       pip3 install web3

def createNewETHWallet():
    wallets = []

    for id in range(1000):
        # 添加一些随机性
        account = Account.create('Random  Seed'+str(id))

        # 私钥
        privateKey = account._key_obj

        # 公钥
        publicKey = privateKey.public_key

        # 地址
        address = publicKey.to_checksum_address()

        wallet = {
            "id": id,
            "address": address,
            "privateKey": privateKey,
            "publicKey": publicKey
        }
        wallets.append(wallet.values())

    return wallets


def saveETHWallet(jsonData):
    with open('wallets.csv', 'w', newline='') as csv_file:
        csv_writer = csv.writer(csv_file)
        csv_writer.writerow(["序号", "钱包地址", "私钥", "公钥"])
        csv_writer.writerows(jsonData)


if __name__ == "__main__":

    print("---- 开始创建钱包 ----")
    # 创建 1000 个随机钱包
    wallets = createNewETHWallet()

    # 保存至 csv 文件
    saveETHWallet(wallets)
    print("---- 完成 ----")