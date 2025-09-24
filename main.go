package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func Add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("🚀 Программа запущена!")

	fmt.Println("Привет, мир!")
	result := Add(2, 3)
	fmt.Println("2 + 3 =", result)

	// 1. Проверяем текущую папку
	dir, _ := os.Getwd()
	fmt.Println("Текущая папка:", dir)

	// 2. Проверяем существует ли файл
	if _, err := os.Stat("data.json"); os.IsNotExist(err) {
		fmt.Println("❌ Файл data.json НЕ НАЙДЕН в текущей папке!")
		fmt.Println("Ищу в:", dir)
		return
	}
	fmt.Println("✅ Файл data.json найден")

	// 3. Читаем файл
	file, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("❌ Ошибка чтения файла:", err)
		return
	}
	fmt.Println("✅ Файл прочитан, размер:", len(file), "байт")

	// 4. Покажем сырое содержимое файла
	fmt.Println("Содержимое файла:")
	fmt.Println(string(file))

	// 5. Парсим JSON
	var users []map[string]interface{}
	if err := json.Unmarshal(file, &users); err != nil {
		fmt.Println("❌ Ошибка парсинга JSON:", err)
		return
	}

	// 6. Выводим пользователей
	fmt.Println("✅ JSON распарсен успешно!")
	fmt.Println("Количество пользователей:", len(users))

	if len(users) == 0 {
		fmt.Println("⚠️  Массив пользователей пустой!")
		return
	}

	fmt.Println("Данные из data.json:")
	for i, user := range users {
		fmt.Printf("Пользователь %d: %v\n", i+1, user)
	}
}
