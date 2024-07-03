// Generics
package main

type MyNumber int

type Number interface {
	~int | float64
}

func Soma [T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

func main() {
	m := map[string]int{"Pedro": 1000, "Rafael": 2000, "Filipe": 3000}
	m2 := map[string]float64{"Pedro": 1000.32, "Rafael": 2000.20, "Filipe": 3000.30}
	m3 := map[string]MyNumber{"Pedro": 100032, "Rafael": 200020, "Filipe": 30000}

	println(Soma(m))
	println((Soma(m2)))
	println((Soma(m3)))
	println((Compara(20,20)))
}

