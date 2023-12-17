// stores/interface.go

package stores

import (
	"github.com/savin8305/ZSASS/models"
	"gofr.dev/pkg/gofr"
)
// Customer interface defines methods for CRUD operations on Customer entities
type Customer interface {
	Get(ctx *gofr.Context, name string) ([]models.Customer, error)
	Create(ctx *gofr.Context, model models.Customer) error
	Delete(ctx *gofr.Context, name string) (int, error)
	List(ctx *gofr.Context) ([]models.Customer, error) // Add List method
}
