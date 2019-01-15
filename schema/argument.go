
package schema

type Argument struct {
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	Type         *OfType      `json:"type"`
	DefaultValue interface{} `json:"defaultValue"`
}