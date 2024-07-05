package main

import "github.com/pedrogutierresbr/pos-goexpert/9-apis/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
