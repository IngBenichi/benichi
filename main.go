package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		// Si le pasas un archivo, lo ejecuta
		archivo := os.Args[1]
		runArchivo(archivo)
	} else {
		// Si no, entra en modo REPL
		fmt.Println("Benichi REPL 🔥")
		repl()
	}
}

func runArchivo(archivo string) {
	file, err := os.Open(archivo)
	if err != nil {
		fmt.Printf("Error al abrir archivo: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		ejecutar(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error leyendo archivo: %v\n", err)
	}
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">>> ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "exit" || line == "quit" {
			break
		}
		ejecutar(line)
	}
}

func ejecutar(line string) {
	if strings.HasPrefix(line, "print ") {
		texto := strings.TrimPrefix(line, "print ")
		texto = strings.Trim(texto, "\"")
		fmt.Println(texto)
	} else {
		fmt.Printf("Comando no reconocido: %s\n", line)
	}
}
