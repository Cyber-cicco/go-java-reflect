package reflect

import (
	"errors"

	sitter "github.com/smacker/go-tree-sitter"
)

type Parameter struct {
    root *sitter.Node
    parent Parameterized
    _type *Type
    identifier string
}

func NewParameter(root *sitter.Node, parent Parameterized, document *Document) (*Parameter, error) {

    if root.Type() != "formal_parameter" {
        return nil, errors.New("Trying to instantiate parameter from unexpected node : " + root.Type())
    }
    _type := root.ChildByFieldName("type")
    name := root.ChildByFieldName("name")

    param := &Parameter{
        root: root,
        parent: parent,
        identifier: name.Content(document.content),
    }

    newType, err := NewType(_type,param, document)

    if err != nil {
        return nil, err
    }

    param._type = newType

    return param, nil
}

func(p *Parameter) GetIdentifier() string {
    return p.identifier
}

func (p *Parameter) GetType() *Type {
    return p._type
}
