package main

import (
	"fmt"
	"os"
)

func main() {
	// Задание 3
	file, err := os.Create("read-only.txt")
	if err != nil {
		fmt.Println("Не удалось создать файл:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Содержимое файла")
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}

	err = file.Chmod(0400)
	if err != nil {
		fmt.Println("Ошибка установки прав доступа:", err)
		return
	}

	file, err = os.Open("read-only.txt")
	if err != nil {
		fmt.Println("Не удалось открыть файл:", err)
		return
	}
	defer file.Close()

	buf := make([]byte, 1024)
	_, err = file.Read(buf)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	fmt.Println("Прочитано из файла:", string(buf))
}
