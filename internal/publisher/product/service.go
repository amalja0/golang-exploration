package product

import (
	postgresadapter "analytic-reporting/internal/publisher/product/adapters/postgres"
	"database/sql"
)

func Init(db *sql.DB) postgresadapter.Repository {
	return postgresadapter.NewRepo(db)
}
