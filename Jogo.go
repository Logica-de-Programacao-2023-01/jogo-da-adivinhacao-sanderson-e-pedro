package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var x [][]int
	for {
		var tentativas int = 1
		y := []int{}
		numero := rand.Intn(100) + 1

		for i := 1; i < 101; i++ {
			var num int
			fmt.Println("Digite um número de 1 a 100: ")
			fmt.Scanf("%d\n", &num)
			y = append(y, num)

			if num < numero {
				fmt.Println("O número certo é maior")
				tentativas++
			} else if num > numero {
				fmt.Println("O número certo é menor")
				tentativas++
			} else {
				fmt.Println("Você acertou!!")
				x = append(x, y)
				fmt.Println("Tentativas: ", tentativas)

				var novamente string
				fmt.Print("Deseja jogar novamente? (n/N = Não || s/S = Sim): ")
				fmt.Scanf("%s\n", &novamente)

				if novamente == "s" || novamente == "S" {
					break
				} else {
					fmt.Println("RESUMO:")
					for j, x2 := range x {
						fmt.Printf("JOGADA: %d | TENTIVAS: %d \n", j+1, len(x2))
					}
					return
				}
			}
		}
	}
}
