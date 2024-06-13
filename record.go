package reflect

import (
	sitter "github.com/smacker/go-tree-sitter"
)

type Record struct {
    root *sitter.Node
    document *Document
}

func NewRecord(node *sitter.Node, d *Document) (*Record, error) {
    root, err := d.NewRootType(node)

	return &Record{
		root:     root,
		document: d,
	}, err
}

func (r *Record) GetDocument() *Document {
    return r.document
}

func (r *Record) GetDeclaredName() string {
    return r.root.ChildByFieldName("name").Content(r.document.content)
}
