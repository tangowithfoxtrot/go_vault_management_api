package schemas

type UnlockRequest struct {
	Password string `json:"password"`
}

type UnlockResponse struct {
	Success bool `json:"success"`
	Data    struct {
		NoColor bool   `json:"noColor"`
		Object  string `json:"object"`
		Title   string `json:"title"`
		Message any    `json:"message"`
	} `json:"data"`
}
