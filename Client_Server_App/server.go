package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)



// Retine datele de configurare
type Config struct {
	ArraySize    int
	MaxGoroutines int
}

var config Config

// Retine nr de goroutines active
var activeGoroutines int32



// Exercitiul 1
func exercise1(data []string) string {
	result := make([]string, len(data[0]))
	for i := range data[0] {
		for _, word := range data {
			if i < len(word) {
				result[i] += string(word[i])
			}
		}
	}
	return strings.Join(result, ",")
}



// Exercitiul 2
func getNumber(str string) int {
	stringNumber := ""
	for _, character := range str {
		if character >= '0' && character <= '9' {
			stringNumber += string(character)
		}
	}

	number, _ := strconv.Atoi(stringNumber)
	return number
}

func isPerfectSquare(number int) bool {
	for i := 1; i*i <= number; i++ {
		if i*i == number {
			return true
		}
	}
	return false
}

func exercise2(data []string) string {
	count := 0
	for _, str := range data {
		number := getNumber(str)
		if isPerfectSquare(number) {
			count++
		}
	}
	return strconv.Itoa(count)
}



// Exercitiul 3
func reverseString(input string) string {
	reversed := ""
	for i := len(input) - 1; i >= 0; i-- {
		reversed += string(input[i])
	}
	return reversed
}

func exercise3(data []string) string {
	result := 0

	for _, numberString := range data {
		reversedString := reverseString(numberString)
		reversedStringInt, _ := strconv.Atoi(reversedString)

		result += reversedStringInt
	}

	return strconv.Itoa(result)
}



// Exercitiul 4
func exercise4(data []string) string {
	if len(data) < 3 {
		return "Array-ul trebuie sa contina cel putin 3 elemente."
	}

	// Extragem limitele intervalului [a,b]
	a, _ := strconv.Atoi(data[0])
	b, _ := strconv.Atoi(data[1])
	// Extragem sirul de numere
	numbers := data[2:]

	var sum, count int
	for _, numStr := range numbers {
		num, _ := strconv.Atoi(numStr)

		digitSum := 0
		numCopy := num
		for numCopy != 0 {
			digitSum += numCopy % 10
			numCopy = numCopy / 10
		}

		if digitSum >= a && digitSum <= b {
			sum += num
			count++
		}
	}

	if count == 0 {
		return "Nu exista numere care sa apartina lui [a,b]"
	}

	average := sum / count
	return strconv.Itoa(average)
}



// Exercitiul 5
func isBinary(str string) bool  {
	for _, character := range str {
		if character != '0' && character != '1' {
			return false
		}
	}
	return true
}

func exercise5(data []string) string {
	var result []string
	for _, str := range data {
		if isBinary(str) {
			value, _ := strconv.ParseInt(str, 2, 64) // converteste din base 2 in int64
			result = append(result, strconv.Itoa(int(value)))
		}
	}

	return strings.Join(result, ",")
}



// Exercitiul 6
func caesar(word string, k int, direction string) string {
	k = k % 26
    if direction == "left" {
        k = -k
    }

	var result string
	for i := 0; i < len(word); i++ {
		character := word[i]
		if character >= 'a' && character <= 'z' {
			result += string('a' + ((int(character - 'a') + k + 26) % 26))
		} else if character >= 'A' && character <= 'Z' {
			result += string('A' + ((int(character - 'A') + k + 26) % 26))
		} else {
			result += string(character)
		}
	}

	return result
}

func exercise6(data []string) string {
	if len(data) < 3 {
		return "Array-ul trebuie sa contina cel putin 3 elemente!"
	}

	k, _ := strconv.Atoi(data[0])

	direction := strings.ToLower(data[1])
	if direction != "left" && direction != "right" {
		return "Directia e invalida. Trebuie sa fie left sau right!"
	}

	var result []string
	for _, word := range data[2:] {
		encoded := caesar(word, k, direction)
		result = append(result, encoded)
	}

	return strings.Join(result, ",")
}



// Exercitiul 7
func exercise7(data []string) string {
	str := data[0]
	decoded := ""

	for i := 0; i < len(str); {
		if i + 2 < len(str) && str[i + 1] >= '0' && str[i + 1] <= '9' { // pt numere cu 2 cifre
			num, _ := strconv.Atoi(str[i : i+2])
			if num >= 1 && num <= 20 {
				decoded += strings.Repeat(string(str[i + 2]), num)
			}
			i += 3
		} else if i+1 < len(str) {
			num, _ := strconv.Atoi(string(str[i]))
			if num >= 1 && num <= 20 {
				decoded += strings.Repeat(string(str[i + 1]), num)
			}
			i += 2
		}
	}

	return decoded
}



// Exercitiul 8
func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i*i <= number; i++ {
		if number % i == 0 {
			return false
		}
	}
	return true
}

func exercise8(data []string) string {
	result := 0

	for _, str := range data {
		number, _ := strconv.Atoi(str)
		if isPrime(number) {
			result += len(str)
		}
	}

	return strconv.Itoa(result)
}



// Exercitiul 9
func exercise9(data []string) string {
	vowels := "aeiouAEIOU"
	result := 0

	for _, str := range data {
		counter := 0
		for i := 0; i < len(str); i += 2 {
			for _, vowel := range vowels {
				if str[i] == byte(vowel) {
					counter++
					break
				}
			}
		}
		if counter % 2 == 0 && counter != 0 {
			result++;
		}
	}

	return strconv.Itoa(result)
}



// Exercitiul 10
func cmmdcTwoNumbers(x int, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}

	if x == y {
		return x
	}

	if x > y {
		return cmmdcTwoNumbers(x - y, y)
	}
	return cmmdcTwoNumbers(x, y - x)
}

func cmmdc(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	result := numbers[0]
	for _, number := range numbers[1:] {
		result = cmmdcTwoNumbers(result, number)
	}
	return result
}

func exercise10(data []string) string {
	var numbers []int
	for _, str := range data {
		isNumber := true
		for i := 0; i < len(str); i++ {
			if str[i] < '0' || str[i] > '9' {
				isNumber = false
				break
			}
		}
		if isNumber {
			number, _ := strconv.Atoi(str)
			numbers = append(numbers, number)
		}
	}
	result := cmmdc(numbers)
	return strconv.Itoa(result)
}



// Exercitiul 11
func transform(number string, k int) int {
	k = k % len(number)
	transformed := number[len(number) - k:] + number[:len(number) - k]
	result, _ := strconv.Atoi(transformed)
	return result
}

func exercise11(data []string) string {
	k, _ := strconv.Atoi(data[0])
	result := 0

	data = data[1:]
	for _, number := range data {
		newNumber := transform(number,k)
		result += newNumber
		fmt.Println(newNumber, result)
	}
	return strconv.Itoa(result)
}



// Exercitiul 12
func exercise12(data []string) string {
	result := 0
	for _, numberStr := range data {
		newNumberStr := string(numberStr[0]) + numberStr
		newNumberInt, _ := strconv.Atoi(newNumberStr)
		result += newNumberInt
	}
	return strconv.Itoa(result)
}



func handleClient(conn net.Conn) {
	// defer face ca aceste functii sa se execute la finalul functiei handleClient
	defer conn.Close()
	defer atomic.AddInt32(&activeGoroutines, -1) // scadem 1 din nr de goroutines active

	reader := bufio.NewReader(conn)

	clientName, _ := reader.ReadString('\n')
	clientName = strings.TrimSpace(clientName)
	fmt.Println("Numele clientului conectat este: ", clientName)
	conn.Write([]byte("Server-ul a primit numele clientului.\n"))

	inputArray, _ := reader.ReadString('\n')
	inputArray = strings.TrimSpace(inputArray)
	data := strings.Split(inputArray, ",")
	fmt.Println("Vectorul trimis de client este: ", data)
	conn.Write([]byte("Server-ul a primit vectorul.\n"))

	exerciseNumber, _ := reader.ReadString('\n')
	exerciseNumber = strings.TrimSpace(exerciseNumber)
	fmt.Println("Numarul exercitiului la care clientul vrea raspuns este: ", exerciseNumber)
	conn.Write([]byte("Server-ul a primit numarul exercitiului.\n"))

	if len(data) > config.ArraySize {
		fmt.Println("Clientul a trimis un array prea mare: ", data)
		conn.Write([]byte("Vectorul trimis e prea mare."))
	} else {
		// Simulare procesare de 7 secunde
		fmt.Println("Se proceseaza...")
		time.Sleep(7 * time.Second)

		var result string

		switch exerciseNumber {
		case "1":
			result = exercise1(data)
		case "2":
			result = exercise2(data)
		case "3":
			result = exercise3(data)
		case "4":
			result = exercise4(data)
		case "5":
			result = exercise5(data)
		case "6":
			result = exercise6(data)
		case "7":
			result = exercise7(data)
		case "8":
			result = exercise8(data)
		case "9":
			result = exercise9(data)
		case "10":
			result = exercise10(data)
		case "11":
			result = exercise11(data)
		case "12":
			result = exercise12(data)
		default:
			result = "Numar de exercitiu invalid!"
		}

		fmt.Println("Server-ul trimite urmatorul raspuns catre client:", result)
		conn.Write([]byte(result + "\n"))
		fmt.Println("Clientul ", clientName, " a primit raspunsul: ", result)
	}
}



func main() {
	// Citim din fisierul config.json
	configurations, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Eroare la a deschide fisierul config.json", err)
		return
	}
	defer configurations.Close()
	json.NewDecoder(configurations).Decode(&config)

	// Ascultam dupa clienti care se conecteaza
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println("Eroare la a porni server-ul", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server-ul a pornit...")
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Eroare la acceptarea conexiunii", err)
			continue
		}
		conn.Write([]byte("Clientul s-a conectat la server.\n"))

		for atomic.LoadInt32(&activeGoroutines) >= int32(config.MaxGoroutines) {
			time.Sleep(2 * time.Second)
			fmt.Println("UN CLIENT ASTEAPTA")
		}

		atomic.AddInt32(&activeGoroutines, 1) // adaugam 1 la nr de goroutines active
		go handleClient(conn)
	}

}
