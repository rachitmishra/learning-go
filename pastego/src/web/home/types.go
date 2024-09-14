package home

import (
	"net/url"

	"github.com/go-playground/form/v4"
	"github.com/google/uuid"
	"rachitmishra.com/pastebin/internal/models"
	"rachitmishra.com/pastebin/internal/validator"
)

type PasteForm struct {
	title   string
	content string
	expires int
	validator.Validator
}

func (p *PasteForm) Title() string {
	return p.title
}
func (p *PasteForm) Content() string {
	return p.content
}
func (p *PasteForm) Expires() int {
	return p.expires
}

func (p *PasteForm) IsValid() bool {
	p.CheckFieldError(validator.NotBlank(p.title), "title", "Title can't be blank")
	p.CheckFieldError(validator.NotBlank(p.content), "content", "Content can't be blank")
	p.CheckFieldError(validator.PermittedInt(p.expires,
		1, 7, 24, 365, 3650), "expires", "Expires value is invalid")
	return p.Valid()
}

type HomeVM struct {
	form *PasteForm
}

func NewHomeVM(pasteForm *PasteForm) HomeVM {
	return HomeVM{
		form: pasteForm,
	}
}

func (h *HomeVM) CheckField(key string) bool {
	if h.form == nil {
		return false
	}
	return h.form.CheckField(key)
}

func (h *HomeVM) GetField(key string) string {
	if h.form == nil {
		return ""
	}
	return h.form.GetField(key)
}

func NewPasteForm(d form.Decoder, values url.Values) (*PasteForm, error) {
	var form PasteForm
	err := d.Decode(&form, values)
	if err != nil {
		return nil, nil
	}
	return &form, nil
}

func (p *PasteForm) ToModel() (*models.Paste, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return models.NewPaste(uid.String(), p.title, p.content, p.expires), nil
}
