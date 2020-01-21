// shippy-service-vessel/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/micro/go-micro"
	pb "github.com/newmind/shippy-service-vessel/proto/vessel"
)

const (
	defaultHost = "mongodb://localhost:27017"
)

func createDummyData(repo repository) {
	vessels := []*Vessel{
		&Vessel{ID: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
		&Vessel{ID: "vesse88", Name: "hello ", MaxWeight: 200000, Capacity: 2},
	}
	for _, v := range vessels {
		repo.Create(context.Background(), v)
	}
}

func main() {
	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	vesselCollection := client.Database("shippy").Collection("vessel")

	repo := &VesselRepository{vesselCollection}
	createDummyData(repo)

	srv := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)

	srv.Init()

	// Register our implementation with
	pb.RegisterVesselServiceHandler(srv.Server(), &handler{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
