package constant

const (
	ProductionEnv = "production"
	StagingEnv    = "staging"
	DevelopmentEnv = "development"
)

// EnvMap is a map of environment
var EnvMap = map[string]bool{
	ProductionEnv: true,
	StagingEnv:    true,
	DevelopmentEnv: true,
}