package types

const (
	GOOGLE = "google"
)

type SocialMedia string

const (
	Google SocialMedia = "google"
)

func (sm SocialMedia) IsValid() bool {
	return sm == Google
}
