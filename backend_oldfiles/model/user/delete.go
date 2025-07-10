package user

import (
	"context"
	"fmt"
)

type UserSoftDelete interface {
	SoftDeleteUser(ctx context.Context, id int64) error
	BatchSoftDeleteUsers(ctx context.Context, ids []int64) error
	RestoreUser(ctx context.Context, id int64) error
	BatchRestoreUsers(ctx context.Context, ids []int64) error
	HardDeleteUser(ctx context.Context, id int64) error
	BatchHardDeleteUsers(ctx context.Context, ids []int64) error
}

func (m *customUserModel) SoftDeleteUser(ctx context.Context, id int64) error {
	existUser, err := m.FindOne(ctx, id)
	if err != nil {
		return fmt.Errorf("user not found")
	}
	existUser.IsDeleted = 1
	err = m.Update(ctx, existUser)
	if err != nil {
		return err
	}
	return nil
}

func (m *customUserModel) BatchSoftDeleteUsers(ctx context.Context, ids []int64) error {
	for _, id := range ids {
		err := m.SoftDeleteUser(ctx, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *customUserModel) RestoreUser(ctx context.Context, id int64) error {
	existUser, err := m.FindOne(ctx, id)
	if err != nil {
		return fmt.Errorf("user not found")
	}
	existUser.IsDeleted = 0
	err = m.Update(ctx, existUser)
	if err != nil {
		return err
	}
	return nil
}

func (m *customUserModel) BatchRestoreUsers(ctx context.Context, ids []int64) error {
	for _, id := range ids {
		err := m.RestoreUser(ctx, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *customUserModel) HardDeleteUser(ctx context.Context, id int64) error {
	existUser, err := m.FindOne(ctx, id)
	if err != nil {
		return fmt.Errorf("user not found")
	}
	if existUser.IsDeleted == 0 {
		return fmt.Errorf("user not deleted")
	}
	err = m.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *customUserModel) BatchHardDeleteUsers(ctx context.Context, ids []int64) error {
	for _, id := range ids {
		err := m.HardDeleteUser(ctx, id)
		if err != nil {
			return err
		}
	}
	return nil
}
