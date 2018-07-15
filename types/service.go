package types

// Service...
type Service interface {
	// Output
	Output() ([]byte, error)
	// Run
	Run() error
}
