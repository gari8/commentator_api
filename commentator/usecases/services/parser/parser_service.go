package parser

type Repository interface {
	ParseFile(body []byte) ([]byte, error)
}

type ParserService struct {
	repo Repository
}

func NewParserService(repo Repository) ParserService {
	return ParserService{repo: repo}
}

func (s ParserService) ParseFile(body []byte) ([]byte, error) {
	return s.repo.ParseFile(body)
}
