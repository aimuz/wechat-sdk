package pay

import (
	"strconv"

	"encoding/xml"
	"github.com/aimuz/wechat-sdk/utils"
)

type (
	// WaxPayNotifyReq WaxPayNotifyReq
	WaxPayNotifyReq struct {
		AppID       string  `xml:"appid"`
		MchID       string  `xml:"mch_id"`
		BankType    string  `xml:"bank_type"`
		CashFee     float64 `xml:"cash_fee"`
		FeeType     string  `xml:"fee_type"`
		IsSubscribe string  `xml:"is_subscribe"`

		NonceStr      string  `xml:"nonce_str"`
		OpenID        string  `xml:"openid"`
		OutTradeNo    string  `xml:"out_trade_no"`
		ResultCode    string  `xml:"result_code"`
		ReturnCode    string  `xml:"return_code"`
		Sign          string  `xml:"sign"`
		TimeEnd       string  `xml:"time_end"`
		TotalFee      float64 `xml:"total_fee"`
		TradeType     string  `xml:"trade_type"`
		TransactionID string  `xml:"transaction_id"`
	}

	// WaxPayNotifyResp WaxPayNotifyResp
	WaxPayNotifyResp struct {
		ReturnCode string `xml:"return_code"`
		ReturnMsg  string `xml:"return_msg"`
	}

	WxPayNotifyReq struct {
		XMLName     xml.Name `xml:"xml" json:"-"`
		Appid       string   `xml:"appid" json:"appid,omitempty"`                 // APPID
		MchID       string   `xml:"mch_id" json:"mch_id,omitempty"`               // 商户号
		DeviceInfo  string   `xml:"device_info" json:"device_info,omitempty"`     // 微信支付分配的终端设备号，
		NonceStr    string   `xml:"nonce_str" json:"nonce_str,omitempty"`         // 随机字符串，不长于32位
		Sign        string   `xml:"sign" json:"-"`                                // 签名
		ResultCode  string   `xml:"result_code" json:"result_code,omitempty"`     // 业务状态
		ErrCode     string   `xml:"err_code" json:"err_code,omitempty"`           // 错误代码
		ErrCodeDes  string   `xml:"err_code_des" json:"err_code_des,omitempty"`   // 错误代码描述
		Openid      string   `xml:"openid" json:"openid,omitempty"`               // 用户OpenID
		IsSubscribe string   `xml:"is_subscribe" json:"is_subscribe,omitempty"`   // 是否关注公众号
		TradeType   string   `xml:"trade_type" json:"trade_type,omitempty"`       // 交易类型
		BankType    string   `xml:"bank_type" json:"bank_type,omitempty"`         // 付款银行
		TotalFee    string   `xml:"total_fee" json:"total_fee,omitempty"`         // 总金额
		FeeType     string   `xml:"fee_type" json:"fee_type,omitempty"`           // 货币种类
		CashFee     string   `xml:"cash_fee" json:"cash_fee,omitempty"`           // 现金支付金额
		CashFeeType string   `xml:"cash_fee_type" json:"cash_fee_type,omitempty"` // 现金支付类型
		CouponFee   string   `xml:"coupon_fee" json:"coupon_fee,omitempty"`       // 代币券金额
		CouponCount string   `xml:"coupon_count" json:"coupon_count,omitempty"`   // 代币券使用数量

		Coupon []WxPayNotifyReqCoupon `xml:"-" json:"-"` // 代币券数组

		TransactionID string `xml:"transaction_id" json:"transaction_id,omitempty"` // 微信支付ID
		OutTradeNo    string `xml:"out_trade_no" json:"out_trade_no,omitempty"`     // 商户支付ID
		Attach        string `xml:"attach" json:"attach,omitempty"`                 // 商家数据包
		TimeEnd       string `xml:"time_end" json:"time_end,omitempty"`             // 支付完成时间

		SubMchID   string `xml:"sub_mch_id" json:"sub_mch_id,omitempty"`   // 商户号
		ReturnCode string `xml:"return_code" json:"return_code,omitempty"` // 返回状态
	}

	WxPayNotifyReqCoupon struct {
		CouponID   string `xml:"coupon_id"`   // 代币券ID
		CouponType string `xml:"coupon_type"` // 代币券类型
		CouponFee  string `xml:"coupon_fee"`  // 单个代币券支付金额
	}
)

//func (m *WxPayNotifyReq) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
//	m.Coupon = make([]WxPayNotifyReqCoupon, 0)
//
//	for {
//		token, err := d.Token()
//		if err != nil {
//			return err
//		}
//
//		switch el := token.(type) {
//		//case xml.StartElement:
//		//	if el.Name.Local == "Image" {
//		//		item := new(Image)
//		//		if err = d.DecodeElement(item, &el); err != nil {
//		//			return err
//		//		}
//		//		mf.Image = *item
//		//	}
//		case xml.EndElement:
//			if el == start.End() {
//				return nil
//			}
//		}
//	}
//	return nil
//}

/**
 * 微信通知验证
 */

// WaxVerifyParams 微信小程序待验证参数
func WaxVerifyParams(req WaxPayNotifyReq) map[string]string {
	verifyParams := make(map[string]string)

	verifyParams["appid"] = req.AppID
	verifyParams["bank_type"] = req.BankType
	verifyParams["cash_fee"] = strconv.FormatFloat(req.CashFee, 'f', 0, 64)
	verifyParams["fee_type"] = req.FeeType
	verifyParams["is_subscribe"] = req.IsSubscribe
	verifyParams["mch_id"] = req.MchID
	verifyParams["nonce_str"] = req.NonceStr
	verifyParams["openid"] = req.OpenID
	verifyParams["out_trade_no"] = req.OutTradeNo
	verifyParams["result_code"] = req.ResultCode
	verifyParams["return_code"] = req.ReturnCode
	verifyParams["time_end"] = req.TimeEnd
	verifyParams["total_fee"] = strconv.FormatFloat(req.TotalFee, 'f', 0, 64)
	verifyParams["trade_type"] = req.TradeType
	verifyParams["transaction_id"] = req.TransactionID

	return verifyParams
}

// WaxpayVerifySign 微信小程序支付签名验证
func WaxpayVerifySign(verifyParams map[string]string, signKey string, sign string) bool {
	signCalc, _ := utils.GenWeChatPaySign(verifyParams, signKey)
	if sign == signCalc {
		return true
	}
	return false
}

// WxVerifyParams Struct2Map
func WxVerifyParams(req *WxPayNotifyReq) map[string]string {
	reqmap, _ := utils.Struct2Map(req)
	return reqmap
}

// VerifySignMd5 验证签名
func VerifySignMd5(verifyParams map[string]string, payKey string, sign string) bool {
	signCalc, _ := utils.GenWeChatPaySign(verifyParams, payKey)
	if sign == signCalc {
		return true
	}
	return false
}
