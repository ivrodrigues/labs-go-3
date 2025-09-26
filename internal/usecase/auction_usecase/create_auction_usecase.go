package auction_usecase

import (
	"context"
	"fullcycle-auction_go/configuration"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/entity/bid_entity"
	"fullcycle-auction_go/internal/internal_error"
	"fullcycle-auction_go/internal/usecase/bid_usecase"
	"time"
)

type AuctionInputDTO struct {
	ProductName string           `json:"product_name" binding:"required,min=1"`
	Category    string           `json:"category" binding:"required,min=2"`
	Description string           `json:"description" binding:"required,min=10,max=200"`
	Condition   ProductCondition `json:"condition" binding:"oneof=0 1 2"`
}

type AuctionOutputDTO struct {
	Id          string           `json:"id"`
	ProductName string           `json:"product_name"`
	Category    string           `json:"category"`
	Description string           `json:"description"`
	Condition   ProductCondition `json:"condition"`
	Status      AuctionStatus    `json:"status"`
	Timestamp   time.Time        `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type WinningInfoOutputDTO struct {
	Auction AuctionOutputDTO          `json:"auction"`
	Bid     *bid_usecase.BidOutputDTO `json:"bid,omitempty"`
}

func NewAuctionUseCase(
	auctionRepositoryInterface auction_entity.AuctionRepositoryInterface,
	bidRepositoryInterface bid_entity.BidEntityRepository,
	cfg *configuration.Conf) AuctionUseCaseInterface {
	auctionUseCase := &AuctionUseCase{
		auctionRepositoryInterface: auctionRepositoryInterface,
		bidRepositoryInterface:     bidRepositoryInterface,
	}

	duration, _ := time.ParseDuration(cfg.AuctionInterval)
	go auctionUseCase.closeExpiredAuctions(context.Background(), duration)
	return auctionUseCase
}

type AuctionUseCaseInterface interface {
	CreateAuction(
		ctx context.Context,
		auctionInput AuctionInputDTO) *internal_error.InternalError

	FindAuctionById(
		ctx context.Context, id string) (*AuctionOutputDTO, *internal_error.InternalError)

	FindAuctions(
		ctx context.Context,
		status AuctionStatus,
		category, productName string) ([]AuctionOutputDTO, *internal_error.InternalError)

	FindWinningBidByAuctionId(
		ctx context.Context,
		auctionId string) (*WinningInfoOutputDTO, *internal_error.InternalError)
}

type ProductCondition int64
type AuctionStatus int64

type AuctionUseCase struct {
	auctionRepositoryInterface auction_entity.AuctionRepositoryInterface
	bidRepositoryInterface     bid_entity.BidEntityRepository
}

func (au *AuctionUseCase) CreateAuction(
	ctx context.Context,
	auctionInput AuctionInputDTO) *internal_error.InternalError {
	auction, err := auction_entity.CreateAuction(
		auctionInput.ProductName,
		auctionInput.Category,
		auctionInput.Description,
		auction_entity.ProductCondition(auctionInput.Condition))
	if err != nil {
		return err
	}

	if err := au.auctionRepositoryInterface.CreateAuction(
		ctx, auction); err != nil {
		return err
	}

	return nil
}

func (au *AuctionUseCase) closeExpiredAuctions(ctx context.Context, duration time.Duration) {
	ticker := time.NewTicker(duration)

	for {
		select {
		case <-ticker.C:
			auctions, err := au.auctionRepositoryInterface.FindAuctions(
				ctx, auction_entity.Active, "", "")
			if err != nil {
				continue
			}

			for i := range auctions {
				if time.Now().After(auctions[i].Timestamp.Add(duration)) {
					auctions[i].Status = auction_entity.Completed
					au.auctionRepositoryInterface.UpdateAuction(ctx, &auctions[i])
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
