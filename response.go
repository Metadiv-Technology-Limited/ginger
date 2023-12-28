package ginger

import "github.com/Metadiv-Technology-Limited/sql/models"

type IResponse interface {
	GetSuccess() bool
	SetSuccess(bool)
	GetDuration() int64
	SetDuration(int64)
	GetPagination() models.IPagination
	SetPagination(models.IPagination)
	GetError() IError
	SetError(IError)
	GetData() any
	SetData(any)
}

type Response struct {
	Success    bool               `json:"success"`
	Duration   int64              `json:"duration"`
	Pagination models.IPagination `json:"pagination"`
	Error      IError             `json:"error"`
	Data       any                `json:"data"`
}

func (r *Response) GetSuccess() bool {
	return r.Success
}

func (r *Response) SetSuccess(success bool) {
	r.Success = success
}

func (r *Response) GetDuration() int64 {
	return r.Duration
}

func (r *Response) SetDuration(duration int64) {
	r.Duration = duration
}

func (r *Response) GetPagination() models.IPagination {
	return r.Pagination
}

func (r *Response) SetPagination(pagination models.IPagination) {
	r.Pagination = pagination
}

func (r *Response) GetError() IError {
	return r.Error
}

func (r *Response) SetError(err IError) {
	r.Error = err
}

func (r *Response) GetData() any {
	return r.Data
}

func (r *Response) SetData(data any) {
	r.Data = data
}
