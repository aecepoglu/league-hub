package graphql

type resolvers struct{}

func (_ *resolvers) Hello() string {
	return "Hello, World!"
}
