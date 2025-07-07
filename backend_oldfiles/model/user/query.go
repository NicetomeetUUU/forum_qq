package user

import (
	"context"
	"fmt"
)

type UsersQuery interface {
	FindUsersByStatus(ctx context.Context, status string) ([]*User, error)
	FindUsersByRole(ctx context.Context, role string) ([]*User, error)
	FindAllActiveUsers(ctx context.Context) ([]*User, error)
}

type UsersCount interface {
	CountUsersByStatus(ctx context.Context, status string) (int64, error)
	CountUsersByRole(ctx context.Context, role string) (int64, error)
	CountAllActiveUsers(ctx context.Context) (int64, error)
}

func (m *customUserModel) FindUsersByStatus(ctx context.Context, status string) ([]*User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE status = ?", m.table)
	var users []*User
	err := m.QueryRowsNoCacheCtx(ctx, &users, query, status)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (m *customUserModel) FindUsersByRole(ctx context.Context, role string) ([]*User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE role = ?", m.table)
	var users []*User
	err := m.QueryRowsNoCacheCtx(ctx, &users, query, role)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (m *customUserModel) FindAllActiveUsers(ctx context.Context) ([]*User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE status = 'active'", m.table)
	var users []*User
	err := m.QueryRowsNoCacheCtx(ctx, &users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (m *customUserModel) CountUsersByStatus(ctx context.Context, status string) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE status = ?", m.table)
	var count int64
	err := m.QueryRowNoCacheCtx(ctx, &count, query, status)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (m *customUserModel) CountUsersByRole(ctx context.Context, role string) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE role = ?", m.table)
	var count int64
	err := m.QueryRowNoCacheCtx(ctx, &count, query, role)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (m *customUserModel) CountAllActiveUsers(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE status = 'active'", m.table)
	var count int64
	err := m.QueryRowNoCacheCtx(ctx, &count, query)
	if err != nil {
		return 0, err
	}
	return count, nil
}
