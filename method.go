package reflect

import (
	"errors"

	"github.com/Cyber-cicco/tree-sitter-query-builder/querier"
	sitter "github.com/smacker/go-tree-sitter"
)

type Method struct {
    root *sitter.Node
    parent TypeElement
    name string
    _type *Type
    diagnostics []error
    document *Document
    parameters []*Parameter
}

//Argument of a given method
type Argument struct {
    root *sitter.Node
}

func (m *Method) GetType() *Type {
    return m._type
}

func NewMethod(node *sitter.Node, parent TypeElement) (*Method, error) {

    _type := node.ChildByFieldName("type")
    paramsNode := node.ChildByFieldName("parameters")
    name := node.ChildByFieldName("name")
    document := parent.GetDocument()
    var paramNodes []*sitter.Node

    if _type  == nil {
        return nil, errors.New("Missing type declaration on method")
    }

    if name  == nil {
        return nil, errors.New("Missing name declaration on method")
    }

    if paramsNode  == nil {
        return nil, errors.New("Missing name declaration on method")
    }

    paramNodes = querier.GetChildrenMatching(node, func(n *sitter.Node) bool {
        return n.Type() == "formal_parameter"
    }, paramNodes)
    method:= &Method{ 
        root: node,
        parent: parent,
        name: name.Content(document.content),
        document: document,
    }
    params := make([]*Parameter, len(paramNodes))

    for i, n := range paramNodes {
        newParam, err := NewParameter()
    }

    newType, err := NewType(_type, method, document)

    if err != nil {
        return nil, err
    }

    method._type = newType
    return method, nil
}


//Search an argument of a method by it's type and it's name
//
//If you pass nil in either of the args, it doesn't try to match it.
//
//Returns an argument and true if there is match, and an argument with
//an nil Tree Sitter Node and false if it wasn't found
func (m *Method) ArgumentSelector(searchedType *Type, name *string) (Argument, bool) {
    return Argument{}, false
}

//Gets the type of an argument
func (a *Argument) GetType() Type {
    return Type{}
}

//Gets the name of an argument
func (a *Argument) GetName() string {
    return ""
}
