package main

import (
	"fmt"
	"tasktracker/tasks"

	"github.com/k0kubun/pp"
)

const mainMenu = ("Выберите действие(введите цифру в консоль):\n1. Добавить задачу.\n2. Обновить задачу.\n3. Удалить задачу.\n4. Поменять статус задачи.\n5. Показать все задачи.\n6. Показать все выполненные задачи.\n7. Показать все невыполненные задачи.\n8. Показать все задачи в процессе выполнения.\n9. Выход из программы.")

func main() {
	store := tasks.NewTaskStore()
	newStore, err := store.NewTask("do")
	if err != nil {
		fmt.Printf("ошибка создания задачи: %v\n", err)
	}
	newStore, err = store.NewTask("do anything")
	if err != nil {
		fmt.Printf("ошибка создания задачи: %v\n", err)
	}
	newStore, err = store.NewTask("do nothing")
	if err != nil {
		fmt.Printf("ошибка создания задачи: %v\n", err)
	}
	newStore, err = store.NewTask("do")
	if err != nil {
		fmt.Printf("ошибка создания задачи: %v\n", err)
	}

	newStore, err = newStore.UpdateTask(4, "qwertyuiop")
	if err != nil {
		fmt.Printf("ошибка обновления задачи: %v\n", err)
	}

	newStore, err = newStore.DeleteTask(2)
	if err != nil {
		fmt.Printf("ошибка обновления задачи: %v\n", err)
	}
	pp.Println(newStore)
}
