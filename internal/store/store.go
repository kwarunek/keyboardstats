package store

import (
	"database/sql"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	dbPath       string
	dumpInterval int
	db           *sql.DB
	localStore   map[string]int
	lastDump     int64
	mu           sync.RWMutex
}

func NewStore(dbPath string, dumpInterval int) (*Store, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	store := &Store{
		dbPath:       dbPath,
		dumpInterval: dumpInterval,
		localStore:   make(map[string]int),
		lastDump:     0,
		db:           db,
	}
	if err := store.Init(); err != nil {
		return nil, err
	}
	return store, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) Init() error {

	sqlStmt := `
	  CREATE TABLE IF NOT EXISTS keyboardstats (
		id integer not null primary key,
		year integer not null,
		week integer not null,
		month integer not null,
		key text not null,
		count integer not null
	  );
	`
	_, err := s.db.Exec(sqlStmt)
	if err != nil {
		return err
	}
	idxYearWeekMonthStmt := `
		CREATE INDEX IF NOT EXISTS year_week_month_idx ON keyboardstats(year, week, month);
	`
	_, err = s.db.Exec(idxYearWeekMonthStmt)
	if err != nil {
		return err
	}
	idxKeyStmt := `
		CREATE INDEX IF NOT EXISTS idx_key ON keyboardstats(key);
	`
	_, err = s.db.Exec(idxKeyStmt)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) Insert(key string, count int, year int, week int, month int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into keyboardstats(year, week, month, key, count) values(?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(year, week, month, key, count)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (s *Store) Query(query string) error {
	return nil
}

func (s *Store) Record(key string) {
	s.mu.Lock()
	s.localStore[key]++
	s.mu.Unlock()
	if time.Now().Unix() > int64(s.lastDump)+int64(s.dumpInterval) {
		s.Dump()
	}
}

func (s *Store) Dump() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	year, week := time.Now().ISOWeek()
	month := int(time.Now().Month())
	for key, count := range s.localStore {
		if err := s.Insert(key, count, year, week, month); err != nil {
			return err
		}
	}
	s.lastDump = time.Now().Unix()
	s.localStore = make(map[string]int)
	return nil
}
