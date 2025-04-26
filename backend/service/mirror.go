package service

import (
	"nlm/config"
	"nlm/constant"
	"nlm/vo"
)

func MirrorHello() vo.MirrorHello {
	return vo.MirrorHello{
		Name:        config.ENV.MIRROR_HELLO_NAME,
		Locale:      config.ENV.MIRROR_HELLO_LOCALE,
		Description: config.ENV.MIRROR_HELLO_DESCRIPTION,
		Maintainer:  config.ENV.MIRROR_HELLO_MAINTAINER,
		Protocol:    "1.0.0",
		RootURL:     config.ENV.ROOT_URL,
		Property:    config.ENV.MIRROR_HELLO_PROPERTY,
		Service: []vo.MirrorHelloService{
			{
				Key:  constant.ServiceKeyHello,
				Path: constant.API_PREFIX + constant.ServicePathHello,
			},
			{
				Key:  constant.ServiceKeyPkgSoftware,
				Path: constant.API_PREFIX + constant.ServicePathPkgSoftware,
			},
			// {
			// 	Key:  constant.ServiceKeyEptToolchain,
			// 	Path: constant.API_PREFIX + constant.ServicePathEptToolchain,
			// },
		},
	}
}

func MirrorPkgSoftware() vo.MirrorPkgSoftware {
	return GetMirrorPkgSoftware()
}
