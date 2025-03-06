package pb

// 注册
type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 登陆
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 登陆响应
type LoginResp struct {
	Token    string `json:"token"`
	Username string `json:"username"`
}

// pub
type PubReq struct {
	Message string `json:"message"`
}
