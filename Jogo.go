package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var tentativas, rodada int
	jogadas := make([][2]int, rodada)
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
		tentativas++
	}
	rodada++

	fmt.Print("Deseja jogar novamente? (n/N = Não || s/S = Sim): ")
	var novamente string
	fmt.Scanln(&novamente)
	for {
		novamatrix := make([][3]int, rodada)
		if novamente != "s" && novamente != "S" {
			novamatrix[rodada][1] = rodada
			novamatrix[rodada][2] = tentativas
			fmt.Println(novamatrix)
			fmt.Println(jogadas)
		} else {
			main()
		}
	}
}
