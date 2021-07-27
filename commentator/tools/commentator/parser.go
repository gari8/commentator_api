package commentator

import (
	"encoding/json"
)

type DataType string

const (
	Signature DataType = "Signature"
	Name DataType = "Name"
	Params DataType = "Params"
	Result DataType = "Result"
	Independent DataType = "Independent"
)

type (
	Parser struct {
		Content string
	}

	Data struct {
		Row int
		Content string
		Type []*DataType
	}

	Request struct {
		Content string `json:"content"`
	}

	Response struct {
		Lines []*Data
	}
)

func ConvertInputDataToStruct(inputs string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(inputs), &req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}

func (p *Parser) Exec() (Response, error) {
	body, err := ConvertInputDataToStruct(p.Content)
	if err != nil {
		return Response{}, err
	}
	an := Analysis{
		Src: body.Content,
	}
	an.Exec()
	var lines []*Data
	for range make([]int, 10) {
		var types []*DataType
		t := Independent
		types = append(types, &t)
		d := Data{
			Row: 1,
			Content: body.Content,
			Type: types,
		}
		lines = append(lines, &d)
	}
	return Response {
		lines,
	}, nil
}