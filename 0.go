package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Bem-vindo ao Jogo da Adivinhação!")
	var allAttempts [][]int

start:
	attempts := make([]int, 0)
	x := rand.Intn(100) + 1

	for i := 1; i < 101; i++ {
		var y int
		fmt.Println("Digite um número entre 1 e 100, por favor: ")
		_, err := fmt.Scan(&y)

		if err != nil {
			fmt.Println("Entrada inválida. Por favor, digite um número inteiro.")
			continue
		}

		attempts = append(attempts, y)

		if y > x {
			fmt.Printf("O número é menor que %d.\n", y)
		} else if y < x {
			fmt.Printf("O número é maior que %d.\n", y)
		} else {
			fmt.Println("Você acertou, parabéns!")

			allAttempts = append(allAttempts, attempts)
			fmt.Printf("Você utlizou %d tentativas para acertar!\n", len(attempts))

			var z string
			fmt.Println("Você deseja jogar novamente? (Digite S para Sim ou N para Não)")
			_, err := fmt.Scan(&z)

			if err != nil {
				fmt.Println("Entrada inválida. Por favor, digite 'S' para jogar novamente ou 'N' para sair.")
				goto start
			}

			if z == "S" || z == "s" {
				goto start
			} else {
				break
			}
		}
	}

	for j, attempts := range allAttempts {
		fmt.Printf("Jogada %d: %d tentativas.\n", j+1, len(attempts))
	}
}
