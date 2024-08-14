package object

type Environment struct {
	store map[string]Object
}

func (env *Environment) Set(key string, value Object) Object {
	env.store[key] = value
	return value
}

func (env *Environment) Get(key string) (Object, bool) {
	value, ok := env.store[key]
	return value, ok
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)

	return &Environment{store: s}
}
