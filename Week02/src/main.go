package main

import (
	"fmt"
	"service"
)

func main() {
	s := service.NewService()
	_, err := s.GetUsernameByUserId(7)
	fmt.Printf("%+v\n", err)
}
