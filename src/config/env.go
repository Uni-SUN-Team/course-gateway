package config

import "os"

func ConfigENV() {
	os.Setenv("PORT", "8080")
	os.Setenv("CONTEXT_PATH", "/course-listenner")
	// Strapi information gateway
	os.Setenv("HOST_STRAPI_SERVICE", "https://api.unisun.dynu.com")
	os.Setenv("PATH_STRAPI_INFORMATION_GATEWAY", "/strapi-information-gateway/api/strapi")
	// Path
	os.Setenv("PATH_STRAPI_ARTICLE", "/api/course")
}
