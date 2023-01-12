package entities

type Product struct {
	BaseEntity
	Name                string
	Config              map[string]interface{}
	MaxFailureCountUser *int
	MaxFailureTimeUser  *int
}
