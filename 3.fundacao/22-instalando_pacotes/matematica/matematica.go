package matematica

func Soma[T int | float64](a, b T) T {
	return a + b
}

var A int = 25

type Carro struct {
	Marca string
}

func (c Carro) Andar() string{
	return "Andou"
}

// no go não existe export default, para exportar, tem que usar maiúsculo