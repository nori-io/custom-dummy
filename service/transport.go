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
	"github.com/nori-io/interfaces/nori/http"
	"github.com/sirupsen/logrus"
)

func Transport(router http.Http, srv Service, log *logrus.Logger) {
	dummyHandler := http.NewServer(
		MakeDummyEndpoint(srv),
		DecodeDummyRequest,
		http.EncodeJSONResponse,
		log,
	)
	dummyPostHandler := http.NewServer(
		MakeDummyPostEndpoint(srv),
		DecodeDummyPostRequest,
		http.EncodeJSONResponse,
		log,
	)

	router.Handle("/dummy", dummyHandler).Methods("GET")
	router.Handle("/dummy", dummyPostHandler).Methods("POST")
}
