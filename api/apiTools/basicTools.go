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

//https://stackoverflow.com/questions/59570978/is-there-a-way-to-check-for-invalid-query-parameters-in-an-http-request
func GetQuery(r *http.Request) (url.Values, bool) {
	query := r.URL.Query()

	allowedQuery := map[string]bool{constants.P_idURL: true, constants.P_nameURL: true, constants.P_scaffolding: true}

	for k := range query {
		if _, ok := allowedQuery[k]; !ok {
			return nil, false
		}
	}

	if query.Has(constants.P_scaffolding) {
		if !(query.Get(constants.P_scaffolding) == "true" || query.Get(constants.P_scaffolding) == "false") {
			return nil, false
		}
	}

	return query, true
}

func GetQueryScaffolding(r *http.Request) url.Values {
	query := r.URL.Query()
	if len(query) == 0 {
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
