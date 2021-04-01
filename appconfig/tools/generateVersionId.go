package tools

import "strconv"

func GenerateVersionId(id string, version int32) string{
	return id + "/v" + strconv.Itoa(int(version))
}