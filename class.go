package reflect

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// Represents a base class in a java project.
type Class struct {
    root *sitter.Node
    document *Document
    parent TypeElement
}

// Creates a new class from sitter node
func NewClass(node *sitter.Node, d *Document, parent *Class) (*Class, error) {

    root, err := d.NewRootType(node)

	return &Class{
		root:     root,
		document: d,
        parent: parent,
	}, err
}

func (c *Class) GetDocument() *Document {
    return c.document
}

func (c *Class) GetParent() TypeElement {
    return c.parent
}

func (c *Class) GetDeclaredName() string {
    return c.root.ChildByFieldName("name").Content(c.document.content)
}

func (c *Class) GetMethods() []*Method {
    return []*Method{}
}


//Search an annotation of a query that goes like this :
//
//Name[param1:value1?,..., paramn:valuen?]
//
//Returns an annotation and true if there is match, and an annotation with
//an nil Tree Sitter Node and false if it wasn't found
func (c *Class) AnnotationSelector(query string) (Annotation, bool) {
    return Annotation{}, false
}

func (c *Class) GetName() string {
    return "todo"
}

//Search an import of the file by matching a regex on it.
//
//Returns a Type and true if there is match, and a Type with
//a nil Tree Sitter Node and false if it wasn't found
func (d *Document) ImportSelector(query string)  {
    return 
}

func (d *Class) ConstructorSelector(query string) *Constructor {
    return &Constructor{}
}

//Finds the first method inside a class
//
//query looks as follows:
//<name>:<return_type?>[type1?,...,typen?]
//
//Returns a Type and true if there is match, and a Type with
//a nil Tree Sitter Node and false if it wasn't found
func (d *Document) MethodSelector(query string) *Method {
    return &Method{}
}

//Finds the first field a class
//
//query looks as follows:
//<name>:<return_type>
//
//Returns a Type and true if there is match, and a Type with
//a nil Tree Sitter Node and false if it wasn't found
func (d *Class) FieldSelector(query string) *Field {
    return &Field{}
}

