package reflect

import (
	"errors"

	"github.com/Cyber-cicco/java-reflect/config"
	"github.com/Cyber-cicco/tree-sitter-query-builder/querier"
	sitter "github.com/smacker/go-tree-sitter"
)

type Document struct {
	root        *sitter.Node
	imports     []*Import
	project     *Project
	path        string
	mainClass   TypeElement
	_package    *Scope
	content     []byte
	diagnostics map[int]error
}

func (p *Project) NewDocument(root *sitter.Node, path string, content []byte) (*Document, bool) {

	var impNodes = []*sitter.Node{}
	var mainClass TypeElement
    var errClass error
    var errPck error
    var identifier *sitter.Node
    var scope *Scope

	impNodes = querier.GetChildrenMatching(root, func(n *sitter.Node) bool {
		return n.Type() == "import_declaration"
	}, impNodes)

	imps := make([]*Import, len(impNodes))
	document := &Document{
		root:        root,
		imports:     imps,
		path:        path,
		content:     content,
		project:     p,
		diagnostics: make(map[int]error),
	}

	for i, n := range impNodes {
		newImp, err := NewImport(n, document)
		if err != nil {
            document.diagnostics[config.DIAG_CLASS] = err
		}
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
		errClass = errors.New("Java file doesn't declare a main Class")
	} else {
		switch match.Type() {
		case "class_declaration":
			mainClass, errClass = NewClass(match, document, nil)
		case "record_declaration":
			mainClass, errClass = NewRecord(match, document, nil)
		case "annotation_type_declaration":
			mainClass, errClass = NewAnnotation(match, document, nil)
		case "enum_declaration":
			mainClass, errClass = NewEnum(match, document, nil)
		case "interface_declaration":
			mainClass, errClass = NewInterface(match, document, nil)
		default:
			panic("Switch statement should be exhaustive")
		}
	}

    if errClass != nil {
        document.diagnostics[config.DIAG_CLASS] = errClass
    }

	document.mainClass = mainClass

	pck := querier.GetFirstMatch(root, func(node *sitter.Node) bool {
		return node.Type() == "package_declaration"
	})


	if pck == nil {
		document.diagnostics[config.DIAG_PCK] = errors.New("Class doesn't have a package declaration")
	} else {
        identifier = querier.GetFirstMatch(pck, func(node *sitter.Node) bool {
            return node.Type() == "scoped_identifier" || node.Type() == "identifier"
        })
    }

	if identifier == nil {
		document.diagnostics[config.DIAG_PCK] = errors.New("found package declaration without identifier")
	} else {
        scope, errPck = p.GetScope(identifier.Content(document.content), document)
    }


	if errPck != nil {
		document.diagnostics[config.DIAG_PCK] = errPck
	}

	document._package = scope

	return document, len(document.diagnostics) == 0
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
