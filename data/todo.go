package data

type Todo struct {
	task        string
	isCompleted bool
}

func NewTodo(_task string) *Todo {
	return &Todo{
		task:        _task,
		isCompleted: false,
	}
}

func (todo *Todo) Edit(_task string) {
	todo.task = _task
}

func (todo *Todo) Complete() {
	todo.isCompleted = true
}

func (todo Todo) GetTask() string {
	return todo.task
}

func (todo Todo) GetIsCompleted() bool {
	return todo.isCompleted
}
