package main

import (
	"fmt"
	"log"
)

func main() {

	n := 0
	fmt.Println("Введите пожалуйста целое число: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Вы ввели число: ,", n)

}
