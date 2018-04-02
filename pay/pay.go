package pay

import (
	"fmt"
	"strings"
	"time"
	"wechat-sdk/utils"
)

type (
	WePay struct {
		AppId     string //微信应用APPId或小程序APPId
		MchId     string //商户号
		PayKey    string //支付密钥
		NotifyUrl string //回调地址
		TradeType string //小程序写"JSAPI",客户端写"APP"
		Body      string //商品描述 必填
	}

	PayRet struct {
		Timestamp string `json:"timestamp,omitempty"` // 时间戳
		NonceStr  string `json:"noncestr,omitempty"`  // 随机字符串
	}

	AppPayRet struct {
		PayRet

		AppId     string `json:"appid,omitempty"`     // 应用ID
		PartnerId string `json:"partnerid,omitempty"` // 微信支付分配的商户号
		PrepayId  string `json:"prepayid,omitempty"`  // 预支付交易会话ID
		Package   string `json:"package,omitempty"`   // 扩展字段 暂填写固定值Sign=WXPay
		Sign      string `json:"sign,omitempty"`      // 签名
	}

	WaxPayRet struct {
		PayRet

		Package  string `json:"package,omitempty"`  // 扩展字段 统一下单接口返回的 prepay_id 参数值，提交格式如：prepay_id=*
		SignType string `json:"signType,omitempty"` // 签名算法，暂支持 MD5
		PaySign  string `json:"paySign,omitempty"`  // 签名
	}
)

// App支付
func (m *WePay) AppPay(totalFee int) (results *AppPayRet, outTradeNo string, err error) {

	outTradeNo = utils.GetTradeNO(m.MchId)

	appUnifiedOrder := &AppUnifiedOrder{
		UnifiedOrder: UnifiedOrder{
			AppId:          m.AppId,
			MchId:          m.MchId,
			NotifyUrl:      m.NotifyUrl,
			TradeType:      m.TradeType,
			SpBillCreateIp: "123.123.123.123", // Ip
			OutTradeNo:     outTradeNo,
			TotalFee:       totalFee,
			Body:           m.Body,
			NonceStr:       utils.RandomNumString(16, 32),
		},
	}

	t, err := utils.Struct2Map(appUnifiedOrder)

	if err != nil {
		return results, outTradeNo, err
	}

	// 获取签名
	sign, err := utils.GenWeChatPaySign(t, m.PayKey)
	if err != nil {
		return results, outTradeNo, err
	}
	appUnifiedOrder.Sign = strings.ToUpper(sign)

	unifiedOrderResp, err := NewUnifiedOrder(appUnifiedOrder)

	if err != nil {
		return results, outTradeNo, err
	}
	results = &AppPayRet{
		PayRet: PayRet{
			Timestamp: fmt.Sprintf("%d", time.Now().Unix()),
			NonceStr:  unifiedOrderResp.NonceStr,
		},
		AppId:     unifiedOrderResp.AppId,
		PartnerId: unifiedOrderResp.MchId,
		PrepayId:  unifiedOrderResp.PrepayId,
		Package:   "Sign=WXPay",
	}

	r, err := utils.Struct2Map(results)

	if err != nil {
		return results, outTradeNo, err
	}

	sign, err = utils.GenWeChatPaySign(r, m.PayKey)
	if err != nil {
		return results, outTradeNo, err
	}
	results.Sign = strings.ToUpper(sign)

	return
}

// 小程序支付
func (m *WePay) WaxPay(totalFee int, openId string) (results *WaxPayRet, outTradeNo string, err error) {

	outTradeNo = utils.GetTradeNO(m.MchId)

	wxaUnifiedOrder := &WxaUnifiedOrder{
		UnifiedOrder: UnifiedOrder{
			AppId:          m.AppId,
			MchId:          m.MchId,
			NotifyUrl:      m.NotifyUrl,
			TradeType:      m.TradeType,
			SpBillCreateIp: "123.123.123.123", // Ip
			OutTradeNo:     outTradeNo,
			TotalFee:       totalFee,
			Body:           m.Body,
			NonceStr:       utils.RandomNumString(16, 32),
		},
		OpenId: openId,
	}

	t, err := utils.Struct2Map(wxaUnifiedOrder)

	if err != nil {
		return results, outTradeNo, err
	}

	// 获取签名
	sign, err := utils.GenWeChatPaySign(t, m.PayKey)
	if err != nil {
		return results, outTradeNo, err
	}
	wxaUnifiedOrder.Sign = strings.ToUpper(sign)

	unifiedOrderResp, err := NewUnifiedOrder(wxaUnifiedOrder)

	if err != nil {
		return results, outTradeNo, err
	}
	results = &WaxPayRet{
		PayRet: PayRet{
			Timestamp: fmt.Sprintf("%d", time.Now().Unix()),
			NonceStr:  unifiedOrderResp.NonceStr,
		},
		Package:  "prepay_id=" + unifiedOrderResp.PrepayId,
		SignType: "MD5",
	}

	r, err := utils.Struct2Map(results)

	if err != nil {
		return results, outTradeNo, err
	}

	sign, err = utils.GenWeChatPaySign(r, m.PayKey)
	if err != nil {
		return results, outTradeNo, err
	}
	results.PaySign = strings.ToUpper(sign)

	return
}

//// 公众号支付
//func (m *WePay) H5Pay(totalFee int, openId string) (results *WaxPayRet, outTradeNo string, err error) {
//
//	return m.WaxPay(totalFee, openId)
//}
//
//// 网页支付
//func (m *WePay) WebPay(totalFee int, openId string) (results *WaxPayRet, outTradeNo string, err error) {
//	return m.WaxPay(totalFee, openId)
//}
