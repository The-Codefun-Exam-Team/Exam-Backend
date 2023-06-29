package utility

import (

)

func Pagination(page int, limit int) (begin int) {
	// page is the current index of the page (1-indexed)
	// limit is the maximum size of the page

	begin = (page - 1) * limit + 1

	return
}
