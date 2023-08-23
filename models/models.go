package models

import (
	"github.com/ditrit/badaas/orm/model"
)

type MyModel struct {
	model.UUIDModel

	Name string
}
