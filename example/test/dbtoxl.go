package test

import (
	"context"
	"database/sql"
	"example/graph/model"
	"fmt"
	"strconv"

	//"log"

	//"strconv"
	//"strings"
	//"time"

	//"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/lib/pq"
	//"golang.org/x/tools/go/analysis/passes/nilfunc"
	//"golang.org/x/tools/go/analysis/passes/nilfunc"
	"github.com/xuri/excelize/v2"
)

const (
	host3     = "localhost"
 	port3     = 5432
 	user3     = "postgres"
 	password3 = "1234"
 	dbname3   = "postgres"
 )

 func Dbtoxl(ctx context.Context,input *model.Inpu)(*model.DataOutput, error) {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host3, port3, user3, password3, dbname3)
 	db, err := sql.Open("postgres", psqlinfo)
 	if err != nil {
 		panic(err)
 	}
 	defer db.Close()


var sta string

type x struct{
	ID sql.NullString
	firstname string
	lastname string
	dob string
	phno string
	bloodgroup string
	address string
	gender string
}

f:=excelize.NewFile()
var xlarr = [30]x{}
s:=f.NewSheet("student")
fmt.Println(s)
f.SetCellValue("student","A2","ID")
f.SetCellValue("student","B2","Firstname")
f.SetCellValue("student","C2","Lastname")
f.SetCellValue("student","D2","Dob")
f.SetCellValue("student","E2","Phno")
f.SetCellValue("student","F2","Bloodgroup")
f.SetCellValue("student","G2","Address")
f.SetCellValue("student","H2","Gender")

if (*input.Inp!=""){
	fetchRow:=`SELECT* FROM pujan_db.student WHERE firstname LIKE '%` + *input.Inp + `%'`
	rows,err := db.Query(fetchRow)
	if err!=nil{
		return nil,err
	}
	var rnumber int = 2	
	for _,column:=range xlarr{
		for rows.Next(){
		err:=rows.Scan(&column.ID,&column.firstname,&column.lastname,&column.dob,&column.phno,&column.bloodgroup,&column.address,&column.gender)
	if err!=nil{
		return nil,err
	}
	if !column.ID.Valid{
		f.SetCellValue("student","A"+strconv.Itoa(rnumber),"")
	}else{
		f.SetCellValue("student","A"+strconv.Itoa(rnumber),column.ID.String)	
	}
	f.SetCellValue("student","B"+strconv.Itoa(rnumber),column.firstname)
	f.SetCellValue("student","C"+strconv.Itoa(rnumber),column.lastname)

	f.SetCellValue("student","D"+strconv.Itoa(rnumber),column.dob)
	f.SetCellValue("student","E"+strconv.Itoa(rnumber),column.phno)
	f.SetCellValue("student","F"+strconv.Itoa(rnumber),column.bloodgroup)
	f.SetCellValue("student","G"+strconv.Itoa(rnumber),column.address)
	f.SetCellValue("student","H"+strconv.Itoa(rnumber),column.gender)

	rnumber++
	}
	
}

}else{fetchRow:=`SELECT* FROM pujan_db.student`
rows,err := db.Query(fetchRow)
if err!=nil{
	return nil,err
}
var rnumber int = 2	
for _,column:=range xlarr{
	for rows.Next(){
	err:=rows.Scan(&column.ID,&column.firstname,&column.lastname,&column.dob,&column.phno,&column.bloodgroup,&column.address,&column.gender)
if err!=nil{
	return nil,err
}
if !column.ID.Valid{
	f.SetCellValue("student","A"+strconv.Itoa(rnumber),"")
}else{
	f.SetCellValue("student","A"+strconv.Itoa(rnumber),column.ID.String)	
}
f.SetCellValue("student","B"+strconv.Itoa(rnumber),column.firstname)
f.SetCellValue("student","C"+strconv.Itoa(rnumber),column.lastname)

f.SetCellValue("student","D"+strconv.Itoa(rnumber),column.dob)
f.SetCellValue("student","E"+strconv.Itoa(rnumber),column.phno)
f.SetCellValue("student","F"+strconv.Itoa(rnumber),column.bloodgroup)
f.SetCellValue("student","G"+strconv.Itoa(rnumber),column.address)
f.SetCellValue("student","H"+strconv.Itoa(rnumber),column.gender)

rnumber++

}

}
}
er:= f.SaveAs("/home/nitsd040/go/DbToXl.xlsx")
if er != nil{
	fmt.Println(er)
	sta = "Failed"
}else{
	sta = "Success"
}

return &model.DataOutput{Status: &sta},nil 
 }