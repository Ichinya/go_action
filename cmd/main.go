package cmd

import (
	"github.com/Ichinya/go_action/apiserver"
	"log"
)

func main() {
	config := apiserver.NewConfig()
	s := apiserver.New(config)

	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}

}
