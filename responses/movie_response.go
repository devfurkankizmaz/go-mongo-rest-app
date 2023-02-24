package responses

type MovieResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"msg"`
	Records map[string]interface{} `json:"records"`
}
