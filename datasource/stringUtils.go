package datasource

import "github.com/volatiletech/null/v8"

func StringToNullableString(value string) null.String {
	if len(value) > 0 {
		return null.StringFrom(value)
	} else {
		return null.NewString("", false)
	}
}

func NullableStringToString(value null.String) string {
	if value.Valid {
		return value.String
	} else {
		return ""
	}
}
