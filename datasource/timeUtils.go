package datasource

import (
	"github.com/volatiletech/null/v8"
	"time"
)

func NullableInt64ToTimeMillis(int64 null.Int64) *time.Time {
	if !int64.Valid {
		return nil
	}

	aTime := time.UnixMilli(int64.Int64)
	return &aTime
}

func NullableIntToTimeMillis(intValue null.Int) *time.Time {
	if !intValue.Valid {
		return nil
	}

	aTime := time.UnixMilli(int64(intValue.Int))
	return &aTime
}

func NullableInt64ToTimeMillisEquals(a null.Int64, b *time.Time) bool {
	aTime := NullableInt64ToTimeMillis(a)
	if aTime == nil && b == nil {
		return true
	}
	if aTime == nil || b == nil {
		return false
	}
	return aTime.Equal(*b)
}

func NullableIntToTimeMillisEquals(a null.Int, b *time.Time) bool {
	aTime := NullableIntToTimeMillis(a)
	if aTime == nil && b == nil {
		return true
	}
	if aTime == nil || b == nil {
		return false
	}
	return aTime.Equal(*b)
}

func TimeMillisToNullableInt64(time *time.Time) null.Int64 {
	if time == nil {
		return null.NewInt64(0, false)
	}
	return null.Int64From(time.UnixMilli())
}

func TimeMillisToNullableInt(time *time.Time) null.Int {
	if time == nil {
		return null.NewInt(0, false)
	}
	return null.IntFrom(int(time.UnixMilli()))
}
