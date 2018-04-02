package login

import (
	"encoding/json"
	"errors"
	"net/url"
	"time"
	"wechat-sdk/common"
	"wechat-sdk/utils"
	"fmt"
)

type (
	WxConfig struct {
		AppId  string `json:"appid"`
		Secret string `json:"secret"`
	}

	WxAccessToken struct {
		//AppId        string `json:"app_id,omitempty"`
		AccessToken  string `json:"access_token,omitempty"`
		ExpiresIn    uint   `json:"expires_in,omitempty"`
		RefreshToken string `json:"refresh_token,omitempty"`
		OpenId       string `json:"openid,omitempty"`
		Scope        string `json:"scope,omitempty"`
		ErrCode      uint   `json:"errcode,omitempty"`
		ErrMsg       string `json:"errmsg,omitempty"`
		ExpiredAt    time.Time
	}

	WxUserInfo struct {
		OpenID     string `json:"openid,omitempty"`     // 授权用户唯一标识
		NickName   string `json:"nickname,omitempty"`   // 普通用户昵称
		Sex        uint32 `json:"sex,omitempty"`        // 普通用户性别，1为男性，2为女性
		Province   string `json:"province,omitempty"`   // 普通用户个人资料填写的省份
		City       string `json:"city,omitempty"`       // 普通用户个人资料填写的城市
		Country    string `json:"country,omitempty"`    // 国家，如中国为CN
		HeadImgUrl string `json:"headimgurl,omitempty"` // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空
		//Privilege  string `json:"privilege"`
		Privilege []string `json:"privilege,omitempty"` // 用户特权信息，json数组，如微信沃卡用户为（chinaunicom）
		UnionId   string   `json:"unionid,omitempty"`   // 普通用户的标识，对当前开发者帐号唯一
		ErrCode   uint     `json:"errcode,omitempty"`
		ErrMsg    string   `json:"errmsg,omitempty"`
	}
)

// 微信APP登录 直接登录获取用户信息
func (m *WxConfig) AppLogin(code string) (wxUserInfo *WxUserInfo, err error) {
	accessToken, err := m.GetWxAccessToken(code)

	if err != nil {
		return wxUserInfo, err
	}

	return accessToken.GetUserInfo()

}

// 微信小程序登录 直接登录获取用户信息
func (m *WxConfig) WexLogin() {

}

// 通过code获取AccessToken
func (m *WxConfig) GetWxAccessToken(code string) (accessToken *WxAccessToken, err error) {

	if code == "" {
		return accessToken, errors.New("getWxAccessToken error: code is null")
	}

	params := url.Values{
		"code":       []string{code},
		"grant_type": []string{"authorization_code"},
	}

	t, err := utils.Struct2Map(m)

	if err != nil {
		return accessToken, err
	}

	for k, v := range t {
		params.Set(k, v)
	}

	if err != nil {
		return accessToken, err
	}

	body, err := utils.NewRequest("GET", common.AccessTokenUrl, []byte(params.Encode()))

	//fmt.Println(string(body))

	if err != nil {
		return accessToken, err
	}
	err = json.Unmarshal(body, &accessToken)
	if err != nil {
		return accessToken, err
	}

	if accessToken.ErrMsg != "" {
		return accessToken, errors.New(accessToken.ErrMsg)
	}

	//accessToken.AppId = m.AppId

	return
}

// 获取用户资料
func (m *WxAccessToken) GetUserInfo() (wxUserInfo *WxUserInfo, err error) {
	if m.AccessToken == "" {
		return nil, errors.New("getWxUserInfo error: accessToken is null")
	}

	if m.OpenId == "" {
		return nil, errors.New("getWxUserInfo error: openID is null")
	}

	params := url.Values{
		"access_token": []string{m.AccessToken},
		"openid":       []string{m.OpenId},
	}

	body, err := utils.NewRequest("GET", common.UserInfoUrl, []byte(params.Encode()))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &wxUserInfo)
	if err != nil {
		return nil, err
	}

	if wxUserInfo.OpenID == "" {
		return wxUserInfo, errors.New(wxUserInfo.ErrMsg)
	}

	return
}

// TODO 重新获取AccessToken
func (m *WxAccessToken) GetRefreshToken(appid string) error {

	if appid == "" {
		return errors.New(common.ErrAppIdEmpty)
	}

	if m.RefreshToken == "" {
		return errors.New(common.ErrRefreshTokenEmpty)
	}

	if m.OpenId == "" {
		return errors.New(common.ErrOpenIdEmpty)
	}

	params := url.Values{
		"appid":         []string{appid},
		"grant_type":    []string{"refresh_token"},
		"refresh_token": []string{m.RefreshToken},
	}

	body, err := utils.NewRequest("GET", common.RefreshTokenUrl, []byte(params.Encode()))

	if err != nil {
		return err
	}

	fmt.Println(string(body))

	err = json.Unmarshal(body, &m)
	if err != nil {
		return err
	}

	if m.ErrMsg != "" {
		return errors.New(m.ErrMsg)
	}

	return nil
}

// TODO 校验AccessToken
func (m *WxAccessToken) CheckAccessToken() (ok bool, err error) {

	if m.AccessToken == "" {
		return ok, errors.New(common.ErrAccessTokenEmpty)
	}

	if m.OpenId == "" {
		return ok, errors.New(common.ErrOpenIdEmpty)
	}

	params := url.Values{
		"openid":       []string{m.OpenId},
		"access_token": []string{m.AccessToken},
	}
	body, err := utils.NewRequest("GET", common.CheckAccessTokenUrl, []byte(params.Encode()))

	if err != nil {
		return ok, err
	}

	result := struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}{}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return ok, err
	}

	if result.ErrMsg != "ok" {
		return ok, errors.New(result.ErrMsg)
	}

	ok = true

	return
}
