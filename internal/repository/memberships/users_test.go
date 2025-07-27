package memberships

import (
	"testing"
	"time"

	"proyek3-catalog-music/internal/models/memberships"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_repository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(t, err)

	type args struct {
		user *memberships.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockfn  func(args args)
	}{
		{
			name: "Create User Success",
			args: args{
				user: &memberships.User{
					Email:        "test@example.com",
					Username:     "testuser",
					PasswordHash: "passwordhash",
				},
			},
			wantErr: false,
			mockfn: func(args args) {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO "users"`).
					WithArgs(
						args.user.Email,
						args.user.Username,
						args.user.PasswordHash,
						"",               // created_by (empty string)
						"",               // updated_by (empty string)
						sqlmock.AnyArg(), // created_at
						sqlmock.AnyArg(), // updated_at
						nil,              // deleted_at (null)
					).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				mock.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockfn != nil {
				tt.mockfn(tt.args)
			}
			r := &repository{
				db: gormdb,
			}
			t.Logf("Test Args: %+v\n", tt.args.user)
			err := r.CreateUser(tt.args.user)
			t.Logf("CreateUser error: %v\n", err)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

// Untuk debug error pada unit test seperti ini, kamu bisa lakukan beberapa langkah berikut:
// 1. Tambahkan log/print pada bagian error untuk melihat detail error-nya.
// 2. Pastikan query dan argumen pada sqlmock sesuai dengan yang dieksekusi oleh GORM.
// 3. Cek output dari mock.ExpectationsWereMet() untuk melihat apakah semua ekspektasi terpenuhi.
// 4. Jalankan test dengan flag verbose: `go test -v`.
// 5. Jika error pada query, print query yang dieksekusi dan argumennya.
// 6. Gunakan t.Logf atau fmt.Printf untuk menampilkan variabel penting di dalam test.

func Test_repository_GetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(t, err)

	type args struct {
		email    string
		username string
		id       int
	}
	tests := []struct {
		name    string
		args    args
		want    *memberships.User
		wantErr bool
		mockfn  func(args args)
	}{
		{
			name: "Get User By Email Success",
			args: args{
				email:    "test@example.com",
				username: "",
				id:       0,
			},
			want: &memberships.User{
				Email:        "test@example.com",
				Username:     "testuser",
				PasswordHash: "passwordhash",
				CreatedBy:    "",
				UpdatedBy:    "",
			},
			mockfn: func(args args) {
				rows := sqlmock.NewRows([]string{"id", "email", "username", "password_hash", "created_by", "updated_by", "created_at", "updated_at", "deleted_at"}).
					AddRow(1, "test@example.com", "testuser", "passwordhash", "", "", time.Now(), time.Now(), nil)
				mock.ExpectQuery(`SELECT .* FROM "users" .*`).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(rows)
			},
		},
		{
			name: "Get User By Username Success",
			args: args{
				email:    "",
				username: "testuser",
				id:       0,
			},
			want: &memberships.User{
				Email:        "test@example.com",
				Username:     "testuser",
				PasswordHash: "passwordhash",
				CreatedBy:    "",
				UpdatedBy:    "",
			},
			mockfn: func(args args) {
				rows := sqlmock.NewRows([]string{"id", "email", "username", "password_hash", "created_by", "updated_by", "created_at", "updated_at", "deleted_at"}).
					AddRow(1, "test@example.com", "testuser", "passwordhash", "", "", time.Now(), time.Now(), nil)
				mock.ExpectQuery(`SELECT .* FROM "users" .*`).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(rows)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockfn != nil {
				tt.mockfn(tt.args)
			}
			r := &repository{
				db: gormdb,
			}
			got, err := r.GetUser(tt.args.email, tt.args.username, tt.args.id)
			t.Logf("Test Args: email=%s, username=%s, id=%d", tt.args.email, tt.args.username, tt.args.id)
			t.Logf("\n")
			if got != nil {
				t.Logf("Test GOT: email=%s, username=%s, id=%d", got.Email, got.Username, got.ID)
			} else {
				t.Logf("Test GOT: nil")
			}
			if err != nil {
				t.Logf("Error: %v", err)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				if got != nil && tt.want != nil {
					assert.Equal(t, tt.want.Email, got.Email)
					assert.Equal(t, tt.want.Username, got.Username)
					assert.Equal(t, tt.want.PasswordHash, got.PasswordHash)
				}
			} else if !tt.wantErr {
				assert.NotNil(t, got)
				assert.Equal(t, tt.want.Email, got.Email)
				assert.Equal(t, tt.want.Username, got.Username)
				assert.Equal(t, tt.want.PasswordHash, got.PasswordHash)
			} else {
				assert.Nil(t, got)
			}
			err = mock.ExpectationsWereMet()
			if err != nil {
				t.Logf("Mock expectations error: %v", err)
			}
			assert.NoError(t, err)
		})
	}
}
