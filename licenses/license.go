package licenses

// LicenseMap is a convenient way of getting a license text from its name
var LicenseMap map[string]string

const (
	// MIT license
	MIT = "MIT"
	// Apache2 license
	Apache2 = "Apache2"
)

func init() {
	LicenseMap = make(map[string]string)
	LicenseMap[MIT] = mit
	LicenseMap[Apache2] = apache2
}
