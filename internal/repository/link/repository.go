package link

import (
	"context"
	"errors"
	"time"

	link_d "github.com/aygumov-g/service-url-shortener-go/internal/domain/link"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, link *link_d.Link) error {
	err := r.db.QueryRow(
		ctx,
		`
		INSERT INTO links (
			original_url,
			custom_code,
			click_count,
			last_accessed_at,
			created_at
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
		`,
		link.OriginalURL,
		link.CustomCode,
		link.ClickCount,
		link.LastAccessedAt,
		link.CreatedAt,
	).Scan(&link.ID)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return link_d.ErrCustomCodeAlreadyExists
			}
		}

		return err
	}

	return nil
}

func (r *repository) Update(ctx context.Context, id int64, now time.Time) error {
	_, err := r.db.Exec(
		ctx,
		`
		UPDATE links
		SET
			last_accessed_at = $2,
			click_count = click_count + 1
		WHERE
			id = $1
		`,
		id,
		now,
	)

	return err
}

func (r *repository) GetByID(ctx context.Context, id int64) (*link_d.Link, error) {
	row := r.db.QueryRow(
		ctx,
		`
		SELECT
			id,
			original_url,
			custom_code,
			click_count,
			last_accessed_at,
			created_at
		FROM links
		WHERE
			id = $1
		`,
		id,
	)

	var link link_d.Link
	if err := row.Scan(
		&link.ID,
		&link.OriginalURL,
		&link.CustomCode,
		&link.ClickCount,
		&link.LastAccessedAt,
		&link.CreatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, link_d.ErrLinkNotFound
		}

		return nil, err
	}

	return &link, nil
}

func (r *repository) GetByCustomCode(ctx context.Context, code string) (*link_d.Link, error) {
	row := r.db.QueryRow(
		ctx,
		`
		SELECT
			id,
			original_url,
			custom_code,
			click_count,
			last_accessed_at,
			created_at
		FROM links
		WHERE
		custom_code = $1
		`,
		code,
	)

	var link link_d.Link
	if err := row.Scan(
		&link.ID,
		&link.OriginalURL,
		&link.CustomCode,
		&link.ClickCount,
		&link.LastAccessedAt,
		&link.CreatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, link_d.ErrLinkNotFound
		}

		return nil, err
	}

	return &link, nil
}
