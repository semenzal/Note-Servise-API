package main

import (
	"context"
	"log"

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

	res_GetNote, err := client.GetNote(context.Background(), &desc.GetNoteRequest{
		Nota: "Hello!",
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Id:", res_GetNote.Id)

	res_GetList, err := client.GetListNote(context.Background(), &desc.GetListNoteRequest{
		AllId: "All Id base",
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("All Id:", res_GetList.Id)

	res_Update, err := client.UpdateNote(context.Background(), &desc.UpdateNoteRequest{
		Update: "New Id",
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("New Id:", res_Update.Id)

	res_Delete, err := client.DeleteNote(context.Background(), &desc.DeleteNoteRequest{
		Delete: "Empty",
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Delete:", res_Delete.Id)
}
