package main

import (
	"fmt"
    "log"
	"os"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)


type Info struct {
	Employee_ID uint64 `pg:"employee_id"`
	Manager_ID  uint64 `pg:"manager_id"`
}

func main(){
 
	 db:=Connect()
	 var Details []Info

	 err:=db.Model(&Details).Select()
	 
	 if err!=nil{
      
		return ;
	 }

	 adjlist:=make(map[uint64][]uint64)

	 for _,detail:=range Details{

		adjlist[detail.Employee_ID]=append(adjlist[detail.Employee_ID],detail.Manager_ID)
	 }

	 var emp_id,manager_id uint64

	 fmt.Println("Please enter the employee id")
	 fmt.Scanf("%v",&emp_id)
	 fmt.Println("Please enter the manager id")
	 fmt.Scanf("%v",&manager_id)

	 visited:=make(map[uint64]bool)
	 iscyclic := detectcycle(visited,adjlist,emp_id,manager_id)
     
	 if iscyclic{
		fmt.Printf("Yes this will result in cycle")
	 }else{
        fmt.Printf("This will not result cycle")
	 }

}

func detectcycle( visited map[uint64]bool,adjlist map[uint64][]uint64,emp_id,manager_id uint64)(bool){
	if visited[manager_id] {
		return true
	}

	visited[emp_id] = true;
	visited[manager_id] = true;

	for _, manager := range adjlist[manager_id] {
		if detectcycle(visited,adjlist,manager_id,manager) {
			return true
		}
	}

	return false

}

func Connect() *pg.DB {

	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "1234",
		Database: "pranav",
	})

	if db == nil {
		log.Print("unable to  connect")
		os.Exit(100)
	}

	log.Print("connection to database was successful")

	
		opts := &orm.CreateTableOptions{
			IfNotExists: true,
		}
	
		if err := db.Model(&Info{}).CreateTable(opts); err != nil {
	
			fmt.Println("error while creating product table:", err)
	
		}
	
		fmt.Println("table was created successfully")

	return db
}

