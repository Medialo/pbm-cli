package utils

const MonorepoVersion = "1.0.0"

func CheckVersion(version string) bool {
	return version == MonorepoVersion
}
