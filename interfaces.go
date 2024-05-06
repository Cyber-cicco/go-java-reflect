package reflect

import sitter "github.com/smacker/go-tree-sitter"

type Annotated interface {
    AnnotationSelector(query string) *Annotation
    GetAnnotations() []*Annotation
}

type Root interface {
    GetRoot() *sitter.Node
}

//Expression that gets back a value of a certain type
//
//Implementers:
//
// - ExpressionStatement
// - ReturnStatement
// - LambdaExpression
// - Method Invocation
//
type TypeProvider interface {
    GetType() *Type
}

//Expression that takes arguments.
//
//Implementers:
//
// - Method
// - Constructor
//
type Argumented interface {
    GetArguments() []*Argument
    ArgumentSelector() *Argument
    ArgumentSelectorAll() []*Argument
}

//Expression that has constructors
//
//Implementers:
//
// - Class
// - Record
// - Enum
//
type Constructed interface {
    GetConstructors() []*Constructor
    ConstructorSelector() *Constructor
    ConstructorSelectorAll() []*Constructor
}

//Expression that has fields
//
//Implementers:
//
// - Class
// - Record
// - Enum
// - Interface
//
type Methoded interface {
    GetMethods() []*Method
    MethodSelector() *Method
    MethodSelectorAll() []*Method
}

//Expression that has fields
//
//Implementers:
//
// - Class
// - Record
// - Enum
// - Annotation
//
type Fielded interface {
    GetFields() []*Field
    FieldSelector() *Field
    FieldSelectorAll() []*Field
}
