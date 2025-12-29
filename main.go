package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"tasktracker/tasks"
)

const mainMenu = ("Выберите действие(введите цифру в консоль):\n1. Добавить задачу.\n2. Обновить задачу.\n3. Удалить задачу.\n4. Поменять статус задачи.\n5. Показать все задачи.\n6. Показать все выполненные задачи.\n7. Показать все невыполненные задачи.\n8. Показать все задачи в процессе выполнения.\n9. Выход из программы.")

func main() {
	filename := "tasks.json"
	store, err := readFile(filename)
	if err != nil {
		log.Fatal("ошибка чтения файла: %w", err)
	}
outherLoop:
	for {
		fmt.Println(mainMenu)
		fmt.Print("Ввод: ")
		input := ScanUserInput()
		switch strings.TrimSpace(input) {
		case "1":
			*store = createTask(*store)
			continue
		// case "2":
		// 	updateTask()
		// 	continue
		// case "3":
		// 	deleteTask()
		// 	continue
		// case "4":
		// 	changeStatus()
		// 	continue
		// case "5":
		// 	getAll()
		// 	continue
		// case "6":
		// 	getAllDone()
		// 	continue
		// case "7":
		// 	getAllTodo()
		// 	continue
		// case "8":
		// 	getAllInProgress()
		// 	continue
		case "9":
			writeFile(filename, *store)
			break outherLoop
		default:
			fmt.Println("Такого пункта нет! Введи число от 1 до 9.")
			continue
		}
	}

	// newStore, err := store.NewTask("do")
	// if err != nil {
	// 	fmt.Printf("ошибка создания задачи: %v\n", err)
	// }
}

func readFile(filename string) (*tasks.TaskStore, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("ошибка при открытии файла, %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	store := tasks.NewTaskStore()
	if err := decoder.Decode(&store); err != nil && !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("ошибка при декодировании файла, %w", err)
	}
	return store, nil
}

func createTask(store tasks.TaskStore) tasks.TaskStore {
	for {
		fmt.Println("Напишите задачу, которую необходимо выполнить: ")
		input := strings.TrimSpace(ScanUserInput())
		if input == "exit" {
			return store
		}
		newStore, err := store.NewTask(input)
		if err != nil {
			fmt.Printf("ошибка создания задачи, %s\n", err)
			continue
		}
		return *newStore
	}
}

// func updateTask(store tasks.TaskStore) {
// 	for {
// 		input := ScanUserInput()
// 	}
// }

// func deleteTask() {
// 	for {
// 		input := ScanUserInput()
// 	}
// }

// func changeStatus() {
// 	for {
// 		input := ScanUserInput()
// 	}
// }

// func getAll() {
// 	for {
// 		input := ScanUserInput()
// 	}
// }

// func getAllDone() {
// 	for {
// 		input := ScanUserInput()
// 	}
// }

// func getAllTodo() {
// 	for {
// 		input := ScanUserInput()
// 	}
// }

// func getAllInProgress() {
// 	for {
// 		input := ScanUserInput()
// 	}
// }

func writeFile(filename string, store tasks.TaskStore) {
	jsonData, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		log.Fatalf("Ошибка сериализации: %v", err)
	}
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		log.Fatalf("Ошибка записи в файл: %v", err)
	}
}

func ScanUserInput() string {
	var input string
	var err error
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("ошибка ввода:", err)
			fmt.Println("Введите еще раз(для выхода напишите exit): ")
			continue
		}
		break
	}
	return input
}
