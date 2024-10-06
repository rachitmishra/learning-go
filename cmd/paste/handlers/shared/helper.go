package shared

import (
	"net/http"
)

func ValidateForm(deps ServiceDeps, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	deps.LogI(r.Form.Encode())
	return nil
}
