package others

import (
	"errors"
	"net/http"
	"strconv"
)

func ParsePageParam(r *http.Request) (int, error) {
	pageStr := r.URL.Query().Get("page")

	if pageStr == "" {
		return 1, nil
	}

	page, err := strconv.Atoi(pageStr)
	if err !=nil || page <= 0 {
		return 0, errors.New("Invalid page parameter")
	}

	return page, nil
}