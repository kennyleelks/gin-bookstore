package models

import "github.com/guregu/null"

type Book struct {
	ID     uint        `json:"id" gorm:"primary_key"`
	Title  string      `json:"title"`
	Author null.String `json:"author"`
	ISBN   null.String `json:"isbn"`
}
