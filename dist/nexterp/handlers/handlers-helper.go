package handlers

import (
	"context"
	"net/http"

	"github.com/otaja/celeritas"
)

func (h *Handlers) Render(w http.ResponseWriter, r *http.Request, tmpl string, variables, data interface{}) error {
	return h.App.Render.Page(w, r, tmpl, variables, data)
}

func (h *Handlers) SessionsPut(ctx context.Context, key string, val interface{}) {
	h.App.Session.Put(ctx, key, val)
}

func (h *Handlers) SessionHas(ctx context.Context, key string) bool {
	return h.App.Session.Exists(ctx, key)
}

func (h *Handlers) SessionGet(ctx context.Context, key string) interface{} {
	return h.App.Session.Get(ctx, key)
}

func (h *Handlers) SessionRemove(ctx context.Context, key string) {
	h.App.Session.Remove(ctx, key)
}

func (h *Handlers) SessionRenew(ctx context.Context) error {
	return h.App.Session.RenewToken(ctx)
}

func (h *Handlers) SessionDestroy(ctx context.Context) error {
	return h.App.Session.Destroy(ctx)
}

func (h *Handlers) RandomString(n int) string {
	return h.App.RandomString(n)
}

func (h *Handlers) Encrypt(text string) (string, error) {
	enc := celeritas.Encryption{Key: []byte(h.App.EncryptionKey)}
	encrypted, err := enc.Encrypt(text)
	if err != nil {
		return "", err
	}

	return encrypted, nil
}

func (h *Handlers) Decrypt(crypot string) (string, error) {
	enc := celeritas.Encryption{Key: []byte(h.App.EncryptionKey)}
	decrypted, err := enc.Decrypt(crypot)
	if err != nil {
		return "", err
	}

	return decrypted, nil
}
