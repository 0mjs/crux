package crux

func New() *App {
	return &App{
		routes: make([]Route, 0),
	}
}
