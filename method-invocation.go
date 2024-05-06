package reflect

import sitter "github.com/smacker/go-tree-sitter"

//Expression representing a method invocaton
type MethodInvocation struct {
    root *sitter.Node
    document *Document
}

//Get the referenced Method
//First search in the methods of the document.
func (m *MethodInvocation) GetMethodReference() *Method {
    return &Method{}
}
