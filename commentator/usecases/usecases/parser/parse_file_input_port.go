package parser

import (
	"commentator/domain"
)

type ParseFileInput struct {
	FileByte []byte `json:"file_byte,omitempty"`
}

type ParseFileInputPort interface {
	ParseFile(input ParseFileInput) domain.Response
}
