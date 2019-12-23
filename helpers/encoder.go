package helpers

import "encoding/json"

func StructureToString(value interface{}) (string, error) {
    var result []byte
    var err error
    result, err = json.Marshal(value)
    return string(result), err
}
