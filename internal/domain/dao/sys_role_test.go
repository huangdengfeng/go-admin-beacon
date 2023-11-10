package dao

import (
	"context"
	"testing"
	"time"
)

func TestSysRoleDao_FindRolesByUid(t *testing.T) {
	DoTransaction(context.Background(), func(ctx context.Context) error {
		roles, err := SysRoleDaoInstance.FindRolesByUid(ctx, 1)
		t.Logf("roles:%+v,error:%v", roles, err)
		return nil
	})
}

func TestSysRolePO_TableName(t *testing.T) {
	type fields struct {
		Id         int32
		Code       string
		Name       string
		Status     int8
		CreateUser int32
		CreateTime time.Time
		UpdateUser int32
		UpdateTime time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SysRolePO{
				Id:         tt.fields.Id,
				Code:       tt.fields.Code,
				Name:       tt.fields.Name,
				Status:     tt.fields.Status,
				CreateUser: tt.fields.CreateUser,
				CreateTime: tt.fields.CreateTime,
				UpdateUser: tt.fields.UpdateUser,
				UpdateTime: tt.fields.UpdateTime,
			}
			if got := s.TableName(); got != tt.want {
				t.Errorf("TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
