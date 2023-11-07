package address_test

import (
	"testing"
	. "unity-tests/address"
)

type TestScenarios struct {
	address        string
	expectedReturn bool
}

func TestShouldBeValidWhenAddressTypeIsRua(t *testing.T) {
	t.Parallel()

	completeAddress := "Rua Grécia"

	if !IsValidBrazilianFormat(completeAddress) {
		t.Errorf("'%s' is not a valid address", completeAddress)
	}
}

func TestShouldBeValidWhenAddressTypeIsAvenida(t *testing.T) {
	t.Parallel()

	completeAddress := "Avenida dos Andradas"

	if !IsValidBrazilianFormat(completeAddress) {
		t.Errorf("'%s' is not a valid address", completeAddress)
	}
}

func TestShouldBeValidWhenAddressTypeIsRodovia(t *testing.T) {
	t.Parallel()

	completeAddress := "Rodovia Fernão Dias"

	if !IsValidBrazilianFormat(completeAddress) {
		t.Errorf("'%s' is not a valid address", completeAddress)
	}
}

func TestShouldBeValidWhenAddressTypeIsPraca(t *testing.T) {
	t.Parallel()

	completeAddress := "Praça Paulo Pinheiro Chagas"

	if !IsValidBrazilianFormat(completeAddress) {
		t.Errorf("'%s' is not a valid address", completeAddress)
	}
}

func TestShouldNotBeValid(t *testing.T) {
	t.Parallel()

	completeAddress := "random address"

	if IsValidBrazilianFormat(completeAddress) {
		t.Errorf("'%s' is not a valid address", completeAddress)
	}
}

func TestShouldValidateMultipleScenarios(t *testing.T) {
	t.Parallel()

	testScenarios := []TestScenarios{
		{"Rua Bologna", true},
		{"Avenida Tulio Fraga", true},
		{"Rodovia Alto do Ipiranga", true},
		{"Praça Sete", true},
		{"random", false},
	}

	for _, testScenario := range testScenarios {
		result := IsValidBrazilianFormat(testScenario.address)

		if result != testScenario.expectedReturn {
			t.Errorf("'%s' is not a valid address", testScenario.address)
		}
	}
}
