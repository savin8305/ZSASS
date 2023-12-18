// stores/interface.go

package stores

import (
	"github.com/savin8305/ZSASS/models"
	"gofr.dev/pkg/gofr"
)
// Car interface defines methods for CRUD operations on Car entities
type Car interface {
	Get(ctx *gofr.Context, name string) ([]models.Car, error)
	Create(ctx *gofr.Context, model models.Car) error
	Delete(ctx *gofr.Context, name string) (int, error)
	List(ctx *gofr.Context) ([]models.Car, error) // Add List method
}
