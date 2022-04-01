package apiTools

import (
	"net/http"
	"net/url"
	"stillasTracker/api/constants"
	"strings"
)

func CreatePath(segments []string) string {
	var finalPath string
	for _, s := range segments {
		finalPath += s + "/"
	}
	finalPath = strings.TrimRight(finalPath, "/")
	return finalPath
}

func GetQuery(r *http.Request) url.Values {
	query := r.URL.Query()

	switch true {
	case query.Has(constants.P_idURL) && len(query) == 1:
		return query
	case query.Has(constants.P_idURL) && query.Has(constants.P_scaffolding) && len(query) == 2:
		return query
	case query.Has(constants.P_scaffolding) && len(query) == 1:
		return query
	default:
		return nil
	}
}

func GetQueryScaffolding(r *http.Request) url.Values {
	query := r.URL.Query()
	if len(query) != 2 {
		return nil
	} else if len(query) == 1 {
		switch true {
		case query.Has("type"):
			return query
		}
	} else {
		switch true {
		case query.Has("type"),
			query.Has("id"):
			return query
		}
	}
	return nil
}
