package grouper

import "github.com/threatcode/osv-scanner/pkg/models"

type IDAliases struct {
	ID      string
	Aliases []string
}

func ConvertVulnerabilityToIDAliases(c []models.Vulnerability) []IDAliases {
	output := []IDAliases{}
	for _, v := range c {
		idAliases := IDAliases{
			ID:      v.ID,
			Aliases: v.Aliases,
		}
		output = append(output, idAliases)
	}

	return output
}
