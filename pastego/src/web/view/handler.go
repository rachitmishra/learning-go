package view

import (
	"fmt"
	"net/http"
	"unicode/utf8"

	"github.com/julienschmidt/httprouter"
	"rachitmishra.com/pastebin/src/web/shared"
	components "rachitmishra.com/pastebin/src/web/shared/components"
)

func Handler(deps shared.ServiceDeps) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := httprouter.ParamsFromContext(r.Context())
		uid := params.ByName("id")
		val, err := deps.DB().Get(uid)
		if err != nil || utf8.RuneCountInString(uid) == 0 {
			components.NotFound().Render(r.Context(), w)
			return
		}
		if val != nil {
			fmt.Fprintf(w, "View %s", val)
		}
	}
}
