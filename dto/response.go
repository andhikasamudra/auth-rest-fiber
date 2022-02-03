package dto

const SuccessMessage = "Success"
const FailedMessage = "Success"

type JSONResponse struct {
	//Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
