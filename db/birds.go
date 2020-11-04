package db

import(

	_ "github.com/lib/pq"
)

func Birds() string {
	
	db := CreateConnection()
	rows, _ := db.Query("SELECT name FROM Birds")
	defer db.Close()	

	for rows.Next() {
            var name string
            _ = rows.Scan(&name)
            return name
     }


	
    return "";
}
