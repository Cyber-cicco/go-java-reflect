package reflect

type Annotated interface {
    AnnotationSelector(query string) Annotation
    GetAnnotations() []Annotation
}

//Expression that gets back a value of a certain type
//implementers:
//
// - ExpressionStatement
// - ReturnStatement
// - LambdaExpression
// - Method Invocation
type TypeProvider interface {
    GetType() *Type
}

