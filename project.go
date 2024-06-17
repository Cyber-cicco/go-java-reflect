package reflect

import (
	"errors"
	"strings"
)

type null struct{}

type ScopeTree struct {
    root *Scope
}

type Scope struct {
	identifier string
	classes    *ClassProvider
    parent *Scope
    children map[*Scope]null
}

type Project struct {
    scopes map[string]*ScopeTree
}

func (p *Project) GetScope(identifier string, classProvider ClassProvider) (*Scope, error) {

    if identifier == "" {
        return nil, errors.New("Empty string passed as a scope")
    }

    ids := strings.Split(identifier, ".")
    exists := true
    scopeTree, ok := p.scopes[]

    for len(ids) > 0 {
        scope, ok := p.
    }

	return &Scope{
	}, nil
}

func (s *Scope) ToString() string {
    return s.identifier
}

func (s *Scope) AddChild()
