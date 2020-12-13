package ent

import (
	"fmt"

	"github.com/godpeny/goserv/ent"
)

var client *ent.Client

func Setup(cli *ent.Client) {
	client = cli
}

func GetClient() (*ent.Client, error) {
	if client == nil {
		return client, fmt.Errorf("no client set")
	}
	return client, nil
}
