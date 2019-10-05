// Copyright 2017 Frédéric Guillot. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package ui // import "miniflux.app/ui"

import (
	"net/http"
	"time"

	"miniflux.app/http/request"
	"miniflux.app/http/response"
	"miniflux.app/http/response/html"
	"miniflux.app/http/route"
	"miniflux.app/storage"
	"miniflux.app/ui/session"
	"miniflux.app/ui/view"
)

func (h *handler) shareGenerate(w http.ResponseWriter, r *http.Request) {
	entryID := request.RouteInt64Param(r, "entryID")
	shareCode, err := h.store.GetEntryShareCode(request.UserID(r), entryID)
	if err != nil {
		html.ServerError(w, r, err)
		return
	}

	html.Redirect(w, r, route.Path(h.router, "share", "shareCode", shareCode))
}

func (h *handler) sharePage(w http.ResponseWriter, r *http.Request) {
	shareCode := request.RouteStringParam(r, "shareCode")
	if shareCode == "" {
		html.NotFound(w, r)
		return
	}

	etag := shareCode
	response.New(w, r).WithCaching(etag, 72*time.Hour, func(b *response.Builder) {
		builder := storage.NewAnonymousQueryBuilder(h.store)
		builder.WithShareCode(shareCode)

		entry, err := builder.GetEntry()
		if err != nil || entry == nil {
			html.NotFound(w, r)
			return
		}

		sess := session.New(h.store, request.SessionID(r))
		view := view.New(h.tpl, r, sess)
		view.Set("entry", entry)

		b.WithHeader("Content-Type", "text/html; charset=utf-8")
		b.WithBody(view.Render("entry"))
		b.Write()
	})
}
