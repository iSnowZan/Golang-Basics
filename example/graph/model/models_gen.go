// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Inpu struct {
	Inp *string `json:"inp_"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AddResult struct {
	N3 *int `json:"n3"`
}

type DataInput struct {
	ID         *string `json:"id"`
	Firstname  string  `json:"firstname"`
	Lastname   *string `json:"lastname"`
	Dob        string  `json:"dob"`
	Phno       *string `json:"phno"`
	Bloodgroup string  `json:"bloodgroup"`
	Address    string  `json:"address"`
	Gender     *string `json:"gender"`
}

type DataOutput struct {
	Status *string `json:"status"`
}

type DataResponse struct {
	ID     *string   `json:"id"`
	Status []*string `json:"Status"`
}

type Ixz struct {
	Output *string `json:"output"`
}

type PieResult struct {
	Status *string `json:"status"`
}

type Sol struct {
	N1 *int    `json:"n1"`
	N2 *int    `json:"n2"`
	A  *string `json:"a"`
}

type StuOutput struct {
	ID     *string `json:"id"`
	Fname  *string `json:"fname"`
	Lname  *string `json:"lname"`
	Dateob *string `json:"dateob"`
	Phone  *string `json:"phone"`
	Blgrp  *string `json:"blgrp"`
	Addr   *string `json:"addr"`
	Gen    *string `json:"gen"`
}

type Xloutput struct {
	Status *string `json:"status"`
}
