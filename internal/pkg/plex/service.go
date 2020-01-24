package plex

import (
	"encoding/json"
	"fmt"
	"os"
)

type Service interface {
	PostListen(interface{})
	GetListen()
}
type service struct {
}

func NewPlexService() service {
	return service{}
}

func (s *service) PostListen(hook interface{}) {
	fmt.Printf("recieved POST hook:\n")
	reporter := json.NewEncoder(os.Stdout)
	reporter.SetEscapeHTML(true)
	reporter.SetIndent("", "    ")
	reporter.Encode(hook)
}

func (s *service) GetListen() {
	fmt.Printf("recieved GET hook\n")
}
