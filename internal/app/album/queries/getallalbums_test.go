package queries

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/ydimitriou/go-clean-architecture-example/internal/domain/album"
	"testing"
)

func TestNewGetAllAlbumsHandler(t *testing.T) {
	type args struct {
		repo album.Repository
	}
	tests := []struct {
		name   string
		args   args
		expRes GetAllAlbumsHandler
	}{
		{
			name: "should return a new GetAllAlbumsHandler",
			args: args{
				repo: &album.MockRepository{},
			},
			expRes: getAllAlbumsHandler{
				repo: &album.MockRepository{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewGetAllAlbumsHandler(tt.args.repo)
			assert.Equal(t, tt.expRes, h)
		})
	}
}

func TestGetAllAlbumsHandler_Handle(t *testing.T) {
	mockUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	type fields struct {
		repo album.Repository
	}
	tests := []struct {
		name   string
		fields fields
		expRes []GetAllAlbumsResult
		expErr error
	}{
		{
			name: "should return empty results when no albums no error",
			fields: fields{
				repo: func() *album.MockRepository {
					mr := album.MockRepository{}
					mr.On("GetAll").Return([]album.Album{}, nil)
					return &mr
				}(),
			},
			expRes: []GetAllAlbumsResult(nil),
			expErr: nil,
		},
		{
			name: "should return one album when no error and one album exist",
			fields: fields{
				repo: func() *album.MockRepository {
					mr := album.MockRepository{}
					albums := []album.Album{{ID: mockUUID, Title: "Foo"}}
					mr.On("GetAll").Return(albums, nil)
					return &mr
				}(),
			},
			expRes: []GetAllAlbumsResult{{ID: mockUUID, Title: "Foo"}},
			expErr: nil,
		},
		{
			name: "should return error when repository fails",
			fields: fields{
				repo: func() *album.MockRepository {
					mr := album.MockRepository{}
					mr.On("GetAll").Return([]album.Album{}, fmt.Errorf("failed to retrieve albums from memory"))
					return &mr
				}(),
			},
			expRes: []GetAllAlbumsResult(nil),
			expErr: fmt.Errorf("failed to retrieve albums from memory"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := getAllAlbumsHandler{tt.fields.repo}
			res, err := h.Handle()
			assert.Equal(t, tt.expErr, err)
			assert.Equal(t, tt.expRes, res)
		})
	}
}
