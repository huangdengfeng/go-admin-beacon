package dao

import (
	"context"
	"go-admin-beacon/internal/infrastructure/config"
	"log"
	"os"
	"reflect"
	"testing"
	"time"
)

func setUp() {
	config.ServerConfigPath = "../../../conf"
	config.Init()
	log.Printf("set up")
}

func tearDown() {
	config.Shutdown()
	log.Printf("tear down")
}

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func TestSysUserDao_FindByUid(t *testing.T) {
	type fields struct {
		dao *dao
	}
	type args struct {
		ctx context.Context
		uid int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SysUserPO
		wantErr bool
	}{
		{
			name:    "正常用例",
			fields:  fields{&dao{getDb}},
			args:    args{ctx: withTxDb(context.Background(), getDb()), uid: 1},
			want:    &SysUserPO{Uid: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SysUserDao{
				dao: tt.fields.dao,
			}
			got, err := s.FindByUid(tt.args.ctx, tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByUid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Uid != tt.want.Uid {
				t.Errorf("FindByUid() got = %v, want %v", got, tt.want)
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("FindByUid() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestSysUserDao_FindByUserName(t *testing.T) {
	type fields struct {
		dao *dao
	}
	type args struct {
		userName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SysUserPO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SysUserDao{
				dao: tt.fields.dao,
			}
			got, err := s.FindByUserName(context.Background(), tt.args.userName)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByUserName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByUserName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSysUserDao_Save(t *testing.T) {
	type fields struct {
		dao *dao
	}
	type args struct {
		po *SysUserPO
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "正常用例",
			fields: fields{&dao{getDb}},
			args: args{&SysUserPO{Name: "test-name", UserName: "test-username", SecretKey: "111", Status: 1,
				CreateTime: time.Now(), CreateUser: 1, UpdateTime: time.Now(), UpdateUser: 1}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SysUserDao{
				dao: tt.fields.dao,
			}
			if uid, err := s.Save(context.Background(), tt.args.po); (err != nil && uid > 0) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSysUserDao_FindByPage(t *testing.T) {
	s := SysUserDao{
		dao: &dao{getDb},
	}
	status := int8(1)
	users, total, err := s.FindByPage(&SysUserPOCondition{
		UserName:  "admin",
		FuzzyName: "管理员",
		Name:      "管理员",
		Status:    &status,
	}, 1, 10)
	if err != nil {
		t.Error(err)
	}
	t.Logf("users:%+v,total:%+v", users, total)
}
