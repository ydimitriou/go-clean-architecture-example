package memory

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/ydimitriou/go-clean-architecture-example/internal/domain/album"
	"testing"
)

func TestNewRepository(t *testing.T) {
	tests := []struct {
		name   string
		expRes album.Repository
	}{
		{
			name: "Should return a memory repository",
			expRes: Repository{
				albums: make(map[string]album.Album),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := NewRepository()
			assert.Equal(t, tt.expRes, mr)
		})
	}
}

func TestRepository_Create(t *testing.T) {
	mockUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	type fields struct {
		albums map[string]album.Album
	}
	type args struct {
		album album.Album
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		expErr error
	}{
		{
			name: "Should add album in memory",
			fields: fields{
				albums: make(map[string]album.Album),
			},
			args: args{
				album: album.Album{ID: mockUUID, Title: "Foo"},
			},
			expErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Repository{albums: tt.fields.albums}
			err := r.Create(tt.args.album)
			assert.Equal(t, tt.expErr, err)
			a, _ := r.GetByID(tt.args.album.ID)
			assert.Equal(t, tt.args.album, *a)
		})
	}
}

func TestRepository_Update(t *testing.T) {
	mockUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	type fields struct {
		albums map[string]album.Album
	}
	type args struct {
		album album.Album
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		expErr error
	}{
		{
			name: "should update title on success",
			fields: fields{
				albums: func() map[string]album.Album {
					mp := make(map[string]album.Album)
					mp[mockUUID.String()] = album.Album{ID: mockUUID, Title: "Foo"}
					return mp
				}(),
			},
			args: args{
				album: album.Album{ID: mockUUID, Title: "Bar"},
			},
			expErr: nil,
		},
		{
			name: "should return error when ID does not exist",
			fields: fields{
				albums: make(map[string]album.Album),
			},
			args: args{
				album: album.Album{ID: mockUUID, Title: "Bar"},
			},
			expErr: fmt.Errorf("update failed, id %v does not exist", mockUUID.String()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := Repository{albums: tt.fields.albums}
			err := mr.Update(tt.args.album)
			assert.Equal(t, tt.expErr, err)
			if err == nil {
				a, _ := mr.GetByID(tt.args.album.ID)
				assert.Equal(t, tt.args.album, *a)
			}
		})
	}
}

func TestRepository_Delete(t *testing.T) {
	mockUUID1 := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	mockUUID2 := uuid.MustParse("ee3fb3e1-d120-4300-85ae-ac6d67deefb7")
	type fields struct {
		albums map[string]album.Album
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		expErr error
	}{
		{
			name: "should delete album from memory when id exists",
			fields: fields{
				albums: func() map[string]album.Album {
					mp := make(map[string]album.Album)
					mp[mockUUID1.String()] = album.Album{ID: mockUUID1, Title: "Foo"}
					return mp
				}(),
			},
			args:   args{id: mockUUID1},
			expErr: nil,
		},
		{
			name: "should return error when id does not exist",
			fields: fields{
				albums: func() map[string]album.Album {
					mp := make(map[string]album.Album)
					mp[mockUUID2.String()] = album.Album{ID: mockUUID2, Title: "Bar"}
					return mp
				}(),
			},
			args:   args{id: mockUUID1},
			expErr: fmt.Errorf("album with id %v not found", mockUUID1.String()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := Repository{albums: tt.fields.albums}
			err := mr.Delete(tt.args.id)
			assert.Equal(t, tt.expErr, err)
			if err == nil {
				assert.Equal(t, len(tt.fields.albums), 0)
			}
		})
	}
}

func TestRepository_GetByID(t *testing.T) {
	mockUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	type fields struct {
		albums map[string]album.Album
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		expErr error
	}{
		{
			name: "should return proper album when id exists in memory",
			fields: fields{
				albums: func() map[string]album.Album {
					mp := make(map[string]album.Album)
					mp[mockUUID.String()] = album.Album{ID: mockUUID, Title: "Foo"}
					return mp
				}(),
			},
			args:   args{id: mockUUID},
			expErr: nil,
		},
		{
			name:   "should return error when id does not exist in memory",
			fields: fields{albums: make(map[string]album.Album)},
			args:   args{id: mockUUID},
			expErr: fmt.Errorf("album with id %v not found", mockUUID.String()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := Repository{albums: tt.fields.albums}
			a, err := mr.GetByID(tt.args.id)
			assert.Equal(t, tt.expErr, err)
			if err == nil {
				assert.Equal(t, tt.fields.albums[tt.args.id.String()], *a)
			}
		})
	}
}

func TestRepository_GetAll(t *testing.T) {
	mockUUID1 := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	mockUUID2 := uuid.MustParse("ee3fb3e1-d120-4300-85ae-ac6d67deefb7")
	type fields struct {
		albums map[string]album.Album
	}

	tests := []struct {
		name   string
		fields fields
		expRes []album.Album
		expErr error
	}{
		{
			name:   "should return 0 albums",
			fields: fields{albums: make(map[string]album.Album)},
			expRes: ([]album.Album)(nil),
			expErr: nil,
		},
		{
			name: "should return 2 albums",
			fields: fields{
				albums: func() map[string]album.Album {
					mp := make(map[string]album.Album)
					mp[mockUUID1.String()] = album.Album{ID: mockUUID1, Title: "Foo"}
					mp[mockUUID2.String()] = album.Album{ID: mockUUID2, Title: "Bar"}
					return mp
				}(),
			},
			expRes: []album.Album{{ID: mockUUID1, Title: "Foo"}, {ID: mockUUID2, Title: "Bar"}},
			expErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := Repository{albums: tt.fields.albums}
			a, err := mr.GetAll()
			assert.Equal(t, tt.expErr, err)
			assert.Equal(t, tt.expRes, a)
		})
	}
}
