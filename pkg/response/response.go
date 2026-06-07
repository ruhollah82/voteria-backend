package response

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Status int    `json:"status"`
	Data   any    `json:"data"`
	Code   string `json:"code"`
	Msg    string `json:"msg"`
}

type ErrorResponse struct {
	Status int               `json:"status"`
	Code   string            `json:"code"`
	Errors map[string]string `json:"errors"`
	Msg    string            `json:"msg"`
}

type JsonResponse struct{}

func NewJSONResponse() JsonResponse {
	return JsonResponse{}
}
func (j JsonResponse) Response(c *gin.Context, status int, responseCode string, msg string, data any, errs map[string]string) {

	if errs == nil {
		r := SuccessResponse{
			Status: status,
			Code:   responseCode,
			Data:   data,
			Msg:    msg,
		}
		c.JSON(200, r)
	} else {
		r := ErrorResponse{
			Status: status,
			Msg:    msg,
			Code:   responseCode,
			Errors: errs,
		}
		c.JSON(200, r)
	}

	slog.Info("request info",
		slog.Int("status", status),
		slog.Any("code", responseCode),
		slog.Any("data", data),
	)
}

func (j JsonResponse) ErrorResponse(c *gin.Context, status int, responseCode string, msg string, errs ...error) {
	if errs == nil {
		panic("ErrorResponse: errs is required in ErrorResponse")
	}
	if len(errs) >= 1 && errs[0] == nil {
		panic("ErrorResponse: first error in errs is nil")
	}

	temp := make(map[string]string)

	for _, err := range errs {
		slog.Error("kdjkfkf", "err", len(errs))
		splited := strings.Split(err.Error(), ": ")
		if len(splited) == 1 {
			temp["non_field"] = splited[0]
		} else {
			temp[splited[0]] = splited[1]
		}
	}

	j.Response(c, status, responseCode, msg, nil, temp)
}

func (j JsonResponse) ServerErrorResponse(c *gin.Context, err error) {
	slog.Error("SERVER ERROR", "error", err.Error())
	j.Response(c, http.StatusInternalServerError, "server_error", "Internal server error", nil, nil)
}

func (j JsonResponse) ServerOrUserErrorResponse(c *gin.Context, status int, msg string, serverErr error, userErrs []error, responseCode string) {
	if status == 0 {
		status = http.StatusBadRequest
	}
	if serverErr != nil {
		j.ServerErrorResponse(c, serverErr)
	} else if userErrs != nil {
		j.ErrorResponse(c, status, responseCode, msg, userErrs...)
	} else {
		panic("func ServerOrUserErrorResponse: serverErr and userErrs are nil")
	}
}

func (j JsonResponse) InvalidJSONErrorResponse(c *gin.Context, err error) {
	j.ErrorResponse(c, http.StatusBadRequest, "invalid_json", "invalid json", errors.New("invalid json"))
}
