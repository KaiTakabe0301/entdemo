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

    QueryGroupWithUsers(context.Background(), client)

}

func QueryGroupWithUsers(ctx context.Context, client *ent.Client) error {
    groups, err := client.Group.
        Query().
        Where(group.HasUsers()).
        All(ctx)
    if err != nil {
        return fmt.Errorf("failed getting groups: %w", err)
    }
    log.Println("groups returned:", groups)
    // Output: (Group(Name=GitHub), Group(Name=GitLab),)
    return nil
}