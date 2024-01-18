package test

import (
	"testing"
)

func TestFindAllProperties(t *testing.T) {

	// Find All Properties
	properties, err := propertyUseCase.FindAll(nil)

	if err != nil {
		t.Error(err)
	}

	t.Log(properties)
}

func TestFindDetailProperty(t *testing.T) {

	// Find Detail Property
	property, err := propertyUseCase.FindDetail("234")

	if err != nil {
		t.Error(err)
	}

	t.Log(property)
}

func TestFindPropertyByPrompt(t *testing.T) {

	// Find Property By Prompt
	properties, err := propertyUseCase.GetPropertyByPrompt("Carikan saya rumah di daerah pancoranmas")

	if err != nil {
		t.Error(err)
	}

	t.Log(properties)
}
