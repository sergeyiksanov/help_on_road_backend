package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/sergeyiksanov/help-on-road/internal/models"
)

type HelpRepository struct {
	rdb *redis.Client
}

func NewHelpRepository(rdb *redis.Client) *HelpRepository {
	return &HelpRepository{rdb: rdb}
}

const redisPrefix = "help_call"
const maxCallsPerUser = 10

func redisKey(userID int64) string {
	return fmt.Sprintf("%s:%d", redisPrefix, userID)
}

func (r *HelpRepository) GetAll(ctx context.Context) ([]*models.HelpCall, error) {
	keys, err := r.rdb.Keys(ctx, redisPrefix+":*").Result()
	if err != nil {
		return nil, err
	}

	var results []*models.HelpCall
	for _, key := range keys {
		val, err := r.rdb.Get(ctx, key).Result()
		if err != nil {
			continue
		}

		var calls []*models.HelpCall
		if err := json.Unmarshal([]byte(val), &calls); err != nil {
			log.Printf("get all: %e", err)
			continue
		}

		results = append(results, calls...)
	}

	return results, nil
}

func (r *HelpRepository) GetByUserId(ctx context.Context, id int64) ([]*models.HelpCall, error) {
	val, err := r.rdb.Get(ctx, redisKey(id)).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var calls []*models.HelpCall
	if err := json.Unmarshal([]byte(val), &calls); err != nil {
		log.Printf("get by u id: %e", err)
		return nil, err
	}

	return calls, nil
}

func (r *HelpRepository) Add(ctx context.Context, m *models.HelpCall) error {
	calls, err := r.GetByUserId(ctx, m.Caller.Id)
	if err != nil {
		return err
	}

	if calls == nil {
		calls = []*models.HelpCall{}
	}

	calls = append([]*models.HelpCall{m}, calls...)

	if len(calls) > maxCallsPerUser {
		calls = calls[:maxCallsPerUser]
	}

	log.Print(calls)
	b, err := json.Marshal(calls)
	if err != nil {
		return err
	}
	log.Print(b)

	return r.rdb.Set(ctx, redisKey(m.Caller.Id), b, 0).Err()
}

func (r *HelpRepository) Delete(ctx context.Context, id int64) error {
	return r.rdb.Del(ctx, redisKey(id)).Err()
}

func (r *HelpRepository) Update(ctx context.Context, userID int64, index int, updatedCall *models.HelpCall) error {
	calls, err := r.GetByUserId(ctx, userID)
	if err != nil {
		return err
	}

	if calls == nil || len(calls) == 0 {
		return fmt.Errorf("no calls found for user %d", userID)
	}

	if index < 0 || index >= len(calls) {
		return fmt.Errorf("invalid index: %d (valid range: 0-%d)", index, len(calls)-1)
	}

	updatedCall.Caller.Id = calls[index].Caller.Id
	calls[index] = updatedCall

	b, err := json.Marshal(calls)
	if err != nil {
		return err
	}

	return r.rdb.Set(ctx, redisKey(userID), b, 0).Err()
}
