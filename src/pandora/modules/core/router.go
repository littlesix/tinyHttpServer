package core

import (
	"bytes"
	"pandora/consts"
)

var (
	routes    []*Router = make([]*Router, 0)
	apiRoutes []*Router = make([]*Router, 0)
)

type Router struct {
	Page IPage
	Rule []byte
}

type RouteMatched struct {
	Page IPage
}

func MatchWebApiRoute(path []byte) *RouteMatched {
	return searchPath(path, apiRoutes)
}
func MatchRoot(path []byte) *RouteMatched {
	return matchRoot(path)
}

func matchRoot(path []byte) *RouteMatched {
	var (
		router *Router
		items  []*Router = routes
	)
	for _, router = range items {
		//fmt.Println("---", string(path), string(router.Rule))
		if bytes.Equal(path, router.Rule) {
			return &RouteMatched{Page: router.Page}
		}

	}

	return nil
}

func searchPath(path []byte, routes []*Router) *RouteMatched {
	if bytes.HasSuffix(path, consts.P_SLASH) {
		path = bytes.TrimSuffix(path, consts.P_SLASH)
	}

	for _, router := range routes {
		if bytes.Equal(path, router.Rule) {
			return &RouteMatched{Page: router.Page}
		}
	}
	return nil

}

func Route(rule string, clz IPage) *Router {
	return addRouteTo(rule, clz)
}

func addRouteTo(rule string, page IPage) *Router {
	r := &Router{
		Rule: []byte(rule),
		Page: page,
	}

	routes = append(routes, r)
	return r
}
