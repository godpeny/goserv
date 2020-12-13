package init

import (
	"context"
	"log"

	md "github.com/godpeny/goserv/configs/model"
	"github.com/godpeny/goserv/ent"
	ent_client "github.com/godpeny/goserv/internal/clients/ent"
)

func InitModel() (*ent.Client, error) {
	client, err := ent.Open(md.DBDriveName, md.DBURL)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
		return client, err
	}
	ent_client.Setup(client)

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, nil
}
