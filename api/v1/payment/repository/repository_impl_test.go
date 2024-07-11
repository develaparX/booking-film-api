package repository

import (
	"bioskuy/api/v1/payment/entity"
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PaymentRepositoryTestSuite struct {
	suite.Suite
	repo       PaymentRepository
	mockDb     *sql.DB
	mockSql    sqlmock.Sqlmock
	ctx        context.Context
	ginContext *gin.Context
}

func (suite *PaymentRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	if err != nil {
		suite.T().Fatal(err)
	}

	suite.mockDb = db
	suite.mockSql = mock
	suite.repo = NewPaymentRepository()
	suite.ctx = context.TODO()
	suite.ginContext = &gin.Context{}
}

func (suite *PaymentRepositoryTestSuite) TearDownTest() {
	suite.mockDb.Close()
}

func TestPaymentRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(PaymentRepositoryTestSuite))
}

func (suite *PaymentRepositoryTestSuite) TestSave_Success() {
	payment := entity.Payment{
		UserID:                 "user-id",
		SeatDetailForBookingID: "seat-detail-id",
		TotalSeat:              5,
		TotalPrice:             100,
	}

	suite.mockSql.ExpectBegin()
	tx, err := suite.mockDb.Begin()
	assert.NoError(suite.T(), err)

	suite.mockSql.ExpectQuery("INSERT INTO payments").
		WithArgs(payment.UserID, payment.SeatDetailForBookingID, payment.TotalSeat, payment.TotalPrice).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("payment-id"))

	savedPayment, err := suite.repo.Save(suite.ctx, tx, payment, suite.ginContext)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "payment-id", savedPayment.ID)
}

func (suite *PaymentRepositoryTestSuite) TestFindAll_Success() {
	suite.mockSql.ExpectBegin()
	tx, err := suite.mockDb.Begin()
	assert.NoError(suite.T(), err)

	query := `SELECT p.id, p.user_id, p.seatdetailforbooking_id, p.total_seat, p.total_price, p.status,
                     sb.id AS seat_booking_id, sb.status AS seat_booking_status,
                     s.id AS seat_id, s.seat_name, s.isAvailable AS seat_isAvailable,
                     sh.id AS showtime_id, sh.show_start, sh.show_end,
                     m.id AS movie_id, m.title AS movie_title, m.description AS movie_description, 
                     m.price AS movie_price, m.duration AS movie_duration, m.status AS movie_status,
                     st.id AS studio_id, st.name AS studio_name
              FROM payments p
              JOIN seat_detail_for_bookings sdfb ON p.seatdetailforbooking_id = sdfb.id
              JOIN seats s ON sdfb.seat_id = s.id
              JOIN seat_bookings sb ON sdfb.seatBooking_id = sb.id
              JOIN showtimes sh ON sb.showtime_id = sh.id
              JOIN movies m ON sh.movie_id = m.id
              JOIN studios st ON sh.studio_id = st.id`

	suite.mockSql.ExpectQuery(query).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "user_id", "seatdetailforbooking_id", "total_seat", "total_price", "status",
			"seat_booking_id", "seat_booking_status", "seat_id", "seat_name", "seat_isAvailable",
			"showtime_id", "show_start", "show_end", "movie_id", "movie_title", "movie_description",
			"movie_price", "movie_duration", "movie_status", "studio_id", "studio_name"}).
			AddRow("payment-id", "user-id", "seat-detail-id", 5, 100, "status",
				"seat-booking-id", "seat-booking-status", "seat-id", "seat-name", true,
				"showtime-id", "show-start", "show-end", "movie-id", "movie-title", "movie-description",
				10, "movie-duration", "movie-status", "studio-id", "studio-name"))

	payments, err := suite.repo.FindAll(suite.ctx, tx, suite.ginContext)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), payments, 1)
	assert.Equal(suite.T(), "payment-id", payments[0].ID)
}

func (suite *PaymentRepositoryTestSuite) TestUpdate_Success() {
	payment := entity.Payment{
		ID:     "payment-id",
		Status: "new-status",
	}

	suite.mockSql.ExpectBegin()
	tx, err := suite.mockDb.Begin()
	assert.NoError(suite.T(), err)

	query := `UPDATE payments SET status = $1 WHERE id = $2`

	suite.mockSql.ExpectExec(query).WithArgs(payment.Status, payment.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	updatedPayment, err := suite.repo.Update(suite.ctx, tx, payment, suite.ginContext)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "new-status", updatedPayment.Status)
}

func (suite *PaymentRepositoryTestSuite) TestDelete_Success() {
	id := "payment-id"

	suite.mockSql.ExpectBegin()
	tx, err := suite.mockDb.Begin()
	assert.NoError(suite.T(), err)

	query := `DELETE FROM payments WHERE id = $1`

	suite.mockSql.ExpectExec(query).WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = suite.repo.Delete(suite.ctx, tx, id, suite.ginContext)
	assert.NoError(suite.T(), err)
}

func (suite *PaymentRepositoryTestSuite) TestSave_InsertError() {
	payment := entity.Payment{
		UserID:                 "user-id",
		SeatDetailForBookingID: "seat-detail-id",
		TotalSeat:              5,
		TotalPrice:             100,
	}

	suite.mockSql.ExpectBegin()
	tx, err := suite.mockDb.Begin()
	assert.NoError(suite.T(), err)

	suite.mockSql.ExpectQuery("INSERT INTO payments").
		WithArgs(payment.UserID, payment.SeatDetailForBookingID, payment.TotalSeat, payment.TotalPrice).
		WillReturnError(errors.New("insert error"))

	savedPayment, err := suite.repo.Save(suite.ctx, tx, payment, suite.ginContext)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "insert error", err.Error())
	assert.Equal(suite.T(), "", savedPayment.ID)
}

func (suite *PaymentRepositoryTestSuite) TestFindByID_QueryError() {
	id := "payment-id"
	suite.mockSql.ExpectBegin()
	tx, err := suite.mockDb.Begin()
	assert.NoError(suite.T(), err)

	query := `SELECT p.id, p.user_id, p.seatdetailforbooking_id, p.total_seat, p.total_price, p.status,
                     sb.id AS seat_booking_id, sb.status AS seat_booking_status,
                     s.id AS seat_id, s.seat_name, s.isAvailable AS seat_isAvailable,
                     sh.id AS showtime_id, sh.show_start, sh.show_end,
                     m.id AS movie_id, m.title AS movie_title, m.description AS movie_description, 
                     m.price AS movie_price, m.duration AS movie_duration, m.status AS movie_status,
                     st.id AS studio_id, st.name AS studio_name
              WHERE p.id = $1`

	suite.mockSql.ExpectQuery(query).WithArgs(id).WillReturnError(errors.New("query error"))

	payment, err := suite.repo.FindByID(suite.ctx, tx, id, suite.ginContext)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "query error", err.Error())
	assert.Equal(suite.T(), entity.Payment{}, payment)
}

func (suite *PaymentRepositoryTestSuite) TestFindAll_QueryError() {
	suite.mockSql.ExpectBegin()
	tx, err := suite.mockDb.Begin()
	assert.NoError(suite.T(), err)

	query := `SELECT p.id, p.user_id, p.seatdetailforbooking_id, p.total_seat, p.total_price, p.status,
                     sb.id AS seat_booking_id, sb.status AS seat_booking_status,
                     s.id AS seat_id, s.seat_name, s.isAvailable AS seat_isAvailable,
                     sh.id AS showtime_id, sh.show_start, sh.show_end,
                     m.id AS movie_id, m.title AS movie_title, m.description AS movie_description, 
                     m.price AS movie_price, m.duration AS movie_duration, m.status AS movie_status,
                     st.id AS studio_id, st.name AS studio_name
              FROM payments p
              JOIN seat_detail_for_bookings sdfb ON p.seatdetailforbooking_id = sdfb.id
              JOIN seats s ON sdfb.seat_id = s.id
              JOIN seat_bookings sb ON sdfb.seatBooking_id = sb.id
              JOIN showtimes sh ON sb.showtime_id = sh.id
              JOIN movies m ON sh.movie_id = m.id
              JOIN studios st ON sh.studio_id = st.id`

	suite.mockSql.ExpectQuery(query).WillReturnError(errors.New("query error"))

	payments, err := suite.repo.FindAll(suite.ctx, tx, suite.ginContext)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "query error", err.Error())
	assert.Nil(suite.T(), payments)
}

func (suite *PaymentRepositoryTestSuite) TestUpdate_UpdateError() {
	payment := entity.Payment{
		ID:     "payment-id",
		Status: "paid",
	}

	suite.mockSql.ExpectBegin()
	tx, err := suite.mockDb.Begin()
	assert.NoError(suite.T(), err)

	query := `UPDATE payments SET status = $1 WHERE id = $2`

	suite.mockSql.ExpectExec(query).
		WithArgs(payment.Status, payment.ID).
		WillReturnError(errors.New("update error"))

	updatedPayment, err := suite.repo.Update(suite.ctx, tx, payment, suite.ginContext)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "update error", err.Error())
	assert.Equal(suite.T(), payment.ID, updatedPayment.ID)
	assert.Equal(suite.T(), "", updatedPayment.Status)
}

func (suite *PaymentRepositoryTestSuite) TestDelete_DeleteError() {
	paymentId := "payment-id"

	suite.mockSql.ExpectBegin()
	tx, err := suite.mockDb.Begin()
	assert.NoError(suite.T(), err)

	query := `DELETE FROM payments WHERE id = $1`

	suite.mockSql.ExpectExec(query).
		WithArgs(paymentId).
		WillReturnError(errors.New("delete error"))

	err = suite.repo.Delete(suite.ctx, tx, paymentId, suite.ginContext)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "delete error", err.Error())
}

func (suite *PaymentRepositoryTestSuite) TestUpdate_Error() {
	payment := entity.Payment{
		ID:     "payment-id",
		Status: "paid",
	}

	suite.mockSql.ExpectBegin()
	tx, err := suite.mockDb.Begin()
	assert.NoError(suite.T(), err)

	query := `UPDATE payments SET status = $1 WHERE id = $2`

	suite.mockSql.ExpectExec(query).
		WithArgs(payment.Status, payment.ID).
		WillReturnError(errors.New("update error"))

	updatedPayment, err := suite.repo.Update(suite.ctx, tx, payment, suite.ginContext)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "update error", err.Error())
	assert.Equal(suite.T(), payment.ID, updatedPayment.ID)
	assert.Equal(suite.T(), "", updatedPayment.Status)
}

func (suite *PaymentRepositoryTestSuite) TestDelete_Error() {
	paymentId := "payment-id"

	suite.mockSql.ExpectBegin()
	tx, err := suite.mockDb.Begin()
	assert.NoError(suite.T(), err)

	query := `DELETE FROM payments WHERE id = $1`

	suite.mockSql.ExpectExec(query).
		WithArgs(paymentId).
		WillReturnError(errors.New("delete error"))

	err = suite.repo.Delete(suite.ctx, tx, paymentId, suite.ginContext)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "delete error", err.Error())
}

func (suite *PaymentRepositoryTestSuite) TestSave_NullOrEmptyValues() {
	payment := entity.Payment{
		UserID:                 "",
		SeatDetailForBookingID: "",
		TotalSeat:              0,
		TotalPrice:             0,
	}

	suite.mockSql.ExpectBegin()
	tx, err := suite.mockDb.Begin()
	assert.NoError(suite.T(), err)

	suite.mockSql.ExpectQuery("INSERT INTO payments").
		WithArgs(payment.UserID, payment.SeatDetailForBookingID, payment.TotalSeat, payment.TotalPrice).
		WillReturnError(errors.New("insert error"))

	savedPayment, err := suite.repo.Save(suite.ctx, tx, payment, suite.ginContext)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "insert error", err.Error())
	assert.Equal(suite.T(), "", savedPayment.ID)
}

func (suite *PaymentRepositoryTestSuite) TestSave_InvalidDataTypes() {
	payment := entity.Payment{
		UserID:                 "user-id",
		SeatDetailForBookingID: "seat-detail-id",
		TotalSeat:              5, // Invalid data type
		TotalPrice:             100,
	}

	suite.mockSql.ExpectBegin()
	tx, err := suite.mockDb.Begin()
	assert.NoError(suite.T(), err)

	suite.mockSql.ExpectQuery("INSERT INTO payments").
		WithArgs(payment.UserID, payment.SeatDetailForBookingID, payment.TotalSeat, payment.TotalPrice).
		WillReturnError(errors.New("invalid data type"))

	savedPayment, err := suite.repo.Save(suite.ctx, tx, payment, suite.ginContext)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "invalid data type", err.Error())
	assert.Equal(suite.T(), "", savedPayment.ID)
}

func (suite *PaymentRepositoryTestSuite) TestFindByID_NonExistentID() {
	id := "non-existent-id"
	suite.mockSql.ExpectBegin()
	tx, err := suite.mockDb.Begin()
	assert.NoError(suite.T(), err)

	query := `SELECT p.id, p.user_id, p.seatdetailforbooking_id, p.total_seat, p.total_price, p.status,
                     sb.id AS seat_booking_id, sb.status AS seat_booking_status,
                     s.id AS seat_id, s.seat_name, s.isAvailable AS seat_isAvailable,
                     sh.id AS showtime_id, sh.show_start, sh.show_end,
                     m.id AS movie_id, m.title AS movie_title, m.description AS movie_description, 
                     m.price AS movie_price, m.duration AS movie_duration, m.status AS movie_status,
                     st.id AS studio_id, st.name AS studio_name
              WHERE p.id = $1`

	suite.mockSql.ExpectQuery(query).WithArgs(id).WillReturnError(errors.New("not found"))

	payment, err := suite.repo.FindByID(suite.ctx, tx, id, suite.ginContext)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "not found", err.Error())
	assert.Equal(suite.T(), entity.Payment{}, payment)
}

func (suite *PaymentRepositoryTestSuite) TestDatabaseConnectionError() {
	suite.mockSql.ExpectBegin().WillReturnError(errors.New("db connection error"))

	_, err := suite.mockDb.Begin()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "db connection error", err.Error())
}
