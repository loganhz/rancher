package mapper

import (
	"encoding/json"
	"github.com/rancher/norman/types"
	"github.com/rancher/norman/types/convert"
	"github.com/rancher/norman/types/values"
	"github.com/sirupsen/logrus"
)

type WorkloadAnnotations struct {
}

func (n WorkloadAnnotations) FromInternal(data map[string]interface{}) {
	v, ok := values.RemoveValue(data, "workloadAnnotations", "field.cattle.io/publicEndpoints")
	if ok {
		if _, ok := data["annotations"]; !ok {
			data["annotations"] = map[string]interface{}{}
		}
		annotations := convert.ToMapInterface(data["annotations"])
		annotations["field.cattle.io/publicEndpoints"] = v
	}
}

func (n WorkloadAnnotations) ToInternal(data map[string]interface{}) error {
	jsonOutput, _ := json.Marshal(r)
	logrus.Info("loganw" + jsonOutput)
	return nil
}

func (n WorkloadAnnotations) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}
