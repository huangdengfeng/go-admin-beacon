package dao

import (
	"fmt"
	"testing"
	"time"
)

func TestSysRoleDao_FindRolesByUid(t *testing.T) {
	var pos *SysRolePO
	var m map[string]*SysRolePO
	var c chan int

	fmt.Printf("pos:%p,m:%p,c:%p", pos, &m, &c)
	s := make([]SysRolePO, 0, 10)
	r := SysRolePO{Name: "test"}
	s = append(s, r)
	r.Name = "test1"
	fmt.Printf(s[0].Name)
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
