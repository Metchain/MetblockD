package utils

import (
	"fmt"
)

func ConvertStringToProtoByte(s string) []byte {
	return []byte(s)
}

func ConvertDomainByteToProtoByte(b [32]byte) []byte {
	nv := []byte{}
	for _, v := range b {
		nv = append(nv, v)
	}

	return nv
}

func ConvertDomainInt8ToProtoInt64(i int8) int64 {
	return int64(i)
}

func ConvertValToString(v float32) string {
	return fmt.Sprintf("%v", v)
}

func ConvertDomainByteToProtoByte32(b []byte) [32]byte {
	nv := [32]byte{}
	for k, v := range b {
		nv[k] = v
	}

	return nv
}
