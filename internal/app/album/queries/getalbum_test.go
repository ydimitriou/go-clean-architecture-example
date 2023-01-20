package queries

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/ydimitriou/go-clean-architecture-example/internal/domain/album"
	"testing"
)

func TestNewGetAlbumHandler(t *testing.T) {
	type args struct {
		repo album.Repository
	}
	tests := []struct {
		name   string
		args   args
		expRes GetAlbumHandler
	}{
		{
			name:   "should return a new GetAlbumHandler",
			args:   args{repo: &album.MockRepository{}},
			expRes: getAlbumHandler{repo: &album.MockRepository{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewGetAlbumHandler(tt.args.repo)
			assert.Equal(t, tt.expRes, h)
		})
	}
}

func TestGetAlbumHandler_Handle(t *testing.T) {
	mockUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	type fields struct {
		repo album.Repository
	}
	type args struct {
		albumRequest GetAlbumRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		expRes *GetAlbumResult
		expErr error
	}{
		{
			name: "should return album when id exists in memory",
			fields: fields{
				repo: func() *album.MockRepository {
					mr := album.MockRepository{}
					mockAlbum := album.Album{
						ID:     mockUUID,
						Title:  "Foo",
						Artist: "Bar",
					}
					mr.On("GetByID", mockUUID).Return(&mockAlbum, nil)
					return &mr
				}(),
			},
			args:   args{albumRequest: GetAlbumRequest{ID: mockUUID}},
			expRes: &GetAlbumResult{ID: mockUUID, Title: "Foo", Artist: "Bar"},
			expErr: nil,
		},
		{
			name: "should return error when memory responds with error",
			fields: fields{
				repo: func() *album.MockRepository {
					mr := album.MockRepository{}
					err := fmt.Errorf("album not found")
					mr.On("GetByID", mockUUID).Return(&album.Album{}, err)
					return &mr
				}(),
			},
			args:   args{albumRequest: GetAlbumRequest{ID: mockUUID}},
			expRes: nil,
			expErr: fmt.Errorf("album not found"),
		},
		{
			name: "should return nil results and nil error when no error and no album exist",
			fields: fields{
				repo: func() *album.MockRepository {
					mr := album.MockRepository{}
					mr.On("GetByID", mockUUID).Return((*album.Album)(nil), nil)
					return &mr
				}(),
			},
			args:   args{albumRequest: GetAlbumRequest{ID: mockUUID}},
			expRes: nil,
			expErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := getAlbumHandler{tt.fields.repo}
			res, err := h.Handle(tt.args.albumRequest)
			assert.Equal(t, tt.expErr, err)
			assert.Equal(t, tt.expRes, res)
		})
	}
}
