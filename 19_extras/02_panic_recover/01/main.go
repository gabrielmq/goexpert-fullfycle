package main

import "fmt"

func main() {
	defer func() {
		// fica monitorando os panics para tomar alguma ação caso ocorram panics
		// para evitar que a aplicação termine abruptamente
		if r := recover(); r != nil {
			fmt.Printf("recovered in main: %v", r)
		}
	}()
	myPanic()
}

func myPanic() {
	panic("something went wrong")
}
