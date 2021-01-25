package database

import "testing"

func TestDefault(t *testing.T) {
	db := Default()

	if nil == db {
		t.Error("Database was nil.")
	}
}
