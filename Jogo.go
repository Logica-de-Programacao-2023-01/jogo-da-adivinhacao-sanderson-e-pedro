package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var tentativas int
	numero := rand.Intn(100) + 1
	y := []int{}
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
			y = append(y, tentativas)
			break
			//Professor se tentativas não incluir o acerto, está certo, se não, só colocar tentativa + 1
		}
		tentativas += 1
	}

	var novamente string
	fmt.Print("Deseja jogar novamente? (n/N = Não || s/S = Sim): ")
	fmt.Scanln(&novamente)

	if novamente != "s" && novamente != "S" {
		fmt.Println("RESUMO:")
		for i := 0; i <= len(y); i++ {
			fmt.Printf("JOGADA: %d | TENTIVAS: %d\f", i+1, y[i])
		}
	} else {
		main()
	}

}
