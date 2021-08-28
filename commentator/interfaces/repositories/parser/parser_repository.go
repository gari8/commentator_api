package parser

type ParserRepository struct{}

func NewParserRepository() ParserRepository {
	return ParserRepository{}
}

func (r ParserRepository) ParseFile(body []byte) ([]byte, error) {
	return body, nil
}
