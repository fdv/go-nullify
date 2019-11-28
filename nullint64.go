package nullify

import (
	"database/sql/driver"
)

// Scan implements the Scanner interface.
func (ni *NullInt64) Scan(value interface{}) error {
	if value == nil {
		ni.Int64, ni.Valid = 0, false
		return nil
	}
	ni.Valid = true
	return nil
}

// Value implements the driver Valuer interface.
func (ni NullInt64) Value() (driver.Value, error) {
	if !ni.Valid {
		return nil, nil
	}
	return ni.Int64, nil
}
