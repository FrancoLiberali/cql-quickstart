package models

import (
	"github.com/FrancoLiberali/cql/model"
)

type MyModel struct {
	model.UUIDModel

	Name string
}
