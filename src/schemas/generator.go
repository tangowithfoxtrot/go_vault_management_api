package schemas

type GeneratorParams struct {
	Length        int
	Uppercase     bool
	Lowercase     bool
	Number        bool
	Special       bool
	Passphrase    bool
	Words         int
	Separator     string
	Capitalize    bool
	IncludeNumber bool
}

type GeneratorResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Object string `json:"object"`
		Data   string `json:"data"`
	} `json:"data"`
}
