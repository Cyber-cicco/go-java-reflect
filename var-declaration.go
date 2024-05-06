package reflect

import sitter "github.com/smacker/go-tree-sitter"



//Declaration of a variable inside a method
type VarDeclaration struct {
    root *sitter.Node
}

func (v *VarDeclaration) GetName() string {
    return ""
}

//Gets the type of a variable
func (v *VarDeclaration) GetType() *Type {
    return &Type{}
}
