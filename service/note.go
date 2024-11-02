package service

import (
	"github.com/depermana12/go-notes/db"
	"github.com/depermana12/go-notes/models"
)

func AddNote(authorId uint, note *models.Note) (*models.Note, error) {
	note.AuthorId = uint(authorId)

	if err := db.GetDB().Create(note).Error; err != nil {
		return nil, err
	}
	return note, nil
}
