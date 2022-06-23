package decorator

import "fmt"

type Route struct{ decorators map[string]Decorator }

func NewRoute() Route { return Route{decorators: make(map[string]Decorator)} }

func (r *Route) Add(d Decorator, path string) { r.decorators[path] = d }

func (r *Route) Exec(path string) {
	d, ok := r.decorators[path]
	if !ok {
		fmt.Println("404 Not Found")
		return
	}

	if err := d.Process(); err != nil {
		fmt.Println("ERROR:", err)
	}
}
