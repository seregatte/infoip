package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	modePtr := flag.String("mode", "full", "Available mode external and full")
	flag.Parse()
	if modePtr != nil && *modePtr == "full" {
		fmt.Printf("Mode %s \n", *modePtr)
	}
	if externalIp, err := external(); err == nil {
		fmt.Printf("External IP: %s", externalIp)
	} else {
		fmt.Errorf("Erro: %s", err)
	}
}

func external() (string, error) {
	response, err := http.Get("http://ipinfo.io/ip")
	if err != nil {
		return "<nil>", err
	}

	defer response.Body.Close()
	if response.StatusCode == 200 { // OK
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		bodyString := string(bodyBytes)
		return bodyString, nil
	} else {
		return "<nil>", fmt.Errorf("Erro http n %s", response.StatusCode)
	}
}
