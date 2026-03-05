package db

func ContainsString(list []string, needle string) bool {
	for _, v := range list {
		if v == needle {
			return true
		}
	}
	return false
}

func FilterOutColumn(cols []string, args []any, colToDrop string) ([]string, []any) {
	if len(cols) != len(args) {
		return cols, args
	}

	nc := make([]string, 0, len(cols))
	na := make([]any, 0, len(args))

	for i := range cols {
		if cols[i] == colToDrop {
			continue
		}
		nc = append(nc, cols[i])
		na = append(na, args[i])
	}

	return nc, na
}
