package person

import "time"

/* This file contains the input struct for creating a new person. This is used to pass the data from the
service to the repository layer, where it will be used to create a new person in the database.*/

type CreatePersonInput struct {
	FirstName string
	LastName  *string
	BirthDate *time.Time
	Gender    *string
}
