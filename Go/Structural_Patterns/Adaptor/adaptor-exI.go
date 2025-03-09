package adaptor

type IHttpClient interface {
	get(url string) string
}

type OldHttpClient struct{}

func (o *OldHttpClient) request(url string) string {
	return "OldHttpClient: " + url
}

type OldHttpClientAdaptor struct {
	Client *OldHttpClient
}

func (o *OldHttpClientAdaptor) Get(url string) string {
	return o.Client.request(url)
}
