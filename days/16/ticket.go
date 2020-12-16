package main

type fields struct {
	data map[string]field
}

type field struct {
	name                              string
	lowerBoundaries, higherBoundaries pair
}

type pair struct {
	lower, higher int
}

func (p *pair) contains(num int) bool {
	if p.lower <= num && p.higher >= num {
		return true
	}

	return false
}

//Contains todo
func (f *field) Contains(num int) bool {
	if f.lowerBoundaries.contains(num) || f.higherBoundaries.contains(num) {
		return true
	}

	return false
}

//ContainsAny todo
func (fs *fields) ContainsAny(num int) bool {
	result := false

	for _, f := range fs.data {
		if f.Contains(num) {
			return true
		}
	}

	return result
}
