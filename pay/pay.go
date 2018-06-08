package pay

import (
	"fmt"
	"github.com/aimuz/wechat-sdk/utils"
	"strings"
	"time"
)

type (
	// WePay 微信支付配置类
	WePay struct {
		AppID     string // 微信应用APPId或小程序APPId
		MchID     string // 商户号
		PayKey    string // 支付密钥
		NotifyURL string // 回调地址
		TradeType string // 小程序写"JSAPI",客户端写"APP"
		Body      string // 商品描述 必填
	}

	// AppRet 返回的基本内容
	AppRet struct {
		Timestamp string `json:"timestamp,omitempty"` // 时间戳
		NonceStr  string `json:"noncestr,omitempty"`  // 随机字符串
	}

	// AppPayRet 下单返回内容
	AppPayRet struct {
		AppRet

		AppID     string `json:"appid,omitempty"`     // 应用ID
		PartnerID string `json:"partnerid,omitempty"` // 微信支付分配的商户号
		PrepayID  string `json:"prepayid,omitempty"`  // 预支付交易会话ID
		Package   string `json:"package,omitempty"`   // 扩展字段 暂填写固定值Sign=WXPay
		Sign      string `json:"sign,omitempty"`      // 签名
	}

	// WaxRet 返回的基本内容
	WaxRet struct {
		Timestamp string `json:"timeStamp,omitempty"` // 时间戳
		NonceStr  string `json:"nonceStr,omitempty"`  // 随机字符串
	}

	// WaxPayRet 微信小程序下单返回内容
	WaxPayRet struct {
		WaxRet

		AppID    string `json:"appId,omitempty"`    // 应用ID
		Package  string `json:"package,omitempty"`  // 扩展字段 统一下单接口返回的 prepay_id 参数值，提交格式如：prepay_id=*
		SignType string `json:"signType,omitempty"` // 签名算法，暂支持 MD5
		PaySign  string `json:"paySign,omitempty"`  // 签名
	}
)

// AppPay App支付
func (m *WePay) AppPay(totalFee int) (results *AppPayRet, outTradeNo string, err error) {

	outTradeNo = utils.GetTradeNO(m.MchID)

	appUnifiedOrder := &AppUnifiedOrder{
		UnifiedOrder: UnifiedOrder{
			AppID:          m.AppID,
			MchID:          m.MchID,
			NotifyURL:      m.NotifyURL,
			TradeType:      m.TradeType,
			SpBillCreateIP: "123.123.123.123", // Ip
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
		AppRet: AppRet{
			Timestamp: fmt.Sprintf("%d", time.Now().Unix()),
			NonceStr:  unifiedOrderResp.NonceStr,
		},
		AppID:     unifiedOrderResp.AppID,
		PartnerID: unifiedOrderResp.MchID,
		PrepayID:  unifiedOrderResp.PrepayID,
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

// WaxPay 小程序支付
func (m *WePay) WaxPay(totalFee int, openID string) (results *WaxPayRet, outTradeNo string, err error) {

	outTradeNo = utils.GetTradeNO(m.MchID)

	wxaUnifiedOrder := &WxaUnifiedOrder{
		UnifiedOrder: UnifiedOrder{
			AppID:          m.AppID,
			MchID:          m.MchID,
			NotifyURL:      m.NotifyURL,
			TradeType:      m.TradeType,
			SpBillCreateIP: "123.123.123.123", // Ip
			OutTradeNo:     outTradeNo,
			TotalFee:       totalFee,
			Body:           m.Body,
			NonceStr:       utils.RandomNumString(16, 32),
		},
		OpenID: openID,
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
		WaxRet: WaxRet{
			Timestamp: fmt.Sprintf("%d", time.Now().Unix()),
			NonceStr:  unifiedOrderResp.NonceStr,
		},
		AppID:    m.AppID,
		Package:  "prepay_id=" + unifiedOrderResp.PrepayID,
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
