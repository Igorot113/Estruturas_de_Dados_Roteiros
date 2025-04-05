package main

import "fmt"

type Task struct {
	name string
	next *Task
}

type TaskList struct {
	head *Task
}

func (t *TaskList) AddTask(name string) {
	newTask := &Task{name: name}
	if t.head == nil {
		t.head = newTask
		return
	}
	current := t.head
	for current.next != nil {
		current = current.next
	}
	current.next = newTask
}

func (t *TaskList) CompleteTask() {
	if t.head != nil {
		t.head = t.head.next
	}

}

func (t *TaskList) ShowTasks() {
	current := t.head
	for current != nil {
		fmt.Print(current.name, " ->")
		current = current.next
	}
	fmt.Println("nil")

}

func main() {
	lista := TaskList{}

	lista.AddTask("Estudar Go")
	lista.AddTask("Fazer compras")
	lista.AddTask("Academia")
	lista.ShowTasks()

	lista.CompleteTask()
	lista.ShowTasks()

	lista.CompleteTask()
	lista.ShowTasks()
}
