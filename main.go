package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	path := flag.String("path", "myapp.log", "The path to the log that should be analyzed")
	level := flag.String("level", "ERROR", "Log level to search for. Options are DEBUG, INFO, ERROR and CRITICAL")

	flag.Parse()

	// Procurando e armazenando o buffer do arquivo em nossa memória
	file, err := os.Open(*path)

	// err me retorna um string
	// exemplo: '2021/04/16 10:52:52 open myapsp.log: no such file or directory'
	if err != nil {
		log.Fatal(err)
	}
	/*
	   'defer' é utilizado para colocar a operação declarada após a palavra
	   em uma pilha de operações que deve ser executada junto ao retorno da função
	*/
	defer file.Close()

	// Utilizando o pacote bufio(Buffer), para ler o arquivo
	r := bufio.NewReader(file)

	// for sem parâmetros é um while
	for {
		// Le o arquivo como uma string quebrando ele em '\n' que é cada linha
		string, err := r.ReadString('\n')
		// Se retornar erro para o loop, nesse exemplo quando ele não acha mais linha ele irá parar
		if err != nil {
			break
		}

		// Com o pacote strings, verifica se existe a palavra ERROR e se sim printa para o usuário
		if strings.Contains(string, *level) {
			fmt.Println(string)
		}
	}

	return
}
