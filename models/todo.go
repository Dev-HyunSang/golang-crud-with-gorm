package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type ToDo struct {
	ToDoUUID    uuid.UUID `json:"todo_uuid" gorm:"column:todo_uuid;"`
	Content     string    `json:"content"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}
