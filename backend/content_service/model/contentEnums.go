package model

import "fmt"

type PostType string
const(
	TypePost  PostType = "Post"
	TypeStory PostType = "Story"
)

func (pt PostType) String() string{
	switch pt {
	case TypePost:
		return "Post"
	case TypeStory:
		return "Story"
	default:
		return fmt.Sprintf("%s", string(pt))
	}
}

func GetPostType(pt string) PostType{
	switch pt {
	case "Post", "post":
		return TypePost
	case "Story", "story":
		return TypeStory
	default:
		return ""
	}
}

type MediaType string
const(
	TypeImage MediaType = "Image"
	TypeVideo           = "Video"
)

func (mt MediaType) String() string{
	switch mt {
	case TypeImage:
		return "Image"
	case TypeVideo:
		return "Video"
	default:
		return fmt.Sprintf("%s", string(mt))
	}
}

func GetMediaType(mt string) MediaType{
	switch mt {
	case "Image":
		return TypeImage
	case "Video":
		return TypeVideo
	default:
		return ""
	}
}

type RequestStatus string
const(
	Pending  RequestStatus = "Pending"
	Accepted               = "Accepted"
	Rejected               = "Rejected"
)

func (mt RequestStatus) String() string{
	switch mt {
	case Pending:
		return "Pending"
	case Accepted:
		return "Accepted"
	case Rejected:
		return "Rejected"
	default:
		return fmt.Sprintf("%s", string(mt))
	}
}

func GetRequestStatus(mt string) RequestStatus{
	switch mt {
	case "Pending":
		return Pending
	case "Accepted":
		return Accepted
	case "Rejected":
		return Rejected
	default:
		return ""
	}
}

type ComplaintCategory string
const(
	Gore       ComplaintCategory = "Gore"
	Nudity                       = "Nudity"
	Violence                     = "Violence"
	Suicide                      = "Suicide"
	FakeNews                     = "Fake News"
	Spam                         = "Spam"
	HateSpeech                   = "Hate Speech"
	Terrorism                    = "Terrorism"
	Harassment                   = "Harassment"
	Other                        = "Other"
)

func (cc ComplaintCategory) String() string{
	switch cc {
	case Gore:
		return "Gore"
	case Nudity:
		return "Nudity"
	case Suicide:
		return "Suicide"
	case FakeNews:
		return "FakeNews"
	case Spam:
		return "Spam"
	case HateSpeech:
		return "HateSpeech"
	case Terrorism:
		return "Terrorism"
	case Harassment:
		return "Harassment"
	case Other:
		return "Other"
	default:
		return fmt.Sprintf("%s", string(cc))
	}
}

func GetComplaintCategory(cc string) ComplaintCategory{
	switch cc {
	case "Gore":
		return Gore
	case "Nudity":
		return Nudity
	case "Suicide":
		return Suicide
	case "FakeNews":
		return FakeNews
	case "Spam":
		return Spam
	case "HateSpeech":
		return HateSpeech
	case "Terrorism":
		return Terrorism
	case "Harassment":
		return Harassment
	case "Other":
		return Other
	default:
		return ""
	}
}