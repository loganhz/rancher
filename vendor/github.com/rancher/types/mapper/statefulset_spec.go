package mapper

import (
	"encoding/json"
	"github.com/rancher/norman/types"
	"github.com/sirupsen/logrus"
)

type StatefulSetSpecMapper struct {
}

func (s StatefulSetSpecMapper) FromInternal(data map[string]interface{}) {
}

func (s StatefulSetSpecMapper) ToInternal(data map[string]interface{}) error {
	jsonOutput, _ := json.Marshal(data)
	logrus.Info("logan22" + string(jsonOutput))
	return nil
}

func (s StatefulSetSpecMapper) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}
