package model

import "time"

type Notes struct {
	NotesId     uint      `gorm:"primaryKey;column:notes_id;autoIncrement" json:"notes_id"`
	Title       string    `gorm:"size:100;not null;column:title" json:"title"`
	Content     string    `gorm:"size:100;not null;column:content" json:"content"`
	CreatedTime time.Time `gorm:"autoCreateTime" json:"time"`
}

type NotesID struct {
	NotesId uint `json:"notes_id" binding:"required"`
}
