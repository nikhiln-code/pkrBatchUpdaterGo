package helper

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	log "github.com/sirupsen/logrus"
)

const url = "https://localhsot:8080/put"

func sendPatchRequest(data string) {
	payload, err := json.Marshal(data)
	if err != nil {
		log.Error("error while parsing the json %s", err)
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Error("error while building the request %s ", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))

}