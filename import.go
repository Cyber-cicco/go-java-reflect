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

	scopeNode := querier.GetFirstMatch(root, func(n *sitter.Node) bool {
		return n.Type() == "scoped_identifier"
	})

    if scopeNode != nil {
        className = scopeNode.ChildByFieldName("name").Content(d.content)
        scope, err = GetScope(scopeNode.ChildByFieldName("scope"), d)

        if err != nil {
            return nil, err
        }
    } else {
        className = querier.QuerySelector(root, "identifier").Content(d.content)
    }

	return &Import{
		root:     root,
		document: d,
		scope:    scope,
        className: className,
	}, nil
}

func (i *Import) ToString() string {
	return i.scope.ToString() + "." + i.className
}

func (i *Import) GetClassName() string {
	return i.className
}
