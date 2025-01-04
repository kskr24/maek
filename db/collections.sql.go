// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: collections.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addNoteToCollection = `-- name: AddNoteToCollection :exec
INSERT INTO collection_notes (collection_id, note_id)
VALUES ($1, $2)
`

type AddNoteToCollectionParams struct {
	CollectionID int64
	NoteID       int64
}

func (q *Queries) AddNoteToCollection(ctx context.Context, arg AddNoteToCollectionParams) error {
	_, err := q.db.Exec(ctx, addNoteToCollection, arg.CollectionID, arg.NoteID)
	return err
}

const deleteCollection = `-- name: DeleteCollection :exec
UPDATE collection
SET deleted = $1,
    updated_by_id = $2,
    updated = $3
WHERE id = $4
  AND workspace_id = $5
`

type DeleteCollectionParams struct {
	Deleted     bool
	UpdatedByID int64
	Updated     int64
	ID          int64
	WorkspaceID int64
}

func (q *Queries) DeleteCollection(ctx context.Context, arg DeleteCollectionParams) error {
	_, err := q.db.Exec(ctx, deleteCollection,
		arg.Deleted,
		arg.UpdatedByID,
		arg.Updated,
		arg.ID,
		arg.WorkspaceID,
	)
	return err
}

const getCollectionByIDAndWorkspace = `-- name: GetCollectionByIDAndWorkspace :one
SELECT id, name, description, created, updated, trashed, deleted, workspace_id, created_by_id, updated_by_id
FROM collection
WHERE id = $1
  AND workspace_id = $2
`

type GetCollectionByIDAndWorkspaceParams struct {
	ID          int64
	WorkspaceID int64
}

func (q *Queries) GetCollectionByIDAndWorkspace(ctx context.Context, arg GetCollectionByIDAndWorkspaceParams) (Collection, error) {
	row := q.db.QueryRow(ctx, getCollectionByIDAndWorkspace, arg.ID, arg.WorkspaceID)
	var i Collection
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Created,
		&i.Updated,
		&i.Trashed,
		&i.Deleted,
		&i.WorkspaceID,
		&i.CreatedByID,
		&i.UpdatedByID,
	)
	return i, err
}

const getCollectionsNameAsc = `-- name: GetCollectionsNameAsc :many
SELECT id, name, description, created, updated, trashed, deleted, workspace_id, created_by_id, updated_by_id
FROM collection
WHERE workspace_id = $1
  AND trashed = $2
  AND deleted = FALSE
  AND (
    name > $4
      OR name = $4 AND id > $5
  )
ORDER BY name ASC,
         id ASC
LIMIT $3
`

type GetCollectionsNameAscParams struct {
	WorkspaceID      int64
	Trashed          bool
	Limit            int64
	LastSortValue    string
	LastCollectionID int64
}

func (q *Queries) GetCollectionsNameAsc(ctx context.Context, arg GetCollectionsNameAscParams) ([]Collection, error) {
	rows, err := q.db.Query(ctx, getCollectionsNameAsc,
		arg.WorkspaceID,
		arg.Trashed,
		arg.Limit,
		arg.LastSortValue,
		arg.LastCollectionID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Collection
	for rows.Next() {
		var i Collection
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Created,
			&i.Updated,
			&i.Trashed,
			&i.Deleted,
			&i.WorkspaceID,
			&i.CreatedByID,
			&i.UpdatedByID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCollectionsNameDesc = `-- name: GetCollectionsNameDesc :many
SELECT id, name, description, created, updated, trashed, deleted, workspace_id, created_by_id, updated_by_id
FROM collection
WHERE workspace_id = $1
  AND trashed = $2
  AND deleted = FALSE
  AND (
    name < $4
      OR name = $4 AND id < $5
  )
ORDER BY name DESC,
         id DESC
LIMIT $3
`

type GetCollectionsNameDescParams struct {
	WorkspaceID      int64
	Trashed          bool
	Limit            int64
	LastSortValue    string
	LastCollectionID int64
}

func (q *Queries) GetCollectionsNameDesc(ctx context.Context, arg GetCollectionsNameDescParams) ([]Collection, error) {
	rows, err := q.db.Query(ctx, getCollectionsNameDesc,
		arg.WorkspaceID,
		arg.Trashed,
		arg.Limit,
		arg.LastSortValue,
		arg.LastCollectionID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Collection
	for rows.Next() {
		var i Collection
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Created,
			&i.Updated,
			&i.Trashed,
			&i.Deleted,
			&i.WorkspaceID,
			&i.CreatedByID,
			&i.UpdatedByID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCollectionsUpdatedAsc = `-- name: GetCollectionsUpdatedAsc :many
SELECT id, name, description, created, updated, trashed, deleted, workspace_id, created_by_id, updated_by_id
FROM collection
WHERE workspace_id = $1
  AND trashed = $2
  AND deleted = FALSE
  AND (
    updated > $4
      OR updated = $4 AND id > $5
  )
ORDER BY updated ASC,
         id ASC
LIMIT $3
`

type GetCollectionsUpdatedAscParams struct {
	WorkspaceID      int64
	Trashed          bool
	Limit            int64
	LastSortValue    int64
	LastCollectionID int64
}

func (q *Queries) GetCollectionsUpdatedAsc(ctx context.Context, arg GetCollectionsUpdatedAscParams) ([]Collection, error) {
	rows, err := q.db.Query(ctx, getCollectionsUpdatedAsc,
		arg.WorkspaceID,
		arg.Trashed,
		arg.Limit,
		arg.LastSortValue,
		arg.LastCollectionID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Collection
	for rows.Next() {
		var i Collection
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Created,
			&i.Updated,
			&i.Trashed,
			&i.Deleted,
			&i.WorkspaceID,
			&i.CreatedByID,
			&i.UpdatedByID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCollectionsUpdatedDesc = `-- name: GetCollectionsUpdatedDesc :many
SELECT id, name, description, created, updated, trashed, deleted, workspace_id, created_by_id, updated_by_id
FROM collection
WHERE workspace_id = $1
  AND trashed = $2
  AND deleted = FALSE
  AND (
    updated < $4
      OR updated = $4 AND id < $5
  )
ORDER BY updated DESC,
         id DESC
LIMIT $3
`

type GetCollectionsUpdatedDescParams struct {
	WorkspaceID      int64
	Trashed          bool
	Limit            int64
	LastSortValue    int64
	LastCollectionID int64
}

func (q *Queries) GetCollectionsUpdatedDesc(ctx context.Context, arg GetCollectionsUpdatedDescParams) ([]Collection, error) {
	rows, err := q.db.Query(ctx, getCollectionsUpdatedDesc,
		arg.WorkspaceID,
		arg.Trashed,
		arg.Limit,
		arg.LastSortValue,
		arg.LastCollectionID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Collection
	for rows.Next() {
		var i Collection
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Created,
			&i.Updated,
			&i.Trashed,
			&i.Deleted,
			&i.WorkspaceID,
			&i.CreatedByID,
			&i.UpdatedByID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInitialCollectionsNameAsc = `-- name: GetInitialCollectionsNameAsc :many
SELECT id, name, description, created, updated, trashed, deleted, workspace_id, created_by_id, updated_by_id
FROM collection
WHERE workspace_id = $1
  AND trashed = $2
  AND deleted = FALSE
ORDER BY name ASC,
         id ASC 
LIMIT $3
`

type GetInitialCollectionsNameAscParams struct {
	WorkspaceID int64
	Trashed     bool
	Limit       int64
}

func (q *Queries) GetInitialCollectionsNameAsc(ctx context.Context, arg GetInitialCollectionsNameAscParams) ([]Collection, error) {
	rows, err := q.db.Query(ctx, getInitialCollectionsNameAsc, arg.WorkspaceID, arg.Trashed, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Collection
	for rows.Next() {
		var i Collection
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Created,
			&i.Updated,
			&i.Trashed,
			&i.Deleted,
			&i.WorkspaceID,
			&i.CreatedByID,
			&i.UpdatedByID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInitialCollectionsNameDesc = `-- name: GetInitialCollectionsNameDesc :many
SELECT id, name, description, created, updated, trashed, deleted, workspace_id, created_by_id, updated_by_id
FROM collection
WHERE workspace_id = $1
  AND trashed = $2
  AND deleted = FALSE
ORDER BY name DESC,
         id DESC
LIMIT $3
`

type GetInitialCollectionsNameDescParams struct {
	WorkspaceID int64
	Trashed     bool
	Limit       int64
}

func (q *Queries) GetInitialCollectionsNameDesc(ctx context.Context, arg GetInitialCollectionsNameDescParams) ([]Collection, error) {
	rows, err := q.db.Query(ctx, getInitialCollectionsNameDesc, arg.WorkspaceID, arg.Trashed, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Collection
	for rows.Next() {
		var i Collection
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Created,
			&i.Updated,
			&i.Trashed,
			&i.Deleted,
			&i.WorkspaceID,
			&i.CreatedByID,
			&i.UpdatedByID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInitialCollectionsUpdatedAsc = `-- name: GetInitialCollectionsUpdatedAsc :many
SELECT id, name, description, created, updated, trashed, deleted, workspace_id, created_by_id, updated_by_id
FROM collection
WHERE workspace_id = $1
  AND trashed = $2
  AND deleted = FALSE
ORDER BY updated ASC,
         id ASC
LIMIT $3
`

type GetInitialCollectionsUpdatedAscParams struct {
	WorkspaceID int64
	Trashed     bool
	Limit       int64
}

func (q *Queries) GetInitialCollectionsUpdatedAsc(ctx context.Context, arg GetInitialCollectionsUpdatedAscParams) ([]Collection, error) {
	rows, err := q.db.Query(ctx, getInitialCollectionsUpdatedAsc, arg.WorkspaceID, arg.Trashed, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Collection
	for rows.Next() {
		var i Collection
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Created,
			&i.Updated,
			&i.Trashed,
			&i.Deleted,
			&i.WorkspaceID,
			&i.CreatedByID,
			&i.UpdatedByID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInitialCollectionsUpdatedDesc = `-- name: GetInitialCollectionsUpdatedDesc :many
SELECT id, name, description, created, updated, trashed, deleted, workspace_id, created_by_id, updated_by_id
FROM collection
WHERE workspace_id = $1
  AND trashed = $2
  AND deleted = FALSE
ORDER BY updated DESC,
         id DESC
LIMIT $3
`

type GetInitialCollectionsUpdatedDescParams struct {
	WorkspaceID int64
	Trashed     bool
	Limit       int64
}

func (q *Queries) GetInitialCollectionsUpdatedDesc(ctx context.Context, arg GetInitialCollectionsUpdatedDescParams) ([]Collection, error) {
	rows, err := q.db.Query(ctx, getInitialCollectionsUpdatedDesc, arg.WorkspaceID, arg.Trashed, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Collection
	for rows.Next() {
		var i Collection
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Created,
			&i.Updated,
			&i.Trashed,
			&i.Deleted,
			&i.WorkspaceID,
			&i.CreatedByID,
			&i.UpdatedByID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertCollection = `-- name: InsertCollection :one
INSERT INTO collection (name, description, created, updated, trashed, deleted, workspace_id, created_by_id, updated_by_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id
`

type InsertCollectionParams struct {
	Name        string
	Description string
	Created     int64
	Updated     int64
	Trashed     bool
	Deleted     bool
	WorkspaceID int64
	CreatedByID int64
	UpdatedByID int64
}

func (q *Queries) InsertCollection(ctx context.Context, arg InsertCollectionParams) (int64, error) {
	row := q.db.QueryRow(ctx, insertCollection,
		arg.Name,
		arg.Description,
		arg.Created,
		arg.Updated,
		arg.Trashed,
		arg.Deleted,
		arg.WorkspaceID,
		arg.CreatedByID,
		arg.UpdatedByID,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const listCollections = `-- name: ListCollections :many
SELECT id, name, description, created, updated, trashed, deleted,
       workspace_id, created_by_id, updated_by_id
FROM collection
WHERE workspace_id = $1
  AND deleted = false
  AND (
    -- Sort by updated timestamp
    CASE WHEN $3 = 'updated' THEN
      CASE WHEN $4 = 'desc' THEN
        CASE WHEN $5::BIGINT > 0 THEN
          (updated, id) < ($5::BIGINT, $6::BIGINT)
        ELSE TRUE END
      ELSE
        CASE WHEN $5::BIGINT > 0 THEN
          (updated, id) > ($5::BIGINT, $6::BIGINT)
        ELSE TRUE END
      END
    -- Sort by name (string)
    WHEN $3 = 'name' THEN
      CASE WHEN $4 = 'desc' THEN
        CASE WHEN $7::TEXT != '' THEN
          (name, id) < ($7::TEXT, $6::BIGINT)
        ELSE TRUE END
      ELSE
        CASE WHEN $7::TEXT != '' THEN
          (name, id) > ($7::TEXT, $6::BIGINT)
        ELSE TRUE END
      END
    ELSE TRUE
    END
  )
ORDER BY
  CASE
    WHEN $3 = 'updated' AND $4 = 'DESC' THEN updated END DESC,
  CASE
    WHEN $3 = 'updated' AND $4 = 'ASC' THEN updated END ASC,
  CASE
    WHEN $3 = 'name' AND $4 = 'DESC' THEN name END DESC,
  CASE
    WHEN $3 = 'name' AND $4 = 'ASC' THEN name END ASC,
  id DESC -- Secondary sort by ID ensures stable ordering
LIMIT $2
`

type ListCollectionsParams struct {
	WorkspaceID   int64
	Limit         int64
	SortBy        pgtype.Text
	SortOrder     pgtype.Text
	CursorUpdated int64
	CursorID      int64
	CursorName    string
}

func (q *Queries) ListCollections(ctx context.Context, arg ListCollectionsParams) ([]Collection, error) {
	rows, err := q.db.Query(ctx, listCollections,
		arg.WorkspaceID,
		arg.Limit,
		arg.SortBy,
		arg.SortOrder,
		arg.CursorUpdated,
		arg.CursorID,
		arg.CursorName,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Collection
	for rows.Next() {
		var i Collection
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Created,
			&i.Updated,
			&i.Trashed,
			&i.Deleted,
			&i.WorkspaceID,
			&i.CreatedByID,
			&i.UpdatedByID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const trashCollection = `-- name: TrashCollection :exec
UPDATE collection
SET trashed = $1,
    updated_by_id = $2,
    updated = $3
WHERE id = $4
  AND workspace_id = $5
`

type TrashCollectionParams struct {
	Trashed     bool
	UpdatedByID int64
	Updated     int64
	ID          int64
	WorkspaceID int64
}

func (q *Queries) TrashCollection(ctx context.Context, arg TrashCollectionParams) error {
	_, err := q.db.Exec(ctx, trashCollection,
		arg.Trashed,
		arg.UpdatedByID,
		arg.Updated,
		arg.ID,
		arg.WorkspaceID,
	)
	return err
}

const updateCollection = `-- name: UpdateCollection :exec
UPDATE collection
SET name          = $1,
    description   = $2,
    updated_by_id = $3,
    updated       = $4
WHERE id = $5
  AND workspace_id = $6
`

type UpdateCollectionParams struct {
	Name        string
	Description string
	UpdatedByID int64
	Updated     int64
	ID          int64
	WorkspaceID int64
}

func (q *Queries) UpdateCollection(ctx context.Context, arg UpdateCollectionParams) error {
	_, err := q.db.Exec(ctx, updateCollection,
		arg.Name,
		arg.Description,
		arg.UpdatedByID,
		arg.Updated,
		arg.ID,
		arg.WorkspaceID,
	)
	return err
}
