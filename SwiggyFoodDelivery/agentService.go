package main

type DeliveryAgentStatus string

const (
	AVAILABLE   DeliveryAgentStatus = "AVAILABLE"
	UNAVAILABLE DeliveryAgentStatus = "UNAVAILABLE"
	ON_DELIVERY DeliveryAgentStatus = "ON_DELIVERY"
)

type DeliveryAgent struct {
	Id       string
	Loaction string
	status   DeliveryAgentStatus
}

type AgentAssignmentStrategy interface {
	Assign() error
}

type NearestAgentAssignment struct{}

func (n *NearestAgentAssignment) Assign() *DeliveryAgent {
	// some logic
	return &DeliveryAgent{
		Id:       "da",
		Loaction: "pincode",
		status:   ON_DELIVERY,
	}
}
