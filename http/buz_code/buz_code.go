package buz_code

type Code uint8

const (
	CODE_OK                    Code = iota //ok 0
	CODE_SERVICE_NETWORK_ERROR             //内网服务请求错误
	CODE_AUTH_FAILED                       //登录态校验失败 1
	CODE_MOD_EXPIRED                       //店铺功能包过期 2
	CODE_MOD_UNBOUGHT                      // 店铺功能包未购买（一次都没买过）3
	CODE_UNAUTHORIZED                      //用户无权限访问该功能包 4
	CODE_SERVER_ERROR                      //服务器内部错误 5
	CODE_INVALID_ARGS                      //参数错误 6
	CODE_NO_COOKIE                         //http头部缺少所需cookie
)

const CODE_IDIOT = 250 // 所有无法识别的code都返回250，代表下游是白痴
