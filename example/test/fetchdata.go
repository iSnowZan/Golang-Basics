package test

import (
	"context"
	"database/sql"
	"example/graph/model"
	"fmt"

	//"log"

	//"strconv"
	//"strings"
	//"time"

	_ "github.com/lib/pq"
	//"golang.org/x/tools/go/analysis/passes/nilfunc"
)

const (
	host1     = "localhost"
	port1     = 5432
	user1     = "postgres"
	password1 = "1234"
	dbname1   = "postgres"
)

func Fetchdata(_ context.Context, input *model.Ixz) ([]*model.StuOutput, error) {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host1, port1, user1, password1, dbname1)
	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if input.Output != nil {
		searchSTD := `SELECT* FROM pujan_db.student WHERE firstname LIKE '%` + *input.Output + `%'`
		r, err := db.Query(searchSTD)
		rentry := []*model.StuOutput{}

		if err != nil {
			return nil, err
		}

		for r.Next() {
			eachval := &model.StuOutput{}
			error1 := r.Scan(&eachval.ID, &eachval.Fname, &eachval.Lname, &eachval.Dateob, &eachval.Phone, &eachval.Blgrp, &eachval.Addr, &eachval.Gen)
			if error1 != nil {
				return nil, error1
			}

			if eachval.ID == nil {
				xyz := " "
				eachval.ID = &xyz
			}

			rentry = append(rentry, eachval)
		}
		return rentry, nil
	}
	searchSTD := `SELECT* FROM pujan_db.student`
	r, err := db.Query(searchSTD)
	rentry := []*model.StuOutput{}

	if err != nil {
		return nil, err
	}

	for r.Next() {
		eachval := &model.StuOutput{}
		error1 := r.Scan(&eachval.ID, &eachval.Fname, &eachval.Lname, &eachval.Dateob, &eachval.Phone, &eachval.Blgrp, &eachval.Addr, &eachval.Gen)
		if error1 != nil {
			return nil, error1
		}
		rentry = append(rentry, eachval)
	}
	return rentry, nil
}
