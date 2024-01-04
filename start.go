package main

import (
	"context"
	"log"
	"time"

	"entdemo/ent"

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

    if err := CreateGraph(context.Background(), client); err != nil {
        log.Fatalf("failed creating graph: %v", err)
    }
}


func CreateGraph(ctx context.Context, client *ent.Client) error {
    // First, create the users.
    a8m, err := client.User.
        Create().
        SetAge(30).
        SetName("Ariel").
        Save(ctx)
    if err != nil {
        return err
    }
    neta, err := client.User.
        Create().
        SetAge(28).
        SetName("Neta").
        Save(ctx)
    if err != nil {
        return err
    }
    // Then, create the cars, and attach them to the users created above.
    err = client.Car.
        Create().
        SetModel("Tesla").
        SetRegisteredAt(time.Now()).
        // Attach this car to Ariel.
        SetOwner(a8m).
        Exec(ctx)
    if err != nil {
        return err
    }
    err = client.Car.
        Create().
        SetModel("Mazda").
        SetRegisteredAt(time.Now()).
        // Attach this car to Ariel.
        SetOwner(a8m).
        Exec(ctx)
    if err != nil {
        return err
    }
    err = client.Car.
        Create().
        SetModel("Ford").
        SetRegisteredAt(time.Now()).
        // Attach this car to Neta.
        SetOwner(neta).
        Exec(ctx)
    if err != nil {
        return err
    }
    // Create the groups, and add their users in the creation.
    err = client.Group.
        Create().
        SetName("GitLab").
        AddUsers(neta, a8m).
        Exec(ctx)
    if err != nil {
        return err
    }
    err = client.Group.
        Create().
        SetName("GitHub").
        AddUsers(a8m).
        Exec(ctx)
    if err != nil {
        return err
    }
    log.Println("The graph was created successfully")
    return nil
}