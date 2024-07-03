// Context --> serve para controlar, no final das contas, o que a aplicação esta vivendo, e se algum tempo estourar, do que estavamos parametrizando
// ele simplesmente pode parar de executar aquela operação

// ele também permite que guardemos informações no contexto (que estipulamos) para que a gente possa resgatar em outras áreas da aplicação
// muito útil em headers http, em chamadas de filas, coisas desse tipo

// exemplo: você vai fazer uma consulta ao banco de dados, esta demorando muito, você já cancela e "olha, eu não vou fazer essa consulta novamente"
// exemplo: você vai fazer uma consulta a uma API, e essa API esta muito lenta, você não quer travar o seu processo, então ai você cancela (usando contexto)
package main

func main() {

}