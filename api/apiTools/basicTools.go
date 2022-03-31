package apiTools

import (
	"net/http"
	"net/url"
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
	if len(query) != 1 {
		return nil
	}
	switch true {
	case query.Has("name"),
		query.Has("id"):
		return query
	}
	return nil
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
