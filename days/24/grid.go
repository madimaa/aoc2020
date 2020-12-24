package main

type grid struct {
	content map[int]map[int]*tile
}

type tile struct {
	x, y      int
	color     bool
	neighbors int
}

func createGrid() *grid {
	return &grid{content: make(map[int]map[int]*tile)}
}

func (g *grid) putIntoGrid(t *tile) {
	if _, ok := g.content[t.x]; !ok {
		g.content[t.x] = make(map[int]*tile)
	}

	g.content[t.x][t.y] = t
}

func (g *grid) get(x, y int) *tile {
	if _, ok := g.content[x]; ok {
		return g.content[x][y]
	}

	return nil
}

func createTile(x, y int, color bool) *tile {
	return &tile{x: x, y: y, color: color}
}

func (t *tile) copy() *tile {
	return &tile{x: t.x, y: t.y, color: t.color}
}

func (t *tile) move(direction string) {
	switch direction {
	case "e":
		t.x++
	case "w":
		t.x--
	case "se":
		t.y--
		t.x++
	case "sw":
		t.y--
	case "ne":
		t.y++
	case "nw":
		t.y++
		t.x--
	default:
		panic("Invalid direction")
	}
}

func (g *grid) flip() *grid {
	newGrid := createGrid()
	for x := range g.content {
		for _, t := range g.content[x] {
			if t.color {
				populateNewGrid(g, newGrid, t)
			}
		}
	}

	updateNewGrid(newGrid)
	return newGrid
}

func populateNewGrid(g, newGrid *grid, t *tile) {
	count := 0
	for x := t.x - 1; x <= t.x+1; x++ {
		for y := t.y - 1; y <= t.y+1; y++ {
			if x == t.x && y == t.y || x == t.x-1 && y == t.y-1 || x == t.x+1 && y == t.y+1 {
				continue
			}

			put := true
			if _, ok := g.content[x][y]; ok {
				if g.content[x][y].color {
					put = false
					count++
				}
			}

			if put {
				if _, ok := newGrid.content[x][y]; ok {
					newGrid.content[x][y].neighbors++
				} else {
					newNeighbor := createTile(x, y, false)
					newNeighbor.neighbors++
					newGrid.putIntoGrid(newNeighbor)
				}
			}
		}
	}

	newTile := createTile(t.x, t.y, true)
	newTile.neighbors = count
	newGrid.putIntoGrid(newTile)
}

func updateNewGrid(g *grid) {
	for x := range g.content {
		for _, t := range g.content[x] {
			if t.color {
				if t.neighbors == 0 || t.neighbors > 2 {
					t.color = false
				}
			} else {
				if t.neighbors == 2 {
					t.color = true
				}
			}
		}
	}
}
