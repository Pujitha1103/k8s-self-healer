package controller

import "fmt"

// Controller is a lightweight stub for the control loop.
type Controller struct{}

// NewController constructs a Controller.
func NewController() *Controller { return &Controller{} }

// Run starts the controller. This is a placeholder.
func (c *Controller) Run() { fmt.Println("controller running") }
