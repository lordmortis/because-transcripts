package datasource

func BoolToInt64(value bool) int64 {
	if value {
		return 1
	} else {
		return 0
	}
}

func Int64ToBool(value int64) bool {
	return value == 1
}
