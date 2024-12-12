package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Conectarea la server prin tcp, functia Dial
	conn, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		fmt.Println("Eroare de conectare la server: ", err)
		return
	}
	// Conexiunea e inchisa cand tot ce e in main() termina de rulat
	defer conn.Close()

	connectionConfirmation, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(strings.TrimSpace(connectionConfirmation))

	// Introducem numele fisierului txt de unde se vor citi datele de input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Fisierul din care sa se citeasca datele de input:")
	fileName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Eroare la citirea de la tastatura a numelui fisierului: ", err)
		return
	}
	fileName = strings.TrimSpace(fileName)

	// Deschidem fisierul txt si citim din el
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Eroare la a deschide fisierul txt: ", err)
		return
	}
	defer file.Close()

	fileReader := bufio.NewReader(file)

	// fiecare fmt.Fprintln sau conn.Write trimite date prin conexiunea tcp
	clientName, _ := fileReader.ReadString('\n')
	clientName = strings.TrimSpace(clientName)
	conn.Write([]byte(clientName + "\n")) //fmt.Fprintln(conn, clientName)

	nameConfirmation, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(strings.TrimSpace(nameConfirmation))

	arrayInput, _ := fileReader.ReadString('\n')
	arrayInput = strings.TrimSpace(arrayInput)
	conn.Write([]byte(arrayInput + "\n")) //fmt.Fprintln(conn, arrayInput)

	arrayConfirmation, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(strings.TrimSpace(arrayConfirmation))

	exerciseNumber, _ := fileReader.ReadString('\n')
	exerciseNumber = strings.TrimSpace(exerciseNumber)
	conn.Write([]byte(exerciseNumber + "\n")) //fmt.Fprintln(conn, exerciseNumber)

	exerciseNumberConfirmation, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(strings.TrimSpace(exerciseNumberConfirmation))

	// Raspunsul primit de la server
	response, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("Raspunsul primit de la server: ", strings.TrimSpace(response))
}
