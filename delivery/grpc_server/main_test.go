package grpc_server

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hengkyawijaya/simple-go/delivery/grpc_server/client"
	"github.com/hengkyawijaya/simple-go/mock"
	"github.com/hengkyawijaya/simple-go/usecase"
)

func Test_grpcServer_Hello(t *testing.T) {
	type fields struct {
		Uc *usecase.Usecase
	}
	type args struct {
		context context.Context
		request *client.HelloRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *client.HelloReply
		wantErr bool
	}{
		{
			name:   "",
			fields: fields{},
			args:   args{},
			want: &client.HelloReply{
				Message: "Hello",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockHelloUC := mock.NewMockHelloUsecase(ctrl)
			mockHelloUC.EXPECT().Hello().Return("Hello")

			g := &grpcServer{
				Uc: &usecase.Usecase{
					HelloUsecase: mockHelloUC,
				},
			}
			got, err := g.Hello(tt.args.context, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hello() got = %v, want %v", got, tt.want)
			}
		})
	}
}
