package util

import (
	"fmt"
	"strings"
)

func SortValidation(sortParam string, availableSorts []string) string {
	var (
		sortBy string
		order  string
	)

	if sortParam != "" {

		if strings.HasPrefix(sortParam, "-") {
			sortBy = sortParam[1:]
			order = "DESC"
		} else {
			sortBy = sortParam
			order = "ASC"
		}

		for _, v := range availableSorts {
			if v == sortBy {
				// valid order
				return fmt.Sprintf("%s %s", sortBy, order)
			}
		}
	}

	// default order
	return "created_at DESC"
}
