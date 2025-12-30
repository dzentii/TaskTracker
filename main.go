package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"tasktracker/tasks"
)

const mainMenu = ("\nВыберите действие(введите цифру в консоль):\n1. Добавить задачу.\n2. Обновить задачу.\n3. Удалить задачу.\n4. Поменять статус задачи.\n5. Показать все задачи.\n6. Показать все выполненные задачи.\n7. Показать все невыполненные задачи.\n8. Показать все задачи в процессе выполнения.\n9. Выход из программы.")

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
		case "2":
			*store = updateTask(*store)
			continue
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
			fmt.Println("Завершение программы...")
			break outherLoop
		default:
			fmt.Println("Такого пункта нет! Введи число от 1 до 9.")
			continue
		}
	}
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
		fmt.Println("\nНапишите задачу, которую необходимо выполнить: ")
		input := strings.TrimSpace(ScanUserInput())
		if input == "exit" {
			return store
		}
		newStore, err := store.NewTask(input)
		if err != nil {
			fmt.Printf("ошибка создания задачи, %s\n", err)
			continue
		}
		fmt.Println("Задача успешно создана!")
		return *newStore
	}
}

func updateTask(store tasks.TaskStore) tasks.TaskStore {
	for {
		fmt.Println("\nНапишите id задачи, которую хотите обновить: ")
		id := strings.TrimSpace(ScanUserInput())
		if id == "exit" {
			return store
		}
		intId, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("id может быть только в виде числа: ", err)
			continue
		}
		task, ok := store.Tasks[intId]
		if !ok {
			fmt.Println("Задачи с таким id нет!")
			continue
		}
		fmt.Printf("\nВы хотите обновить задачу с id %d?\nЗадача: %s\nСтатус: %s\nСоздана: %s\nОбновлена: %s\n", task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		for {
			fmt.Println("1. Да.\n2. Нет")
			input := strings.TrimSpace(ScanUserInput())
			if input == "1" {
				fmt.Println("\nВведите новое описание задачи: ")
				description := strings.TrimSpace(ScanUserInput())
				newStore, err := store.UpdateTask(intId, description)
				if err != nil {
					fmt.Println("ошибка обновления задачи: ", err)
					continue
				}
				fmt.Println("Задача успешно обновлена!")
				return *newStore
			} else if input == "2" {
				break
			} else {
				fmt.Println("Введите 1 или 2: ")
				continue
			}
		}
	}
}

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
