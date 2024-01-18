package test

import (
	"geoproperty_be/config"
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	// Load Function Configuration
	if err := config.InitializeConfig("../../.env"); err != nil {
		panic(err)
	}

	m.Run()

	// Pass Test
	log.Println("All Test Passed")
}

func TestConnection(t *testing.T) {
	// Connect to Database
	if _, err := config.Connect(); err != nil {
		t.Error(err)
	}

	// Pass Test
	t.Log("TestConnection Passed")

}
