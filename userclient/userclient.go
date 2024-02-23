package userclient

import (
	"context"
	"strconv"
	"zg5/day1/rpc/userrpc/model/mysql"
	pb "zg5/day1/rpc/userrpc/userrpc"
)

type UserRpcClient struct {
	pb.UnimplementedStreamGreeterServer
}

func (c *UserRpcClient) AccordingToConditionGet(ctx context.Context, in *pb.AccordingToConditionGetReq) (*pb.AccordingToConditionGetResp, error) {
	user, err := mysql.AccordingToConditionGetUser(in.Condition)
	if err != nil {
		return &pb.AccordingToConditionGetResp{}, err
	}
	userProto := &pb.User{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
		Phone:    user.Phone,
		Email:    user.Email,
	}
	return &pb.AccordingToConditionGetResp{
		Data: userProto,
	}, nil
}

func (c *UserRpcClient) AccordingToConditionGets(ctx context.Context, in *pb.AccordingToConditionGetsReq) (*pb.AccordingToConditionGetsResp, error) {
	users, err := mysql.AccordingToConditionGetUsers(in.Condition)
	if err != nil {
		return &pb.AccordingToConditionGetsResp{}, err
	}
	var data []*pb.User
	for _, v := range users {
		user := &pb.User{
			Id:       v.Id,
			Username: v.Username,
			Password: v.Password,
			Phone:    v.Phone,
			Email:    v.Email,
		}
		data = append(data, user)
	}
	return &pb.AccordingToConditionGetsResp{
		Data: data,
	}, nil
}

func (c *UserRpcClient) InsertUser(ctx context.Context, in *pb.InsertUserReq) (*pb.InsertUserResp, error) {
	user, err := mysql.InsertUser(mysql.User{
		Username: in.Data.Username,
		Password: in.Data.Password,
		Phone:    in.Data.Phone,
		Email:    in.Data.Email,
	})
	if err != nil {
		return &pb.InsertUserResp{}, err
	}
	return &pb.InsertUserResp{
		UserId: user.Id,
	}, nil
}

func (c *UserRpcClient) UpdateUser(ctx context.Context, in *pb.UpdateUserReq) (*pb.UpdateUserResp, error) {
	user, err := mysql.AccordingToConditionGetUser(map[string][]byte{"id": []byte(strconv.Itoa(int(in.Data.Id)))})
	if err != nil {
		return &pb.UpdateUserResp{}, err
	}
	user.Username = in.Data.Username
	user.Password = in.Data.Password
	user.Email = in.Data.Email
	user.Phone = in.Data.Phone
	err = mysql.UpdateUser(user)
	if err != nil {
		return &pb.UpdateUserResp{}, err
	}
	return &pb.UpdateUserResp{
		Status: 1,
	}, nil
}

func (c *UserRpcClient) DeletedUser(ctx context.Context, in *pb.DeletedUserReq) (*pb.DeletedUserResp, error) {
	err := mysql.DeletedUser(mysql.User{
		Id: in.UserId,
	})
	if err != nil {
		return &pb.DeletedUserResp{}, err
	}
	return &pb.DeletedUserResp{
		Status: 1,
	}, nil
}
