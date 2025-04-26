package constant

type ServiceKeys string

const (
	ServiceKeyHello        ServiceKeys = "HELLO"
	ServiceKeyEptToolchain ServiceKeys = "EPT_TOOLCHAIN"
	ServiceKeyPkgSoftware  ServiceKeys = "PKG_SOFTWARE"
)

const (
	ServicePathHello        = "/api/hello"
	ServicePathEptToolchain = "/api/ept/toolchain"
	ServicePathPkgSoftware  = "/api/pkg/software"
)
