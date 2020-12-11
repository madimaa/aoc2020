package array2d

/*
This solution is slower than [][]type on less than ~50k entries
but significantly faster on 10m+ data
*/

//Array2D - 2 dimensional array
type Array2D struct {
	data []interface{}
	x    int
	y    int
}

//Create - create 2d array
func Create(x, y int) *Array2D {
	return &Array2D{data: make([]interface{}, x*y), x: x, y: y}
}

//Put - add element to the array
func (a2d *Array2D) Put(x, y int, value interface{}) {
	a2d.data[x*a2d.y+y] = value
}

//Get - get value from x y index
func (a2d *Array2D) Get(x, y int) interface{} {
	return a2d.data[x*a2d.y+y]
}

//GetSize - returns the size of X and size of Y
func (a2d *Array2D) GetSize() (int, int) {
	return a2d.x, a2d.y
}

//GetPosition - returns the position of the element
func (a2d *Array2D) GetPosition(x, y int) int {
	return x*a2d.y + y
}

//Copy - returns a copy of the original array
func (a2d *Array2D) Copy() *Array2D {
	newArray := &Array2D{data: make([]interface{}, a2d.x*a2d.y), x: a2d.x, y: a2d.y}
	for index, val := range a2d.data {
		newArray.data[index] = val
	}

	return newArray
}

//Equals - true if the arrays are identical
func (a2d *Array2D) Equals(other *Array2D) bool {
	if a2d.x != other.x || a2d.y != other.y {
		return false
	}

	for index, val := range a2d.data {
		if val != other.data[index] {
			return false
		}
	}

	return true
}
