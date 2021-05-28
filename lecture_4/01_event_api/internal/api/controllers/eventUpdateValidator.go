package controllers

import "github.com/superbet-group/code-cadets-2021/lecture_3/03_project/controller/internal/api/controllers/models"

// EventUpdateValidator validates event update requests.
type EventUpdateValidator interface {
	EventUpdateIsValid(eventUpdateRequestDto models.EventUpdateRequestDto) bool
}
