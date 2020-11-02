package matrices

// interface to define basic functionalities
type matrix interface {
	Get(i, j int) (int, error)
	Set(i, j, val int) error
	String() string
}
