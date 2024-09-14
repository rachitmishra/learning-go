package create

import (
	"fmt"
	"net/http"

	"rachitmishra.com/pastebin/src/web/home"
	"rachitmishra.com/pastebin/src/web/shared"
)

func Handler(a shared.ServiceDeps) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		createPaste(a, w, r)
	}
}

func createPaste(
	deps shared.ServiceDeps,
	w http.ResponseWriter,
	r *http.Request,
) {
	err := shared.ValidateForm(deps, r)
	if err != nil {
		deps.ClientError(w, http.StatusBadRequest)
	}
	form, err := home.NewPasteForm(*deps.FormDecoder(), r.PostForm)
	if err != nil {
		deps.ClientError(w, http.StatusBadRequest)
		return
	}
	if !form.Valid() {
		homeVm := home.NewHomeVM(form)
		home.Content(homeVm).Render(r.Context(), w)
		return
	}

	model, err := form.ToModel()
	if err != nil {
		deps.ServerError(w, err)
		return
	}
	_, err = deps.DB().Insert(model.Uid, model)
	if err != nil {
		deps.ServerError(w, err)
		return
	}
	redirect := fmt.Sprintf("/paste/%s", model.Uid)
	deps.LogI(redirect)
	w.Header().Add("HX-Replace-Url", redirect)
	http.Redirect(w, r, redirect, http.StatusSeeOther)
}
