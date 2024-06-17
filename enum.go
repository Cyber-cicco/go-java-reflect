package reflect

import (
	"errors"

	sitter "github.com/smacker/go-tree-sitter"
)

type Enum struct {
	root     *sitter.Node
	document *Document
	parent   TypeElement
}

func NewEnum(node *sitter.Node, d *Document, parent TypeElement) (*Enum, error) {

	if node.Type() != "enum_declaration" {
        return nil, errors.New("Unexpected type for enum declaration : " + node.Type())
    }

	return &Enum{
		root:     node,
		document: d,
		parent:   parent,
	}, nil
}

func (e *Enum) GetDeclaredName() string {
	return e.root.ChildByFieldName("name").Content(e.document.content)
}

func (e *Enum) GetDocument() *Document {
	return e.document
}

func (e *Enum) GetParent() TypeElement {
    return e.parent
}
