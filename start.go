package main

import (
	"context"
	"fmt"
	"log"

	"entdemo/ent"
	"entdemo/ent/user"

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

    a8m, err := client.User.Query().Where(user.NameEQ("a8m")).All(context.Background())
    if err != nil {
        log.Fatalf("failed querying user: %v", err)
    }


    QueryCarUsers(context.Background(), a8m[1])
}


func QueryCarUsers(ctx context.Context, a8m *ent.User) error {
    cars, err := a8m.QueryCars().All(ctx)
    if err != nil {
        return fmt.Errorf("failed querying user cars: %w", err)
    }
    // Query the inverse edge.
    for _, c := range cars {
        owner, err := c.QueryOwner().Only(ctx)
        if err != nil {
            return fmt.Errorf("failed querying car %q owner: %w", c.Model, err)
        }
        log.Printf("car %q owner: %q\n", c.Model, owner.Name)
    }
    return nil
}