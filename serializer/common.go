package serializer

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	ReqCode string      `json:"reqCode"`
}
