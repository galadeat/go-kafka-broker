package protocol

func IsSuportedVersion(v int16) bool {
	return v >= 0 && v <= 4 // Supprots versions up to 4
}
