package btree

import (
  "encoding/json";
  "github.com/jenazads/goutils";
)

func assertSerializationImplementation() {
  var _ goutils.JSONSerializer = (*BTree)(nil)
  var _ goutils.JSONDeserializer = (*BTree)(nil)
}

// ToJSON return JSON format of elements
func (t *BTree) ToJSON() ([]byte, error) {
  elements := make(map[string]interface{})
  it := t.Iterator()
  for it.Next() {
    elements[goutils.ToString(it.Key())] = it.Value()
  }
  return json.Marshal(&elements)
}

// FromJSON Convert to JSON format the elements
func (t *BTree) FromJSON(data []byte) error {
  elements := make(map[string]interface{})
  err := json.Unmarshal(data, &elements)
  if err == nil {
    t.Clear()
    for key, value := range elements {
      t.Put(key, value)
    }
  }
  return err
}
