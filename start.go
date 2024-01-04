package main

import (
	"context"
	"fmt"
	"log"

	"entdemo/ent"
	"entdemo/ent/group"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=entdemo password=password sslmode=disable")
    if err != nil {
        log.Fatalf("failed opening connection to postgres: %v", err)
    }
    defer client.Close()
    // Run the auto migration tool.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }

    QueryGithub(context.Background(), client)

}

func QueryGithub(ctx context.Context, client *ent.Client) error {
    cars, err := client.Group.
        Query().
        Where(group.Name("GitHub")). // (Group(Name=GitHub),)
        QueryUsers().                // (User(Name=Ariel, Age=30),)
        QueryCars().                 // (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
        All(ctx)
    if err != nil {
        return fmt.Errorf("failed getting cars: %w", err)
    }
    log.Println("cars returned:", cars)
    // Output: (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
    return nil
}