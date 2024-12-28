package crux

import (
	"net/http"
	"strings"
)

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &Context{
		Response:    w,
		Request:     r,
		PathParams:  make(map[string]string),
		QueryParams: r.URL.Query(),
	}

	path := r.URL.Path

	for _, route := range a.routes {
		if matchPath(path, route, ctx) && route.method == r.Method {
			route.handler(ctx)
			return
		}
	}
	http.NotFound(w, r)
}

func matchPath(path string, route Route, ctx *Context) bool {
	pathParts := strings.Split(strings.Trim(path, "/"), "/")
	routeParts := strings.Split(strings.Trim(route.path, "/"), "/")

	if len(pathParts) != len(routeParts) {
		return false
	}

	for i, part := range routeParts {
		if strings.HasPrefix(part, ":") {
			paramName := strings.TrimPrefix(part, ":")
			ctx.PathParams[paramName] = pathParts[i]
			continue
		}
		if part != pathParts[i] {
			return false
		}
	}
	return true
}
