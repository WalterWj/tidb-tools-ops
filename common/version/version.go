package version

var (
	// bild hash
	build string
	// GO version
	goVersion string
	// build time
	buildTime string
	// version
	version string
)

func NewToolsVersion() map[string]string {
	return map[string]string{
		"build":     build,
		"goVersion": goVersion,
		"buildTime": buildTime,
		"version":   version,
	}

}
