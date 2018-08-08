package main

import (
	"errors"

	"github.com/Henry-Sarabia/prismata"
)

type Player struct {
	InitialRating prismata.Rating
	FinalRating   prismata.Rating
	RatingDelta   []float64
	ScoreDelta    int
}

func getPlayer(rep *prismata.Replay, n int) (*Player, error) {
	if n != 1 && n != 2 {
		return nil, errors.New("n must be 1 or 2")
	}
	p := Player{}
	r := rep.RatingInfo

	p.InitialRating = r.InitialRatings[n-1]
	p.FinalRating = r.FinalRatings[n-1]
	p.RatingDelta = r.RatingChanges[n-1]
	p.ScoreDelta = r.ScoreChanges[n-1]

	return &p, nil
}

func GetPlayerOne(r *prismata.Replay) (*Player, error) {
	return getPlayer(r, 1)
}

func GetPlayerTwo(r *prismata.Replay) (*Player, error) {
	return getPlayer(r, 2)
}
