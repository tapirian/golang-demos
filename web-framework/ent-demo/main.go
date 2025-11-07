package main

import (
	"context"
	"ent_demo/ent"
	"ent_demo/ent/user"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	_, err = CreateUser(ctx, client)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}

	_, err = QueryUser(ctx, client, "James")
	if err != nil {
		log.Fatalf("failed querying user: %v", err)
	}

	_, err = UpdateUser(ctx, client)
	if err != nil {
		log.Fatalf("failed querying user: %v", err)
	}

	_, err = QueryUser(ctx, client, "James Bond")
	if err != nil {
		log.Fatalf("failed querying user: %v", err)
	}

	_, err = DeleteUser(ctx, client)
	if err != nil {
		log.Fatalf("failed querying user: %v", err)
	}
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetName("James").
		SetAge(19).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	log.Println("User is Created:", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client, userName string) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name(userName)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	log.Println("User returned:", u)
	return u, nil
}

func UpdateUser(ctx context.Context, client *ent.Client) (int, error) {
	u, err := client.User.
		Update().
		Where(user.Name("James")).
		SetAge(20).
		SetName("James Bond").
		Save(ctx)
	if err != nil {
		return 0, err
	}

	log.Println("User is Updated, affected rows count:", u)
	return u, nil
}

func DeleteUser(ctx context.Context, client *ent.Client) (int, error) {
	u, err := client.User.
		Delete().
		Where(user.Name("James")).
		Exec(ctx)
	if err != nil {
		return 0, err
	}

	log.Println("User is Deleted, affected rows count:", u)
	return u, nil
}
