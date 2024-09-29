package testing

import (
	"log"
	"math/rand"
	"time"
	"votingapi/src/postgres"
)

func StressTest() {
	time.Sleep(5 * time.Second)
	db, err := postgres.Database()
	if err != nil {
		log.Fatal(err)
		return
	}

	for i := range 100000 {
		go func() {

			tx, err := db.Beginx()
			if err != nil {
				log.Fatal(err)
				return
			}

			_, err = tx.Exec("INSERT INTO test_insert (number1, str) VALUES ($1, 'test');", rand.Int()%10)
			if err != nil {
				log.Fatal(err)
				return
			}

			_, err = tx.Exec("SELECT * FROM test_insert limit 100 offset $1;", i)
			if err != nil {
				log.Fatal(err)
				return
			}

			err = tx.Commit()
			if err != nil {
				log.Fatal(err)
				return
			}
			if i%1000 == 0 {
				log.Println(i)
			}
		}()
	}

}
