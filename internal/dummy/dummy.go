package dummy

import (
	"github.com/nori-io/common/v4/pkg/domain/registry"
	"github.com/nori-io/interfaces/nori/http"
)

type Instance struct {
	Http http.Http
}

func New(r registry.Registry) (*Instance, error) {

	dummyNew, err := http.GetHttp(r)

	if err != nil {
		return nil, err
	}

	return &Instance{
		Http: dummyNew,
	}, nil
}
