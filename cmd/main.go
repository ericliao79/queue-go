package main

import (
	"fmt"
	"github.com/ericliao79/queue-go"
	"net/http"
)

func main() {
	count := 0

	queue_go.Run()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("eric")
		job := queue_go.Job{
			Uuid: "fewfpw",
			PayLoad: func(uuid string, data interface{}) {
				fmt.Println(uuid)
			},
		}
		queue_go.JobQueue <- job
		count++
	})

	http.ListenAndServe(":8081", nil)
}
