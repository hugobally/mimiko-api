// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqltypes

type KnotInput struct {
	TrackID  *string `json:"trackId"`
	Level    *int    `json:"level"`
	Visited  *bool   `json:"visited"`
	SourceID *uint   `json:"sourceId"`
}

type MapInput struct {
	Title      *string `json:"title"`
	Public     bool    `json:"public"`
	FlagshipID *string `json:"flagshipID"`
}

type MapsFilter struct {
	Author *uint `json:"author"`
}

type MutationResult struct {
	Success bool `json:"success"`
	Count   int  `json:"count"`
}

type SpotifyAuthToken struct {
	AccessToken string `json:"accessToken"`
	TokenExpiry string `json:"tokenExpiry"`
}