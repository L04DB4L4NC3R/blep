package main

import (
	"fmt"
	"log"

	"github.com/raindog308/gorqlite"
)

func main() {
	conn, err := gorqlite.Open("http://")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	res, err := conn.QueryOne("CREATE TABLE logs(log_id INT PRIMARY KEY AUTO_INCREMENT, timestamp VARCHAR(25), log VARCHAR(200), host VARCHAR(20))")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(res)

	// res, err := conn.WriteOne("INSERT INTO foo(name) VALUES('BAR')")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Printf("last insert id was %d\n", res.RowsAffected)

	// ress, err := conn.QueryOne("SELECT * FROM foo")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(ress)

}
