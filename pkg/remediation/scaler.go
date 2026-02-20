package remediation

import "fmt"

// Scaler provides simple scaling remediation actions.
type Scaler struct{}

// NewScaler constructs a Scaler.
func NewScaler() *Scaler { return &Scaler{} }

// ScaleUp is a stub that simulates scaling a Deployment.
func (s *Scaler) ScaleUp(name string, replicas int) error {
    fmt.Printf("Scaling %s to %d replicas\n", name, replicas)
    return nil
}
