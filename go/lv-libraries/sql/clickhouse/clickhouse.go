package clickhouse

// PostgreSQL Constants.
const (
	ClickhouseDsnProtocolPostfix           = "://"
	ClickhouseDsnUsernamePasswordDelimiter = ":"
	ClickhouseDsnUsernameHostDelimiter     = "@"
	ClickhouseDsnHostPortDelimiter         = ":"
	ClickhouseDsnHostDatabaseDelimiter     = "/"
	ClickhouseDsnParametersPrefix          = "?"
)

// MakeClickhouseDsn Function returns a Connection String for PostgreSQL
// according to the Documentation at:
// "https://github.com/mailru/go-clickhouse".
// Format Reference:
// schema://user:password@host[:port]/database?param1=value1&...&paramN=valueN
func MakeClickhouseDsn(
	protocol string, // Obligatory Parameter.
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

	dsn = protocol + ClickhouseDsnProtocolPostfix
	if len(username) > 0 {
		if len(password) > 0 {
			dsn = dsn + username + ClickhouseDsnUsernamePasswordDelimiter +
				password + ClickhouseDsnUsernameHostDelimiter
		} else {
			dsn = dsn + username + ClickhouseDsnUsernameHostDelimiter
		}
	}

	dsn = dsn + host + ClickhouseDsnHostPortDelimiter + port

	if len(database) > 0 {
		dsn = dsn + ClickhouseDsnHostDatabaseDelimiter + database
	}

	if len(parameters) > 0 {
		dsn = dsn + ClickhouseDsnParametersPrefix + parameters
	}

	return dsn
}
