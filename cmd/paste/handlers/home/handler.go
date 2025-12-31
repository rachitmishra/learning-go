package home

import (
	"net/http"

	"rachitmishra.com/pastebin/cmd/web/shared"
)

func Handler(deps shared.ViewDeps) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// if r.URL.Path != "/" {
		// 	components.NotFound().Render(ctx, w)
		// 	return
		// }

		homeVm := NewHomeVM(nil)
		err := Home(homeVm).Render(ctx, w)

		if err != nil {
			deps.ServerError(w, err)
			return
		}
	}
}
