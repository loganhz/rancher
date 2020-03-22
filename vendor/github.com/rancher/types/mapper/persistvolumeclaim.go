package mapper

import (
	"encoding/json"
	"github.com/rancher/norman/types"
	"github.com/sirupsen/logrus"
)

type PersistVolumeClaim struct {
}

func (p PersistVolumeClaim) FromInternal(data map[string]interface{}) {
}

func (p PersistVolumeClaim) ToInternal(data map[string]interface{}) error {
	jsonOutput, _ := json.Marshal(r)
	logrus.Info("loganpvc" + jsonOutput)
	return nil
}

func (p PersistVolumeClaim) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}
