# wechat-sdk
[![Go Report Card](https://goreportcard.com/badge/github.com/aimuz/wechat-sdk)](https://goreportcard.com/report/github.com/aimuz/wechat-sdk)
[![Build Status](https://travis-ci.org/aimuz/wechat-sdk.svg?branch=master)](https://travis-ci.org/aimuz/wechat-sdk)

最全最好用的微信SDK，支持APP，小程序，H5，Web登录支付，企业付款等功能


## 快速开始
以下是APP和小程序支付简单例子
```go
import "github.com/aimuz/wechat-sdk/pay"

wePay := &WePay{
	AppId:     "xxx",
	MchId:     "xxx",
	PayKey:    "xxx",
	NotifyUrl: "xxx",
	TradeType: "xxx", // APP支付填写`APP`,小程序支付填写`JSAPI`
	Body:      "xxx",
}

# APP支付
results, outTradeNo, err := wePay.AppPay(100) // 金额，以分为单位

# 小程序支付
results, outTradeNo, err := wePay.WaxPay(100, "open_id") // 金额，以分为单位；open_id为获取的用户的open_id
```

## 使用
### 小程序支付通知
```go
waxNotify := pay.WaxPayNotifyReq{}
ctx.ReadXML(&waxNotify)
verifyParams := pay.WaxVerifyParams(waxNotify)
valid := pay.WaxpayVerifySign(verifyParams, appKey, waxNotify.Sign) //appKey 为自己在微信支付后台设置的API密钥

resp := new(pay.WaxPayNotifyResp)

if valid {
	// 业务处理逻辑···
	resp.ReturnCode = "SUCCESS"
	resp.ReturnMsg = "OK"
} else {
	// 错误处理逻辑···
	resp.ReturnCode = "FAIL"
	resp.ReturnMsg = "Verify Failed"
}
```

### 发送普通红包
```go
wx := &WePay{
	AppID:      "xx",
	PayKey:     "xx",
	MchID:      "xx",
	TradeType:  "xx",
	CertFile:   "xx", // 证书路径
	keyFile:    "xx", // 证书秘钥路径
	RootCaFile: "xx", // 根证书路径
}

billNO, redPackResp, err := wx.SendRedPack(totalAmount, openID, sendName, wishing, actName, remark)

```

#### APP支付

##### APP简单使用

## 讨论

[issues](https://github.com/aimuz/wechat-sdk/issues)

[论坛](https://kezhan.io/t/aimuz.wechat-sdk)

## 支持功能

- [x] APP支付
- [x] APP登录
- [x] H5登录
- [x] 小程序登录
- [x] 小程序支付
- [ ] Web登录
- [ ] 公众号支付
- [ ] 扫码支付
- [ ] 刷卡支付
- [ ] 企业付款
- [x] 现金红包
   - [x] 发送红包
   - [ ] 裂变红包
