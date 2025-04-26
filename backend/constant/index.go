package constant

const API_PREFIX = "/api"

type ServiceKeys string

const (
	ServiceKeyHello        ServiceKeys = "HELLO"
	ServiceKeyEptToolchain ServiceKeys = "EPT_TOOLCHAIN"
	ServiceKeyPkgSoftware  ServiceKeys = "PKG_SOFTWARE"
)

const (
	ServicePathHello        = "/hello"
	ServicePathEptToolchain = "/ept/toolchain"
	ServicePathPkgSoftware  = "/pkg/software"
)
