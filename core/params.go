package crux

func (c *Context) Param(name string) string {
	return c.PathParams[name]
}

func (c *Context) Query(name string) string {
	return c.QueryParams.Get(name)
}
