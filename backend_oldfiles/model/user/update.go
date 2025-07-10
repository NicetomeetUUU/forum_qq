package user

import (
	"context"
	"fmt"
)

type UserUpdate interface {
	UpdateUser(ctx context.Context, id int64, data *User) error
}

type UserBatchUpdate interface {
	UpdataUsersStatus(ctx context.Context, ids []int64, status string) error
	UpdataUsersRole(ctx context.Context, ids []int64, role string) error
}

func (m *customUserModel) UpdateUser(ctx context.Context, id int64, data *User) error {
	existUser, err := m.FindOne(ctx, id)
	if err != nil {
		return fmt.Errorf("user not found")
	}
	existUser.Status = data.Status
	err = m.Update(ctx, existUser)
	if err != nil {
		return err
	}
	return nil
}

func (m *customUserModel) UpdataUsersStatus(ctx context.Context, ids []int64, status string) error {
	for _, id := range ids {
		err := m.UpdateUser(ctx, id, &User{Status: status})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *customUserModel) UpdataUsersRole(ctx context.Context, ids []int64, role string) error {
	return nil
}
