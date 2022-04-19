package counter

import "strconv"

// Counter implements increase number.
type Counter int

// Increase adds number + 1.
func (v *Counter) Increase() {
	(*v)++
}

// String converts the integer to string.
func (v *Counter) String() string {
	return strconv.Itoa(int(*v))
}
