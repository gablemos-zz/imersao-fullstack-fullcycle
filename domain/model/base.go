package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init(){
	govalidator.SetFieldsRequiredByDefault(true)
}

type Base struct {
	ID string `json:"id" valid:"uuid"`
	CreatedAt time.Time `json:"createdAt" valid:"-"`
	UpdatedAt time.Time `json:"updatedAt" valid:"-"`
}