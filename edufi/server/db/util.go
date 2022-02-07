package db

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// createID creates a schema ID from a table id
func (P *Postgres) createID(table string, id int64) (*string, error) {
	if id < 0 {
		return nil, fmt.Errorf("invalid table id")
	}
	s := fmt.Sprintf("%s-%d", table, id)

	return &s, nil
}

// getID retrieves a table id from a schema ID
func (P *Postgres) getID(id string, table string) (*int64, error) {
	s := strings.Split(id, "-")
	if len(s) != 2 {
		log.Printf("getID: id not found")
		return nil, fmt.Errorf("id not found")
	}
	itemID := s[1]
	i, err := strconv.ParseInt(itemID, 10, 64)
	if err != nil {
		log.Printf("getID: %v", err)
		return nil, err
	}
	return &i, nil
}
