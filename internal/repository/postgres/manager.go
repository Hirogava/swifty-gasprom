package postgres

import (
	"database/sql"
	"fmt"
	"sync"
)

type Manager struct {
	Conn *sql.DB
	WG   *sync.WaitGroup
	MU   *sync.RWMutex
}

func NewManager(driverName string, sourceName string) *Manager {
	db, err := sql.Open(driverName, sourceName)
	if err != nil {
		panic(fmt.Sprintf("couldn't connect to the database: %v", err))
	}

	if err = db.Ping(); err != nil {
		panic(fmt.Sprintf("the database is not responding: %v", err))
	}

	return &Manager{
		Conn: db,
		WG:   &sync.WaitGroup{},
		MU:   &sync.RWMutex{},
	}
}

func (manager *Manager) Close() {
	if manager.Conn != nil {
		manager.Conn.Close()
		manager.Conn = nil
	}
}
