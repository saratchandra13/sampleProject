package dbmysql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/ShareChat/service-template/config"
	"github.com/ShareChat/service-template/pkg/domain/entity"
	"github.com/ShareChat/service-template/pkg/domain/entity/enterr"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type ReviewStore struct {
	config *config.Store
	db     *sql.DB
}

func NewReviewStore(config *config.Store) *ReviewStore {
	host := config.DataSources.Database.Mysql.Host
	password := config.DataSources.Database.Mysql.Pass
	dbName := config.DataSources.Database.Mysql.DbName
	dataSourceName := host + ":" + password + "@/" + dbName
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("failed to create db instance", err)
	}
	return &ReviewStore{
		config: config,
		db:     db,
	}
}

func generateReviewId() string {
	return uuid.New().String()
}

func (rs *ReviewStore) AddReview(review *entity.Review) (string, error) {
	reviewId := generateReviewId()
	stmt, err := rs.db.Prepare("INSERT INTO beer_review(beer_id, review_id, review_meta, rating, user_id)" +
		"VALUES(?, ?, ?, ?, ?)")

	// this is an example to demonstrate how to create custom error
	if err != nil {
		customErr := enterr.NewCustomError(enterr.AddReviewFailed, "failed", errors.Wrap(err, "some message"))
		return "", customErr
	}
	defer stmt.Close()
	if _, err := stmt.Exec(review.BeerId, reviewId, review.Meta, review.Rating, review.UserId); err != nil {
		return "", err
	}
	return reviewId, nil
}

type reviewSchema struct {
	BeerId     string
	ReviewId   string
	ReviewMeta string
	Rating     int8
	UserId     string
}

func (r *reviewSchema) PrettyPrint() {
	marshalVal, _ := json.MarshalIndent(r, "", " ")
	fmt.Println(string(marshalVal))
}

func parseEntity(dbMeta *reviewSchema) *entity.Review {
	review := entity.NewReview()
	review.Id = dbMeta.ReviewId
	review.BeerId = dbMeta.BeerId
	review.UserId = dbMeta.UserId
	review.Rating = dbMeta.Rating
	review.Meta = dbMeta.ReviewMeta
	return review
}

func (rs *ReviewStore) ListReview(beerId string) ([]*entity.Review, error) {
	stmt, err := rs.db.Prepare("SELECT * FROM beer_review where beer_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(beerId)
	if err != nil {
		return nil, err
	}
	var sch = reviewSchema{}
	var reviewList = make([]*entity.Review, 0)
	for rows.Next() {
		rows.Scan(&sch.BeerId, &sch.ReviewId, &sch.ReviewMeta, &sch.Rating, &sch.UserId)
		reviewList = append(reviewList, parseEntity(&sch))
		sch.PrettyPrint()
	}

	return reviewList, nil
}

func (rs *ReviewStore) Close() {
	rs.db.Close()
}
