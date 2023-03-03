package validator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type BookData struct {
	Title  string
	Author string
	Rating int
}

func (b BookData) Validate() error {
	return validation.ValidateStruct(&b,
		// Street cannot be empty, and the length must between 5 and 50
		validation.Field(&b.Title, validation.Required.Error("veuillez entrez  le titre"), validation.Length(5, 50).Error("le nombre de caracters minimum accepte est 5")),
		// City cannot be empty, and the length must between 5 and 50
		validation.Field(&b.Author, validation.Required.Error("veuillez entrer le nom de l'auteur"), validation.Length(5, 50)),
		// State cannot be empty, and must be a string consisting of two letters in upper case
		validation.Field(&b.Rating, validation.Required.Error("veuillez entrer la note")),
	)
}
