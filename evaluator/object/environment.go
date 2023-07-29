package object

type Environment struct {
	register map[string]Object
	parent   *Environment
}

func NewEnvironment() *Environment {
	return &Environment{
		register: make(map[string]Object),
	}
}

// Get an object  from the environment by name
func (e *Environment) Get(name string) (Object, bool) {
	val, ok := e.register[name]
	if !ok && e.parent != nil {
		return e.parent.Get(name)
	}
	return val, ok
}

// Set/update an object on the environment
func (e *Environment) Set(name string, val Object) Object {
	e.register[name] = val
	return val
}

// NewScope creates a new child scope with the current env as its parent scope
func (e *Environment) NewScope() *Environment {
	return &Environment{
		register: make(map[string]Object),
		parent:   e,
	}
}
