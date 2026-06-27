package pokeapi

import (
	"net/http"

	pokecache "github.com/william-jce/pokedex/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	Cache      *pokecache.Cache
}
