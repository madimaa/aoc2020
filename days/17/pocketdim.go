package main

//Pocketdim pocket dimension structure
type Pocketdim struct {
	content map[int]map[int]map[int]*Conwaycube
}

//Conwaycube conway cube structure
type Conwaycube struct {
	x, y, z   int
	neighbors int
	status    rune
}

//CreatePocketDimension returns a new pocket dimension
func CreatePocketDimension() *Pocketdim {
	return &Pocketdim{content: make(map[int]map[int]map[int]*Conwaycube)}
}

//PutIntoDimension put a cube into the dimension
func (dim *Pocketdim) PutIntoDimension(cube *Conwaycube) {
	if _, ok := dim.content[cube.x]; !ok {
		dim.content[cube.x] = make(map[int]map[int]*Conwaycube)
	}

	if _, ok := dim.content[cube.x][cube.y]; !ok {
		dim.content[cube.x][cube.y] = make(map[int]*Conwaycube)
	}

	dim.content[cube.x][cube.y][cube.z] = cube
}

//CreateConwayCube returns a new conway cube
func CreateConwayCube(x, y, z int, status rune) *Conwaycube {
	return &Conwaycube{x: x, y: y, z: z, neighbors: 0, status: status}
}

//CycleDimension update dimension status
func CycleDimension(dim *Pocketdim) *Pocketdim {
	newDim := CreatePocketDimension()
	for x := range dim.content {
		for y := range dim.content[x] {
			for z := range dim.content[x][y] {
				populateNewDim(dim, newDim, dim.content[x][y][z])
			}
		}
	}

	updateNewDim(newDim)
	return newDim
}

func populateNewDim(dim, newDim *Pocketdim, cube *Conwaycube) {
	count := 0
	for x := cube.x - 1; x <= cube.x+1; x++ {
		for y := cube.y - 1; y <= cube.y+1; y++ {
			for z := cube.z - 1; z <= cube.z+1; z++ {
				if x == cube.x && y == cube.y && z == cube.z {
					continue
				}

				if _, ok := dim.content[x][y][z]; ok {
					count++
				} else {
					if _, ok := newDim.content[x][y][z]; ok {
						newDim.content[x][y][z].neighbors++
					} else {
						newNeighbor := CreateConwayCube(x, y, z, '.')
						newNeighbor.neighbors++
						newDim.PutIntoDimension(newNeighbor)
					}
				}
			}
		}
	}

	newCube := CreateConwayCube(cube.x, cube.y, cube.z, '#')
	newCube.neighbors = count
	newDim.PutIntoDimension(newCube)
}

func updateNewDim(dim *Pocketdim) {
	for x := range dim.content {
		for y := range dim.content[x] {
			for _, cube := range dim.content[x][y] {
				switch cube.status {
				case '.':
					if cube.neighbors == 3 {
						cube.status = '#'
					}
				case '#':
					if cube.neighbors != 2 && cube.neighbors != 3 {
						cube.status = '.'
					}
				}
			}
		}
	}
}
