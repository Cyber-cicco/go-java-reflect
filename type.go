package reflect

import (
	"errors"

	"github.com/Cyber-cicco/tree-sitter-query-builder/querier"
	sitter "github.com/smacker/go-tree-sitter"
)

type Type struct {
	identifier string
	document   *Document
	scope      *sitter.Node
	parent    TypeProvider
	primitive  bool
}

func (t *Type) GetDeclaredName() string {
	return t.identifier
}

// Creates a type from a type argument.
// Finds it's scope from the import or from the prefix of the type.
func NewType(node *sitter.Node, parent TypeProvider, document *Document) (*Type, error) {

	var identifier string
	var scope *sitter.Node = nil

	if node.Type() != "type_identifier" &&
		node.Type() != "scoped_type_identifier" &&
		node.Type() != "integral_type" &&
		node.Type() != "floating_point_type" &&
		node.Type() != "boolean_type" {
		return nil, errors.New("non type_identifier passed as type")
	}

	if node.Type() == "scoped_type_identifier" {
        identifier = querier.BreadthFirstMatch(node, func(n *sitter.Node) bool {
            return node.Type() == "type_identifier"
        }).Content(document.content)
        scope = querier.GetFirstMatch(node, func(n *sitter.Node) bool {
            return node.Type() == "scoped_type_identifier"
        })
	} else {
		identifier = node.Content(document.content)
        if node.Type() != "integral_type" &&
		node.Type() != "floating_point_type" &&
		node.Type() != "boolean_type" {
            document.ImportSelector("")
        }
	}

	return &Type{
		identifier: identifier,
		document:   document,
        parent: parent,
	}, nil
}

func getClassFromPackageName(node sitter.Node) string {
	return ""
}
