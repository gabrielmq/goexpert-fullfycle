package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
)

type DBConfig struct {
	DB       string `json:"db"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

var dbConfig DBConfig

func MarshalConfig(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(data, &dbConfig); err != nil {
		panic(err)
	}
}

// FSNotify é uma lib para recarregar arquivos que contenham variaveis de ambiente
// em casos em que essas variaveis tenham seus valores alterados enquanto a aplicação esta em executação
// essa lib evita que tenhamos que reiniciar a aplicação para pegar os novos valores das variaveis
func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	MarshalConfig("config.json")

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Println("event: ", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					MarshalConfig("config.json")
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("error:", err)
			}
		}
	}()

	if err := watcher.Add("config.json"); err != nil {
		panic(err)
	}
	<-done
}
