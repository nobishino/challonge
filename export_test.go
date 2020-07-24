package challonge

func SetBaseUrl(url string) func() {
	original := baseUrl
	baseUrl = url
	return func() {
		baseUrl = original
	}
}
