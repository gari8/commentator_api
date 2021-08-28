package parser

import (
	"commentator/domain"
	parserService "commentator/usecases/services/parser"
	"net/http"
)

type ParseFileInterctor struct {
	outputPort    ParseFileOutputPort
	parserService parserService.ParserService
}

func NewParseFileInterctor(
	outputPort ParseFileOutputPort,
	parserService parserService.ParserService) ParseFileInterctor {
	return ParseFileInterctor{
		outputPort:    outputPort,
		parserService: parserService,
	}
}

func (i ParseFileInterctor) ParseFile(input ParseFileInput) domain.Response {
	body, err := i.parserService.ParseFile(input.FileByte)
	if err != nil {
		return domain.Response{
			Code: http.StatusInternalServerError,
			Err:  err,
		}
	}
	return i.outputPort.CreateResponse(body)
}
