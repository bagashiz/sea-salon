package server

import "net/http"

// parseForm parses the form values from an HTTP request then returns them as a map.
func parseForm(r *http.Request) (map[string]string, error) {
	if err := r.ParseForm(); err != nil {
		return nil, err
	}

	formValues := make(map[string]string, 0)
	for key, value := range r.PostForm {
		formValues[key] = value[0]
	}

	return formValues, nil
}
