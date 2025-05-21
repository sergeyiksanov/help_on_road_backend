package repositories

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type tokenClaims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

type TokensRepository struct {
	rdb       *redis.Client
	secretKey []byte
}

func NewTokensRepository(r *redis.Client, secretKey string) *TokensRepository {
	return &TokensRepository{
		rdb:       r,
		secretKey: []byte(secretKey),
	}
}

func (r *TokensRepository) CreateToken(ctx context.Context, userID int64) (string, error) {
	claims := tokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: strconv.FormatInt(userID, 10),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(r.secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	err = r.rdb.Set(ctx, "token:"+tokenString, userID, 0).Err()
	if err != nil {
		return "", fmt.Errorf("failed to save token in Redis: %w", err)
	}

	return tokenString, nil
}

func (r *TokensRepository) GetUserTokens(ctx context.Context, userID int64) ([]string, error) {
	var cursor uint64
	var tokens []string

	for {
		keys, nextCursor, err := r.rdb.Scan(ctx, cursor, "token:*", 100).Result()
		if err != nil {
			return nil, fmt.Errorf("failed to scan tokens: %w", err)
		}

		for _, key := range keys {
			tokenUserID, err := r.rdb.Get(ctx, key).Result()
			if err != nil {
				continue
			}

			if tokenUserID == strconv.FormatInt(userID, 10) {
				token := key[6:]
				tokens = append(tokens, token)
			}
		}

		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}

	return tokens, nil
}

func (r *TokensRepository) DeleteToken(ctx context.Context, token string) error {
	exists, err := r.rdb.Exists(ctx, "token:"+token).Result()
	if err != nil {
		return fmt.Errorf("failed to check token existence: %w", err)
	}
	if exists == 0 {
		return errors.New("token not found")
	}

	_, err = r.rdb.Del(ctx, "token:"+token).Result()
	if err != nil {
		return fmt.Errorf("failed to delete token: %w", err)
	}

	return nil
}

func (r *TokensRepository) GetUserIDByToken(ctx context.Context, token string) (int64, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return r.secretKey, nil
	})

	if err != nil {
		return -1, fmt.Errorf("invalid token: %w", err)
	}

	if !parsedToken.Valid {
		return -1, errors.New("token is not valid")
	}

	claims, ok := parsedToken.Claims.(*tokenClaims)
	if !ok {
		return -1, errors.New("invalid token claims")
	}

	userID, err := r.rdb.Get(ctx, "token:"+token).Result()
	if err == redis.Nil {
		return -1, errors.New("token not found in storage")
	} else if err != nil {
		return -1, fmt.Errorf("failed to get user ID from Redis: %w", err)
	}

	if userID != claims.Subject {
		return -1, errors.New("token user ID mismatch")
	}
	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return -1, errors.New("token user ID mismatch")
	}

	return id, nil
}
