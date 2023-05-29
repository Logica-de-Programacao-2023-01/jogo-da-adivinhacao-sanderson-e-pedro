package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var tentativas int = 0
	var rodada int
	numero := rand.Intn(10) + 1
	y := make([]int, 0)
	for {

		var num int
		fmt.Println("Digite um número de 1 a 100: ")
		fmt.Scanln(&num)
		tentativas++
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
	}
	rodada++

	var novamente string
	fmt.Print("Deseja jogar novamente? (n/N = Não || s/S = Sim): ")
	fmt.Scanln(&novamente)
	if novamente != "s" && novamente != "S" {
		fmt.Println("RESUMO:")
	} else {
		main()
	}

	for i, s := range y {
		fmt.Printf("JOGADA: %d | TENTIVAS: %d\f", i, s)
		i++
	}

}
