package reflect

import sitter "github.com/smacker/go-tree-sitter"

type Field struct {
    root *sitter.Node
}

//Gets the type of the field
func (f* Field) GetType() *Type {
    return &Type{}
}

//Returns the name of the field
func (f *Field) GetName() string {
    return ""
}

func (f *Field) Final() bool {
    return false
}

func (f *Field) Static() bool {
    return false
}
func (f *Field) GetRoot() *sitter.Node {
    return f.root
}

func (f * Field) GetConstructor() *TypeProvider {
    return nil
}
