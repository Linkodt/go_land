package message

const (
	LoginMesType 		= 	"LoginMes"
	LoginResMesType 	= 	"LoginResMes"
	RegisterMesType 	= 	"RegisterMes"
)

type Message struct{
	Type string  `json:"type"`
	Data string  `json:"data"`
}

type LoginMes struct {
	UserId int    		`json:"userid"`	 // 用户id
	UserPwd string 	 `json:"userpwd"`		// 用户密码
	Username string 	`json:"username"`	 // 用户名
}

type LoginResMes struct{
	Code int 	`json:"code"`	// 返回状态码500 表示该用户未注册 200表示登录成功
	Error string  `json:"error"`	// 返回错误信息
}

type RegisterMes struct{
	UserId int    		`json:"userid"`	 // 用户id
	UserPwd string 	 `json:"userpwd"`		// 用户密码
	// Username string 	`json:"username"`	 // 用户名
}