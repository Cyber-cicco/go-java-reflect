package reflect

import sitter "github.com/smacker/go-tree-sitter"

type Type struct {
    Identifier string
    Scope *sitter.Node
}

func (t *Type) GetDeclaredName() string {
    return t.Identifier
}

func GetType(node *sitter.Node) *Type {
    return &Type{}
}

func getClassFromPackageName(node sitter.Node) string {
    return ""
}
