package pay

import (
	"encoding/xml"
	"github.com/aimuz/wechat-sdk/utils"
	"github.com/aimuz/wechat-sdk/common"
	"errors"
	"fmt"
)

/*
redpack 红包
*/

type (
	RedPack interface {
		Send(payKey, certFile, keyFile, rootCaFile string) (*RedPackResp, error)
	}

	SendRedPackReq struct {
		XMLName      xml.Name `xml:"xml"`
		NonceStr     string   `xml:"nonce_str,omitempty" json:"nonce_str"`           // NonceStr 随机字符串
		Sign         string   `xml:"sign,omitempty" json:"sign"`                     // Sign 签名
		MchBillNo    string   `xml:"mch_billno,omitempty" json:"mch_billno"`         // MchBillNo 商户订单号
		MchID        string   `xml:"mch_id,omitempty" json:"mch_id"`                 // MchID 商户号
		WxAppID      string   `xml:"wxappid,omitempty" json:"wxappid"`               // WxAppID 公众账号APPID
		SendName     string   `xml:"send_name,omitempty" json:"send_name"`           // SendName 商户名称，发送者名称
		ReOpenID     string   `xml:"re_openid,omitempty" json:"re_openid"`           // ReOpenID 用户OpenID
		TotalAmount  int64    `xml:"total_amount,omitempty" json:"total_amount"`     // TotalAmount 发送金额大小
		TotalNum     int      `xml:"total_num,omitempty" json:"total_num"`           // TotalNum 发送数量
		Wishing      string   `xml:"wishing,omitempty" json:"wishing"`               // Wishing 红包祝福语
		ClientIP     string   `xml:"client_ip,omitempty" json:"client_ip"`           // ClientIP 服务器ID
		ActName      string   `xml:"act_name,omitempty" json:"act_name"`             // ActName 活动名称
		Remark       string   `xml:"remark,omitempty" json:"remark"`                 // Remark 备注
		SceneID      string   `xml:"scene_id,omitempty" json:"scene_id"`             // SceneID 场景ID
		RiskInfo     string   `xml:"risk_info,omitempty" json:"risk_info"`           // RiskInfo 活动信息
		ConsumeMchID string   `xml:"consume_mch_id,omitempty" json:"consume_mch_id"` // ConsumeMchID 资金授权商户号
	}

	RedPackResp struct {
		ReturnMsg   string `xml:"return_msg"`
		MchID       string `xml:"mch_id"`
		WxAppID     string `xml:"wxappid"`
		ReOpenid    string `xml:"re_openid"`
		TotalAmount string `xml:"total_amount"`
		ReturnCode  string `xml:"return_code"`
		ResultCode  string `xml:"result_code"`
		ErrCode     string `xml:"err_code"`
		ErrCodeDes  string `xml:"err_code_des"`
		MchBillNo   string `xml:"mch_billno"`

		// 以下发送裂变红包的时候才会用到
		SendTime   string `xml:"send_time"`
		SendListID string `xml:"send_listid"`
	}
)

// SendRedPack 简单调用方法
func (m *WePay) SendRedPack(totalAmount int64, openID, sendName, wishing, actName, remark string) (string, *RedPackResp, error) {
	mchBillNo := utils.GetBillNo(m.MchID, 28)
	req := &SendRedPackReq{
		NonceStr:    utils.RandomString(32),
		MchBillNo:   mchBillNo,
		MchID:       m.MchID,
		WxAppID:     m.AppID,
		SendName:    sendName,
		ReOpenID:    openID,
		TotalAmount: totalAmount,
		TotalNum:    1,
		Wishing:     wishing,
		ClientIP:    "123.123.123.123",
		ActName:     actName,
		Remark:      remark,
	}
	return m.sendRedPack(req)
}

//func (m *WePay) SendRedPackByMap( map[string]interface{}) (string, *RedPackResp, error) {
//
//}

// SendRedPackByStruct 自定义发送红包参数
func (m *WePay) SendRedPackByStruct(req *SendRedPackReq) (string, *RedPackResp, error) {
	return m.sendRedPack(req)
}

func (m *WePay) sendRedPack(req *SendRedPackReq) (string, *RedPackResp, error) {
	resp, err := req.Send(m.PayKey, m.CertFile, m.keyFile, m.RootCaFile)
	if err != nil {
		return req.MchBillNo, resp, err
	}

	err = resp.CheckErr()
	if err != nil {
		return req.MchBillNo, resp, err
	}

	return req.MchBillNo, resp, nil
}

// SendRedPack 发送普通红包
func (m *SendRedPackReq) Send(payKey string, certFile, keyFile, rootCaFile string) (*RedPackResp, error) {

	tMap, err := utils.Struct2Map(m)
	if err != nil {
		return nil, err
	}

	fmt.Println(tMap)

	m.Sign, err = utils.GenWeChatPaySign(tMap, payKey)
	if err != nil {
		return nil, err
	}
	fmt.Println(m.Sign)

	data, err := xml.Marshal(m)
	if err != nil {
		return nil, err
	}

	request, err := utils.NewCertRequest(certFile, keyFile, rootCaFile)
	if err != nil {
		return nil, err
	}

	body, err := request.NewRequest("POST", common.SendRedPackURL, data)
	if err != nil {
		return nil, err
	}

	resp := new(RedPackResp)
	err = xml.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CheckErr 检查微信是否返回错误信息
func (m *RedPackResp) CheckErr() error {
	if m.ResultCode != "SUCCESS" || m.ReturnCode != "SUCCESS" {
		return errors.New(fmt.Sprintf("[%s,%s]%s", m.ResultCode, m.ErrCode, m.ErrCodeDes))
	}

	return nil
}
