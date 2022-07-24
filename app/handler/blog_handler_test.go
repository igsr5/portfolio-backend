package handler

import (
	"encoding/json"
	"portfolio-backend/domain"
	mock_domain "portfolio-backend/domain/mock"
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	blogs_pb "github.com/igsr5/portfolio-proto/go/lib/blogs"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestBatchGetBlogs(t *testing.T) {
	tests := []struct {
		name           string
		request        events.APIGatewayProxyRequest
		mockFn         func(mr *mock_domain.MockBlogRepository)
		wantStatusCode int
		wantBlogs      []*blogs_pb.Blog
	}{
		{
			name: "get all blogs",
			request: events.APIGatewayProxyRequest{
				HTTPMethod: "GET",
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
			mockFn: func(mr *mock_domain.MockBlogRepository) {
				mr.EXPECT().GetBlogs(gomock.Any()).Return([]*domain.Blog{
					{
						Id:           "aaa",
						Title:        "Hello World",
						PostedAt:     time.Date(2020, time.December, 10, 23, 1, 10, 0, time.Local),
						SiteUrl:      "https://example.com",
						ThumbnailUrl: "https://example.com/thumbnail.png",
						ServiceName:  "Qiita",
					},
					{
						Id:           "bbb",
						Title:        "Hello World",
						PostedAt:     time.Date(2020, time.December, 10, 23, 1, 10, 0, time.Local),
						SiteUrl:      "https://example.com",
						ThumbnailUrl: "https://example.com/thumbnail.png",
						ServiceName:  "Qiita",
					},
				}, nil)
			},
			wantStatusCode: 200,
			wantBlogs: []*blogs_pb.Blog{
				{
					Id:           "aaa",
					Title:        "Hello World",
					PostedAt:     timestamppb.New(time.Date(2020, time.December, 10, 23, 1, 10, 0, time.Local)),
					SiteUrl:      "https://example.com",
					ThumbnailUrl: "https://example.com/thumbnail.png",
					ServiceName:  "Qiita",
				},
				{
					Id:           "bbb",
					Title:        "Hello World",
					PostedAt:     timestamppb.New(time.Date(2020, time.December, 10, 23, 1, 10, 0, time.Local)),
					SiteUrl:      "https://example.com",
					ThumbnailUrl: "https://example.com/thumbnail.png",
					ServiceName:  "Qiita",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			mbr := mock_domain.NewMockBlogRepository(mockCtrl)
			tt.mockFn(mbr)

			h := NewBlogHandler(mbr)

			res, err := h.BatchGetBlogs(tt.request)
			if err != nil {
				t.Errorf("failed to BatchGetBlogs. err: %v", err)
			}

			if res.StatusCode != tt.wantStatusCode {
				t.Errorf("bad status code. got: %d, want %d", res.StatusCode, tt.wantStatusCode)
			}

			var resBody blogs_pb.BatchGetBlogsResponse
			err = json.Unmarshal([]byte(res.Body), &resBody)
			if err != nil {
				t.Fatalf("failed to unmarshal response body. %v", err)
			}

			if !cmp.Equal(resBody.Blogs, tt.wantBlogs, protocmp.Transform()) {
				t.Fatalf("bad blogs. diff: %v", cmp.Diff(resBody.Blogs, tt.wantBlogs))
			}
		})
	}
}
