package wellgo

import (
	"sync"
	"regexp"
)

var (
	router *Router
)

type Router struct {
	bindings      *sync.Map
	regexpRouters []*RegexpRouter
}

type RegexpRouter struct {
	path       string
	regex      *regexp.Regexp
	controller ControllerInterface
}

func InitRouter() {
	router = GetRouterInstance()
}

func GetRouterInstance() *Router {
	if router == nil {
		router = &Router{
			bindings:      &sync.Map{},
			regexpRouters: make([]*RegexpRouter, 0),
		}
	}
	return router
}

func (r *Router) Register(path string, controller ControllerInterface) {
	r.bindings.Store(path, controller)
}

func (r *Router) RegexpRegister(regexPath string, controller ControllerInterface) {
	compiledRegex, err := regexp.Compile(regexPath)
	Assert(err == nil, NewWException(ErrSystemError, GetErrorCode(ErrSystemError)))

	r.regexpRouters = append(r.regexpRouters, &RegexpRouter{
		path:       regexPath,
		regex:      compiledRegex,
		controller: controller,
	})
}

// TODO support regex
func (r *Router) Match(path string) (ControllerInterface, error) {
	if controller, found := r.matchBindings(path); found {
		return controller, nil
	}
	if controller, found := r.matchRegexpBindings(path); found {
		return controller, nil
	}

	return nil, ErrInterfaceNotFound
}

func (r *Router) matchBindings(path string) (ControllerInterface, bool) {
	if controller, found := r.bindings.Load(path); found {
		return controller.(ControllerInterface), true
	} else {
		return nil, false
	}
}
func (r *Router) matchRegexpBindings(path string) (ControllerInterface, bool) {
	for i := 0; i < len(r.regexpRouters); i++ {
		if r.regexpRouters[i].regex.MatchString(path) {
			return r.regexpRouters[i].controller, true
		}
	}
	return nil, false
}
