package main

import (
	"context"
	"io"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/superchris/ent-ogen-example/ent/enttest"
	"github.com/superchris/ent-ogen-example/ent/ogent"
)

func TestAdd(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	p := client.Person.Create().
		SetFirstName("Chris").
		SetLastName("Schmich").
		SetSalary(12300).
		SetEmail("bob@foo.bar").
		SetBirthDate(time.Now())
	_, error := p.Save(context.Background())
	if error != nil {
		t.Fatal(error)
	}

	srv, err := ogent.NewServer(ogent.NewOgentHandler(client))
	require.NoError(t, err)

	ts := httptest.NewServer(srv)
	defer ts.Close()
	
	resp, err := ts.Client().Get(ts.URL + "/people")
	require.NoError(t, err)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	assert.Contains(t, string(body), "Chris")
}
