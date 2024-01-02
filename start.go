package main

import (
	"context"
	"fmt"
	"log"

	"entdemo/ent"
	"entdemo/ent/car"
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


    QueryCars(context.Background(), a8m[0])
    QueryCars(context.Background(), a8m[1])
}


func QueryCars(ctx context.Context, a8m *ent.User) error {
    cars, err := a8m.QueryCars().All(ctx)
    if err != nil {
        return fmt.Errorf("failed querying user cars: %w", err)
    }
    log.Println("returned cars:", cars)

    // What about filtering specific cars.
    ford, err := a8m.QueryCars().
        Where(car.Model("Ford")).
        Only(ctx)
    if err != nil {
        return fmt.Errorf("failed querying user cars: %w", err)
    }
    log.Println(ford)
    return nil
}