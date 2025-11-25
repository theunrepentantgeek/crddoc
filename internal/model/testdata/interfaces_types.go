package v1

import "io"

// Greeter is an interface for things that can greet.
type Greeter interface {
	// Greet returns a greeting message.
	Greet() string
}

// Speaker is an interface for things that can speak.
type Speaker interface {
	// Speak returns what the speaker says.
	Speak(message string) string
	// Volume returns the volume level.
	Volume() int
}

// MultiTalent is an interface combining multiple capabilities.
type MultiTalent interface {
	Greeter
	Speaker
	// Perform does something special.
	Perform() error
}

// Robot is a type that implements Greeter.
type Robot struct {
	Name        string `json:"name"`
	Model       string `json:"model"`
	VolumeLevel int    `json:"volumeLevel"`
}

// Greet returns a greeting from the robot.
func (r Robot) Greet() string {
	return "Hello, I am " + r.Name
}

// Human is a type that implements both Greeter and Speaker.
type Human struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
}

// Greet returns a greeting from the human.
func (h *Human) Greet() string {
	return "Hi, my name is " + h.Name
}

// Speak makes the human say something.
func (h *Human) Speak(message string) string {
	return h.Name + " says: " + message
}

// Volume returns the human's speaking volume.
func (h *Human) Volume() int {
	return 50
}

// Type assertions to document interface implementations
var _ Greeter = Robot{}
var _ Greeter = &Human{}
var _ Speaker = &Human{}

// Type assertion for an external interface
var _ io.Reader = &CustomReader{}

// CustomReader is a type that implements io.Reader.
type CustomReader struct {
	Data []byte `json:"data"`
}

// Read implements io.Reader.
func (r *CustomReader) Read(p []byte) (n int, err error) {
	copy(p, r.Data)
	return len(r.Data), nil
}
