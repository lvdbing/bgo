package convert

import "strconv"

type StrTo string

func (s StrTo) String() string {
	return string(s)
}

func (s StrTo) Int() int {
	v, _ := strconv.Atoi(s.String())
	return v
}

func (s StrTo) Int64() int64 {
	v, _ := strconv.ParseInt(s.String(), 10, 64)
	return v
}
