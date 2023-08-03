package graph

// import database "command-line-arguments/Users/leandroluna/Documents/Github/FullCycle/SystemCommunications/GraphQL/internal/database/category.go"
import database "github.com/SystemCommunications/GraphQL/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoryDB *database.Category
}
