package main

import (
	"fmt"
	"github.com/facebookgo/inject"
	"net/http"
	"os"
)

type NameAPI struct {
	HTTPTrasport http.RoundTripper `inject:""`
}

func (a *NameAPI) Name(id uint64) interface{} {
	return "Spock"
}

type PlanetAPI struct {
	HTTPTransport http.RoundTripper `inject:""`
}

func (a *PlanetAPI) Planet(id uint64) interface{} {
	return "Vulcan"
}

type HomePlanetRenderApp struct {
	NameAPI   *NameAPI   `inject:""`
	PlanetAPI *PlanetAPI `index:""`
}

func (a *HomePlanetRenderApp) Render(id uint64) string {
	return fmt.Sprintf(
		"%s is from the planet %s.",
		a.NameAPI.Name(id),
		a.PlanetAPI.Planet(id),
	)
}

func main() {
	var g inject.Graph

	var a HomePlanetRenderApp
	err := g.Provide(
		&inject.Object{Value: &a},
		&inject.Object{Value: http.DefaultTransport},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := g.Populate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(a.Render(42))
}
