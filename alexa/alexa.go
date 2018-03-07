package alexa

// Endpoint are objects for alexas smart home skills
type Endpoint struct {
	EndpointID        string   `json:"endpointId"`
	ManufacturerName  string   `json:"manufacturerName,omitempty"`
	ModelName         string   `json:"modelName,omitempty"`
	FriendlyName      string   `json:"friendlyName,omitempty"`
	Description       string   `json:"description,omitempty"`
	DisplayCategories []string `json:"displayCategories,omitempty"`
	// possible values: 'turnOn', 'turnOff', ...
	Capabilities *[]Capability     `json:"capabilities,omitempty"`
	Cookie       map[string]string `json:"cookie,omitempty"`
	Scope        *Scope            `json:"scope,omitempty"`
}

// Scope for respones authentication
type Scope struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}

// Capability describes endpoints use
type Capability struct {
	Type       string              `json:"type"`
	Interface  string              `json:"interface"`
	Version    string              `json:"version"`
	Properties *CapabilityProperty `json:"properties,omitempty"`
}

// CapabilityProperty used for describing capabilities
type CapabilityProperty struct {
	Supported           []map[string]string `json:"supported,omitempty"`
	ProactivelyReported bool                `json:"proactivelyReported"`
	Retrievable         bool                `json:"retrievable"`
}

// Header objects for alexas requests and responses
type Header struct {
	Namespace        string `json:"namespace"`
	Name             string `json:"name"`
	PayloadVersion   string `json:"payloadVersion"`
	MessageID        string `json:"messageId,omitempty"`
	CorrelationToken string `json:"correlationToken,omitempty"`
}

// Response object for alexa skills.
// header: {...}, payload: {...}
type Response struct {
	Context *Context `json:"context,omitempty"`
	Event   Event    `json:"event"`
}

// Request object for alexa skill requests.
type Request struct {
	Directive Directive `json:"directive"`
}

// Directive for alexa skill requests. Payload is unspecified at this point.
type Directive struct {
	Header   Header   `json:"header"`
	Endpoint Endpoint `json:"endpoint,omitempty"`
	Payload  Payload  `json:"payload,omitempty"`
}

// Event is a wrapper for responses
type Event struct {
	Header   Header    `json:"header"`
	Endpoint *Endpoint `json:"endpoint,omitempty"`
	Payload  Payload   `json:"payload,omitempty"`
}

// Payload object for alexa respones.
// payload: {...}
type Payload interface {
}

// ColorPayload object for alexa responses.
type ColorPayload struct {
	Hue        float64 `json:"hue"`
	Saturation float64 `json:"saturation"`
	Brightness float64 `json:"brightness"`
}

// Context for responses.
type Context struct {
	Properties []IContextProperty `json:"properties"`
}

// IContextProperty for context property differentation.
type IContextProperty interface{}

// ContextProperty type for context objects in responses. Subclassing is necessary.
type ContextProperty struct {
	Namespace                 string      `json:"namespace"`
	Name                      string      `json:"name"`
	Value                     interface{} `json:"value"`
	TimeOfSample              string      `json:"timeOfSample"`
	UncertaintyInMilliseconds int         `json:"uncertaintyInMilliseconds"`
}

// DiscoveryResponsesPayload for finding devices.
type DiscoveryResponsesPayload struct {
	Endpoints []Endpoint `json:"endpoints"`
}

// EmptyPayload for responses without payload data.
type EmptyPayload struct{}

// ErrorPayload for return error responses.
type ErrorPayload struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// Channel for responses in Changechannel Controller
type Channel struct {
	Number            string `json:"number"`
	CallSign          string `json:"callSign"`
	AffiliateCallSign string `json:"affiliateCallSign"`
}
