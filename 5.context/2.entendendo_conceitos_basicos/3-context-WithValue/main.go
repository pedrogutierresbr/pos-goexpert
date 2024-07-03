// Context withValue

// Use com responsabilidade, pois vc não sabe se você tem essa info, se a info foi passada de forma obrigatoria
// o recomendado é sempre passar os dados via parametro para as funções
package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "pedroso123")
	bookHotel(ctx, "Copacabana Palace")
}

func bookHotel(ctx context.Context, name string) {
	token := ctx.Value("token")
	fmt.Println(token)
	fmt.Println(name)
}

// IMPORTANTE: toda vez que utilizar contexto em uma função que está criando, mesmo que tenha outros parametros depois,
// por CONVEÇÃO, as variáveis de contexto sempre vão em primeiro lugar na sua função

// na pós, o professor que ele normalmente utiliza passar via contexto, informações de métricas, telemetria (mas isso é ele)
