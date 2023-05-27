package main

import (
	"fmt"
	"math/rand"
)

func jogo() {

}

func main() {

	//rand.Seed(time.Now().Unix())
	var tentativas int
	jogadas := [1][1]int{}
	numero := rand.Intn(100) + 1
	for {
		var num int
		fmt.Println("Digite um número de 1 a 100: ")
		fmt.Scanln(&num)
		if num < numero {
			fmt.Println("O número certo é maior")
		} else if num > numero {
			fmt.Println("O número certo é menor")
		} else {
			fmt.Println("Você acertou!!")
			fmt.Println("Tentativas: ", tentativas)
			break
			//Professor se tentativas não incluir o acerto, está certo, se não, só colocar tentativa + 1
		}
		tentativas += 1

	}
	fmt.Print("Deseja jogar novamente? (s/n): ")
	var novamente string
	fmt.Scanln(&novamente)
	for {
		if novamente != "s" && novamente != "S" {
			fmt.Println("Foi de americanes")
		} else {
			main()
		}
	}
	fmt.Println("\n--- Estatísticas das Jogadas ---")
	for i, jogada := range jogadas {
		fmt.Printf("Jogada %d: %d tentativa(s)\n", i+1, len(jogada))
	}
}
