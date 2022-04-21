package converters

import "encoding/json"

func MapToString(mapData map[string]string) string {
	out, _ := json.Marshal(mapData)

	return string(out)
}

func StringToMap(stringData string) map[string]string {
	var out map[string]string
	_ = json.Unmarshal([]byte(stringData), &out)

	return out
}

func UrlValuesToString(urlValuesData map[string][]string) string {
	out, _ := json.Marshal(urlValuesData)

	return string(out)
}

func StringToUrlValues(stringData string) map[string][]string {
	var out map[string][]string
	_ = json.Unmarshal([]byte(stringData), &out)

	return out
}
