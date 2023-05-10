package response

type ApiResponseModel struct {
	Code int         `json:"code"`
	Mssg string      `json:"mssg"`
	Data interface{} `json:"data"`
}
