package main

import (
	"fmt"
	"os"
)

func main() {
	i := 0
	for i < 500 {
		file, err := os.Create(fmt.Sprintf("./tmp/file%d.txt", i))
		if err != nil {
			panic(err)
		}
		defer file.Close()
		file.WriteString("Hello, World!")
		i++
	}
}

// snipet para criar arquivos
// rode ele antes de iniciar o experimentos
