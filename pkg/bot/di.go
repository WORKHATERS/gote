package bot

import "reflect"

type Dependencies struct {
	services map[reflect.Type]reflect.Value
}

func NewDependencies() *Dependencies {
	return &Dependencies{services: make(map[reflect.Type]reflect.Value)}
}

func (dd *Dependencies) Provide(obj any) {
	dd.services[reflect.TypeOf(obj)] = reflect.ValueOf(obj)
}

func Resolve[T any](dd *Dependencies) (T, bool) {
	var zero T
	t := reflect.TypeOf((*T)(nil)).Elem()
	v, ok := dd.services[t]
	if !ok {
		return zero, false
	}
	return v.Interface().(T), true
}
