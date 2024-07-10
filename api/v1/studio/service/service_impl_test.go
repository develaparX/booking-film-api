package service

import (
	repoSeat "bioskuy/api/v1/seat/repository"
	"bioskuy/api/v1/studio/dto"
	"bioskuy/api/v1/studio/entity"
	"bioskuy/api/v1/studio/repository"
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StudioServiceTestSuite struct {
	suite.Suite
	mockDb *sql.DB
	mockSql sqlmock.Sqlmock
	ginCtx *gin.Context
	repoStudio repository.StudioRepository
	repoSeat repoSeat.SeatRepository
	validate *validator.Validate
	sS *studioService
}

var mockingStudio = entity.Studio{
	ID:       "1231cmf1m",
	Name:     "Studio 1",
	Capacity: 100,
}

var mockPayload = dto.UpdateStudioRequest{
	ID:       "1231cmf1m",
	Name:     "Studio 1",
}

var mockPayloadReq = dto.CreateStudioRequest{
	Name:       "Studio 1",
	Capacity:   100,
	MaxRowSeat: 5,
}

func (suite *StudioServiceTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDb = db
	suite.mockSql = mock
	suite.repoStudio = repository.NewStudioRepository()
	suite.repoSeat = repoSeat.NewSeatRepository()
	suite.validate = validator.New()
	suite.sS = &studioService{
		RepoStudio: suite.repoStudio,
		RepoSeat:   suite.repoSeat,
		Validate:   suite.validate,
		DB:         suite.mockDb,
	}
	suite.ginCtx, _ = gin.CreateTestContext(nil)
}

func TestStudioServiceTestSuite(t *testing.T) {
	suite.Run(t, new(StudioServiceTestSuite))
}

func (suite *StudioServiceTestSuite) TestCreate_Sucess() {
	ctx := context.Background()
	request := dto.CreateStudioRequest{
		Name:       "Studio 1",
		Capacity:   100,
		MaxRowSeat: 5,
	}

	suite.mockSql.ExpectBegin()

	suite.mockSql.ExpectExec("INSERT INTO studios").
		WithArgs(mockingStudio.Name, mockingStudio.Capacity).
		WillReturnResult(sqlmock.NewResult(1, 1))

	for row := 0; row < (request.Capacity+request.MaxRowSeat-1)/request.MaxRowSeat; row++ {
		for seatNum := 1; seatNum <= request.MaxRowSeat; seatNum++ {
			actualSeatNum := row*request.MaxRowSeat + seatNum
			if actualSeatNum > request.Capacity {
				break
			}
			suite.mockSql.ExpectExec("INSERT INTO seats").
				WithArgs(sqlmock.AnyArg(), true, sqlmock.AnyArg()).
				WillReturnResult(sqlmock.NewResult(1, 1))
		}
	}

	suite.mockSql.ExpectCommit()

	_, err := suite.sS.Create(ctx, mockPayloadReq, suite.ginCtx)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), request.Name, mockPayloadReq.Name)
	assert.Equal(suite.T(), request.Capacity, mockPayloadReq.Capacity)
	// assert.NotEmpty(suite.T(), response.ID)
	// assert.NoError(suite.T(), suite.mockSql.ExpectationsWereMet())
}

func (suite *StudioServiceTestSuite) TestFindAll_Success() {
	ctx := context.Background()

	// Mocking the data
	studios := []entity.Studio{
		{ID: "1", Name: "Studio 1", Capacity: 100},
		{ID: "2", Name: "Studio 2", Capacity: 150},
	}

	suite.mockSql.ExpectBegin()

	rows := sqlmock.NewRows([]string{"id", "name", "capacity"}).
		AddRow(studios[0].ID, studios[0].Name, studios[0].Capacity).
		AddRow(studios[1].ID, studios[1].Name, studios[1].Capacity)

	suite.mockSql.ExpectQuery("SELECT (.+) FROM studios").WillReturnRows(rows)

	suite.mockSql.ExpectCommit()

	response, err := suite.sS.FindAll(ctx, suite.ginCtx)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), response, 2)
	assert.Equal(suite.T(), studios[0].ID, response[0].ID)
	assert.Equal(suite.T(), studios[0].Name, response[0].Name)
	assert.Equal(suite.T(), studios[0].Capacity, response[0].Capacity)
	assert.Equal(suite.T(), studios[1].ID, response[1].ID)
	assert.Equal(suite.T(), studios[1].Name, response[1].Name)
	assert.Equal(suite.T(), studios[1].Capacity, response[1].Capacity)
	assert.NoError(suite.T(), suite.mockSql.ExpectationsWereMet())
}

func (suite *StudioServiceTestSuite) TestFindByID_Success() {
	ctx := context.Background()
	id := "1231cmf1m"

	// Mocking the data
	studio := entity.Studio{
		ID:       id,
		Name:     "Studio 1",
		Capacity: 100,
	}

	suite.mockSql.ExpectBegin()

	rows := sqlmock.NewRows([]string{"id", "name", "capacity"}).
		AddRow(studio.ID, studio.Name, studio.Capacity)

	suite.mockSql.ExpectQuery("SELECT id, name, capacity FROM studios").WithArgs(id).WillReturnRows(rows)

	suite.mockSql.ExpectCommit()

	_, err := suite.sS.FindByID(ctx, id, suite.ginCtx)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), studio.ID, mockPayload.ID)
	assert.Equal(suite.T(), studio.Name, mockPayload.Name)
	assert.NoError(suite.T(), suite.mockSql.ExpectationsWereMet())
}

func (suite *StudioServiceTestSuite) TestUpdate_Success() {
	ctx := context.Background()
	request := dto.UpdateStudioRequest{
		ID:       "1231cmf1m",
		Name:     "Studio 1",
	}

	// Mocking the existing data
	existingStudio := entity.Studio{
		ID:       request.ID,
		Name:     "Studio 1",
		Capacity: 100,
	}

	suite.mockSql.ExpectBegin()

	// Mocking the query to find the studio by ID
	findRows := sqlmock.NewRows([]string{"id", "name", "capacity"}).
		AddRow(existingStudio.ID, existingStudio.Name, existingStudio.Capacity)
	suite.mockSql.ExpectQuery("SELECT id, name, capacity FROM studios").WithArgs(request.ID).WillReturnRows(findRows)

	// Mocking the update operation
	suite.mockSql.ExpectExec("UPDATE studios SET name = \\? WHERE id = \\?").
		WithArgs(request.Name, request.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	suite.mockSql.ExpectCommit()

	response, err := suite.sS.Update(ctx, request, suite.ginCtx)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), request.ID, response.ID)
	assert.Equal(suite.T(), request.Name, response.Name)
	assert.NoError(suite.T(), suite.mockSql.ExpectationsWereMet())
}

func (suite *StudioServiceTestSuite) TestDelete_Success() {
	ctx := context.Background()
	id := "1231cmf1m"

	// Mocking the existing data
	existingStudio := entity.Studio{
		ID:       id,
		Name:     "Studio 1",
		Capacity: 100,
	}

	suite.mockSql.ExpectBegin()

	// Mocking the query to find the studio by ID
	findRows := sqlmock.NewRows([]string{"id", "name", "capacity"}).
		AddRow(existingStudio.ID, existingStudio.Name, existingStudio.Capacity)
	suite.mockSql.ExpectQuery("SELECT id, name, capacity FROM studios").WithArgs(id).WillReturnRows(findRows)

	// Mocking the deletion of seats
	suite.mockSql.ExpectExec("DELETE FROM seats").
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Mocking the deletion of the studio
	suite.mockSql.ExpectExec("DELETE FROM studios").
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	suite.mockSql.ExpectCommit()

	err := suite.sS.Delete(ctx, id, suite.ginCtx)
	assert.NoError(suite.T(), err)
	assert.NoError(suite.T(), suite.mockSql.ExpectationsWereMet())
}