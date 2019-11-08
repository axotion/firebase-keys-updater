package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const baseURL = "https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com"
const filename = "keys.json"

func main() {
	minutesPtr := flag.Int("every", 60, "Refresh file with public tokens every N minutes")
	flag.Parse()
	log.Printf("I will refresh tokens every %d minutes", *minutesPtr)
	updateKeys(*minutesPtr)
}


func updateKeys(everyMinutes int) {

	defer func() {
		if r := recover() ; r != nil {
			time.Sleep(time.Second * 5)
			log.Println(r)
			updateKeys(everyMinutes)
		}
	}()

	for {
		rawResponse, err := http.Get(baseURL)
		checkError(err)
		response, _ := ioutil.ReadAll(rawResponse.Body)
		ioutil.WriteFile(filename, response, 0644)
		log.Println("Refreshed")
		time.Sleep(time.Minute * time.Duration(everyMinutes))
	}
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}