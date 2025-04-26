package constant

const API_PREFIX = "/api"

type ServiceKeys string

const (
	ServiceKeyHello        ServiceKeys = "HELLO"
	ServiceKeyEptToolchain ServiceKeys = "EPT_TOOLCHAIN"
	ServiceKeyPkgSoftware  ServiceKeys = "PKG_SOFTWARE"
	ServiceKeyRedirect     ServiceKeys = "REDIRECT"
)

const (
	ServicePathHello            = "/hello"
	ServicePathEptToolchain     = "/ept/toolchain"
	ServicePathPkgSoftware      = "/pkg/software"
	ServicePathRedirect         = "/pkg/redirect/:scope/:software/:file_name"
	ServicePathRedirectTemplate = "/pkg/redirect/{scope}/{software}/{file_name}"
)
