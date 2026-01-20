package main

import (
	"authentication/data"
	"os"
	"testing"
)

var testApp Config

// TestMain set testing environment up
func TestMain(m *testing.M) {
	repo := data.NewPostgresTestRepository(nil)
	testApp.Repo = repo
	os.Exit(m.Run())
}
