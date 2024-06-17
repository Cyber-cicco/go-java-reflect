package reflect

import sitter "github.com/smacker/go-tree-sitter"

type Constructor struct {
    root *sitter.Node
    document *Document
}


func (c *Constructor) GetArguments() []*Parameter {
    return []*Parameter{}
}
func (c *Constructor) ArgumentSelector() *Parameter {
    return &Parameter{}
}
func (c *Constructor) ArgumentSelectorAll() []*Parameter {
    return []*Parameter{}
}
