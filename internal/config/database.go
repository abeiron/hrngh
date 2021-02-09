package config

// Database represents the instance of the SQL database.
//
// `username`: Represents the login name for the database connection.
// `passphrase`: Represents the passphrase for the database connection.
type Database struct {
	username string
	passphrase string
}

func DbConfig() Database {
	return Database {
		"admin",
		"PASSPHRASE",
	}
}

// Connect initiates a connection to the database.
func (db *Database) Connect() error {
	return nil
}
