package common

const (
	UnifiedOrderUrl = "https://api.mch.weixin.qq.com/pay/unifiedorder" //微信统一下单
)

// https://open.weixin.qq.com/cgi-bin/showdocument?action=dir_list&t=resource/res_list&verify=1&id=open1419317853&token=&lang=zh_CN
const (
	//https://api.weixin.qq.com/sns/oauth2/access_token?appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code
	AccessTokenUrl      = "https://api.weixin.qq.com/sns/oauth2/access_token"  //code获取access_token
	RefreshTokenUrl     = "https://api.weixin.qq.com/sns/oauth2/refresh_token" //重新获取access_token
	UserInfoUrl         = "https://api.weixin.qq.com/sns/userinfo"             //通过access_token获取userInfo
	CheckAccessTokenUrl = "https://api.weixin.qq.com/sns/auth"                 //检验授权凭证（access_token）是否有效
)
