package reflect

import (
	"errors"

	"github.com/Cyber-cicco/tree-sitter-query-builder/querier"
	sitter "github.com/smacker/go-tree-sitter"
)

type Document struct {
	root      *sitter.Node
	imports   []*Import
    project   Project
	path      string
	mainClass TypeElement
	_package  *Scope
	content   []byte
}

func (p *Project) NewDocument(root *sitter.Node, path string, content []byte) (*Document, error) {

	var impNodes = []*sitter.Node{}
	var mainClass TypeElement
	var err error

	impNodes = querier.GetChildrenMatching(root, func(n *sitter.Node) bool {
		return n.Type() == "import_declaration"
	}, impNodes)

	imps := make([]*Import, len(impNodes))
	document := &Document{
		root:    root,
		imports: imps,
		path:    path,
		content: content,
	}

	for i, n := range impNodes {
		newImp, newErr := NewImport(n, document)
		err = newErr
		imps[i] = newImp
	}

	match := querier.GetFirstMatch(root, func(node *sitter.Node) bool {
		return node.Type() == "class_declaration" ||
			node.Type() == "record_declaration" ||
			node.Type() == "annotation_type_declaration" ||
			node.Type() == "enum_declaration" ||
			node.Type() == "interface_declaration"
	})

	if match == nil {
		err = errors.New("Java file doesn't declare a main Class")
	}

	switch match.Type() {
	case "class_declaration":
		mainClass, err = NewClass(match, document, nil)
	case "record_declaration":
		mainClass, err = NewRecord(match, document, nil)
	case "annotation_type_declaration":
		mainClass, err = NewAnnotation(match, document, nil)
	case "enum_declaration":
		mainClass, err = NewEnum(match, document, nil)
	case "interface_declaration":
		mainClass, err = NewInterface(match, document, nil)
	default:
		panic("Switch statement should be exhaustive")
	}

	document.mainClass = mainClass

	pck := querier.GetFirstMatch(root, func(node *sitter.Node) bool {
		return node.Type() == "package_declaration"
	})

	if pck == nil {
		return nil, errors.New("Class doesn't have a package declaration")
	}

	identifier := querier.GetFirstMatch(pck, func(node *sitter.Node) bool {
		return node.Type() == "scoped_identifier" || node.Type() == "identifier"
	})

	if identifier == nil {
		return nil, errors.New("found package declaration without identifier")
	}

	scope, err := p.GetScope(identifier.Content(document.content), document)
	document._package = scope

	return document, err
}

// Returns all imports object of a given java document
func (d *Document) GetImports() []*Import {
	return d.imports
}

func (d *Document) GetMainClass() TypeElement {
	return d.mainClass
}

func (d *Document) GetPackage() *Scope {
    return d._package
}

func (d *Document) GetImportByIdentifier(identifier string) (*Import, error) {

	var match *Import

	for _, imp := range d.imports {
		if imp.className == identifier {
			match = imp
			break
		}
	}

	if match == nil {
		return nil, errors.New("import not found for identifier " + identifier)
	}

	return match, nil
}
