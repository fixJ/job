package store

type Factory interface {
	task() Task
}
