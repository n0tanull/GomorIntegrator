package user

import (
	"context"
	"github.com/gofiber/fiber/v3/log"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"simpleauth/pkg/melog"
)

type repo struct {
	store *pgx.Conn
	mel   melog.Logger
}

func NewRepo(mel melog.Logger) Repo {
	store, err := Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return &repo{
		store: store,
		mel:   mel,
	}
}
func Connect(ctx context.Context) (*pgx.Conn, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	conn, err := pgx.Connect(ctx, string("host="+os.
		Getenv("DB_HOST")+" port="+os.
		Getenv("DB_PORT")+" user="+os.
		Getenv("DB_USER")+" password="+os.
		Getenv("DB_PASSWORD")+" dbname="+os.
		Getenv("DB_NAME")+" sslmode=disable"))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (r *repo) AddSocial(telegram string, ctx context.Context) (string, error) {
	t, err := r.store.Begin(ctx)
	if err != nil {
		return "", err
	}
	_, err = t.Exec(ctx, "INSERT INTO clients (telegram) VALUES ($1)", telegram)
	if err != nil {
		_ = t.Rollback(ctx)
		return "", err
	}
	return telegram, t.Commit(ctx)
}

func (r *repo) AddUser(email, password string, ctx context.Context) (string, error) {

	t, err := r.store.Begin(ctx)
	if err != nil {
		return "", err
	}
	_, err = t.Exec(ctx, "INSERT INTO users (email, password) VALUES ($1, $2)", email, password)
	if err != nil {
		_ = t.Rollback(ctx)
		return "", err
	}
	return email, t.Commit(ctx)
}

func (r *repo) GetUser(id string, ctx context.Context) (string, error) {
	t, err := r.store.Begin(ctx)
	if err != nil {
		return "", err
	}
	var email string
	err = t.QueryRow(ctx, "SELECT email FROM users WHERE id = $1", id).Scan(&email)
	if err != nil {
		_ = t.Rollback(ctx)
		return "", err
	}
	return email, t.Commit(ctx)
}

func (r *repo) GetUsers(ctx context.Context) ([]string, error) {
	t, err := r.store.Begin(ctx)
	if err != nil {
		return nil, err
	}
	var email string
	rows, err := t.Query(ctx, "SELECT email FROM users")
	if err != nil {
		_ = t.Rollback(ctx)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&email)
		if err != nil {
			_ = t.Rollback(ctx)
			return nil, err
		}
	}
	return []string{email}, t.Commit(ctx)
}
