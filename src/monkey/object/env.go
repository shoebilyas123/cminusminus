package object

type Environment struct {
	store      map[string]Object
	outerScope *Environment
}

func (env *Environment) Set(key string, value Object) Object {
	env.store[key] = value
	return value
}

func (env *Environment) Get(key string) (Object, bool) {
	value, ok := env.store[key]

	if !ok && env.outerScope != nil {
		value, ok = env.outerScope.Get(key)
	}
	return value, ok
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)

	return &Environment{store: s}
}

func NewClosure(outerScope *Environment) *Environment {
	return &Environment{outerScope: outerScope, store: make(map[string]Object)}
}
