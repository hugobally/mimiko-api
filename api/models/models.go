package models

type MutationResult struct {
	Success bool
	Count   int
}

type KnotInput struct {
	TrackId *string
	Level   *int32
	Visited *bool
}

type MapInput struct {
	Title      *string
	Public     *bool
	FlagshipId *string
}

type MapsFilter struct {
	UserId *string
	Offset int32
	Limit  int32
}
