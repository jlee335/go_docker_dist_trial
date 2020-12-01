//This will recieve user input and send request package to the cache server.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	UID := r.FormValue("name")
	num_str := r.FormValue("address")
	send_back_request(w, UID, num_str)
}

//concurrent function that sends back request to waiting user!
func send_back_request(w http.ResponseWriter, UID string, num string) {
	// :calc-server
	if num == "" {
		num = "1"
	}
	if UID == "" {
		UID = "1"
	}
	req, err := http.NewRequest("GET", "http://calc-server:8081", nil)
	if err != nil {
		fmt.Fprintf(w, "crucial error")
		return
	}
	q := req.URL.Query()
	q.Add("query", num)
	q.Add("UID", UID)
	//req.URL.RawQuery = req.URL.RawQuery + q.Encode()

	fmt.Fprintf(w, "request sent to enc::")
	fmt.Fprintf(w, req.URL.RawQuery)
	fmt.Println(req.URL)
	c := http.DefaultClient
	resp, err := c.Do(req)

	if err != nil {
		fmt.Fprintf(w, "Server is down!!: =>")
		fmt.Fprintf(w, req.URL.String())
		return
	}

	responseData, _ := ioutil.ReadAll(resp.Body)

	fmt.Fprintf(w, "\nGOT RESPONSE\n")
	fmt.Fprintf(w, string(responseData))

}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // New code
	http.Handle("/", fileServer)                        // New code
	fmt.Printf("Starting server at port 8080\n")

	http.HandleFunc("/form", formHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
