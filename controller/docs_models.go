package controller

// _ResponsePostList 帖子列表接口响应数据
type _ResponsePostList struct {
	Code    int           `json:"code"`    // 业务响应状态码
	Message string        `json:"message"` // 提示信息
	Data    []interface{} `json:"data"`    // 数据
}
