package buz_code

type Code uint

const CODE_OK Code = 0
const (
	CODE_SERVICE_NETWORK_ERROR Code = iota + 5000 //内网服务请求错误 1
	CODE_AUTH_FAILED                              //登录态校验失败 2
	CODE_MOD_EXPIRED                              //店铺功能包过期 3
	CODE_MOD_UNBOUGHT                             // 店铺功能包未购买（一次都没买过）4
	CODE_UNAUTHORIZED                             //用户无权限访问该功能包 5
	CODE_SERVER_ERROR                             //服务器内部错误 6
	CODE_INVALID_ARGS                             //参数错误 7
	CODE_NO_COOKIE                                //http头部缺少所需cookie
)

const CODE_IDIOT = 250 // 所有无法识别的code都返回250，代表下游是白痴
