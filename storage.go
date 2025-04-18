package main

import (
	"database/sql"

	_"github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=gobank dbname=bank password=gobank sslmode=disable" 
	db, err := sql.Open("postgres", connStr)
	
	if err != nil {
		return nil, err 
	}
	
	if err := db.Ping(); err != nil {
		return nil, err 
	}

	return &PostgresStore{
		db: db,
	}, nil 
}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `create table if not exists accounts(
		account_id serial primary key, 
		first_name varchar(25),
		last_name  varchar(25),
		account_number serial,
		balance bigint,
		create_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}
func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}






