package models

import (
	"math/rand"
	"time"
)

// GeneratorPrice object for generate Current course
type GeneratorPrice struct {
	Comp       *Company
	LastCourse *Course
}

// Course object current course
type Course struct {
	Comp *Company
	Bid  int64
	Ask  int64
	Time time.Time
}

// NewCourse Init first course for company
func NewCourse(comp *Company) *Course {
	bad := rand.Int63n(int64(comp.MaxBid-comp.MinBid)) + int64(comp.MinBid)
	ask := bad + rand.Int63n(int64(comp.MaxAskBidDiff))
	return &Course{
		Comp: comp,
		Bid:  bad,
		Ask:  ask,
		Time: time.Now(),
	}
}

// NewGenerator Constructor
func NewGenerator(comp *Company, initCourse *Course) *GeneratorPrice {
	return &GeneratorPrice{
		Comp:       comp,
		LastCourse: initCourse,
	}
}

func (g *GeneratorPrice) GenerateCourse() *Course {
	newBidDiff := rand.Int63n(int64(g.Comp.MaxBidDifferent))
	if g.LastCourse.Bid+newBidDiff >= int64(g.Comp.MaxBid) {
		g.LastCourse.Bid -= newBidDiff
	} else {
		g.LastCourse.Bid += newBidDiff
	}
	newAskDiff := rand.Int63n(int64(g.Comp.MaxAskDifferent))
	if g.LastCourse.Ask+newAskDiff >= int64(g.Comp.MaxAsk) {
		g.LastCourse.Ask -= newAskDiff
	} else {
		g.LastCourse.Ask += newAskDiff
	}
	return g.LastCourse
}
