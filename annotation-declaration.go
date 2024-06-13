package reflect

import sitter "github.com/smacker/go-tree-sitter"

type AnnotationDeclaration struct {
	root     *sitter.Node
	document *Document
}

func NewAnnotation(node *sitter.Node, d *Document) (*AnnotationDeclaration, error) {
	root, err := d.NewRootType(node)

	return &AnnotationDeclaration{
		root:     root,
		document: d,
	}, err
}

func (a *AnnotationDeclaration) GetDocument() *Document {
	return a.document
}

func (a *AnnotationDeclaration) GetDeclaredName() string {
	return a.root.ChildByFieldName("name").Content(a.document.content)
}
