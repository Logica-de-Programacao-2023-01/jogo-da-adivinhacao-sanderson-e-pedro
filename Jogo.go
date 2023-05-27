package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var tentativas, rodada int
	jogadas := []int{}
	rodadas := []int{}
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
	jogadas = append(jogadas, tentativas)
	rodadas = append(rodadas, rodada)
	fmt.Print("Deseja jogar novamente? (n/N = Não || s/S = Sim): ")
	var novamente string
	fmt.Scanln(&novamente)
	if novamente != "s" && novamente != "S" {
		fmt.Println("Resumo: ")
		for y := range jogadas {
			fmt.Printf("TENTATIVA: %d ", y)
			for x := range jogadas {
				fmt.Printf("RODADA: %d", x)
			}
		}
	} else {
		main()
	}

}
