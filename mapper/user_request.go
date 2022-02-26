package mapper

type DataRequest struct {
	Name string `json:"name" validation:"required"`
	Age  int  `json:"age" validation:"required_notNull"`
	MobileNuber int `json:"mobNo" validation:"required"`
	Message string `json:"msg" validation:"required_notNull"`
}

type DataResponse struct {
	ReqID int `json:"reqId"`
	ReqStatus ReqStatus `json:"reqStatus"`
}
type ReqStatus struct {
	Status string `json:"status"`
	ErrorCode string `json:"error"`
	ErrorMsg string `json:"msg"`
}
