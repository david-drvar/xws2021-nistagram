package saga

import "encoding/json"

const (
	UserChannel           string = "UserChannel"
	RecommendationChannel string = "RecommendationChannel"
	ReplyChannel          string = "ReplyChannel"
	ServiceUser           string = "User"
	ServiceRecommendation string = "Recommendation"
	ActionStart           string = "Start"
	ActionDone            string = "DoneMsg"
	ActionError           string = "ErrorMsg"
	ActionRollback        string = "RollbackMsg"
)

type Message struct {
	Service       string `json:"service"`
	SenderService string `json:"sender_service"`
	Action        string `json:"action"`
	UserId        string `json:"user_id"`
	Ok            bool   `json:"ok"`
}

func (m Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}
