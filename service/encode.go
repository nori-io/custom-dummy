// Copyright (C) 2018 The Nori Authors info@nori.io
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU Lesser General Public
// License as published by the Free Software Foundation; either
// version 3 of the License, or (at your option) any later version.
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program; if not, see <http://www.gnu.org/licenses/>.
package service

import (
	"context"
	"encoding/json"
	rest "github.com/cheebo/gorest"
	// norihttp "github.com/nori-io/nori-common/transport/http"
	norihttp "github.com/nori-io/common/v4/pkg/errors"
	"net/http"
)

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(norihttp.Errorer); ok && e.Error() != nil {
		EncodeError(ctx, e.Error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// encode errors from business-logic
func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var msg string

	switch err.(type) {
	case rest.ErrResp:
		er := err.(rest.ErrResp)
		w.WriteHeader(er.Meta.ErrCode)
		json.NewEncoder(w).Encode(er)
	case rest.ErrFieldResp:
		er := err.(rest.ErrFieldResp)
		w.WriteHeader(er.Meta.ErrCode)
		json.NewEncoder(w).Encode(er)
	default:
		code := http.StatusInternalServerError

		w.WriteHeader(code)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": msg,
		})
	}
}
