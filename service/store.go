package service

import (
	"database/sql"
	"github.com/spf13/viper"
	"log"
	"math/rand"
	_ "modernc.org/sqlite"
	"os"
	"sync"
)

type Store struct {
	mu sync.RWMutex
	db *sql.DB
}

// Clean architecture
func NewStore() *Store {
	return &Store{db: NewDBInstance()}
}

// Clean architecture 2
func NewDBInstance() *sql.DB {
	sqlDB, err := sql.Open("sqlite", viper.GetString("db"))
	if err != nil {
		log.Fatal(err)
	}

	dbinit, err := os.ReadFile("../configs/dbinit")
	if err != nil {
		log.Fatal(err)
	}
	if _, err = sqlDB.Exec(string(dbinit)); err != nil {
		log.Fatal(err)
	}
	log.Print("Successfully connxted to DB")
	return sqlDB

}

// Func to save the url into the DB
func (s *Store) Save(url string) (string, error) {
	str := Gen()
	var u string
	s.mu.Lock()
	defer s.mu.Unlock()
	ok := s.db.QueryRow(`SELECT url FROM map WHERE key = ?`, str).Scan(&u)
	if ok == sql.ErrNoRows {

		_, err := s.db.Exec("INSERT INTO map(key, url) VALUES(?, ?)", str, url)
		if err != nil {
			return "", err
		}

		return str, nil
	} else {
		log.Print("Resolving collision")
		s.Save(url)
	}
	return "", nil
}

func (s *Store) Load(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var url string
	err := s.db.QueryRow("SELECT url FROM map WHERE key = ?", key).Scan(&url)
	if err != nil {
		if err == sql.ErrNoRows {
			// no such key
			return "", false
		}
		log.Fatal(err)
	}

	return url, true
}

func Gen() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, viper.GetInt("KeyLen"))
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
