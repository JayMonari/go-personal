package main

type Dep struct {
	Prereq int
	Job    int
}

type JobNode struct {
	Job      int
	Prereqs  []*JobNode
	Visited  bool
	Visiting bool
}

func TopologicalSort(jobs []int, deps []Dep) []int {
	return getOrderedJobs(createJobGraph(jobs, deps))
}

func createJobGraph(jobs []int, deps []Dep) *JobGraph {
	graph := NewJobGraph(jobs)
	for _, dep := range deps {
		graph.AddPrereq(dep.Job, dep.Prereq)
	}
	return graph
}

func getOrderedJobs(graph *JobGraph) []int {
	orderedJobs := []int{}
	nodes := graph.Nodes
	for len(nodes) != 0 {
		node := nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		containsCycle := depthFirstTraverse(node, &orderedJobs)
		if containsCycle {
			return []int{}
		}
	}
	return orderedJobs
}

func depthFirstTraverse(node *JobNode, orderedJobs *[]int) (isCyclic bool) {
	switch {
	case node.Visited:
		return false
	case node.Visiting:
		return true
	}

	node.Visiting = true
	for _, prereqNode := range node.Prereqs {
		if depthFirstTraverse(prereqNode, orderedJobs) {
			return true
		}
	}
	node.Visited = true
	node.Visiting = false
	*orderedJobs = append(*orderedJobs, node.Job)
	return false
}

type JobGraph struct {
	Nodes []*JobNode
	Graph map[int]*JobNode
}

func NewJobGraph(jobs []int) *JobGraph {
	g := &JobGraph{Graph: map[int]*JobNode{}}
	for _, job := range jobs {
		g.AddNode(job)
	}
	return g
}

func (g *JobGraph) AddPrereq(job, prereq int) {
	jn := g.GetNode(job)
	jn.Prereqs = append(jn.Prereqs, g.GetNode(prereq))
}

func (g *JobGraph) AddNode(job int) {
	g.Graph[job] = &JobNode{Job: job}
	g.Nodes = append(g.Nodes, g.Graph[job])
}

func (g *JobGraph) GetNode(job int) *JobNode {
	if _, ok := g.Graph[job]; !ok {
		g.AddNode(job)
	}
	return g.Graph[job]
}

// TODO(jay): Move this to work for breadth-first search.
// type JobNode struct {
// 	Job          int
// 	Deps         []*JobNode
// 	NumOfPrereqs int
// }
//
// func TopologicalSort(jobs []int, deps []Dep) []int {
// 	return getOrderedJobs(createJobGraph(jobs, deps))
// }
//
// func createJobGraph(jobs []int, deps []Dep) *JobGraph {
// 	graph := NewJobGraph(jobs)
// 	for _, dep := range deps {
// 		graph.AddDep(dep.Prereq, dep.Job)
// 	}
// 	return graph
// }
//
// func getOrderedJobs(graph *JobGraph) []int {
// 	orderedJobs := []int{}
// 	nodesWithNoPrereqs := []*JobNode{}
// 	for _, node := range graph.Nodes {
// 		if node.NumOfPrereqs == 0 {
// 			nodesWithNoPrereqs = append(nodesWithNoPrereqs, node)
// 		}
// 	}
// 	for len(nodesWithNoPrereqs) > 0 {
// 		node := nodesWithNoPrereqs[len(nodesWithNoPrereqs)-1]
// 		nodesWithNoPrereqs = nodesWithNoPrereqs[:len(nodesWithNoPrereqs)-1]
// 		orderedJobs = append(orderedJobs, node.Job)
// 		removeDeps(node, &nodesWithNoPrereqs)
// 	}
// 	for _, node := range graph.Nodes {
// 		if node.NumOfPrereqs > 0 {
// 			return []int{}
// 		}
// 	}
// 	return orderedJobs
// }
//
// func removeDeps(node *JobNode, nodesWithNoPrereqs *[]*JobNode) {
// 	for len(node.Deps) > 0 {
// 		dep := node.Deps[len(node.Deps)-1]
// 		node.Deps = node.Deps[:len(node.Deps)-1]
// 		dep.NumOfPrereqs--
// 		if dep.NumOfPrereqs == 0 {
// 			*nodesWithNoPrereqs = append(*nodesWithNoPrereqs, dep)
// 		}
// 	}
// }
//
// type JobGraph struct {
// 	Nodes []*JobNode
// 	Graph map[int]*JobNode
// }
//
// func NewJobGraph(jobs []int) *JobGraph {
// 	g := &JobGraph{
// 		Graph: map[int]*JobNode{},
// 	}
// 	for _, job := range jobs {
// 		g.AddNode(job)
// 	}
// 	return g
// }
//
// func (g *JobGraph) AddDep(job, dep int) {
// 	jobNode, depNode := g.GetNode(job), g.GetNode(dep)
// 	jobNode.Deps = append(jobNode.Deps, depNode)
// 	depNode.NumOfPrereqs++
// }
//
// func (g *JobGraph) AddNode(job int) {
// 	g.Graph[job] = &JobNode{Job: job}
// 	g.Nodes = append(g.Nodes, g.Graph[job])
// }
//
// func (g *JobGraph) GetNode(job int) *JobNode {
// 	if _, found := g.Graph[job]; !found {
// 		g.AddNode(job)
// 	}
// 	return g.Graph[job]
// }

// Test Case 1
//
// {
//   "jobs": [1, 2, 3, 4],
//   "deps": [
//     [1, 2],
//     [1, 3],
//     [3, 2],
//     [4, 2],
//     [4, 3]
//   ]
// }
//
// Test Case 2
//
// {
//   "jobs": [1, 2, 3, 4, 5, 6, 7, 8],
//   "deps": [
//     [3, 1],
//     [8, 1],
//     [8, 7],
//     [5, 7],
//     [5, 2],
//     [1, 4],
//     [1, 6],
//     [1, 2],
//     [7, 6]
//   ]
// }
//
// Test Case 3
//
// {
//   "jobs": [1, 2, 3, 4, 5, 6, 7, 8],
//   "deps": [
//     [3, 1],
//     [8, 1],
//     [8, 7],
//     [5, 7],
//     [5, 2],
//     [1, 4],
//     [6, 7],
//     [1, 2],
//     [7, 6]
//   ]
// }
//
// Test Case 4
//
// {
//   "jobs": [1, 2, 3, 4, 5, 6, 7, 8],
//   "deps": [
//     [3, 1],
//     [8, 1],
//     [8, 7],
//     [5, 7],
//     [5, 2],
//     [1, 4],
//     [1, 6],
//     [1, 2],
//     [7, 6],
//     [4, 6],
//     [6, 2],
//     [2, 3]
//   ]
// }
//
// Test Case 5
//
// {
//   "jobs": [1, 2, 3, 4, 5, 6, 7, 8],
//   "deps": [
//     [1, 2],
//     [2, 3],
//     [3, 4],
//     [4, 5],
//     [5, 6],
//     [6, 7],
//     [7, 8],
//     [8, 1]
//   ]
// }
//
// Test Case 6
//
// {
//   "jobs": [1, 2, 3, 4, 5, 6, 7, 8, 9],
//   "deps": [
//     [1, 2],
//     [2, 3],
//     [3, 4],
//     [4, 5],
//     [5, 6],
//     [7, 6],
//     [7, 8],
//     [8, 1]
//   ]
// }
//
// Test Case 7
//
// {
//   "jobs": [1, 2, 3, 4, 5, 6, 7, 8],
//   "deps": [
//     [1, 2],
//     [3, 5],
//     [4, 6],
//     [3, 6],
//     [1, 7],
//     [7, 8],
//     [1, 8],
//     [2, 8]
//   ]
// }
//
// Test Case 8
//
// {
//   "jobs": [1, 2, 3, 4, 5, 6, 7, 8],
//   "deps": [
//     [1, 2],
//     [1, 3],
//     [1, 4],
//     [1, 5],
//     [1, 6],
//     [1, 7],
//     [2, 8],
//     [3, 8],
//     [4, 8],
//     [5, 8],
//     [6, 8],
//     [7, 8]
//   ]
// }
//
// Test Case 9
//
// {
//   "jobs": [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12],
//   "deps": [
//     [1, 2],
//     [1, 3],
//     [1, 4],
//     [1, 5],
//     [1, 6],
//     [1, 7],
//     [2, 8],
//     [3, 8],
//     [4, 8],
//     [5, 8],
//     [6, 8],
//     [7, 8],
//     [2, 3],
//     [2, 4],
//     [5, 4],
//     [7, 6],
//     [6, 2],
//     [6, 3],
//     [6, 5],
//     [5, 9],
//     [9, 8],
//     [8, 0],
//     [4, 0],
//     [5, 0],
//     [9, 0],
//     [2, 0],
//     [3, 9],
//     [3, 10],
//     [10, 11],
//     [11, 12],
//     [2, 12]
//   ]
// }
//
// Test Case 10
//
// {
//   "jobs": [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12],
//   "deps": [
//     [1, 2],
//     [1, 3],
//     [1, 4],
//     [1, 5],
//     [1, 6],
//     [1, 7],
//     [2, 8],
//     [3, 8],
//     [4, 8],
//     [5, 8],
//     [6, 8],
//     [7, 8],
//     [2, 3],
//     [2, 4],
//     [5, 4],
//     [7, 6],
//     [6, 2],
//     [6, 3],
//     [6, 5],
//     [5, 9],
//     [9, 8],
//     [8, 0],
//     [4, 0],
//     [5, 0],
//     [9, 0],
//     [2, 0],
//     [3, 9],
//     [3, 10],
//     [10, 11],
//     [11, 12],
//     [12, 2]
//   ]
// }
//
// Test Case 11
//
// {
//   "jobs": [1, 2, 3, 4, 5],
//   "deps": []
// }
//
// Test Case 12
//
// {
//   "jobs": [1, 2, 3, 4, 5],
//   "deps": [
//     [1, 4],
//     [5, 2]
//   ]
// }
