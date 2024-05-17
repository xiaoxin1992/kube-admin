package pkg

func Page(page, size, total int) (offset int, limits int) {
	offset = (page - 1) * size
	limits = size * page
	if limits >= total {
		limits = total
	}
	if total == 0 {
		offset = 0
	}
	if offset > total {
		offset = total
	}
	return offset, limits
}
