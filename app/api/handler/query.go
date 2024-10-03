package handler

import (
	"blog-api/internal/repository"
	"log"
	"net/url"
	"strconv"
)

func parsePaginationQuery(query url.Values) (result repository.Pagination) {
	var err error
	var limit = query.Get("limit")
	var page = query.Get("page")

	if result.Limit, err = strconv.Atoi(limit); limit != "" && err != nil {
		log.Print("error parsing query limit", err.Error())

	}

	if result.Page, err = strconv.Atoi(page); page != "" && err != nil {
		log.Print("error parsing query page", err.Error())
	}

	return
}
