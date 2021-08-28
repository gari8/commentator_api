package main

import (
	parserController "commentator/interfaces/controllers/parser"
	parserPresenter "commentator/interfaces/presenters/parser"
	parserRepository "commentator/interfaces/repositories/parser"
	parserService "commentator/usecases/services/parser"
	parserUsecase "commentator/usecases/usecases/parser"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	repo := parserRepository.NewParserRepository()
	service := parserService.NewParserService(repo)
	presenter := parserPresenter.NewParseFilePresenter()
	interactor := parserUsecase.NewParseFileInterctor(presenter, service)
	controller := parserController.NewParserController(interactor)
	lambda.Start(controller.ParseFile)
}
