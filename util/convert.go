package util

import "fmt"

func ConvertToInt(s any) (int, error) {
	var count int
	switch v := s.(type) {
	case int:
		return v, nil
	case int64:
		return int(v), nil
	case float64:
		count = int(v)
		if fmt.Sprintf("%d", count) != fmt.Sprintf("%v", v) {
			return 0, fmt.Errorf("the value must not be a float64: %v", s)
		}
		return count, nil
	default:
		return 0, fmt.Errorf("the value must be an int: %v", s)
	}
}
