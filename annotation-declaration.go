package reflect

import (
	"errors"

	sitter "github.com/smacker/go-tree-sitter"
)

type AnnotationDeclaration struct {
	root     *sitter.Node
	document *Document
    parent *Class
}

func NewAnnotation(node *sitter.Node, d *Document, parent *Class) (*AnnotationDeclaration, error) {

	if node.Type() != "annotation_type_declaration" {
        return nil, errors.New("Unexpected type for enum declaration : " + node.Type())
    }

	return &AnnotationDeclaration{
		root:     node,
		document: d,
        parent: parent,
	}, nil
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
