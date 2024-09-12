package entity

import (
	"time"
)

/*
Since variables have a 0 default value, you should usually start your enums on a non-zero value.
There are cases where using the zero value makes sense, for example when the zero value case is the
desirable default behavior.
*/

type Status uint8

const (
	Pending Status = iota + 1
	InProgress
	Completed
)

const (
	PendingStr    = "pending"
	InProgressStr = "in_progress"
	CompletedStr  = "completed"
)

func (s Status) String() string {
	switch s {
	case Pending:
		return PendingStr
	case InProgress:
		return InProgressStr
	case Completed:
		return CompletedStr
	}

	return ""
}

func MapToStatusTaskEntity(status string) Status {
	switch status {
	case PendingStr:
		return Pending
	case InProgressStr:
		return InProgress
	case CompletedStr:
		return Completed
	}

	return Status(0)
}

type Task struct {
	ID          uint
	UserID      uint
	Title       string
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
