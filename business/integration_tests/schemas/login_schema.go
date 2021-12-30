package schemas

type TokenStruct struct {
	Token   string `json:"access_token"`
	ExpTime int64  `json:"expTime"`
}

func LoginSuccessSchema() interface{} {
	return `{
	"type": "object",
	"properties": {
			"access_token": {
				"type": "string"
			},
			"expTime": {
				"type": "integer"
			}
		},
	"required": ["access_token", "expTime"]
	}`
}
