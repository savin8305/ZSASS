package handlers

import (
	"fmt"

	"github.com/savin8305/ZSASS/models"
	"github.com/savin8305/ZSASS/stores"
	"gofr.dev/pkg/gofr"
)

type handler struct {
	store stores.Customer
}

// New is factory function for handler layer
//
//nolint:revive // handler should not be used without proper initilization with required dependency
func New(c stores.Customer) handler {
	return handler{store: c}
}

func (h handler) Get(ctx *gofr.Context) (interface{}, error) {
	name := ctx.Param("name")

	resp, err := h.store.Get(ctx, name)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var c models.Customer

	err := ctx.Bind(&c)
	if err != nil {
		return nil, err
	}

	err = h.store.Create(ctx, c)
	if err != nil {
		return nil, err
	}

	return "New Customer Added!", nil
}

func (h handler) Delete(ctx *gofr.Context) (interface{}, error) {
	name := ctx.Param("name")

	deleteCount, err := h.store.Delete(ctx, name)
	if err != nil {
		return nil, err
	}

	return fmt.Sprintf("%v Customers Deleted!", deleteCount), nil
}

// Add the following methods to the handler struct

func (h handler) List(ctx *gofr.Context) (interface{}, error) {
	resp, err := h.store.List(ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h handler) Update(ctx *gofr.Context) (interface{}, error) {
	name := ctx.Param("name")

	var updateData models.Customer
	err := ctx.Bind(&updateData)
	if err != nil {
		return nil, err
	}

	err = h.store.Update(ctx, name, updateData)
	if err != nil {
		return nil, err
	}

	return "Customer Updated!", nil
}
