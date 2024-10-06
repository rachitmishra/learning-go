package login

import (
	"net/http"

	"rachitmishra.com/pastebin/cmd/web/shared"
)

func Handler(deps shared.ViewDeps) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		homeVm := NewLoginVM(nil)
		err := Login(homeVm).Render(ctx, w)

		if err != nil {
			deps.ServerError(w, err)
			return
		}
	}
}
