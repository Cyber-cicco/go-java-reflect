package reflect

import (
	"errors"

	"github.com/Cyber-cicco/java-reflect/utils"
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

func (d *Document) NewRootType(n *sitter.Node) (*sitter.Node, error) {
	nameNode := n.ChildByFieldName("name")

	if nameNode == nil {
		return nil, errors.New("Main class doesn't have a name field")
	}
	name := nameNode.Content(d.content) + ".java"

	if name != utils.GetFileNameFromUrl(d.path) {
		return nil, errors.New("Main class's name doesn't match the file's name")
	}

	return n, nil
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

func (d *Document) GetMainClass() (TypeElement, error) {

	match := querier.GetFirstMatch(d.root, func(node *sitter.Node) bool {
		return node.Type() == "class_declaration" ||
			node.Type() == "record_declaration"
	})

	if match == nil {
		return nil, errors.New("Java file doesn't declare a main Class")
	}

	switch match.Type() {
	case "class_declaration":
		return NewClass(match, d, nil)
	case "record_declaration":
		return NewRecord(match, d, nil)
	case "annotation_type_declaration":
		return NewAnnotation(match, d,nil)
	case "enum_declaration":
		return NewEnum(match, d, nil)
	case "interface_declaration":
		return NewInterface(match, d, nil)
	}
	panic("Switch statement should be exhaustive")

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
