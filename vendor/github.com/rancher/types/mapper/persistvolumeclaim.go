package mapper

import (
	"encoding/json"
	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/values"
	"github.com/sirupsen/logrus"
)

type PersistVolumeClaim struct {
}

func (p PersistVolumeClaim) FromInternal(data map[string]interface{}) {
}

func (p PersistVolumeClaim) ToInternal(data map[string]interface{}) error {
	jsonOutput, _ := json.Marshal(data)
	if v, ok := values.GetValue(data, "storageClassId"); ok && v == nil {
		values.PutValue(data, "", "storageClassId")
	}
	jsonOutput, _ = json.Marshal(data)
	logrus.Info("loganpvc1" + string(jsonOutput))
	values.RemoveValue(data, "storageClassId")
	jsonOutput, _ = json.Marshal(data)
	logrus.Info("loganpvc2" + string(jsonOutput))
	return nil
}

func (p PersistVolumeClaim) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}
