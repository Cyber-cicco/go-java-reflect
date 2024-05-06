package reflect

import sitter "github.com/smacker/go-tree-sitter"

type Document struct {
    Root *sitter.Node
}

func (d *Document) GetMethods() []*Method {
    return []*Method{}
}

//Search an annotation of a query that goes like this :
//
//Name[param1:value1?,..., paramn:valuen?]
//
//Returns an annotation and true if there is match, and an annotation with
//an nil Tree Sitter Node and false if it wasn't found
func (d *Document) AnnotationSelector(query string) (Annotation, bool) {
    return Annotation{}, false
}

//Search an import of the file by matching a regex on it.
//
//Returns a Type and true if there is match, and a Type with
//a nil Tree Sitter Node and false if it wasn't found
func (d *Document) ImportSelector(query string) *Type {
    return &Type{}
}

//Finds the first inside a class
//
//query looks as follows:
//<name>:<return_type?>[type1?,...,typen?]
//
//Returns a Type and true if there is match, and a Type with
//a nil Tree Sitter Node and false if it wasn't found
func (d *Document) MethodSelector(query string) *Method {
    return &Method{}
}

//Returns all imports object of a given java document
func (d *Document) GetImports() []Type {
    return []Type{}
}

