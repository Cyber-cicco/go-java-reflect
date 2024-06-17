package reflect

import (
	"errors"
	"strings"
)

type Scope struct {
	identifier string
	classes    []ClassProvider
	parent     *Scope
	children   map[string]*Scope
}

type Project struct {
	scopes map[string]*Scope
}

func (p *Project) GetScope(identifier string, classProvider ClassProvider) (*Scope, error) {
    //mut
	var curr *Scope

	if identifier == "" {
		return nil, errors.New("Empty string passed as a scope")
	}

	//mut
	ids := strings.Split(identifier, ".")

	for len(ids) > 0 {
        var ok bool
		curr, ok = p.scopes[ids[0]]

		if !ok {
			//mutation
			curr = p.addIdsToScope(ids, curr)
			break
		}

		//mutation
		ids = ids[1:]
	}

    if classProvider != nil {
        curr.classes = append(curr.classes, classProvider)
    }

	return curr, nil
}

// Unpure function
//
// Takes a list of scope identifiers, and a current node
// For every identifier, creates a scope, and adds the following
// idenfier as it's child.
func (p *Project) addIdsToScope(ids []string, curr *Scope) *Scope {

	for len(ids) > 0 {
		newScope := newScope(curr, ids[0])

		if curr != nil {
			curr.children[newScope.identifier] = newScope
		} else {
			p.scopes[newScope.identifier] = newScope
		}

		//mutation
		curr = newScope
		//mutation
		ids = ids[1:]
	}
    return curr
}

func NewProject() *Project {
    scopes := make(map[string]*Scope)
    return &Project{
        scopes: scopes,
    }
}

func (s *Scope) ToString() string {
	return s.identifier
}

func (s *Scope) GetFullScope() string {
    buffer := []string{}
    //mut
    curr := s
    for curr != nil {
        buffer = append([]string{curr.identifier}, buffer...)
        //mutation
        curr = curr.parent
    }
    return strings.Join(buffer, ".")
}

func newScope(parent *Scope, identifier string) *Scope {
	return &Scope{
		identifier: identifier,
		parent:     parent,
		classes:    []ClassProvider{},
		children:   map[string]*Scope{},
	}
}
