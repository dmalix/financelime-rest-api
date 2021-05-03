/* Copyright © 2020. Financelime, https://financelime.com. All rights reserved.
   Author: DmAlix. Contacts: <dmalix@financelime.com>, <dmalix@yahoo.com>
   License: GNU General Public License v3.0, https://www.gnu.org/licenses/gpl-3.0.html */

package system

import (
	authorization2 "github.com/dmalix/financelime-authorization/authorization"
	"github.com/dmalix/financelime-authorization/utils/router"
	"net/http"
)

func Router(mux *http.ServeMux, handler API, middleware authorization2.APIMiddleware) {

	mux.Handle("/v1/version",
		router.Group(
			router.EndPoint(router.Point{Method: http.MethodGet, Handler: handler.version()}),
		))
}