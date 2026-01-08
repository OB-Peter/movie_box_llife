package data

import (
	"database/sql"
	"errors"
)

// Define custom error types for database operations
var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// Models struct wraps all the different model types
type Models struct {
	Movies      MovieModel
	Permissions PermissionModel
	Users       UserModel
	Tokens      TokenModel
}

// NewModels returns a Models struct containing the initialized model instances
func NewModels(db *sql.DB) Models {
	return Models{
		Movies:      MovieModel{DB: db},
		Users:       UserModel{DB: db},
		Tokens:      TokenModel{DB: db},
		Permissions: PermissionModel{DB: db},
	}
}
