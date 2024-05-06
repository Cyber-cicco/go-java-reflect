package reflect

import sitter "github.com/smacker/go-tree-sitter"

//Generic expression statement
type ExpressionStatement struct {
    root *sitter.Node
}

func (e *ExpressionStatement) GetType() *Type {
    return &Type{}
}

