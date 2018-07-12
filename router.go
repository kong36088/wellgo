package wellgo

import (
	"sync"
	"errors"
)

var (
	router *Router
)

type Router struct {
	bindings *sync.Map
}

func GetRouterInstance() *Router {
	if router == nil {
		router = &Router{
			bindings: &sync.Map{},
		}
	}
	return router
}

func (r *Router) Register(path string, controller *Controller) {
	r.bindings.Store(path, controller)
}

// TODO support regex
func (r *Router) Match(path string) (*Controller, error) {
	if controller, found := r.bindings.Load(path); !found {
		return nil, errors.New("path not found")
	} else {
		return controller.(*Controller), OK
	}
}
