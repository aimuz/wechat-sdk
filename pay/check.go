package pay

import (
	"strconv"
	"github.com/mailbaoer/wechat-sdk/utils"
)

type (
	WaxPayNotifyReq struct {
		AppID         string  `xml:"appid"`
		MchID         string  `xml:"mch_id"`
		BankType      string  `xml:"bank_type"`
		CashFee       float64 `xml:"cash_fee"`
		FeeType       string  `xml:"fee_type"`
		IsSubscribe   string  `xml:"is_subscribe"`

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
)

/**
 * 微信通知验证
 */

 // 微信小程序待验证参数
 func WaxVerifyParams(req WaxPayNotifyReq) ( map[string]string){
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

// 微信小程序支付签名验证
func WaxpayVerifySign(verifyParams map[string]string, signKey string, sign string) bool {
    signCalc,_ := utils.GenWeChatPaySign(verifyParams, signKey)
    if sign == signCalc {
        return true
	}
    return false
}