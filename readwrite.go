//sGolang program to read and write the files
package main

// importing the packages
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type JSON struct {
	jmap map[string]interface{}
}

func New() (newObject JSON) {
	m := make(map[string]interface{})
	newObject.jmap = m
	return newObject
}
func AddLogs() {
	fmt.Printf("Writing to a file in Go lang\n")
	file, err := os.Create("/data/logs.txt")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()
	len, err := file.WriteString("This is a read/write file")
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}
func ReadLogs() {
	fmt.Printf("\n\nReading a file in Go lang\n")
	fileName := "/data/logs.txt"
	data, err := ioutil.ReadFile("/data/logs.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", fileName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)
}
func WriteData(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	fmt.Println("@@")
	AddLogs()
	response := "writing logs to logs.txt file"
	// response.Put("message", "Logs Added Successfully")
	w.Write([]byte(response))
}
func ReadData(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	ReadLogs()
	response := "reading logs from logs.txt file"
	//response.Put("message", "Logs Readed Successfully")
	w.Write([]byte(response))

}
func handleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/write", WriteData).Methods("POST")
	router.HandleFunc("/read", ReadData).Methods("GET")
	log.Fatal(http.ListenAndServe(":8089", router))
}

// main function
func main() {
	handleRequest()
}
