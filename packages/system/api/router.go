/* Copyright © 2020. Financelime, https://financelime.com. All rights reserved.
   Author: DmAlix. Contacts: <dmalix@financelime.com>, <dmalix@yahoo.com>
   License: GNU General Public License v3.0, https://www.gnu.org/licenses/gpl-3.0.html */

package api

import (
	"github.com/dmalix/financelime-rest-api/packages/system"
	"github.com/dmalix/financelime-rest-api/utils/router"
	"net/http"
)

func Router(mux *http.ServeMux, service system.Service, middleware ...func(http.Handler) http.Handler) {

	handler := NewHandler(service)

	mux.Handle("/system/dist",
		router.Group(
			router.EndPoint(router.Point{Method: http.MethodGet, Handler: handler.Dist()}),
			middleware,
		))
}
