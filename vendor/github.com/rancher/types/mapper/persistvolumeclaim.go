package mapper

import (
	"github.com/rancher/norman/types"
)

type PersistVolumeClaim struct {
}

func (p PersistVolumeClaim) FromInternal(data map[string]interface{}) {
}

func (p PersistVolumeClaim) ToInternal(data map[string]interface{}) error {
	return nil
}

func (p PersistVolumeClaim) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}
