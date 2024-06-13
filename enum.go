package reflect

import sitter "github.com/smacker/go-tree-sitter"

type Enum struct {
    root *sitter.Node
    document *Document
}

func NewEnum(node *sitter.Node, d *Document) (*Enum, error) {

    n, err := d.NewRootType(node)
    
    return &Enum{
        root: n,
        document: d,
    }, err
}

func (e *Enum) GetDeclaredName() string {
    return e.root.ChildByFieldName("name").Content(e.document.content)
}

func (e *Enum) GetDocument() *Document {
    return e.document
}
