package reflect

import (
	"errors"

	"github.com/Cyber-cicco/tree-sitter-query-builder/querier"
	sitter "github.com/smacker/go-tree-sitter"
)

type Document struct {
	root    *sitter.Node
	path    string
	content []byte
}

func NewDocument(root *sitter.Node, path string, content []byte) *Document {
	return &Document{
		root:    root,
		path:    path,
		content: content,
	}
}

// Returns all imports object of a given java document
func (d *Document) GetImports() []*Import {
	nodes := []*sitter.Node{}
	nodes = querier.GetChildrenMatching(d.root, func(n *sitter.Node) bool {
		return n.Type() == "import_declaration"
	}, nodes)
	imps := make([]*Import, len(nodes))
	for i, n := range nodes {
		newImp, err := NewImport(n, d)
		if err != nil {
			panic("A node filtered to be an import declaration should be an import declaration")
		}
		imps[i] = newImp
	}
	return imps
}

func (d *Document) GetMainClass() (RootElement, error) {

	match := querier.GetFirstMatch(d.root, func(node *sitter.Node) bool {
		return node.Type() == "class_declaration"
	})

	if match == nil {
		return nil, errors.New("Java file doesn't declare a main Class")
	}

    return NewClass(match, d)
}

func (d *Document) GetPackage() (string, error) {
	match := querier.GetFirstMatch(d.root, func(node *sitter.Node) bool {
		return node.Type() == "package_declaration"
	})

	if match == nil {
		return "", errors.New("Class doesn't have a package declaration")
	}

	identifier := querier.GetFirstMatch(match, func(node *sitter.Node) bool {
		return node.Type() == "scoped_identifier"
	})

	if identifier == nil {
		return "", errors.New("found package declaration without scoped identifier")
	}

	return identifier.Content(d.content), nil

}
