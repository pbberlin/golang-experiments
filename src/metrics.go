package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type ReceivedEvent struct {
	Name    string            `json:"name"`
	Payload map[string]string `json:"payload"`
}

type StoredEvent struct {
	Name      string
	Payload   string
	Timestamp time.Time
}

func process(receivedEvent ReceivedEvent) (StoredEvent, error) {
	var storedEvent StoredEvent
	var err error

	if storedPayload, err := json.Marshal(receivedEvent.Payload); err == nil {
		log.Println(string(storedPayload))

		storedEvent.Name = receivedEvent.Name
		storedEvent.Payload = string(storedPayload)
		storedEvent.Timestamp = time.Now()
	}

	return storedEvent, err
}

func handler(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var receivedEvent ReceivedEvent

	if err := decoder.Decode(&receivedEvent); err != nil {
		log.Println("Could not decode JSON")
		return
	}

	if storedEvent, err := process(receivedEvent); err != nil {
		log.Printf("Error turning a received event into a stored event: %v\n", receivedEvent)
	} else {
		log.Printf("%v\n", storedEvent)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
