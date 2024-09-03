package constants

const (
	ErrLoadEnvFile             = "failed to load .env file"
	ErrConnectMySQL            = "failed to connect to MySQL"
	ErrPingDatabase            = "failed to ping the db"
	ErrMigration               = "failed to migration"
	ErrInvalidPostID           = "Invalid post ID"
	ErrPostNotFound            = "Post not found"
	ErrFetchPosts              = "Failed to fetch posts"
	ErrGetDBConnection         = "failed to get db connection"
	ErrCreateMySQLDriver       = "failed to create MySQL driver"
	ErrMigrationFailed         = "migration failed"
	ErrCreateMigrationInstance = "failed to create migration instance"
)
