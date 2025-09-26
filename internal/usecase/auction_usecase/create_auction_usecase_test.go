package auction_usecase

import (
	"context"
	"fullcycle-auction_go/configuration"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/internal_error"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type AuctionRepositoryMock struct {
	auctions []auction_entity.Auction
}

func (m *AuctionRepositoryMock) CreateAuction(
	ctx context.Context,
	auctionEntity *auction_entity.Auction) *internal_error.InternalError {
	m.auctions = append(m.auctions, *auctionEntity)
	return nil
}

func (m *AuctionRepositoryMock) FindAuctions(
	ctx context.Context,
	status auction_entity.AuctionStatus,
	category, productName string) ([]auction_entity.Auction, *internal_error.InternalError) {
	var filteredAuctions []auction_entity.Auction
	for _, auction := range m.auctions {
		if auction.Status == status {
			filteredAuctions = append(filteredAuctions, auction)
		}
	}
	return filteredAuctions, nil
}

func (m *AuctionRepositoryMock) FindAuctionById(
	ctx context.Context, id string) (*auction_entity.Auction, *internal_error.InternalError) {
	for _, auction := range m.auctions {
		if auction.Id == id {
			return &auction, nil
		}
	}
	return nil, internal_error.NewNotFoundError("auction not found")
}

func (m *AuctionRepositoryMock) UpdateAuction(
	ctx context.Context,
	auctionEntity *auction_entity.Auction) *internal_error.InternalError {
	for i, auction := range m.auctions {
		if auction.Id == auctionEntity.Id {
			m.auctions[i] = *auctionEntity
			break
		}
	}
	return nil
}

func TestAuctionUseCase_CloseExpiredAuctions(t *testing.T) {
	cfg := &configuration.Conf{
		AuctionInterval: "1s",
	}

	repo := &AuctionRepositoryMock{}
	NewAuctionUseCase(repo, nil, cfg)

	auction, _ := auction_entity.CreateAuction("Test Product", "Test Category", "Test Description", auction_entity.New)
	repo.CreateAuction(context.Background(), auction)

	time.Sleep(2 * time.Second)

	auctions, _ := repo.FindAuctions(context.Background(), auction_entity.Completed, "", "")
	assert.Equal(t, 1, len(auctions))
	assert.Equal(t, auction_entity.Completed, auctions[0].Status)
}
