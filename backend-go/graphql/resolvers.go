package graphql

// RootResolver .
type RootResolver struct{}

// Movies query resolver
func (*RootResolver) Movies(args struct {
	rating float64
	limit  int
}) {

}
