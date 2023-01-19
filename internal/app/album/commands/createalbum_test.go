package commands

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/ydimitriou/go-clean-architecture-example/internal/domain/album"
	timeMock "github.com/ydimitriou/go-clean-architecture-example/internal/pkg/time"
	uuidMock "github.com/ydimitriou/go-clean-architecture-example/internal/pkg/uuid"
	"testing"
	"time"
)

func TestCreateAlbumHandler_Handle(t *testing.T) {
	mockUUID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")
	mockTime, _ := time.Parse("yyyy-MM-dd", "2022-01-24")
	type fields struct {
		uuidProvider uuidMock.Provider
		timeProvider timeMock.Provider
		repo         album.Repository
	}
	type args struct {
		req CreateAlbumRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		expErr error
	}{
		{
			name: "should not return error when success",
			fields: fields{
				uuidProvider: func() *uuidMock.MockProvider {
					mp := uuidMock.MockProvider{}
					mp.On("NewUUID").Return(mockUUID)
					return &mp
				}(),
				timeProvider: func() *timeMock.MockProvider {
					mp := timeMock.MockProvider{}
					mp.On("Now").Return(mockTime)
					return &mp
				}(),
				repo: func() *album.MockRepository {
					a := album.Album{
						ID:          mockUUID,
						Title:       "Foo",
						Artist:      "Bar",
						Price:       6.24,
						Description: "Foo Bar",
						CreatedAt:   mockTime,
					}
					mr := album.MockRepository{}
					mr.On("Create", a).Return(nil)

					return &mr
				}(),
			},
			args: args{
				req: CreateAlbumRequest{
					Title:       "Foo",
					Artist:      "Bar",
					Price:       6.24,
					Description: "Foo Bar",
				},
			},
			expErr: nil,
		},
		{
			name: "should return error when repo returns error",
			fields: fields{
				uuidProvider: func() *uuidMock.MockProvider {
					mp := uuidMock.MockProvider{}
					mp.On("NewUUID").Return(mockUUID)
					return &mp
				}(),
				timeProvider: func() *timeMock.MockProvider {
					tp := timeMock.MockProvider{}
					tp.On("Now").Return(mockTime)
					return &tp
				}(),
				repo: func() *album.MockRepository {
					a := album.Album{
						ID:          mockUUID,
						Title:       "Foo",
						Artist:      "Bar",
						Price:       3.24,
						Description: "Foo Bar",
						CreatedAt:   mockTime,
					}
					mr := album.MockRepository{}
					mr.On("Create", a).Return(errors.New("repo error"))
					return &mr
				}(),
			},
			args: args{
				req: CreateAlbumRequest{
					Title:       "Foo",
					Artist:      "Bar",
					Price:       3.24,
					Description: "Foo Bar",
				},
			},
			expErr: errors.New("repo error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := createAlbumHandler{
				uuidProvider: tt.fields.uuidProvider,
				timeProvider: tt.fields.timeProvider,
				repo:         tt.fields.repo,
			}

			err := h.Handle(tt.args.req)
			assert.Equal(t, tt.expErr, err)
		})
	}
}
