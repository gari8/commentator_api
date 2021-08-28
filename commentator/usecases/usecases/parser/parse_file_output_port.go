package parser

import "commentator/domain"

type ParseFileOutputPort interface {
	CreateResponse(body []byte) domain.Response
}
