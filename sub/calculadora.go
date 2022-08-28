package sub

import "fmt"

// A função fazOperacao declara as variaveis n1, n2 e operacao. Busca o INPUT dos valores para essas variaveis
// Envia essas variaveis preenchidas para a funcao calcula, pega o resultadp e o status e printa no terminal
func FazOperacao() {

	// declara as variaveis
	var n1 int
	var n2 int
	var operacao string
	// Linhas 14, 16 e 18 ele escreve os textos em string. Nas linhas 15, 17 e 19 ele recebe o texto e coloca eles nas variaveis n1, n2 e operacao, respectivamente.
	fmt.Println("Digite o valor de n1:")
	fmt.Scanln(&n1)
	fmt.Println("Digite o valor de n2:")
	fmt.Scanln(&n2)
	fmt.Println("Digite a operação:")
	fmt.Scanln(&operacao)
	//Cria duas variaveis que recebem o retorno da func calcula e printa ele
	resultadoDoCalculo, status := calcula(n1, n2, operacao)
	fmt.Println(resultadoDoCalculo)
	fmt.Println(status)

}

// A funcao calcula recebe como parametro dois numeros e uma string, verifica utilizando uma condicional qual
// o tipo de operacao, e a partir disso retorna o resultado da operacao como um INT e o status do calculo como
// uma string
func calcula(n1, n2 int, operacao string) (int, string) {
	// se operacao for declarado como soma ele retorna a linha 35
	// se operacao for declarado como subtracao ele retorna a linha 37
	// se operacao for declarado como qualquer outra coisa ele retorna a linha 39
	if operacao == "soma" {
		return n1 + n2, "sucesso"
	} else if operacao == "subtracao" || operacao == "subtração" {
		return n1 - n2, "sucesso"
	} else {
		return 0, "erro"
	}

}
