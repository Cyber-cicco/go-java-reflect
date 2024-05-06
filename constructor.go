package reflect

import sitter "github.com/smacker/go-tree-sitter"

type Constructor struct {
    root *sitter.Node
    document *Document
}


func (c *Constructor) GetArguments() []*Argument {
    return []*Argument{}
}
func (c *Constructor) ArgumentSelector() *Argument {
    return &Argument{}
}
func (c *Constructor) ArgumentSelectorAll() []*Argument {
    return []*Argument{}
}
