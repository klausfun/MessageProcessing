package repository

import (
	"MessageProcessing/models"
	"database/sql"
	"errors"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"log"
	"testing"
)

func TestMessagePostgres_Create(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := NewMessagePostgres(db)

	type args struct {
		message models.Message
	}
	type mockBehavior func(args args, id int)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		id           int
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{
				message: models.Message{
					Content: "test content",
				},
			},
			id: 2,
			mockBehavior: func(args args, id int) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO message").
					WithArgs(args.message.Content).
					WillReturnRows(rows)
			},
		},
		{
			name: "Empty Fields",
			args: args{
				message: models.Message{},
			},
			mockBehavior: func(args args, id int) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(id).RowError(0, errors.New("some error"))
				mock.ExpectQuery("INSERT INTO message").
					WithArgs(args.message.Content).
					WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args, testCase.id)

			got, err := r.Create(testCase.args.message)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.id, got)
			}
		})
	}
}

func TestPostPostgres_GetCurMessages(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := NewMessagePostgres(db)

	testTable := []struct {
		name    string
		mock    func()
		want    []models.Message
		wantErr bool
	}{
		{
			name: "OK",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "content"}).
					AddRow(1, "content1").
					AddRow(2, "content2").
					AddRow(3, "content3")

				mock.ExpectQuery("SELECT (id, content) FROM message WHERE status_id=1").WillReturnRows(rows)
			},
			want: []models.Message{
				{1, "content1"},
				{2, "content2"},
				{3, "content3"},
			},
		},
		{
			name: "No Records",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "content"})

				mock.ExpectQuery("SELECT id, content FROM message WHERE status_id=1").WillReturnRows(rows)
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()

			got, err := r.GetCurMessages()
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.want, got)
			}
		})
	}
}

func TestPostPostgres_GetCompMessages(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := NewMessagePostgres(db)

	testTable := []struct {
		name    string
		mock    func()
		want    []models.Message
		wantErr bool
	}{
		{
			name: "OK",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "content"}).
					AddRow(1, "content1").
					AddRow(2, "content2").
					AddRow(3, "content3")

				mock.ExpectQuery("SELECT (id, content) FROM message WHERE status_id=2").WillReturnRows(rows)
			},
			want: []models.Message{
				{1, "content1"},
				{2, "content2"},
				{3, "content3"},
			},
		},
		{
			name: "No Records",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "content"})

				mock.ExpectQuery("SELECT id, content FROM message WHERE status_id=2").WillReturnRows(rows)
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()

			got, err := r.GetCompMessages()
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.want, got)
			}
		})
	}
}

func TestPostPostgres_UpdateStatus(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := NewMessagePostgres(db)

	testTable := []struct {
		name    string
		mock    func()
		input   int
		wantErr bool
	}{
		{
			name: "OK",
			mock: func() {
				mock.ExpectExec("UPDATE message SET status_id = 2 WHERE id = (.+)").
					WithArgs(5).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: 5,
		},
		{
			name: "Not Found",
			mock: func() {
				mock.ExpectExec("UPDATE message SET status_id = 2 WHERE id = (.+)").
					WithArgs(5).WillReturnError(sql.ErrNoRows)
			},
			input:   5,
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()

			err := r.UpdateStatus(testCase.input)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
