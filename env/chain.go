package env

import "github.com/kode4food/ale/data"

type chainedNamespace struct {
	child  Namespace
	parent Namespace
}

func newChild(parent Namespace, n data.Name) Namespace {
	e := parent.Environment()
	return chain(parent, e.New(n))
}

func chain(parent Namespace, child Namespace) *chainedNamespace {
	return &chainedNamespace{
		parent: parent,
		child:  child,
	}
}

func (ns *chainedNamespace) Environment() *Environment {
	return ns.child.Environment()
}

func (ns *chainedNamespace) Domain() data.Name {
	return ns.child.Domain()
}

func (ns *chainedNamespace) Declared() []data.Name {
	return ns.child.Declared()
}

func (ns *chainedNamespace) Declare(n data.Name) Entry {
	return ns.child.Declare(n)
}

func (ns *chainedNamespace) Resolve(n data.Name) (Entry, bool) {
	if e, ok := ns.child.Resolve(n); ok {
		return e, true
	}
	return ns.parent.Resolve(n)
}
