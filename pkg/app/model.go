package app

type Response struct {
	Code int         `json:"code"` // 状态码
	Msg  string      `json:"msg"`  // 消息
	Data interface{} `json:"data"` // 数据集
}

type Page struct {
	List      interface{} `json:"list"`      //数据集
	Total     int         `json:"total"`     //总数
	PageIndex int         `json:"pageIndex"` //当前页
	PageSize  int         `json:"pageSize"`  // 分页大侠
}
