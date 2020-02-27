package todolist

import "time"

const (
	STATUS_PENDING = 0
	STATUS_DONE = 1
)

type Task struct {
	Id 			int			`json:"id"`
	Uid 		int			`json:"-"`
	Title 		string 		`json:"title"`
	Status 		int			`json:"status"`
	Updated_At	time.Time	`json:"-"`
	Created_At  time.Time	`json:"-"`
}