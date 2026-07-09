package models

import "time"

type TickerType int

const (
	TickerInterval TickerType = 0
	TickerTime     TickerType = 1
)

type Task struct {
	Name    string
	Prefix  string
	Handler func()

	TickerType           TickerType
	Interval             time.Duration
	Hour, Minute, Second int
}

func (task *Task) NameWithPrefix() string {
	return task.Prefix + task.Name
}
