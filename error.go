package serpstat

type SerpstatError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
