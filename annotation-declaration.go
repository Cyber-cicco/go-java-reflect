package reflect

import sitter "github.com/smacker/go-tree-sitter"

type AnnotationDeclaration struct {
	root     *sitter.Node
	document *Document
    parent *Class
}

func NewAnnotation(node *sitter.Node, d *Document, parent *Class) (*AnnotationDeclaration, error) {
	root, err := d.NewRootType(node)

	return &AnnotationDeclaration{
		root:     root,
		document: d,
        parent: parent,
	}, err
}

func (a *AnnotationDeclaration) GetDocument() *Document {
	return a.document
}

func (a *AnnotationDeclaration) GetDeclaredName() string {
	return a.root.ChildByFieldName("name").Content(a.document.content)
}

func (a *AnnotationDeclaration) GetParent() TypeElement {
    return a.parent
}
