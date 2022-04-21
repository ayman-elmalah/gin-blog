package static

func GetStaticPaths() map[string]string {
	return map[string]string{
		"/assets": "./public/assets",
	}
}
