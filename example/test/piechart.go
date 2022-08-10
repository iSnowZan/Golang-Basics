	package test

	import (
		"os"
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
		"github.com/go-echarts/go-echarts/v2/charts"
		"github.com/go-echarts/go-echarts/v2/opts"
		"github.com/go-echarts/go-echarts/v2/types"

	)

	const (
		host4     = "localhost"
		port4    = 5432
		user4     = "postgres"
		password4 = "1234"
		dbname4   = "postgres"
	)

	type stud struct {
		name string
		marks int
	}

	func Piechart(ctx context.Context)(*model.PieResult, error) {
		psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host4, port4, user4, password4, dbname4)
		db, err := sql.Open("postgres", psqlinfo)
		if err != nil {
			panic(err)
		}
		defer db.Close()


	var sta string

	studobj := [5]stud{}

	studpie:=make([]opts.PieData,0)

	fetchrow:=`SELECT name , marks FROM pujan_db.pie`

	rows,err:=db.Query(fetchrow)
	if err!=nil{
		return nil,err
	}

	for _,s := range studobj{
		for rows.Next(){
			err:=rows.Scan(&s.name,&s.marks)
				if err!=nil{
					return nil,err
				}	
			studpie=append(studpie, opts.PieData{
				Name : s.name,
				Value : s.marks,
			})
				
		}
	}

	pie := charts.NewPie()

	pie.SetGlobalOptions(charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeChalk}),
	charts.WithTitleOpts(opts.Title{Title:"Pie Graph : Distribution in Marks"}),
	)

	pie.AddSeries("Mark",studpie)

	f,e:=os.Create("Pie.html")

	if e!=nil{
		sta="Failed"
	}else{
		sta="Success"
	}


	pie.Render(f)

	return &model.PieResult{Status: &sta},nil
	}