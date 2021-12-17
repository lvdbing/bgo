package convert

import "strconv"

func Stoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func Stoi64(s string) int64 {
	v, _ := strconv.ParseInt(s, 10, 64)
	return v
}

func Itos(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func I64tos(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Stof(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

func Stof32(s string) float32 {
	v, _ := strconv.ParseFloat(s, 32)
	return float32(v)
}

func Ftos(f float64) string {
	return strconv.FormatFloat(f, 'E', -1, 64)
}

func F32tos(f float32) string {
	return strconv.FormatFloat(float64(f), 'E', -1, 32)
}
