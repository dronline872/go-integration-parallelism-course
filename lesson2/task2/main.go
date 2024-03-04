package main

import (
	"fmt"
	"os"
)

func main() {
	// Задание 2
	file, err := os.Open("log.txt")
	if err != nil {
		fmt.Println("Не удалось открыть файл:", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Ошибка получения информации о файле:", err)
		return
	}

	if stat.Size() == 0 {
		fmt.Println("Файл пуст")
		return
	}

	buf := make([]byte, stat.Size())
	_, err = file.Read(buf)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	fmt.Println(string(buf))
}
