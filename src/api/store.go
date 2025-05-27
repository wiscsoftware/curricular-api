package api

import (
	"curricular-api/model"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/rs/xid"
	"math/rand"
	"time"
)

type DataStore struct {
	students []model.Student
	teachers []model.Teacher
	classes  []model.Class
	courses  []model.Course
}

func NewDataStore(size int) DataStore {

	ds := DataStore{}

	var students []model.Student
	var teachers []model.Teacher
	var classes []model.Class
	var courses []model.Course

	idNames := map[int]string{0: "campusId", 1: "emplId", 2: "libraryId"}
	terms := []model.Term{model.Spring, model.Summer, model.Fall}

	residency := []model.ResidencyStatus{model.Resident, model.NonResident, model.Undermined}

	for i := 0; i < size; i++ {

		var studentIdentities []model.Identity
		for _, idName := range idNames {
			guid := xid.New()
			studentIdentities = append(studentIdentities, model.Identity{
				Name:  idName,
				Value: guid.String(),
			})
		}

		var studentAddresses []model.Address
		studentAddress1 := gofakeit.Address()
		studentAddress2 := gofakeit.Address()
		studentAddresses = append(studentAddresses,
			model.Address{
				AddressLine1: studentAddress1.Street,
				City:         studentAddress1.City,
				State:        studentAddress1.State,
				Country:      studentAddress1.Country,
				ZipCode:      studentAddress1.Zip,
			}, model.Address{
				AddressLine1: studentAddress2.Street,
				City:         studentAddress2.City,
				State:        studentAddress2.State,
				Country:      studentAddress2.Country,
				ZipCode:      studentAddress2.Zip,
			})

		dob := gofakeit.Date()
		students = append(students, model.Student{
			Identities: studentIdentities,
			Addresses:  studentAddresses,
			FirstName:  gofakeit.FirstName(),
			LastName:   gofakeit.LastName(),
			Birthdate:  dob,
			AgeInYears: time.Now().Year() - dob.Year(),
			Residency:  residency[rand.Intn(len(residency))],
		})

		var teacherIdentities []model.Identity
		for _, idName := range idNames {
			guid := xid.New()
			teacherIdentities = append(teacherIdentities, model.Identity{
				Name:  idName,
				Value: guid.String(),
			})
		}

		teachers = append(teachers, model.Teacher{
			Identities: teacherIdentities,
			FirstName:  gofakeit.FirstName(),
			LastName:   gofakeit.LastName(),
		})

		crs := allCourses()
		cls := crs[rand.Intn(len(crs))]
		classAddress := gofakeit.Address()
		classes = append(classes, model.Class{
			Id:     xid.New().String(),
			Name:   cls.name,
			Credit: model.NewCredit(cls.credit),
			Location: model.Address{
				AddressLine1: classAddress.Street,
				City:         classAddress.City,
				State:        classAddress.State,
				Country:      classAddress.Country,
				ZipCode:      classAddress.Zip,
			},
			Time: gofakeit.FutureDate(),
		})

		courses = append(courses, model.Course{
			Id:       xid.New().String(),
			Name:     cls.name,
			Term:     terms[rand.Intn(len(terms))],
			Credit:   model.NewCredit(cls.credit),
			Teachers: teachers[:rand.Intn(len(teachers))],
		})

	}
	ds.students = students
	ds.teachers = teachers
	ds.classes = classes
	ds.courses = courses

	return ds
}

func (d DataStore) Students() []model.Student {
	return d.students
}

func (d DataStore) Teachers() []model.Teacher {
	return d.teachers
}

func (d DataStore) Classes() []model.Class {
	return d.classes
}

func (d DataStore) Courses() []model.Course {
	return d.courses
}

type course struct {
	name   string
	credit float64
}

type subject struct {
	name    string
	courses []course
}

func allCourses() []course {

	subjects := map[int]subject{
		0: {
			name: "Physics",
			courses: []course{
				{
					name:   "PHYSICS 103 — GENERAL PHYSICS",
					credit: 4.0,
				},
				{
					name:   "PHYSICS 104 — GENERAL PHYSICS",
					credit: 4.0,
				},
				{
					name:   "PHYSICS 106 — PHYSICS OF SPORTS",
					credit: 3.0,
				},
				{
					name:   "PHYSICS 120 — SPECIAL TOPICS IN PHYSICS",
					credit: 3.0,
				},
				{
					name:   "PHYSICS 201 — GENERAL PHYSICS",
					credit: 5.0,
				},
			},
		},
		1: {
			name: "Mathematics",
			courses: []course{
				{
					name:   "MATH 96 — PREPARATORY ALGEBRA",
					credit: 3.0,
				},
				{
					name:   "MATH 112 — ALGEBRA",
					credit: 3.0,
				},
				{
					name:   "MATH 113 — TRIGONOMETRY",
					credit: 3.0,
				},
				{
					name:   "MATH 114 — ALGEBRA AND TRIGONOMETRY",
					credit: 5.0,
				},
				{
					name:   "MATH 211 — SURVEY OF CALCULUS 1",
					credit: 4.0,
				},
				{
					name:   "MATH 321 — APPLIED MATHEMATICAL ANALYSIS",
					credit: 3.0,
				},
			},
		},
		3: {
			name: "Music",
			courses: []course{
				{
					name:   "MUSIC 34 — STUDY ABROAD: MUSIC PERFORMANCE ENSEMBLE",
					credit: 1.0,
				},
				{
					name:   "CONCERT BAND",
					credit: 1.0,
				},
				{
					name:   "THE SYMPHONY",
					credit: 3.0,
				},
				{
					name:   "BASIC CONCEPTS OF MUSIC 2",
					credit: 3.0,
				},
				{
					name:   "MUSIC 240 — INTERPLAY BETWEEN MUSIC, ART, AND SOCIETY",
					credit: 3.0,
				},
			},
		},
	}

	var courses []course

	for _, sub := range subjects {
		courses = append(courses, sub.courses...)
	}

	return courses
}
