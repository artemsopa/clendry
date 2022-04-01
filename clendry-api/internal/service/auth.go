package service

import (
	"github.com/artomsopun/clendry/clendry-api/internal/repository"
	"github.com/artomsopun/clendry/clendry-api/pkg/auth"
	"github.com/artomsopun/clendry/clendry-api/pkg/hash"
	"time"
)

type AuthService struct {
	repoUser     repository.Users
	repoSession  repository.Sessions
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager

	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewAuthsService(repoUser repository.Users, repoSession repository.Sessions,
	hasher hash.PasswordHasher, tokenManager auth.TokenManager,
	accessTTL, refreshTTL time.Duration) *AuthService {
	return &AuthService{
		repoUser:        repoUser,
		repoSession:     repoSession,
		hasher:          hasher,
		tokenManager:    tokenManager,
		accessTokenTTL:  accessTTL,
		refreshTokenTTL: refreshTTL,
	}
}
