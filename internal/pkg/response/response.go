package response

import (
	"encoding/json"
	"net/http"

	"go-template/internal/pkg/status"
	"go-template/internal/pkg/utils"
)

type (
	appError interface {
		Error() string
		Code() string
		Message() string

		Status() int
	}

	// BaseResponse is standard response of the app which include code, message, data and meta,...
	BaseResponse struct {
		status.Status
		Data interface{} `json:"data"`
		Meta string      `json:"meta"`
	}

	baseResponse BaseResponse

	// IDResponse is a response helper that has ID
	IDResponse struct {
		ID string `json:"id"`
	}
)

// JSON write status and JSON data to http response writer
func JSON(w http.ResponseWriter, status int, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(b)
}

// Error write error, status to http response writer with JSON format: {"code": status, "message": error}
func Error(w http.ResponseWriter, err error, status int) {
	if appError, ok := err.(appError); ok {
		JSON(w, appError.Status(), appError)
		return
	}
	JSON(w, status, map[string]interface{}{
		"code":    status,
		"message": err.Error(),
	})
}

// MarshalJSON implement encoding/json.Marshaler interface.
// It will automatically set AppError to Success if AppError is nil
func (rs BaseResponse) MarshalJSON() ([]byte, error) {
	var v = baseResponse(rs)
	if v.Status.Status() == 0 {
		v.Status = utils.Gen("").Success
	}
	return json.Marshal(v)
}
