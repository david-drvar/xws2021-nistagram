package model

var (
	areCampaignsOngoing = false
)

func GetAreCampaignsOngoing() bool {
	return areCampaignsOngoing
}

func SetAreCampaignsOngoing(value bool){
	areCampaignsOngoing = value
}