package reflect

import sitter "github.com/smacker/go-tree-sitter"

type Enum struct {
	root     *sitter.Node
	document *Document
	parent   TypeElement
}

func NewEnum(node *sitter.Node, d *Document, parent TypeElement) (*Enum, error) {

	n, err := d.NewRootType(node)

	return &Enum{
		root:     n,
		document: d,
		parent:   parent,
	}, err
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
