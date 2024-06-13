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
    root *sitter.Node
    document *Document
    scope *sitter.Node
}

// Create an import by checking if the root node is indeed an import
// declaration
func NewImport(root *sitter.Node, d *Document) (*Import, error) {
    if (root.Type() != "import_declaration") {
        return nil, errors.New("import node needs to be of type 'import_declaration', '" + root.Type() + "' was given")
    }
    return &Import{
        root: root,
        document: d,
    }, nil
}

func (i *Import) getMainScope() *sitter.Node {
    return querier.GetFirstMatch(i.root, func(n *sitter.Node) bool {
        return n.Type() == "scoped_identifier"
    })
}

func (i *Import) ToString() string {
    return i.getMainScope().Content(i.document.content)
}

