package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"portfolio-backend/domain"

	"github.com/aws/aws-lambda-go/events"
	rss_feeds_pb "github.com/igsr5/portfolio-proto/go/lib/blogs/rss_feed"
)

type rssFeedHandlerImpl struct {
	rssFeedRepository domain.RSSFeedRepository
}

func NewRSSFeedHandler(rssFeedRepository domain.RSSFeedRepository) domain.RSSFeedHandler {
	return rssFeedHandlerImpl{
		rssFeedRepository: rssFeedRepository,
	}
}

func (p rssFeedHandlerImpl) BatchGetRSSFeeds(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctx := context.Background()

	rssFeeds, err := p.rssFeedRepository.GetRSSFeeds(ctx)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("failed to get rss feeds errors: %#v", err),
		}, err
	}

	resBody := rss_feeds_pb.BatchGetRSSFeedsResponse{
		RssFeeds: make([]*rss_feeds_pb.RSSFeed, len(rssFeeds)),
	}
	for i, rssFeed := range rssFeeds {
		resBody.RssFeeds[i] = &rss_feeds_pb.RSSFeed{
			Id:  rssFeed.Id,
			Url: rssFeed.Url,
		}
	}

	resBodyStr, err := json.Marshal(resBody)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("failed json.Unmarshal() with errors: %#v", err),
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(resBodyStr),
	}, nil
}

func (p rssFeedHandlerImpl) GetRSSFeed(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctx := context.Background()

	id := request.PathParameters["id"]

	rssFeed, err := p.rssFeedRepository.GetRSSFeed(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return events.APIGatewayProxyResponse{
				StatusCode: 404,
				Body:       "rss feed not found",
			}, nil

		}
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("failed to get rss feeds errors: %#v", err),
		}, err
	}

	resBody := rss_feeds_pb.GetRSSFeedResponse{
		RssFeed: &rss_feeds_pb.RSSFeed{
			Id:  rssFeed.Id,
			Url: rssFeed.Url,
		}}

	resBodyStr, err := json.Marshal(resBody)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("failed json.Unmarshal() with errors: %#v", err),
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(resBodyStr),
	}, nil
}

func (p rssFeedHandlerImpl) CreateRSSFeed(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctx := context.Background()

	params := rss_feeds_pb.CreateRSSFeedRequest{}
	err := json.Unmarshal([]byte(request.Body), &params)
	if err != nil {
		log.Printf("failed json.Unmarshal() with errors: %#v", err)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("failed json.Unmarshal() with errors: %#v", err),
		}, err
	}

	if params.GetUrl() == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "bad request body. json field `url` is must be specifed.",
		}, nil
	}

	exists, err := p.rssFeedRepository.IsExistsUrl(ctx, params.GetUrl())
	if err != nil {
		log.Printf("failed to IsExistsUrl() with errors: %#v", err)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("failed to check whether exists url in rss_feeds table. with errors: %#v", err),
		}, err
	}

	if exists {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "bad request body. specifed `url` is already exists.",
		}, nil
	}

	err = p.rssFeedRepository.CreateRSSFeed(ctx, params)
	if err != nil {
		log.Printf("failed create rss feed with errors: %#v", err)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("failed create rss feed with errors: %#v", err),
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func getURL(body string) (string, error) {
	b := struct {
		Url string `json:"url"`
	}{}

	err := json.Unmarshal([]byte(body), &b)

	return b.Url, err
}

func (p rssFeedHandlerImpl) DeleteRSSFeed(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ctx := context.Background()

	id := request.PathParameters["id"]

	err := p.rssFeedRepository.DeleteRSSFeed(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return events.APIGatewayProxyResponse{
				StatusCode: 404,
				Body:       "rss feed not found",
			}, nil

		}
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       fmt.Sprintf("failed to get rss feeds errors: %#v", err),
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}
