package nullify

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

type TestStruct struct {
	NullInt     NullInt64   `json:"nullInt" db:"null_int"`
	NullStr     NullString  `json:"nullString" db:"null_string"`
	NullBoolean NullBool    `json:"nullBool" db:"null_bool"`
	NullFloat   NullFloat64 `json:"nullFloat" db:"null_float"`
	NullTimer   NullTime    `json:"nullTime" db:"null_time"`
}

type ValidationStruct struct {
	NullInt     int64     `json:"nullInt" db:"null_int"`
	NullStr     string    `json:"nullString" db:"null_string"`
	NullBoolean bool      `json:"nullBool" db:"null_bool"`
	NullFloat   float64   `json:"nullFloat" db:"null_float"`
	NullTimer   time.Time `json:"nullTime" db:"null_time"`
}

func TestNullInt(t *testing.T) {
	var testStruct TestStruct
	var validationStruct ValidationStruct
	var JSONString []byte
	var err error

	Convey("Given a null int64", t, func() {
		Convey("Given some JSON", func() {
			JSONString = []byte(`{ "nullInt": null }`)

			err = json.Unmarshal(JSONString, &testStruct)
			So(testStruct.NullInt.Int64, ShouldEqual, 0)
			So(testStruct.NullInt.Valid, ShouldBeTrue)
			So(err, ShouldBeNil)

			JSONString, err = json.Marshal(testStruct)
			So(err, ShouldBeNil)
			err = json.Unmarshal(JSONString, &validationStruct)
			So(validationStruct.NullInt, ShouldEqual, 0)
		})
	})

	Convey("Given a valid int64", t, func() {
		Convey("Given some JSON", func() {
			JSONString = []byte(`{ "nullInt": 1337 }`)

			err = json.Unmarshal(JSONString, &testStruct)
			So(testStruct.NullInt.Int64, ShouldEqual, 1337)
			So(testStruct.NullInt.Valid, ShouldBeTrue)
			So(err, ShouldBeNil)

			JSONString, err = json.Marshal(testStruct)
			So(err, ShouldBeNil)
			err = json.Unmarshal(JSONString, &validationStruct)
			So(validationStruct.NullInt, ShouldEqual, 1337)
		})
	})
}

func TestNullString(t *testing.T) {
	var testStruct TestStruct
	var validationStruct ValidationStruct
	var JSONString []byte
	var err error

	Convey("Given a null string", t, func() {
		Convey("Given some JSON", func() {
			JSONString = []byte(`{ "nullString": null }`)

			err = json.Unmarshal(JSONString, &testStruct)
			So(testStruct.NullStr.String, ShouldEqual, "")
			So(testStruct.NullStr.Valid, ShouldBeTrue)
			So(err, ShouldBeNil)

			JSONString, err = json.Marshal(testStruct)
			So(err, ShouldBeNil)
		})
	})

	Convey("Given a valid string", t, func() {
		Convey("Given some JSON", func() {
			JSONString = []byte(`{ "nullString": "1337" }`)

			err = json.Unmarshal(JSONString, &testStruct)
			So(testStruct.NullStr.String, ShouldEqual, "1337")
			So(testStruct.NullStr.Valid, ShouldBeTrue)
			So(err, ShouldBeNil)

			JSONString, err = json.Marshal(testStruct)
			So(err, ShouldBeNil)

			err = json.Unmarshal(JSONString, &validationStruct)
			So(validationStruct.NullStr, ShouldEqual, "1337")
		})
	})
}

func TestNullBool(t *testing.T) {
	var testStruct TestStruct
	var validationStruct ValidationStruct
	var JSONString []byte
	var err error

	Convey("Given a null boolean", t, func() {
		Convey("Given some JSON", func() {
			JSONString = []byte(`{ "nullBool": null }`)

			err = json.Unmarshal(JSONString, &testStruct)
			So(testStruct.NullBoolean.Bool, ShouldBeFalse)
			So(testStruct.NullBoolean.Valid, ShouldBeTrue)
			So(err, ShouldBeNil)

			JSONString, err = json.Marshal(testStruct)
			So(err, ShouldBeNil)
		})
	})

	Convey("Given a valid boolean", t, func() {
		Convey("Given some JSON", func() {
			JSONString = []byte(`{ "nullBool": true }`)

			err = json.Unmarshal(JSONString, &testStruct)
			So(testStruct.NullBoolean.Bool, ShouldBeTrue)
			So(testStruct.NullBoolean.Valid, ShouldBeTrue)
			So(err, ShouldBeNil)

			JSONString, err = json.Marshal(testStruct)
			So(err, ShouldBeNil)

			err = json.Unmarshal(JSONString, &validationStruct)
			So(validationStruct.NullBoolean, ShouldBeTrue)

		})
	})
}

func TestNullFloat(t *testing.T) {
	var testStruct TestStruct
	var validationStruct ValidationStruct
	var JSONString []byte
	var err error

	Convey("Given a null Float", t, func() {
		Convey("Given some JSON", func() {
			JSONString = []byte(`{ "nullFloat": null }`)

			err = json.Unmarshal(JSONString, &testStruct)
			So(testStruct.NullFloat.Float64, ShouldEqual, 0)
			So(testStruct.NullFloat.Valid, ShouldBeTrue)
			So(err, ShouldBeNil)

			JSONString, err = json.Marshal(testStruct)
			So(err, ShouldBeNil)
		})
	})

	Convey("Given a valid Float", t, func() {
		Convey("Given some JSON", func() {
			JSONString = []byte(`{ "nullFloat": 1.337 }`)

			err = json.Unmarshal(JSONString, &testStruct)
			So(testStruct.NullFloat.Float64, ShouldEqual, 1.337)
			So(testStruct.NullFloat.Valid, ShouldBeTrue)
			So(err, ShouldBeNil)

			JSONString, err = json.Marshal(testStruct)
			So(err, ShouldBeNil)

			err = json.Unmarshal(JSONString, &validationStruct)
			So(validationStruct.NullFloat, ShouldEqual, 1.337)
		})
	})
}
