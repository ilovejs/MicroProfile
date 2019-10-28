package db

import (
	"context"
	"github.com/ilovejs/profile/schema"
)

type Repository interface {
	CreateProfile(ctx context.Context, p schema.Profile) (int, error)
	UpdateProfile(ctx context.Context, profile schema.Profile, pid int) error
	ListProfiles(ctx context.Context, skip uint64, take uint64) ([]schema.Profile, error)
	Close()
}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}

func CreateProfile(ctx context.Context, profile schema.Profile) (int, error) {
	return impl.CreateProfile(ctx, profile)
}

func UpdateProfile(ctx context.Context, profile schema.Profile, pid int) error {
	return impl.UpdateProfile(ctx, profile, pid)
}

func ListProfiles(ctx context.Context, skip uint64, take uint64) ([]schema.Profile, error) {
	return impl.ListProfiles(ctx, skip, take)
}
