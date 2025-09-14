package link

import (
	"context"
	"fmt"
	"main-service/internal/models"
	"time"

	"github.com/redis/go-redis/v9"
)

type LinkRepository struct {
	client *redis.Client
}

func NewLinkRepository(client *redis.Client) *LinkRepository {
	return &LinkRepository{client: client}
}

func (r *LinkRepository) Create(link *models.Link) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := r.client.Set(ctx, link.OriginLink, link.PseudoLink, 0).Err(); err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := r.client.Set(ctx, link.PseudoLink, link.OriginLink, 0).Err(); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (r *LinkRepository) Find(originLink string) (link *models.Link, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pseudoLink, err := r.client.Get(ctx, originLink).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil // key not found
		}
		return nil, err
	}

	return &models.Link{
		OriginLink: originLink,
		PseudoLink: pseudoLink,
	}, nil
}

func (r *LinkRepository) Delete(originLink string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pseudoLink, err := r.client.Get(ctx, originLink).Result()
	if err != nil {
		if err != redis.Nil {
			return err
		}
	}
	if err := r.client.Del(ctx, originLink).Err(); err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := r.client.Del(ctx, pseudoLink).Err(); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil

}

func (r *LinkRepository) Exist(originLink string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	exists, err := r.client.Exists(ctx, originLink).Result()
	return exists > 0, err
}
