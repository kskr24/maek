-- name: InsertCollection :one
INSERT INTO collection (name, description, created, updated, trashed, deleted, workspace_id, created_by_id, updated_by_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id;

-- name: AddNoteToCollection :exec
INSERT INTO collection_notes (collection_id, note_id)
VALUES ($1, $2);

-- name: GetCollectionByIDAndWorkspace :one
SELECT id, name, description, created, updated, trashed, deleted, workspace_id, created_by_id, updated_by_id
FROM collection
WHERE id = $1
  AND workspace_id = $2;

-- name: UpdateCollection :exec
UPDATE collection
SET name          = $1,
    description   = $2,
    updated_by_id = $3,
    updated       = $4
WHERE id = $5
  AND workspace_id = $6;

-- name: TrashCollection :exec
UPDATE collection
SET trashed = $1,
    updated_by_id = $2,
    updated = $3
WHERE id = $4
  AND workspace_id = $5;

-- name: DeleteCollection :exec
UPDATE collection
SET deleted = $1,
    updated_by_id = $2,
    updated = $3
WHERE id = $4
  AND workspace_id = $5;

-- name: GetInitialCollectionsUpdatedAsc :many
SELECT id, name, description, created, updated, trashed, deleted, workspace_id, created_by_id, updated_by_id
FROM collection
WHERE workspace_id = $1
  AND trashed = $2
  AND deleted = FALSE
ORDER BY updated ASC,
         id ASC
LIMIT $3;

-- name: GetCollectionsUpdatedAsc :many
SELECT id, name, description, created, updated, trashed, deleted, workspace_id, created_by_id, updated_by_id
FROM collection
WHERE workspace_id = $1
  AND trashed = $2
  AND deleted = FALSE
  AND (
    updated > @last_sort_value
      OR updated = @last_sort_value AND id > @last_collection_id
  )
ORDER BY updated ASC,
         id ASC
LIMIT $3;


-- name: GetInitialCollectionsUpdatedDesc :many
SELECT id, name, description, created, updated, trashed, deleted, workspace_id, created_by_id, updated_by_id
FROM collection
WHERE workspace_id = $1
  AND trashed = $2
  AND deleted = FALSE
ORDER BY updated DESC,
         id DESC
LIMIT $3;

-- name: GetCollectionsUpdatedDesc :many
SELECT id, name, description, created, updated, trashed, deleted, workspace_id, created_by_id, updated_by_id
FROM collection
WHERE workspace_id = $1
  AND trashed = $2
  AND deleted = FALSE
  AND (
    updated < @last_sort_value
      OR updated = @last_sort_value AND id < @last_collection_id
  )
ORDER BY updated DESC,
         id DESC
LIMIT $3;

-- name: GetInitialCollectionsNameAsc :many
SELECT id, name, description, created, updated, trashed, deleted, workspace_id, created_by_id, updated_by_id
FROM collection
WHERE workspace_id = $1
  AND trashed = $2
  AND deleted = FALSE
ORDER BY name ASC,
         id ASC 
LIMIT $3;

-- name: GetCollectionsNameAsc :many
SELECT id, name, description, created, updated, trashed, deleted, workspace_id, created_by_id, updated_by_id
FROM collection
WHERE workspace_id = $1
  AND trashed = $2
  AND deleted = FALSE
  AND (
    name > @last_sort_value
      OR name = @last_sort_value AND id > @last_collection_id
  )
ORDER BY name ASC,
         id ASC
LIMIT $3;

-- name: GetInitialCollectionsNameDesc :many
SELECT id, name, description, created, updated, trashed, deleted, workspace_id, created_by_id, updated_by_id
FROM collection
WHERE workspace_id = $1
  AND trashed = $2
  AND deleted = FALSE
ORDER BY name DESC,
         id DESC
LIMIT $3;

-- name: GetCollectionsNameDesc :many
SELECT id, name, description, created, updated, trashed, deleted, workspace_id, created_by_id, updated_by_id
FROM collection
WHERE workspace_id = $1
  AND trashed = $2
  AND deleted = FALSE
  AND (
    name < @last_sort_value
      OR name = @last_sort_value AND id < @last_collection_id
  )
ORDER BY name DESC,
         id DESC
LIMIT $3;

-- name: ListCollections :many
SELECT id, name, description, created, updated, trashed, deleted,
       workspace_id, created_by_id, updated_by_id
FROM collection
WHERE workspace_id = $1
  AND deleted = false
  AND (
    -- Sort by updated timestamp
    CASE WHEN @sort_by = 'updated' THEN
      CASE WHEN @sort_order = 'desc' THEN
        CASE WHEN @cursor_updated::BIGINT > 0 THEN
          (updated, id) < (@cursor_updated::BIGINT, @cursor_id::BIGINT)
        ELSE TRUE END
      ELSE
        CASE WHEN @cursor_updated::BIGINT > 0 THEN
          (updated, id) > (@cursor_updated::BIGINT, @cursor_id::BIGINT)
        ELSE TRUE END
      END
    -- Sort by name (string)
    WHEN @sort_by = 'name' THEN
      CASE WHEN @sort_order = 'desc' THEN
        CASE WHEN @cursor_name::TEXT != '' THEN
          (name, id) < (@cursor_name::TEXT, @cursor_id::BIGINT)
        ELSE TRUE END
      ELSE
        CASE WHEN @cursor_name::TEXT != '' THEN
          (name, id) > (@cursor_name::TEXT, @cursor_id::BIGINT)
        ELSE TRUE END
      END
    ELSE TRUE
    END
  )
ORDER BY
  CASE
    WHEN @sort_by = 'updated' AND @sort_order = 'DESC' THEN updated END DESC,
  CASE
    WHEN @sort_by = 'updated' AND @sort_order = 'ASC' THEN updated END ASC,
  CASE
    WHEN @sort_by = 'name' AND @sort_order = 'DESC' THEN name END DESC,
  CASE
    WHEN @sort_by = 'name' AND @sort_order = 'ASC' THEN name END ASC,
  id DESC -- Secondary sort by ID ensures stable ordering
LIMIT $2;

