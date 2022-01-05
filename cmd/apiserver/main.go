package main

import (
	"encoding/json"
	"flag"
	"github.com/ziyadovea/golang-http-rest-api/internal/app/apiserver"
	"log"
	"os"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.json", "path to config apiserver file")
}

func main() {

	flag.Parse()

	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	config := apiserver.NewConfig()
	err = json.Unmarshal(data, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
