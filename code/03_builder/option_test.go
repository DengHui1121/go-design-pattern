package _3_builder

import (
	"reflect"
	"testing"
)

func TestNewResourcePollConfig(t *testing.T) {
	type args struct {
		name string
		opts []ResourcePoolConfigOptFunc
	}
	tests := []struct {
		name    string
		args    args
		want    *ResourcePoolConfig
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "test1",
			args: args{
				name: "1",
				opts: []ResourcePoolConfigOptFunc{WithMaxTotal(5), WithMaxIdle(4), WithMinIdle(1)},
			},
			want: &ResourcePoolConfig{name: "1", maxTotal: 5, maxIdle: 4, minIdle: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewResourcePollConfig(tt.args.name, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewResourcePollConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResourcePollConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
