package repository

import (
	"context"
	"database/sql"
	"portfolio-backend/domain"
	"portfolio-backend/infra/models"

	rss_feeds_pb "github.com/igsr5/portfolio-proto/go/lib/blogs/rss_feed"
	"github.com/lithammer/shortuuid/v3"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type rssFeedRepositoryImpl struct {
	db *sql.DB
}

func NewRSSFeedRepository(db *sql.DB) domain.RSSFeedRepository {
	return rssFeedRepositoryImpl{
		db: db,
	}
}

func (r rssFeedRepositoryImpl) CreateRSSFeed(ctx context.Context, input rss_feeds_pb.CreateRSSFeedRequest) error {
	feedUrl := input.GetUrl()
	if feedUrl == "" {
		return errors.New("url must not be blank.")
	}

	rssFeed := models.RSSFeed{
		ID:  shortuuid.New(),
		URL: feedUrl,
	}

	err := rssFeed.Insert(ctx, r.db, boil.Infer())
	if err != nil {
		return errors.Wrap(err, "failed to insert rss feed to db")
	}

	return nil
}
