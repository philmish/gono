package api

import (
	"github.com/philmish/gono/api/endpoints"
)

func Serve(addr, env string) error {
    router := endpoints.NewRouter()
    return router.Run(addr)
}
