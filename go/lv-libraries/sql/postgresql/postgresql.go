package postgresql

// PostgreSQL Constants.
const (
	PostgresqlDsnPrefix                    = "postgresql://"
	PostgresqlDsnUsernamePasswordDelimiter = ":"
	PostgresqlDsnUsernameHostDelimiter     = "@"
	PostgresqlDsnHostPortDelimiter         = ":"
	PostgresqlDsnHostDatabaseDelimiter     = "/"
	PostgresqlDsnParametersPrefix          = "?"
)

// MakePostgresqlDsn Function returns a Connection String for PostgreSQL
// according to the Documentation at:
// "https://www.postgresql.org/docs/10/libpq-connect.html".
// Format Reference:
// postgresql://[user[:password]@][netloc][:port][,...][/dbname][?param1=value1&...]
func MakePostgresqlDsn(
	host string, // Obligatory Parameter.
	port string, // Obligatory Parameter.
	database string, // Optional Parameter.
	username string, // Optional Parameter.
	password string, // Optional Parameter.

	// Key-Value List without the '?' Prefix.
	// Optional Parameter.
	parameters string,
) string {

	var dsn string

	dsn = PostgresqlDsnPrefix
	if len(username) > 0 {
		if len(password) > 0 {
			dsn = dsn + username + PostgresqlDsnUsernamePasswordDelimiter +
				password + PostgresqlDsnUsernameHostDelimiter
		} else {
			dsn = dsn + username + PostgresqlDsnUsernameHostDelimiter
		}
	}

	dsn = dsn + host + PostgresqlDsnHostPortDelimiter + port

	if len(database) > 0 {
		dsn = dsn + PostgresqlDsnHostDatabaseDelimiter + database
	}

	if len(parameters) > 0 {
		dsn = dsn + PostgresqlDsnParametersPrefix + parameters
	}

	return dsn
}
