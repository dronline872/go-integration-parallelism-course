package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	// Задание 1
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("Не удалось создать файл:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Введите сообщение: ")
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			break
		}
		line := fmt.Sprintf("%s %s\n", time.Now().Format("2006-01-02 15:04:05"), input)
		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Ошибка записи в файл:", err)
		}
	}
}
