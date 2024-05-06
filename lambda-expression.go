package reflect

import sitter "github.com/smacker/go-tree-sitter"


type LambdaExpression struct {
    root *sitter.Node
}

func (l *LambdaExpression) GetType() *Type {
    return &Type{}
}

