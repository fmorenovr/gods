package bstree

import (
  "encoding/json";
  "github.com/jenazads/goutils";
)

func assertSerializationImplementation() {
  var _ goutils.JSONSerializer = (*BSTree)(nil)
  var _ goutils.JSONDeserializer = (*BSTree)(nil)
}

// ToJSON return JSON format of elements
func (t *BSTree) ToJSON() ([]byte, error) {
  elements := make(map[string]interface{})
  it := t.Iterator()
  for it.Next() {
    elements[goutils.ToString(it.Key())] = it.Value()
  }
  return json.Marshal(&elements)
}

// FromJSON Convert to JSON format the elements
func (t *BSTree) FromJSON(data []byte) error {
  elements := make(map[string]interface{})
  err := json.Unmarshal(data, &elements)
  if err == nil {
    t.Clear()
    for key, value := range elements {
      t.Insert(key, value)
    }
  }
  return err
}
