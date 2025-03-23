package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type Perguntas struct {
	Texto   string
	Opcoes  []string
	Questao int
}


type StatusJogo struct {
	Name      string
	Pontuacao int
	Questoes  []Perguntas
}


func (g *StatusJogo) init() {
	fmt.Println("Bem-vindo(a) ao quiz de super-heróis!")
	fmt.Print("Digite seu nome: ")

	reader := bufio.NewReader(os.Stdin)
	nome, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler o nome:", err)
		return
	}

	g.Name = strings.TrimSpace(nome)
	fmt.Println("\nVamos ao jogo,", g.Name)
}


func (g *StatusJogo) ProcessoCSV() {

	f, err := os.Open("herois.csv")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo CSV:", err)
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ';'

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Erro ao ler o arquivo CSV:", err)
		return
	}

	
	for index, record := range records {
		if index == 0 {
			continue 
		}


		resposta, err := strconv.Atoi(record[5])
		if err != nil {
			fmt.Println("Erro", record)
			continue
		}

		questao := Perguntas{
			Texto:   record[0],
			Opcoes:  record[1:5],
			Questao: resposta,
		}
		g.Questoes = append(g.Questoes, questao)
	}

	fmt.Println("Perguntas carregadas com sucesso! Total:", len(g.Questoes))
}

func (g *StatusJogo) Jogar() {
	fmt.Println("\nIniciando o quiz!\n")

	for i, questao := range g.Questoes {
		fmt.Printf("Pergunta %d: %s\n", i+1, questao.Texto)

		
		for j, opcao := range questao.Opcoes {
			fmt.Printf("%d) %s\n", j+1, opcao)
		}

		fmt.Print("Sua resposta: ")
		var respostaUsuario int
		_, err := fmt.Scan(&respostaUsuario)
		if err != nil {
			fmt.Println("Entrada inválida. Pulei essa pergunta.")
			continue
		}

		
		if respostaUsuario == questao.Questao {
			fmt.Println("✅ Resposta correta!")
			g.Pontuacao++
		} else {
			fmt.Printf("❌ Resposta errada! A correta era: %d\n", questao.Questao)
		}

		fmt.Println("---------------------------")
	}

	
	fmt.Printf("\nFim do jogo, %s! Sua pontuação final foi: %d/%d\n", g.Name, g.Pontuacao, len(g.Questoes))
}


func main() {
	jogo := &StatusJogo{}

	
	jogo.init()
	jogo.ProcessoCSV()

	
	if len(jogo.Questoes) > 0 {
		jogo.Jogar()
	} else {
		fmt.Println("Nenhuma pergunta disponível.")
	}
}
