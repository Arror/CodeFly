package reader

import "github.com/samuel/go-thrift/parser"

// ReadThrift Read thrift file via input infomation
func ReadThrift(input string) (map[string]*parser.Thrift, error) {

	p := parser.Parser{}

	ts, _, err := p.ParseFile(input)

	if err != nil {
		return nil, err
	}

	return ts, nil
}
