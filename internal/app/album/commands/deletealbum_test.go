package commands

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/ydimitriou/go-clean-architecture-example/internal/domain/album"
	"testing"
)

func TestNewDeleteAlbumHandler(t *testing.T) {
	type args struct {
		repo album.Repository
	}
	tests := []struct {
		name   string
		args   args
		expRes DeleteAlbumHandler
	}{
		{
			name:   "should return a new delete album handler",
			args:   args{repo: &album.MockRepository{}},
			expRes: deleteAlbumHandler{repo: &album.MockRepository{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewDeleteAlbumHandler(tt.args.repo)
			assert.Equal(t, tt.expRes, h)
		})
	}
}

func TestDeleteAlbumHandler_Handle(t *testing.T) {
	mockUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	type fields struct {
		repo album.Repository
	}
	type args struct {
		req DeleteAlbumRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		expErr error
	}{
		{
			name: "should return nil when delete is successful",
			fields: fields{
				repo: func() *album.MockRepository {
					mr := album.MockRepository{}
					getRes := album.Album{ID: mockUUID, Title: "foo"}
					mr.On("GetByID", mockUUID).Return(&getRes, nil)
					mr.On("Delete", mockUUID).Return(nil)
					return &mr
				}(),
			},
			args:   args{req: DeleteAlbumRequest{ID: mockUUID}},
			expErr: nil,
		},
		{
			name: "should return error when get by id fails",
			fields: fields{
				repo: func() *album.MockRepository {
					mr := album.MockRepository{}
					err := fmt.Errorf("failed to get album with id %v", mockUUID.String())
					mr.On("GetByID", mockUUID).Return((*album.Album)(nil), err)
					return &mr
				}(),
			},
			args:   args{req: DeleteAlbumRequest{ID: mockUUID}},
			expErr: fmt.Errorf("failed to get album with id %v", mockUUID.String()),
		},
		{
			name: "should return error when get by id returns nil album",
			fields: fields{
				repo: func() *album.MockRepository {
					mr := album.MockRepository{}
					mr.On("GetByID", mockUUID).Return((*album.Album)(nil), nil)
					return &mr
				}(),
			},
			args:   args{req: DeleteAlbumRequest{ID: mockUUID}},
			expErr: fmt.Errorf("album with id %v does not exist", mockUUID.String()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := deleteAlbumHandler{repo: tt.fields.repo}
			err := h.Handle(tt.args.req)
			assert.Equal(t, tt.expErr, err)
		})
	}
}
