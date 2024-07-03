// Comando para worksapace

go work init (pacotes que quero usar...)

// exemplo dessa aula
// na pasta 4 (pasta "pai" do projeto)
// go work init ./math ./sistema


ai vc pode rodar seu go run, pois o go ira identificar que é um workspace e estara tudo certo


-----------------------------------------------------------------------------------------------
caso entre em sistema, e adicione um outro pacote tipo uuid:

problema: caso você faço um go mod tidy pra baoxar outras bibliotecas, ele vai quebrar, devido essa
que esta sendo gerenciada pelo workspace não estar publicada

solução:

1 - publica o package
2 - go mod tidy -e (ele vai ignorar os pacotes que não conseguiu achar)

é o que temos no momento