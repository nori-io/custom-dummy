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
)

type Service interface {
	Dummy(ctx context.Context, req interface{}) (resp *DummyResponse)
	DummyPost(ctx context.Context, req interface{}) (resp *DummyResponse)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) Dummy(ctx context.Context, req interface{}) (resp *DummyResponse) {
	e := DummyResponse{
		Message: "Dummy message",
		Err:     nil,
	}
	return &e
}

func (s *service) DummyPost(ctx context.Context, req interface{}) (resp *DummyResponse) {
	r := req.(DummyPostRequest)
	e := DummyResponse{
		Message: r.Message,
		Err:     nil,
	}
	return &e
}
