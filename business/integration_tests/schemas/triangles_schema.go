package schemas

func TrianglesSchema() interface{} {
	return `{
	"type": "object",
	"properties": {
			"id": {
				"type": "string"
			},
			"side1": {
				"type": "integer"
			},
			"side2": {
				"type": "integer"
			},
			"side3": {
				"type": "integer"
			},
			"type": {
				"type": "string"
			}
		},
	"required": ["id", "side1", "side2", "side3", "type"]
	}`
}

func ListTrianglesSchema() interface{} {
	return `{
	"type": "array",
	"items": {
		"type": "object",
		"properties": {
			"id": {
				"type": "string"
			},
			"side1": {
				"type": "integer"
			},
			"side2": {
				"type": "integer"
			},
			"side3": {
				"type": "integer"
			},
			"type": {
				"type": "string"
			}
		}
	},
	"required": ["id", "side1", "side2", "side3", "type"]
	}`
}
