package controller

import (
    "fmt"
    "time"
)

// Controller is a lightweight stub for the control loop.
type Controller struct{}

// NewController constructs a Controller.
func NewController() *Controller { return &Controller{} }

// Run starts the controller and keeps it alive (placeholder).
func (c *Controller) Run() {
    fmt.Println("controller running")
    for {
        time.Sleep(30 * time.Second)
    }
}
