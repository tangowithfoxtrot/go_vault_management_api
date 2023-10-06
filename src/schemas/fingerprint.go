package schemas

type FingerprintResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Object string `json:"object"`
		Data   string `json:"data"`
	} `json:"data"`
}
