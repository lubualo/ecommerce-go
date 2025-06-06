package routers

import (
	"github.com/lubualo/ecommerce-go/handlers"
)

type CategoryRouter struct{}

func (router *CategoryRouter) Get(user string, id string, query map[string]string) (int, string) {
	return 405, Get + " not implemented"
}

func (router *CategoryRouter) Post(body string, user string) (int, string) {
	return handlers.PostCategory(body, user)
}

func (router *CategoryRouter) Put(body string, user string, id string) (int, string) {
	return 405, Put + "not implemented"
}

func (router *CategoryRouter) Delete(user string, id string) (int, string) {
	return 405, Delete + "not implemented"
}
