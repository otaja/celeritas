package middleware

import (
	"github.com/otaja/nexterp/data"

	"github.com/otaja/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
