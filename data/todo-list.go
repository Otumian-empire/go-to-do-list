package data

type TodoList struct {
	tasks []Todo
}

func NewTodoList() *TodoList {
	return &TodoList{}
}

func (list *TodoList) Add(todo Todo) {
	list.tasks = append(list.tasks, todo)
}

func (list *TodoList) Read() []Todo {
	return list.tasks
}

func (list *TodoList) ReadOne(index int) Todo {
	return list.tasks[index]
}

func (list *TodoList) Update(index int, todo Todo) {
	list.tasks[index] = todo
}

/* func (list *TodoList) Delete(index int) error {

	if index > len(list.Read()) {
		return fmt.Errorf("enter a valid numeric index from the TODO List")
	}

	// split the todo list at the index and combine the other 2 parts
	list.tasks = append(list.tasks[:index], list.tasks[index+1:]...)
	return nil
} */

func (list *TodoList) Delete(index int) {
	// split the todo list at the index and combine the other 2 parts
	list.tasks = append(list.tasks[:index], list.tasks[index+1:]...)

}
