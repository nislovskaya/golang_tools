package response

import (
	jsoniter "github.com/json-iterator/go"
	errors "github.com/nislovskaya/golang_tools/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

type response struct {
	writer http.ResponseWriter
	logger *logrus.Entry
}

type Response interface {
	Ok(data interface{})
	Created(data interface{})
	NoContent()

	BadRequest(message string)
	Unauthorized(message string)
	Forbidden(message string)
	NotFound(message string)
	Conflict(message string)

	InternalServerError(message string)
}

func New(w http.ResponseWriter, logger *logrus.Entry) Response {
	return &response{
		writer: w,
		logger: logger,
	}
}

func (r *response) Ok(data interface{}) {
	r.json(http.StatusOK, data)
}

func (r *response) Created(data interface{}) {
	r.json(http.StatusCreated, data)
}

func (r *response) NoContent() {
	r.json(http.StatusNoContent, nil)
}

func (r *response) BadRequest(message string) {
	r.json(http.StatusBadRequest, &errors.Error{
		Code:    http.StatusBadRequest,
		Message: message,
	})
}

func (r *response) Unauthorized(message string) {
	r.json(http.StatusUnauthorized, &errors.Error{
		Code:    http.StatusUnauthorized,
		Message: message,
	})
}

func (r *response) Forbidden(message string) {
	r.json(http.StatusForbidden, &errors.Error{
		Code:    http.StatusForbidden,
		Message: message,
	})
}

func (r *response) NotFound(message string) {
	r.json(http.StatusNotFound, &errors.Error{
		Code:    http.StatusNotFound,
		Message: message,
	})
}

func (r *response) Conflict(message string) {
	r.json(http.StatusConflict, &errors.Error{
		Code:    http.StatusConflict,
		Message: message,
	})
}

func (r *response) InternalServerError(message string) {
	r.json(http.StatusInternalServerError, &errors.Error{
		Code:    http.StatusInternalServerError,
		Message: message,
	})
}

func (r *response) json(statusCode int, data interface{}) {
	r.writer.Header().Set("Content-Type", "application/json")
	r.writer.WriteHeader(statusCode)
	if err := jsoniter.NewEncoder(r.writer).Encode(data); err != nil {
		r.logger.WithError(err).Error("failed to encode json")
	}
}
