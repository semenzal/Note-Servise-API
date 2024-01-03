package main

import (
	"context"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"

	desc "github.com/semenzal/note-service-api/pkg/note_v1"
	"google.golang.org/grpc"
)

const address = "localhost:50051"

func main() {
	con, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("didn't connect: %s", err.Error())
	}
	defer con.Close()

	client := desc.NewNoteServiceClient(con)

	res, err := client.Create(context.Background(), &desc.CreateRequest{
		Title:  "Wow!",
		Text:   "I'm surprised!",
		Author: "Semen",
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Id:", res.Id)

	resGetNote, err := client.Get(context.Background(), &desc.GetRequest{
		Id: 2,
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Id:", resGetNote.String())

	resGetList, err := client.List(context.Background(), &empty.Empty{})
	if err != nil {
		log.Println(err.Error())
	}
	for _, note := range resGetList.Notes {

		fmt.Println(note.Title)
		fmt.Println(note.Text)
		fmt.Println(note.Author)
	}
	log.Println("All Id:", resGetList.Notes)

	resUpdate, err := client.Update(context.Background(), &desc.UpdateRequest{
		Title:  "New Wow!",
		Text:   "New I'm surprised!",
		Author: "New Semen",
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("New Id:", resUpdate.String())

	resDelete, err := client.Delete(context.Background(), &desc.DeleteRequest{
		Id: 12,
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Delete:", resDelete.String())
}
