package parsers

import "strconv"

func GetFloat(cfg map[string]any, key string, def float64) (float64, bool) {
	if cfg == nil {
		return def, false
	}
	v, ok := cfg[key]
	if !ok {
		return def, false
	}

	switch t := v.(type) {
	case float64:
		return t, true
	case float32:
		return float64(t), true
	case int:
		return float64(t), true
	case int64:
		return float64(t), true
	case uint64:
		return float64(t), true
	case string:
		f, err := strconv.ParseFloat(t, 64)
		if err == nil {
			return f, true
		}
	}

	return def, false
}

func GetInt(cfg map[string]any, key string, def int) (int, bool) {
	if cfg == nil {
		return def, false
	}
	v, ok := cfg[key]
	if !ok {
		return def, false
	}

	switch t := v.(type) {
	case int:
		return t, true
	case int64:
		return int(t), true
	case uint64:
		return int(t), true
	case float64:
		return int(t), true
	case string:
		i, err := strconv.Atoi(t)
		if err == nil {
			return i, true
		}
	}

	return def, false
}
