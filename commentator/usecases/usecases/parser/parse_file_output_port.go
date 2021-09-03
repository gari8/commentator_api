package parser

import "commentator/domain"

type ParseFileOutputPort interface {
	CreateResponse(body string) domain.Response
}
