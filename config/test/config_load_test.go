package test

import (
	"geoproperty_be/config"
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	// Load Function Configuration
	if err := config.InitializeConfig("../../.env"); err != nil {
		t.Error(err)
	}

	// Check Each Environment Variable is Exist
	if _, exist := os.LookupEnv("DB_HOST"); !exist {
		t.Error("Environment Variable DB_HOST is not Exist")
	}

	if _, exist := os.LookupEnv("DB_PORT"); !exist {
		t.Error("Environment Variable DB_PORT is not Exist")
	}

	// Pass Test
	t.Log("TestGetEnv Passed")
}
