package main

import (
    "context"
    "log"
    "net/http"

    "entgo.io/ent/dialect/sql/schema"
    "github.com/superchris/ent-ogen-example/ent"
    "github.com/superchris/ent-ogen-example/ent/ogent"
    _ "github.com/lib/pq"
)

func main() {
    // Create ent client.
    client, err := ent.Open("postgres", "postgresql://root@localhost:26257/defaultdb?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    // Run the migrations.
    if err := client.Schema.Create(context.Background(), schema.WithAtlas(true)); err != nil {
        log.Fatal(err)
    }
    // Start listening.
    srv, err := ogent.NewServer(ogent.NewOgentHandler(client))
    if err != nil {
        log.Fatal(err)
    }
    if err := http.ListenAndServe(":8082", srv); err != nil {
        log.Fatal(err)
    }
}