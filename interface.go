package reflect

import (
	sitter "github.com/smacker/go-tree-sitter"
)

type Interface struct {
    root *sitter.Node
    document *Document
}

func NewInterface(node *sitter.Node, d *Document) (*Interface, error) {

    root, err := d.NewRootType(node)

	return &Interface{
		root:     root,
		document: d,
	}, err
}

func (i *Interface) GetDeclaredName() string {
    return i.root.ChildByFieldName("name").Content(i.document.content)
}

func (i *Interface) GetDocument() *Document {
    return i.document
}
