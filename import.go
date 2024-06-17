package reflect

import (
	"errors"

	"github.com/Cyber-cicco/tree-sitter-query-builder/querier"
	sitter "github.com/smacker/go-tree-sitter"
)

// An import contains only a node.
// It should be passed as a parameter to a
// function that can convert it into a Class.
type Import struct {
	root      *sitter.Node
	document  *Document
	static    bool
	className string
	scope     *Scope
}

// Create an import by checking if the root node is indeed an import
// declaration
func NewImport(root *sitter.Node, d *Document) (*Import, error) {

    var scope *Scope = nil
    var err error
    var className string


	if root.Type() != "import_declaration" {
		return nil, errors.New("import node needs to be of type 'import_declaration', '" + root.Type() + "' was given")
	}

    error := querier.QuerySelector(root, "ERROR")

    if error != nil {
        return nil, errors.New("ERROR was found in import and shouldn't be used as an adress")
    }

	scopeNode := querier.GetFirstMatch(root, func(n *sitter.Node) bool {
		return n.Type() == "scoped_identifier"
	})
    scopeIdentifier := scopeNode.ChildByFieldName("scope").Content(d.content)

    if scopeNode != nil {
        className = scopeNode.ChildByFieldName("name").Content(d.content)
        scope, err = d.project.GetScope(scopeIdentifier, d)

        if err != nil {
            return nil, err
        }
    } else {
        className = querier.QuerySelector(root, "identifier").Content(d.content)
    }


    imp := &Import{
		root:     root,
		document: d,
		scope:    scope,
        className: className,
	}
    return imp, nil
}

func (i *Import) ToString() string {
	return i.scope.GetFullScope() + "." + i.className
}

func (i *Import) GetClassName() string {
	return i.className
}
