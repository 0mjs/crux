package crux

import "strings"

type Group struct {
	prefix string
	app    *App
}

func (g *Group) Group(prefix string) *Group {
	fullPrefix := g.prefix + "/" + strings.Trim(prefix, "/")
	return &Group{
		prefix: fullPrefix,
		app:    g.app,
	}
}

func (a *App) Group(prefix string) *Group {
	return &Group{
		prefix: strings.Trim(prefix, "/"),
		app:    a,
	}
}

func (g *Group) GET(path string, handler interface{}) {
	fullPath := "/" + g.prefix + "/" + strings.Trim(path, "/")
	g.app.GET(fullPath, handler)
}

func (g *Group) POST(path string, handler interface{}) {
	fullPath := "/" + g.prefix + "/" + strings.Trim(path, "/")
	g.app.POST(fullPath, handler)
}

func (g *Group) PUT(path string, handler interface{}) {
	fullPath := "/" + g.prefix + "/" + strings.Trim(path, "/")
	g.app.PUT(fullPath, handler)
}

func (g *Group) DELETE(path string, handler interface{}) {
	fullPath := "/" + g.prefix + "/" + strings.Trim(path, "/")
	g.app.DELETE(fullPath, handler)
}

func (g *Group) PATCH(path string, handler interface{}) {
	fullPath := "/" + g.prefix + "/" + strings.Trim(path, "/")
	g.app.PATCH(fullPath, handler)
}

func (g *Group) HEAD(path string, handler interface{}) {
	fullPath := "/" + g.prefix + "/" + strings.Trim(path, "/")
	g.app.HEAD(fullPath, handler)
}

func (g *Group) OPTIONS(path string, handler interface{}) {
	fullPath := "/" + g.prefix + "/" + strings.Trim(path, "/")
	g.app.OPTIONS(fullPath, handler)
}
