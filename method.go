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

func (m *Method) GetType() *Type {
    return m._type
}

func NewMethod(node *sitter.Node, parent TypeElement) (*Method, error) {

    if node.Type() != "method_declaration" {
        return nil, errors.New("trying to create a method from a unappropriate node : " + node.Type())
    }

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

    paramNodes = querier.QuerySelectorAll(node, "formal_parameter", paramNodes)
    method:= &Method{ 
        root: node,
        parent: parent,
        name: name.Content(document.content),
        document: document,
    }
    params := make([]*Parameter, len(paramNodes))

    for i, n := range paramNodes {
        newParam, err := NewParameter(n, method, document)
        if err != nil {
            return nil, err
        }
        params[i] = newParam
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
func (m *Method) ArgumentSelector() *Parameter {
    return &Parameter{}
}

func (m *Method) ArgumentSelectorAll() []*Parameter {
    return []*Parameter{}
}

func (m *Method) GetArguments() []*Parameter {
    return m.parameters
}

