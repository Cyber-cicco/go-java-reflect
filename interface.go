package reflect

import (
	"errors"

	sitter "github.com/smacker/go-tree-sitter"
)

type Interface struct {
	root     *sitter.Node
	document *Document
	parent   TypeElement
}

func NewInterface(node *sitter.Node, d *Document, parent TypeElement) (*Interface, error) {

    if node.Type() != "interface_declaration" {
        return nil, errors.New("Unexpected node type for interface declaration : " + node.Type())
    }

	return &Interface{
		root:     node,
		document: d,
		parent:   parent,
	}, nil
}

func (i *Interface) GetDeclaredName() string {
	return i.root.ChildByFieldName("name").Content(i.document.content)
}

func (i *Interface) GetDocument() *Document {
	return i.document
}

func (i *Interface) GetParent() TypeElement {
	return i.parent
}
