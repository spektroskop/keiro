package keiro

import "net/http"

type Route struct {
	http.Handler
	Parts []string
}

func (route *Route) match(parts []string) map[string]string {
	if len(parts) != len(route.Parts) {
		return nil
	}

	params := make(map[string]string)
	for i, part := range route.Parts {
		if part != parts[i] && part[0] != ':' {
			return nil
		} else if part[0] == ':' {
			params[part[1:]] = parts[i]
		}
	}

	return params
}
