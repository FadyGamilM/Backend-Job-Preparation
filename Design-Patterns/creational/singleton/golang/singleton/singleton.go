package singleton

import "sync"

// the type is private so client code cannot instantiate it
type singleton struct {
	// whatever the fields we need to define here
	val int
}

var instance *singleton

// to manage synconrization
var once sync.Once

func GetSingletonInstance() *singleton {
	// check if this instance is not created before
	if instance != nil {
		// if its created => return the existing instance
		return instance
	} else {
		// if its not created before => create a new instance and return it
		once.Do(
			func() {
				instance = &singleton{
					val: 0,
				}
			})
		return instance
	}
}

func (s *singleton) Val() int {
	return s.val
}

func (s *singleton) SetVal(v int) {
	s.val = v
}
