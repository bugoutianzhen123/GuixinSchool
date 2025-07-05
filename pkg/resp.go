package pkg

import "fmt"


type Resp struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data any `json:"data"`
}

var (
	SuccessResp = Resp{Code: 200, Msg: "成功"}
	ParamErrResp = Resp{Code: 10001, Msg: "参数错误"}
	AuthResp = Resp{Code: 10002, Msg: "认证失败"}
)


func WithMsg(resp Resp, msg string) Resp {
	resp.Msg = fmt.Sprintf("%s: %s", resp.Msg, msg)
	return resp
}

func WithData(resp Resp, data any) Resp {
	resp.Data = data
	return resp
}