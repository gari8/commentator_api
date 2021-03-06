package parser

import (
	"commentator/domain"
	"net/http"
)

type ParseFilePresenter struct{}

func NewParseFilePresenter() ParseFilePresenter {
	return ParseFilePresenter{}
}

func (p ParseFilePresenter) CreateResponse(body string) domain.Response {
	return domain.Response{
		Code:   http.StatusOK,
		Object: body,
	}
}
