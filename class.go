package reflect

import (
	"errors"

	sitter "github.com/smacker/go-tree-sitter"
)

// Represents a base class in a java project.
type Class struct {
	root       *sitter.Node
	document   *Document
	identifier string
	parent     TypeElement
}

// Creates a new class from sitter node
func NewClass(node *sitter.Node, d *Document, parent *Class) (*Class, error) {

    var err error

	if node.Type() != "class_declaration" {
		return nil, errors.New("Unexpected type for class declaration : " + node.Type())
	}

    identifier := node.ChildByFieldName("name").Content(d.content)

	return &Class{
		root:     node,
		document: d,
		parent:   parent,
        identifier: identifier,
	}, err
}

func (c *Class) GetDocument() *Document {
	return c.document
}

func (c *Class) GetParent() TypeElement {
	return c.parent
}

func (c *Class) GetDeclaredName() string {
	return c.identifier
}

func (c *Class) GetMethods() []*Method {
	return []*Method{}
}

// Search an annotation of a query that goes like this :
//
// Name[param1:value1?,..., paramn:valuen?]
//
// Returns an annotation and true if there is match, and an annotation with
// an nil Tree Sitter Node and false if it wasn't found
func (c *Class) AnnotationSelector(query string) (Annotation, bool) {
	return Annotation{}, false
}

func (c *Class) GetName() string {
	return "todo"
}

func (d *Class) ConstructorSelector(query string) *Constructor {
	return &Constructor{}
}

// Finds the first method inside a class
//
// query looks as follows:
// <name>:<return_type?>[type1?,...,typen?]
//
// Returns a Type and true if there is match, and a Type with
// a nil Tree Sitter Node and false if it wasn't found
func (d *Document) MethodSelector(query string) *Method {
	return &Method{}
}

// Finds the first field a class
//
// query looks as follows:
// <name>:<return_type>
//
// Returns a Type and true if there is match, and a Type with
// a nil Tree Sitter Node and false if it wasn't found
func (d *Class) FieldSelector(query string) *Field {
	return &Field{}
}
