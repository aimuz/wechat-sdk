# wechat-sdk
最全最好用的微信SDK，支持APP，小程序，H5，Web登录支付，企业付款等功能

## 快速开始
以下是APP支付简单例子
```go
wePay := &WePay{
	AppId:     "xxx",
	MchId:     "xxx",
	PayKey:    "xxx",
	NotifyUrl: "xxx",
	TradeType: "xxx",
	Body:      "xxx",
}

results, outTradeNo, err := wePay.AppPay(100))
```

## 使用


#### APP支付

##### APP简单使用

## 支持功能

- [x] APP支付
- [x] APP登录
- [x] H5登录
- [x] 小程序登录
- [ ] 小程序支付
- [ ] Web登录
- [ ] 公众号支付
- [ ] 扫码支付
- [ ] 刷卡支付
- [ ] 企业付款
- [ ] 现金红包
