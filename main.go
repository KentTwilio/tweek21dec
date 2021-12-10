package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("[main]\n")

	http.HandleFunc("/hello", hello)

	fmt.Printf("[main] listening...\n")
	http.ListenAndServe(":8888", nil)

	fmt.Printf("[main] end\n")
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("[hello]\n")

	printHeaders(req)

	fmt.Fprintf(w, "hello world\n")
}

func printHeaders(req *http.Request) {
	fmt.Printf("[printHeaders]\n")

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Printf("%v: %v\n", name, h)
		}
	}
}

// modules
// seating:
//			bar:
//				available seats
//				waiting list
//			tables:
//				time slots
//					reserved
//		Load - load the state
//		Save - save the state
//		Reserve -
// users:
//			name
//			id
//			phone
// 		Create
//		Delete
// commands:
//		help [<command>]
//		reserve <table/bar> <# of people> <time>
//		show
//		cancel <need a way to uniquely identify reservations>
// interactor:
//		help [<command>]
//		reserve - collect each parameter all at once or in order
//			res table - restart process
//		show
//		cancel - <parameter> or show reservations
// system:
//		command
//			advance time
//			remove from bar wait list
