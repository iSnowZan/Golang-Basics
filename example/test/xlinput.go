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

	//"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/lib/pq"
	//"golang.org/x/tools/go/analysis/passes/nilfunc"
	"github.com/xuri/excelize/v2"
)

const (
	host2     = "localhost"
	port2     = 5432
	user2     = "postgres"
	password2 = "1234"
	dbname2   = "postgres"
)

func Xlinput(ctx context.Context) (*model.Xloutput, error) {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host2, port2, user2, password2, dbname2)
	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var sta string

	f, err := excelize.OpenFile("/home/nitsd040/go/exl.xlsx")
	if err != nil {
		return nil, err
	}

	defer f.Close()

	rows, err := f.GetRows("student")
	if err != nil {
		fmt.Println(err)
	}

	for row, col := range rows {
		if row == 0 {
			continue
		}
		if len(col) > 1 {
			
			if col[0] == "" {
				que := `insert into pujan_db.student("firstname","lastname","dob","phno","bloodgroup","address","gender")values($1,$2,$3,$4,$5,$6,$7)`
				_, err := db.Query(que, col[1], col[2], col[3], col[4], col[5], col[6], col[7])

				if err != nil {
					sta = "Failed to insert"

				}
				sta = "Success"
			} else if col[0] != "" {
				upd := `update pujan_db.student set firstname=$1,lastname=$2,dob=$3,phno=$4,bloodgroup=$5,address=$6,gender=$7 where id = $8`
				_, err = db.Exec(upd, col[1], col[2], col[3], col[4], col[5], col[6], col[7], col[0])

				if err != nil {
					sta = "Failed to update"
				}
				sta = "Success"

			}
		} else {
			del := `delete from pujan_db.student where id =$1`
			_, err = db.Exec(del, col[0])

			if err != nil {
				sta = "Failed to delete"
			}
			sta = "Success"
		}
	}

	return &model.Xloutput{Status: &sta}, nil

}
