package model

import "time"


type Todo struct {
	//使用的ID属性
	ID int `json:"id" db:"id"`
	//使用的Title属性
	Title string `json:"title" db:"title"`
	//使用的Description
	Description string `json:"description" db:"description"`
	//是否完成
	Completed bool `json:"completed" db:"completed"`
	//通用
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}