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

    QueryArielCars(context.Background(), client)

}

func QueryArielCars(ctx context.Context, client *ent.Client) error {
    // Get "Ariel" from previous steps.
    a8m := client.User.
        Query().
        Where(
            user.HasCars(),
            user.Name("Ariel"),
        ).
        OnlyX(ctx)
    cars, err := a8m.                       // Get the groups, that a8m is connected to:
            QueryGroups().                  // (Group(Name=GitHub), Group(Name=GitLab),)
            QueryUsers().                   // (User(Name=Ariel, Age=30), User(Name=Neta, Age=28),)
            QueryCars().                    // (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
            Where(                          //
                car.Not(                    //  Get Neta and Ariel cars, but filter out
                    car.Model("Mazda"),     //  those who named "Mazda"
                ),                          //
            ).                              //
            All(ctx)
    if err != nil {
        return fmt.Errorf("failed getting cars: %w", err)
    }
    log.Println("cars returned:", cars)
    // Output: (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Ford, RegisteredAt=<Time>),)
    return nil
}