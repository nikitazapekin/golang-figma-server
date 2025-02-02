package models

import (
	"database/sql"
	"errors"
	"time"
)

type Draft struct {
	ID             int
	Name           string
	Description    string
	Likes          int
	CreatedAt      time.Time
	AuthorID       int
	AuthorUsername string
	AuthorEmail    string
}

func CreateDraft(db *sql.DB, name, description string, authorID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO drafts (draft_name, draft_description, draft_author) 
		VALUES ($1, $2, $3)`, name, description, authorID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func GetDraftByID(db *sql.DB, draftID int) (*Draft, error) {
	var draft Draft
	err := db.QueryRow(`
		SELECT draft_id, draft_name, draft_description, likes, created_at, draft_author 
		FROM drafts WHERE draft_id = $1`, draftID).Scan(
		&draft.ID, &draft.Name, &draft.Description, &draft.Likes, &draft.CreatedAt, &draft.AuthorID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("draft not found")
		}
		return nil, err
	}

	return &draft, nil
}


func GetAllDrafts(db *sql.DB) ([]Draft, error) {
	rows, err := db.Query(`
		SELECT 
			d.draft_id, 
			d.draft_name, 
			d.draft_description, 
			d.likes, 
			d.created_at, 
			d.draft_author, 
			u.username AS author_username, 
				u.email AS author_email
		FROM drafts d
		JOIN users u ON d.draft_author = u.id
	`)
 
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drafts []Draft
	for rows.Next() {
		var draft Draft

		err := rows.Scan(&draft.ID, &draft.Name, &draft.Description, &draft.Likes, &draft.CreatedAt, &draft.AuthorID, &draft.AuthorUsername, &draft.AuthorEmail)
		if err != nil {
			return nil, err
		}
		drafts = append(drafts, draft)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return drafts, nil
}

func UpdateDraft(db *sql.DB, draftID int, name, description string) error {
	_, err := db.Exec(`
		UPDATE drafts 
		SET draft_name = $1, draft_description = $2 
		WHERE draft_id = $3`, name, description, draftID)
	return err
}

func DeleteDraft(db *sql.DB, draftID int) error {
	_, err := db.Exec(`
		DELETE FROM drafts WHERE draft_id = $1`, draftID)
	return err
}


func GetPersonalDraftsById(db *sql.DB, userID int) ([]Draft, error) {
	rows, err := db.Query(`
		SELECT 
			d.draft_id, 
			d.draft_name, 
			d.draft_description, 
			d.likes, 
			d.created_at, 
			d.draft_author, 
			u.username AS author_username, 
			u.email AS author_email
		FROM drafts d
		JOIN users u ON d.draft_author = u.id
		WHERE d.draft_author = $1
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drafts []Draft
	for rows.Next() {
		var draft Draft
		err := rows.Scan(
			&draft.ID, &draft.Name, &draft.Description, &draft.Likes, 
			&draft.CreatedAt, &draft.AuthorID, &draft.AuthorUsername, &draft.AuthorEmail,
		)
		if err != nil {
			return nil, err
		}
		drafts = append(drafts, draft)
	}
 
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return drafts, nil
}
