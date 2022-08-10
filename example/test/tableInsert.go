package test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"example/graph/model"

	//"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

func TableInsert(ctx context.Context, input *model.DataInput) (*model.DataResponse, error) {
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	resultdata := &model.DataResponse{}
	_, e := IsValidDateWithDateObject(input.Dob)
	log.Println("Inside func number", *input.Phno)
	// var phoneNumber int
	// phoneNumber, _ = strconv.Atoi(*input.Phno)
	// if err != nil {
	// 	panic(err)
	// }
	validPhn := checkPhno(*input.Phno)
	input.Bloodgroup = strings.ToUpper(input.Bloodgroup)
	validBloodGroup := checkBlood(input.Bloodgroup)
	*input.Gender = strings.ToUpper(*input.Gender)
	validGender := checkGender(*input.Gender)
	if e == nil && validPhn && validBloodGroup && validGender && input.ID == nil {
		insertStd := `insert into pujan_db.Student("firstname","lastname","dob","phno","bloodgroup","address","gender")values($1,$2,$3,$4,$5,$6,$7)RETURNING ID`
		log.Println("Write Query")
		err := db.QueryRowContext(ctx, insertStd, input.Firstname, input.Lastname, input.Dob, input.Phno, input.Bloodgroup, input.Address, input.Gender).Scan(&input.ID)
		if err != nil {
			log.Println("error in query", err)
			resultdata.ID = nil
			var output1 string = "Failed to insert data"
			resultdata.Status = append(resultdata.Status, &output1)
			return resultdata, err
		}
		log.Println("continue")
		resultdata.ID = input.ID
		var output2 string = "Success - inserted into table"
		resultdata.Status = append(resultdata.Status, &output2)
		return resultdata, nil
	} else if input.ID != nil {
		updateStd := `update pujan_db.student set firstname=$1, lastname=$2, dob=$3,phno=$4,bloodgroup=$5,address=$6,gender=$7 where id=$8`
		_, err := db.Exec(updateStd, input.Firstname, input.Lastname, input.Dob, input.Phno, input.Bloodgroup, input.Address, input.Gender, input.ID)
		if err != nil {
			resultdata.ID = nil
			var output1 string = "Failed to update data"
			resultdata.Status = append(resultdata.Status, &output1)
			return resultdata, err
		}
		resultdata.ID = input.ID
		var output2 string = "Success - Updated into table"
		resultdata.Status = append(resultdata.Status, &output2)
		return resultdata, nil
	}

	resultdata.ID = nil
	if e != nil {
		var err1 string = "Failed Date Format"
		resultdata.Status = append(resultdata.Status, &err1)
	}
	if !validPhn {
		var err2 string = "Failed Phone Number Format"
		resultdata.Status = append(resultdata.Status, &err2)
	}

	if !validBloodGroup {
		var err3 string = "Failed Blood Group Format"
		resultdata.Status = append(resultdata.Status, &err3)
	}

	if !validGender {
		var err4 string = "Failed Gender Format"
		resultdata.Status = append(resultdata.Status, &err4)
	}
	return resultdata, nil
}
func IsValidDateWithDateObject(date string) (time.Time, error) {
	const (
		layoutISO = "2006-01-02"
	)
	timeobject, err := time.Parse(layoutISO, date)
	if err != nil {
		return timeobject, err
	}
	return timeobject, nil
}
func checkPhno(phno string) bool {
	log.Println("Inside func number")
	//m := strconv.Itoa(phno)
	log.Println("string of num is", phno)
	if len(phno) == 10 {
		log.Println("true")
		return true
	}
	return false
}
func checkBlood(blood string) bool {

	bloodmap := make(map[string]string)

	bloodmap["Group1"] = "A+"
	bloodmap["Group2"] = "A-"
	bloodmap["Group3"] = "B+"
	bloodmap["Group4"] = "B-"
	bloodmap["Group5"] = "AB+"
	bloodmap["Group6"] = "AB-"
	bloodmap["Group7"] = "O+"
	bloodmap["Group8"] = "O+"

	for _, bgroup := range bloodmap {
		if blood == bgroup {
			return true
		}
	}
	return false
}

func checkGender(gender string) bool {
	gendermap := make(map[string]string)
	gendermap["type1"] = "MALE"
	gendermap["type2"] = "FEMALE"
	for _, gendtype := range gendermap {
		if gender == gendtype {
			return true
		}
	}
	return false
}