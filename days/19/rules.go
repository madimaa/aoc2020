package main

//Rule a container to store the rules
type Rule struct {
	letter  rune
	subrule []Subrule
}

//Subrule a container to store the subrules
type Subrule struct {
	subrules []int
}

func (r Rule) isLast() bool {
	if r.subrule == nil {
		return true
	}

	return false
}
