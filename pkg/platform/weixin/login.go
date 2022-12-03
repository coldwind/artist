package weixin

type AuthData struct {
	Errcode    int    `json:"errcode"`
	Openid     string `json:"openid"`
	Unionid    string `json:"unionid"`
	SessionKey string `json:"session_key"`
	Errmsg     string `json:"errmsg"`
}

func GetUserInfo(code string) *AuthData {
	// url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	// api := fmt.Sprintf(url, icfg.YamlConfig.WXAppid, icfg.YamlConfig.WXAppSecret, code)
	// data := utils.HttpGet(api)

	authData := new(AuthData)
	// json.Unmarshal(data, authData)

	return authData
}
