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

    QueryUser(context.Background(), client)
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
    u, err := client.User.
        Query().
        Where(user.Name("a8m")).
        // `Only` fails if no user found,
        // or more than 1 user returned.
        Only(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed querying user: %w", err)
    }
    log.Println("user returned: ", u)
    return u, nil
}