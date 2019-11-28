package nullify

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

// NullInt64 is an alias for sql.NullInt64 data type
type NullInt64 struct {
	sql.NullInt64
}

// NullString is an alias for sql.NullString data type
type NullString struct {
	sql.NullString
}

// NullBool is an alias for sql.NullBool data type
type NullBool struct {
	sql.NullBool
}

// NullFloat64 is an alias for sql.NullFloat64 data type
type NullFloat64 struct {
	sql.NullFloat64
}

// NullTime is an alias for mysql.NullTime data type
type NullTime struct {
	mysql.NullTime
}
