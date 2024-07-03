// Trabalhando com JSON
package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int  `json:"n"`
	Saldo int	`json:"s"`
}

func main() {
	// Pegando uma struct e transformando para JSON
	conta := Conta{Numero: 1, Saldo: 150}
	res, err := json.Marshal(conta) // faz a conversão para JSON --> eu guardo o valor
	if err != nil {
		println(err)
	}
	// println(res) // retorna em bytes por padrão
	println(string(res))

	// Encoder
	err = json.NewEncoder(os.Stdout).Encode(conta) // pego o valor fazendo o processo de serialização entregando pra alguém
	if err != nil {
		println(err)
	}

	// Pegando um JSON e transformando na struct
	// jsonPuro := []byte(`{"Numero": 2, "Saldo": 200}`)
	// var contaX Conta
	// err = json.Unmarshal(jsonPuro, &contaX)
	// if err != nil {
	// 	println(err)
	// }
	// println(contaX.Saldo)

	jsonPuro := []byte(`{"n": 2, "s": 200}`) // ai vc usa tag lá na struct
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}
	println(contaX.Saldo)
}