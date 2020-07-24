package challonge

func SetBaseURL(url string) func() {
	original := baseURL
	baseURL = url
	return func() {
		baseURL = original
	}
}
