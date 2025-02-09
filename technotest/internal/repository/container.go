package repository

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"technotest/internal/entity"
	"technotest/pkg/postgres"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type ContainerRepository struct {
	*postgres.Client
}

func NewContainerRepository(cl *postgres.Client) *ContainerRepository {
	return &ContainerRepository{cl}
}

func (c *ContainerRepository) UpdateContainersStatus(
	ctx context.Context,
	containers []entity.ContainerStatus,
) error {
	qb := psql.Insert("containers").
		Columns("id", "image", "state", "status", "name").
		Suffix("ON CONFLICT (id) DO UPDATE SET image = EXCLUDED.image, " +
			"state = EXCLUDED.state, " +
			"status = EXCLUDED.status, " +
			"name = EXCLUDED.name")

	for _, container := range containers {
		qb = qb.Values(
			container.ID,
			container.Image,
			container.State,
			container.Status,
			container.Name,
		)
	}

	sql, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("qb.ToSql: %w", err)
	}

	if _, err = c.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}

func (c *ContainerRepository) GetContainersStatus(ctx context.Context) ([]entity.ContainerStatus, error) {
	qb := psql.Select("id", "image", "state", "status", "name").From("containers")

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("qb.ToSql: %w", err)
	}

	rows, err := c.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	defer rows.Close()

	var containers []entity.ContainerStatus
	for rows.Next() {
		var container entity.ContainerStatus

		if err = rows.Scan(
			&container.ID,
			&container.Image,
			&container.State,
			&container.Status,
			&container.Name,
		); err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		containers = append(containers, container)
	}

	return containers, nil
}
