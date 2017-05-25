package sealer

import (
	"go/ast"
	"go/token"
	"io/ioutil"
)

type PositionHelper struct {
	set     *token.FileSet
	content map[string][]byte
}

func NewPositionHelper(set *token.FileSet) *PositionHelper {
	return &PositionHelper{
		set:     set,
		content: make(map[string][]byte),
	}
}

func (l *PositionHelper) getFileContent(filename string) ([]byte, error) {
	content, exists := l.content[filename]
	if !exists {
		var err error
		content, err = ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		l.content[filename] = content
	}
	return content, nil
}

func (l *PositionHelper) GetContentByPosition(node ast.Node) (string, error) {
	content, err := l.getFileContent(l.set.Position(node.Pos()).Filename)
	if err != nil {
		return "", err
	}

	return string(content[node.Pos()-1 : node.End()-1]), nil
}

func (l *PositionHelper) GetStringPosition(pos token.Pos) string {
	return l.set.Position(pos).String()
}
