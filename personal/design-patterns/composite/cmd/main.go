package main

import (
	"composite"
	"fmt"
)

func main() {
	p := composite.Project{Name: "Upload images"}
	t1 := composite.Task{Name: "Mockup", Anchor: "UX Dev", Total: 1000}
	t2 := composite.Task{Name: "Markup", Anchor: "Web Dev"}
	t3 := composite.Task{Name: "JS", Anchor: "Frontend Dev", Total: 1500}
	t4 := composite.Task{Name: "API Backend", Anchor: "Backend Dev"}
	t5 := composite.Task{Name: "Database", Anchor: "DBA Dev", Total: 4000}

  st21 := composite.SubTask{Name: "HTML", Total: 300}
  st22 := composite.SubTask{Name: "CSS", Total: 700}
  t2.Add(&st21)
  t2.Add(&st22)

  st41 := composite.SubTask{Name: "Authentication", Total: 1000}
  st42 := composite.SubTask{Name: "DB connection", Total: 700}
  t4.Add(&st41)
  t4.Add(&st42)

	p.Add(&t1)
	p.Add(&t2)
	p.Add(&t3)
	p.Add(&t4)
	p.Add(&t5)
	fmt.Println(&p)
}
