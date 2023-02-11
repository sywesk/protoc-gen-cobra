package flag

import (
	"encoding/json"
	"strconv"
)

func ParseBoolE(val string) (bool, error) { return strconv.ParseBool(val) }

func ParseInt32E(val string) (int32, error) {
	if i, err := strconv.ParseInt(val, 10, 32); err != nil {
		return 0, err
	} else {
		return int32(i), nil
	}
}

func ParseInt64E(val string) (int64, error) { return strconv.ParseInt(val, 10, 64) }

func ParseUint32E(val string) (uint32, error) {
	if i, err := strconv.ParseUint(val, 10, 32); err != nil {
		return 0, err
	} else {
		return uint32(i), nil
	}
}

func ParseUint64E(val string) (uint64, error) { return strconv.ParseUint(val, 10, 64) }

func ParseFloat32E(val string) (float32, error) {
	if i, err := strconv.ParseFloat(val, 32); err != nil {
		return 0, err
	} else {
		return float32(i), nil
	}
}

func ParseFloat64E(val string) (float64, error) { return strconv.ParseFloat(val, 64) }

func ParseStringE(val string) (string, error) { return val, nil }

func ParseJsonE[T any](val string) (*T, error) {
	v := new(T)
	if err := json.Unmarshal([]byte(val), v); err != nil {
		return nil, err
	} else {
		return v, nil
	}
}
