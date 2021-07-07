package model

import "fmt"

type UserRole string

const (
	Basic UserRole = "Basic"
	Agent          = "Agent"
)

func (r UserRole) String() string {
	switch r {
	case Basic:
		return "Basic"
	case Agent:
		return "Agent"
	default:
		return fmt.Sprintf("%s", string(r))
	}
}

func GetUserRole(r string) UserRole {
	switch r {
	case "Basic":
		return Basic
	case "Agent":
		return Agent
	default:
		return ""
	}
}
