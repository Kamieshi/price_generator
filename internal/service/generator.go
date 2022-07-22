// Package service Business logic generator
package service

import (
	"math/rand"
	"time"

	"priceGenerator/internal/models"
)

// Generator Work with courses
type Generator struct {
	LastCourse *models.Course
}

// NewGenerator Constructor
func NewGenerator(comp *models.Company) *Generator {
	ask := rand.Int63n(1000)
	bid := ask - rand.Int63n(100)
	initCourse := models.Course{
		Comp: comp,
		Bid:  bid,
		Ask:  ask,
		Time: time.Now(),
	}
	return &Generator{
		LastCourse: &initCourse,
	}
}

// GenerateCourse Generate new course
func (g *Generator) GenerateCourse() {
	diff := rand.Int63n(10)
	if rand.Intn(10)%2 == 0 && g.LastCourse.Bid-diff >= 0 {
		diff *= -1
	}
	g.LastCourse.Ask += diff
	g.LastCourse.Bid += diff
}
