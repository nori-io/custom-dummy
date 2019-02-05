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
package main

import (
	"context"

	"github.com/nori-io/nori-common/meta"

	cfg "github.com/nori-io/nori-common/config"

	iplugin "github.com/nori-io/nori-common/plugin"
	"github.com/nori-io/dummy/service"
)

type plugin struct {
	instance service.Service
}

var (
	Plugin plugin
)

func (p *plugin) Init(_ context.Context, configManager cfg.Manager) error {
	configManager.Register(p.Meta())
	return nil
}

func (p *plugin) Instance() interface{} {
	return p.instance
}

func (p plugin) Meta() meta.Meta {
	return &meta.Data{
		ID: meta.ID{
			ID:      "nori/custom/dummy",
			Version: "1.0",
		},
		Author: meta.Author{
			Name: "Nori",
			URI:  "https://nori.io",
		},
		Core: meta.Core{
			VersionConstraint: ">=1.0, <2.0",
		},
		Dependencies: []meta.Dependency{meta.HTTP.Dependency("")},
		Description: meta.Description{
			Name:        "Nori InMemory Cache",
			Description: "InMemory Cache",
		},
		Interface: meta.Custom,
		License: meta.License{
			Title: "",
			Type:  "GPLv3",
			URI:   "https://www.gnu.org/licenses/",
		},
		Tags: []string{"dummy", "rest", "api"},
	}
}

func (p *plugin) Start(ctx context.Context, registry iplugin.Registry) error {
	if p.instance == nil {
		p.instance = service.NewService()
		http, _ := registry.Http()
		service.Transport(http, p.instance, registry.Logger(p.Meta()))
	}
	return nil
}

func (p *plugin) Stop(_ context.Context, _ iplugin.Registry) error {
	p.instance = nil
	return nil
}
