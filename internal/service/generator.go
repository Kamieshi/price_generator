// Package service Business logic generator
package service

import (
	"math/rand"
	"time"

	"github.com/Kamieshi/price_generator/internal/models"
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
		Time: time.Now().Format("2006-01-02T15:04:05.000TZ-07:00"),
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
	g.LastCourse.Time = time.Now().Format("2006-01-02T15:04:05.000TZ-07:00")
	g.LastCourse.Ask += diff
	g.LastCourse.Bid += diff
}
func (g *Generator) GenerateAddCourse() {
	diff := rand.Int63n(10)
	g.LastCourse.Time = time.Now().Format("2006-01-02T15:04:05.000TZ-07:00")
	g.LastCourse.Ask += diff
	g.LastCourse.Bid += diff
}
