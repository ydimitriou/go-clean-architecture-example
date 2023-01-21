package commands

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/ydimitriou/go-clean-architecture-example/internal/domain/album"
	"testing"
)

func TestNewUpdateAlbumHandler(t *testing.T) {
	type args struct {
		repo album.Repository
	}
	tests := []struct {
		name   string
		args   args
		expRes UpdateAlbumHandler
	}{
		{
			name:   "should return a new update album handler",
			args:   args{repo: &album.MockRepository{}},
			expRes: updateAlbumHandler{repo: &album.MockRepository{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewUpdateAlbumHandler(tt.args.repo)
			assert.Equal(t, tt.expRes, h)
		})
	}
}

func TestUpdateAlbumHandler_Handle(t *testing.T) {
	mockUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	type fields struct {
		repo album.Repository
	}
	type args struct {
		req UpdateAlbumRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		expErr error
	}{
		{
			name: "should return nil when update successful",
			fields: fields{
				repo: func() *album.MockRepository {
					mr := album.MockRepository{}
					getRes := album.Album{ID: mockUUID, Title: "foo"}
					updateReq := album.Album{ID: mockUUID, Title: "update foo"}
					mr.On("GetByID", mockUUID).Return(&getRes, nil)
					mr.On("Update", updateReq).Return(nil)
					return &mr
				}(),
			},
			args: args{
				req: UpdateAlbumRequest{ID: mockUUID, Title: "update foo"},
			},
			expErr: nil,
		},
		{
			name: "should return error when get by id fails",
			fields: fields{
				repo: func() *album.MockRepository {
					mr := album.MockRepository{}
					err := fmt.Errorf("failed to access memory")
					mr.On("GetByID", mockUUID).Return((*album.Album)(nil), err)
					return &mr
				}(),
			},
			args: args{
				req: UpdateAlbumRequest{ID: mockUUID, Title: "update foo"},
			},
			expErr: fmt.Errorf("failed to access memory"),
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
			args: args{
				req: UpdateAlbumRequest{ID: mockUUID, Title: "update foo"},
			},
			expErr: fmt.Errorf("update failed, album with id %v does not exist", mockUUID),
		},
		{
			name: "should return error when update fails",
			fields: fields{
				repo: func() *album.MockRepository {
					mr := album.MockRepository{}
					getRes := album.Album{ID: mockUUID, Title: "foo"}
					updateReq := album.Album{ID: mockUUID, Title: "update foo"}
					err := fmt.Errorf("failed to update album")
					mr.On("GetByID", mockUUID).Return(&getRes, nil)
					mr.On("Update", updateReq).Return(err)
					return &mr
				}(),
			},
			args: args{
				req: UpdateAlbumRequest{ID: mockUUID, Title: "update foo"},
			},
			expErr: fmt.Errorf("failed to update album"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := updateAlbumHandler{tt.fields.repo}
			err := h.Handle(tt.args.req)
			assert.Equal(t, tt.expErr, err)
		})
	}
}
