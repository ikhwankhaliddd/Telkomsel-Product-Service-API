package pagination

func CalculateOffset(page, limit int) int {
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 1
	}

	return (page - 1) * limit
}
