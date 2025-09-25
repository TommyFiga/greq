package types

type ResponseData struct {
	Protocol string
	Status   string
	Headers  map[string][]string
	Body     []byte
}