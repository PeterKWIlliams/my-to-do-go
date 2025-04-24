package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/PeterKWIlliams/my-to-do-go/internal/config"
	"github.com/PeterKWIlliams/my-to-do-go/internal/database"
	"github.com/PeterKWIlliams/my-to-do-go/utils"
)

func TestFormatProjectData(t *testing.T) {
	cfg := &config.Config{}
	svc := NewService(cfg)

	tests := []struct {
		name     string
		input    map[string]string
		args     []string
		op       ProjectOperation
		existing *database.Project
		want     *database.Project
		wantErr  bool
	}{
		{
			name: "valid create project",
			input: map[string]string{
				"directory":      "/test/dir",
				"targetDuration": "2h30m",
			},
			args: []string{"test-project"},
			op:   CreateOperation,
			want: &database.Project{
				Name:             "test-project",
				WorkingDirectory: utils.ToNullString(ptr("/test/dir")),
				TargetDuration:   utils.ToNullInt64(ptr(int64(9000))),
			},

			wantErr: false,
		},

		{
			name: "valid update project",
			input: map[string]string{
				"directory":      "/new/dir",
				"targetDuration": "2h30m",
			},
			args: []string{},
			op:   UpdateOperation,
			existing: &database.Project{
				ID:               "this is my id",
				Name:             "test-project",
				TargetDuration:   utils.ToNullInt64(ptr(int64(3000))),
				WorkingDirectory: utils.ToNullString(ptr("/old/dir")),
			},
			want: &database.Project{
				Name:             "test-project",
				WorkingDirectory: utils.ToNullString(ptr("/new/dir")),
				TargetDuration:   utils.ToNullInt64(ptr(int64(9000))),
			},

			wantErr: false,
		},
		{
			name: "invalid create project",
			input: map[string]string{
				"directory":      "/new/dir",
				"targetDuration": "2h30m",
			},
			args: []string{},
			op:   CreateOperation,
			existing: &database.Project{
				ID:               "this is my id",
				TargetDuration:   utils.ToNullInt64(ptr(int64(3000))),
				WorkingDirectory: utils.ToNullString(ptr("/old/dir")),
			},
			want: &database.Project{
				Name:             "test-project",
				WorkingDirectory: utils.ToNullString(ptr("/new/dir")),
				TargetDuration:   utils.ToNullInt64(ptr(int64(9000))),
			},

			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := svc.FormatProjectData(tt.input, tt.args, tt.op, tt.existing)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want.Name, got.Name)
			assert.Equal(t, tt.want.WorkingDirectory, got.WorkingDirectory)
			assert.Equal(t, tt.want.TargetDuration, got.TargetDuration)
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}
