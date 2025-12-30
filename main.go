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
		log.Fatalf("ошибка чтения файла: %s", err)
	}
outherLoop:
	for {
		fmt.Println(mainMenu)
		fmt.Print("Ввод: ")
		input := ScanUserInput()
		switch strings.TrimSpace(input) {
		case "1":
			createTask(store)
			continue
		case "2":
			updateTask(store)
			continue
		case "3":
			deleteTask(store)
			continue
		case "4":
			changeStatus(store)
			continue
		case "5":
			getAll(store)
			continue
		case "6":
			getAllDone(store)
			continue
		case "7":
			getAllTodo(store)
			continue
		case "8":
			getAllInProgress(store)
			continue
		case "9":
			writeFile(filename, store)
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

func createTask(store *tasks.TaskStore) {
	for {
		fmt.Println("\nНапишите задачу, которую необходимо выполнить: ")
		input := strings.TrimSpace(ScanUserInput())
		if input == "exit" {
			return
		}

		if _, err := store.NewTask(input); err != nil {
			fmt.Printf("ошибка создания задачи, %s\n", err)
			continue
		}
		fmt.Println("Задача успешно создана!")
		return
	}
}

func updateTask(store *tasks.TaskStore) {
	for {
		fmt.Println("\nНапишите id задачи, которую хотите обновить: ")
		id := strings.TrimSpace(ScanUserInput())
		if id == "exit" {
			return
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
				if _, err := store.UpdateTask(intId, description); err != nil {
					fmt.Println("ошибка обновления задачи: ", err)
					continue
				}
				fmt.Println("Задача успешно обновлена!")
				return
			} else if input == "2" {
				return
			} else {
				fmt.Println("Введите 1 или 2: ")
				continue
			}
		}
	}
}

func deleteTask(store *tasks.TaskStore) {
	for {
		fmt.Println("\nНапишите id задачи, которую хотите удалить: ")
		id := strings.TrimSpace(ScanUserInput())
		if id == "exit" {
			return
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
		fmt.Printf("\nВы хотите удалить задачу с id %d?\nЗадача: %s\nСтатус: %s\nСоздана: %s\nОбновлена: %s\n", task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		for {
			fmt.Println("1. Да.\n2. Нет")
			input := strings.TrimSpace(ScanUserInput())
			if input == "1" {
				fmt.Println("Удаление задачи...")

				if _, err := store.DeleteTask(task.ID); err != nil {
					fmt.Println("ошибка удаления задачи:", err)
					continue
				}
				fmt.Println("Задача удалена!")
				return
			} else if input == "2" {
				return
			} else {
				fmt.Println("Введите 1 или 2: ")
				continue
			}
		}
	}
}

func changeStatus(store *tasks.TaskStore) {
	for {
		fmt.Println("\nНапишите id задачи, для которой хотите обновить статус: ")
		id := strings.TrimSpace(ScanUserInput())
		if id == "exit" {
			return
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
		fmt.Printf("\nВы хотите обновить статус для задачи с id %d?\nЗадача: %s\nСтатус: %s\nСоздана: %s\nОбновлена: %s\n", task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		for {
			fmt.Println("1. Да.\n2. Нет")
			input := strings.TrimSpace(ScanUserInput())
			if input == "1" {
				fmt.Println("\nВыберите новый статус: \n1. Необходимо выполнить.\n2. В процессе выполнения\n3. Выполнена.")
				var status string
				for {
					inputStatus := strings.TrimSpace(ScanUserInput())
					switch inputStatus {
					case "1":
						status = tasks.StatusTodo
					case "2":
						status = tasks.StatusInProgress
					case "3":
						status = tasks.StatusDone
					default:
						fmt.Println("Ошибка! Введите номер статуса!")
						continue
					}
					break
				}
				if _, err := store.ChangeStatus(intId, status); err != nil {
					fmt.Println("ошибка обновления статуса: ", err)
					continue
				}
				fmt.Println("Статус успешно обновлен!")
				return
			} else if input == "2" {
				return
			} else {
				fmt.Println("Введите 1 или 2: ")
				continue
			}
		}
	}
}

func getAll(store *tasks.TaskStore) {
	for _, task := range store.Tasks {
		fmt.Printf("\nID задачи: %d\nОписание задачи: %s\nСтатус выполнения: %s\nДата создания: %s\nДата обновления: %s\n", task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
	}
}

func getAllDone(store *tasks.TaskStore) {
	for _, task := range store.Tasks {
		if task.Status == tasks.StatusDone {
			fmt.Printf("\nID задачи: %d\nОписание задачи: %s\nСтатус выполнения: %s\nДата создания: %s\nДата обновления: %s\n", task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		}
	}
}

func getAllTodo(store *tasks.TaskStore) {
	for _, task := range store.Tasks {
		if task.Status == tasks.StatusTodo {
			fmt.Printf("\nID задачи: %d\nОписание задачи: %s\nСтатус выполнения: %s\nДата создания: %s\nДата обновления: %s\n", task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		}
	}
}

func getAllInProgress(store *tasks.TaskStore) {
	for _, task := range store.Tasks {
		if task.Status == tasks.StatusInProgress {
			fmt.Printf("\nID задачи: %d\nОписание задачи: %s\nСтатус выполнения: %s\nДата создания: %s\nДата обновления: %s\n", task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
		}
	}
}

func writeFile(filename string, store *tasks.TaskStore) {
	jsonData, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		log.Fatalf("Ошибка сериализации: %v", err)
	}

	if err = os.WriteFile(filename, jsonData, 0644); err != nil {
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
