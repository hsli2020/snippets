// https://pibigstar.github.io/2019/09/15/golang对接支付宝支付/

// 网页扫码支付

package main

import (
	"fmt"
	"github.com/smartwalle/alipay"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

var (
	// appId
	appId = ""

	// 应用公钥
	aliPublicKey = ""

	// 应用私钥
	privateKey = ""
	client, _ = alipay.New(appId, aliPublicKey, privateKey, false)
)

func init() {
	client.LoadAppPublicCert("应用公钥证书")
	client.LoadAliPayPublicCert("支付宝公钥证书")
	client.LoadAliPayRootCert("支付宝根证书")
}

//网站扫码支付
func WebPageAlipay() {
	pay := alipay.AliPayTradePagePay{}

	// 支付宝回调地址（需要在支付宝后台配置）
	// 支付成功后，支付宝会发送一个POST消息到该地址
	pay.NotifyURL = "http://www.pibigstar/alipay"

	// 支付成功之后，浏览器将会重定向到该 URL
	pay.ReturnURL = "http://localhost:8088/return"

	//支付标题
	pay.Subject = "支付宝支付测试"

	//订单号，一个订单号只能支付一次
	pay.OutTradeNo = time.Now().String()

	//销售产品码，与支付宝签约的产品码名称,目前仅支持FAST_INSTANT_TRADE_PAY
	pay.ProductCode = "FAST_INSTANT_TRADE_PAY"

	//金额
	pay.TotalAmount = "0.01"

	url, err := client.TradePagePay(pay)
	if err != nil {
		fmt.Println(err)
	}
	payURL := url.String()

	//这个 payURL 即是用于支付的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	fmt.Println(payURL)

	//打开默认浏览器
	payURL = strings.Replace(payURL,"&","^&",-1)
	exec.Command("cmd", "/c", "start",payURL).Start()
}

//手机客户端支付
func WapAlipay() {
	pay := alipay.AliPayTradeWapPay{}

	// 支付成功之后，支付宝将会重定向到该 URL
	pay.ReturnURL = "http://localhost:8088/return"

	//支付标题
	pay.Subject = "支付宝支付测试"

	//订单号，一个订单号只能支付一次
	pay.OutTradeNo = time.Now().String()

	//商品code
	pay.ProductCode = time.Now().String()

	//金额
	pay.TotalAmount = "0.01"

	url, err := client.TradeWapPay(pay)
	if err != nil {
		fmt.Println(err)
	}
	payURL := url.String()

	//这个 payURL 即是用于支付的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	fmt.Println(payURL)

	//打开默认浏览器
	payURL = strings.Replace(payURL,"&","^&",-1)
	exec.Command("cmd", "/c", "start",payURL).Start()
}

func main() {
	//生成支付URL
	WapAlipay()

	//支付成功之后的返回URL页面
	http.HandleFunc("/return", func(rep http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		ok, err := client.VerifySign(req.Form)
		if err == nil && ok {
			rep.Write([]byte("支付成功"))
		}
	})

	//支付成功之后的通知页面
	http.HandleFunc("/alipay", func(rep http.ResponseWriter, req *http.Request) {
		var noti, _ = client.GetTradeNotification(req)
		if noti != nil {
			fmt.Println("支付成功")
			//修改订单状态。。。。
		} else {
			fmt.Println("支付失败")
		}
		alipay.AckNotification(rep) // 确认收到通知消息
	})
	
	fmt.Println("server start....")
	http.ListenAndServe(":8088", nil)
}
