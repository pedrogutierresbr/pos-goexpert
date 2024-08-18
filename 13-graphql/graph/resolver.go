package graph

import "github.com/pedrogutierresbr/pos-goexpert/12-graphql/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryDB *database.Category
	CouseDb    *database.Course
}
