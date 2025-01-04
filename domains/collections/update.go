package collections

import (
	"context"
	"errors"
	"strings"

	"github.com/bluele/go-timecop"
	"github.com/karngyan/maek/db"
)

type UpdateCollectionRequest struct {
	ID          int64
	Name        string
	Description string
	WorkspaceID int64
	UpdatedByID int64
}

func UpdateCollection(ctx context.Context, req *UpdateCollectionRequest) error {
	if req.ID == 0 {
		return errors.New("id is required")
	}

	if req.WorkspaceID == 0 {
		return errors.New("workspace is required")
	}

	if req.UpdatedByID == 0 {
		return errors.New("updated by is required")
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Description = strings.TrimSpace(req.Description)

	return db.Q.UpdateCollection(ctx, db.UpdateCollectionParams{
		Name:        req.Name,
		Description: req.Description,
		UpdatedByID: req.UpdatedByID,
		Updated:     timecop.Now().Unix(),
		ID:          req.ID,
		WorkspaceID: req.WorkspaceID,
	})
}
