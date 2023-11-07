package address

import (
	"slices"
	"strings"
)

func IsValidBrazilianFormat(address string) bool {
	addressFormats := []string{"rua", "avenida", "rodovia", "praça"}

	addressToLower := strings.ToLower(address)

	splittedAddress := strings.Split(addressToLower, " ")

	return slices.Contains(addressFormats, splittedAddress[0])
}
