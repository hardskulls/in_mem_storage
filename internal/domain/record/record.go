package record

import "time"

// CreatedAt represents the time when a record was added to the database.
type CreatedAt = time.Time

// Author of a record in the database.
type Author = string

// ID is a key used to identify a record in the database.
type ID = string

func New[D any](data D, author Author) Record[D] {
	return Record[D]{
		data: data,
		metaData: MetaData{
			author:    author,
			createdAt: time.Now(),
		},
	}
}

func (r *Record[D]) Update(data D) {
	r.data = data
}

// Record is a value object representing a record in the database.
type Record[D any] struct {
	data     D
	metaData MetaData
}

func (r Record[D]) Data() D {
	return r.data
}

// MetaData is a value object representing the metadata of a record in the database.
type MetaData struct {
	author    Author
	createdAt CreatedAt
}
