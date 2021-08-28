package parser

import (
	parserUsecase "commentator/usecases/usecases/parser"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gari8/shuttle"

	"github.com/aws/aws-lambda-go/events"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("error: No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("error: Non 200 Response found")
)

const (
	boundary = "---------------------------"
)

type ParserController interface {
	ParseFile(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

type parserController struct {
	parseFileInputPort parserUsecase.ParseFileInputPort
}

func NewParserController(
	parseFileInputPort parserUsecase.ParseFileInputPort) ParserController {
	return parserController{parseFileInputPort: parseFileInputPort}
}

func (c parserController) getParseFileParams(request events.APIGatewayProxyRequest) (parserUsecase.ParseFileInput, error) {
	sh := shuttle.New(request.Body, boundary)
	content := []byte(sh.Launch("file"))
	return parserUsecase.ParseFileInput{
		FileByte: content,
	}, nil
}

func (c parserController) ParseFile(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	resp, err := http.Get(DefaultHTTPGetAddress)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if resp.StatusCode != 200 {
		return events.APIGatewayProxyResponse{}, ErrNon200Response
	}

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if len(ip) == 0 {
		return events.APIGatewayProxyResponse{}, ErrNoIP
	}

	input, err := c.getParseFileParams(request)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	fmt.Println("reached")

	// 通信開始
	response := c.parseFileInputPort.ParseFile(input)

	responseBody, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(responseBody),
		StatusCode: 200,
	}, nil
}
