// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: relation/relation.proto

package relation

import (
	context "context"
	common "github.com/clubo-app/protobuf/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RelationServiceClient is the client API for RelationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelationServiceClient interface {
	FriendRequest(ctx context.Context, in *FriendRequestRequest, opts ...grpc.CallOption) (*common.SuccessIndicator, error)
	AcceptFriend(ctx context.Context, in *AcceptFriendRequest, opts ...grpc.CallOption) (*common.SuccessIndicator, error)
	RemoveFriend(ctx context.Context, in *RemoveFriendRequest, opts ...grpc.CallOption) (*common.SuccessIndicator, error)
	GetFriendRelation(ctx context.Context, in *GetFriendRelationRequest, opts ...grpc.CallOption) (*FriendRelation, error)
	GetFriendsOfUser(ctx context.Context, in *GetFriendsOfUserRequest, opts ...grpc.CallOption) (*PagedFriendRelations, error)
	GetFriendCount(ctx context.Context, in *GetFriendCountRequest, opts ...grpc.CallOption) (*GetFriendCountResponse, error)
	GetManyFriendCount(ctx context.Context, in *GetManyFriendCountRequest, opts ...grpc.CallOption) (*GetManyFriendCountResponse, error)
	FavorParty(ctx context.Context, in *FavorPartyRequest, opts ...grpc.CallOption) (*FavoriteParty, error)
	DefavorParty(ctx context.Context, in *FavorPartyRequest, opts ...grpc.CallOption) (*common.SuccessIndicator, error)
	GetFavoritePartiesByUser(ctx context.Context, in *GetFavoritePartiesByUserRequest, opts ...grpc.CallOption) (*PagedFavoriteParties, error)
	GetFavorisingUsersByParty(ctx context.Context, in *GetFavorisingUsersByPartyRequest, opts ...grpc.CallOption) (*PagedFavoriteParties, error)
}

type relationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationServiceClient(cc grpc.ClientConnInterface) RelationServiceClient {
	return &relationServiceClient{cc}
}

func (c *relationServiceClient) FriendRequest(ctx context.Context, in *FriendRequestRequest, opts ...grpc.CallOption) (*common.SuccessIndicator, error) {
	out := new(common.SuccessIndicator)
	err := c.cc.Invoke(ctx, "/relation.RelationService/FriendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) AcceptFriend(ctx context.Context, in *AcceptFriendRequest, opts ...grpc.CallOption) (*common.SuccessIndicator, error) {
	out := new(common.SuccessIndicator)
	err := c.cc.Invoke(ctx, "/relation.RelationService/AcceptFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) RemoveFriend(ctx context.Context, in *RemoveFriendRequest, opts ...grpc.CallOption) (*common.SuccessIndicator, error) {
	out := new(common.SuccessIndicator)
	err := c.cc.Invoke(ctx, "/relation.RelationService/RemoveFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetFriendRelation(ctx context.Context, in *GetFriendRelationRequest, opts ...grpc.CallOption) (*FriendRelation, error) {
	out := new(FriendRelation)
	err := c.cc.Invoke(ctx, "/relation.RelationService/GetFriendRelation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetFriendsOfUser(ctx context.Context, in *GetFriendsOfUserRequest, opts ...grpc.CallOption) (*PagedFriendRelations, error) {
	out := new(PagedFriendRelations)
	err := c.cc.Invoke(ctx, "/relation.RelationService/GetFriendsOfUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetFriendCount(ctx context.Context, in *GetFriendCountRequest, opts ...grpc.CallOption) (*GetFriendCountResponse, error) {
	out := new(GetFriendCountResponse)
	err := c.cc.Invoke(ctx, "/relation.RelationService/GetFriendCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetManyFriendCount(ctx context.Context, in *GetManyFriendCountRequest, opts ...grpc.CallOption) (*GetManyFriendCountResponse, error) {
	out := new(GetManyFriendCountResponse)
	err := c.cc.Invoke(ctx, "/relation.RelationService/GetManyFriendCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) FavorParty(ctx context.Context, in *FavorPartyRequest, opts ...grpc.CallOption) (*FavoriteParty, error) {
	out := new(FavoriteParty)
	err := c.cc.Invoke(ctx, "/relation.RelationService/FavorParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) DefavorParty(ctx context.Context, in *FavorPartyRequest, opts ...grpc.CallOption) (*common.SuccessIndicator, error) {
	out := new(common.SuccessIndicator)
	err := c.cc.Invoke(ctx, "/relation.RelationService/DefavorParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetFavoritePartiesByUser(ctx context.Context, in *GetFavoritePartiesByUserRequest, opts ...grpc.CallOption) (*PagedFavoriteParties, error) {
	out := new(PagedFavoriteParties)
	err := c.cc.Invoke(ctx, "/relation.RelationService/GetFavoritePartiesByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetFavorisingUsersByParty(ctx context.Context, in *GetFavorisingUsersByPartyRequest, opts ...grpc.CallOption) (*PagedFavoriteParties, error) {
	out := new(PagedFavoriteParties)
	err := c.cc.Invoke(ctx, "/relation.RelationService/GetFavorisingUsersByParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationServiceServer is the server API for RelationService service.
// All implementations must embed UnimplementedRelationServiceServer
// for forward compatibility
type RelationServiceServer interface {
	FriendRequest(context.Context, *FriendRequestRequest) (*common.SuccessIndicator, error)
	AcceptFriend(context.Context, *AcceptFriendRequest) (*common.SuccessIndicator, error)
	RemoveFriend(context.Context, *RemoveFriendRequest) (*common.SuccessIndicator, error)
	GetFriendRelation(context.Context, *GetFriendRelationRequest) (*FriendRelation, error)
	GetFriendsOfUser(context.Context, *GetFriendsOfUserRequest) (*PagedFriendRelations, error)
	GetFriendCount(context.Context, *GetFriendCountRequest) (*GetFriendCountResponse, error)
	GetManyFriendCount(context.Context, *GetManyFriendCountRequest) (*GetManyFriendCountResponse, error)
	FavorParty(context.Context, *FavorPartyRequest) (*FavoriteParty, error)
	DefavorParty(context.Context, *FavorPartyRequest) (*common.SuccessIndicator, error)
	GetFavoritePartiesByUser(context.Context, *GetFavoritePartiesByUserRequest) (*PagedFavoriteParties, error)
	GetFavorisingUsersByParty(context.Context, *GetFavorisingUsersByPartyRequest) (*PagedFavoriteParties, error)
	mustEmbedUnimplementedRelationServiceServer()
}

// UnimplementedRelationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRelationServiceServer struct {
}

func (UnimplementedRelationServiceServer) FriendRequest(context.Context, *FriendRequestRequest) (*common.SuccessIndicator, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FriendRequest not implemented")
}
func (UnimplementedRelationServiceServer) AcceptFriend(context.Context, *AcceptFriendRequest) (*common.SuccessIndicator, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptFriend not implemented")
}
func (UnimplementedRelationServiceServer) RemoveFriend(context.Context, *RemoveFriendRequest) (*common.SuccessIndicator, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFriend not implemented")
}
func (UnimplementedRelationServiceServer) GetFriendRelation(context.Context, *GetFriendRelationRequest) (*FriendRelation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendRelation not implemented")
}
func (UnimplementedRelationServiceServer) GetFriendsOfUser(context.Context, *GetFriendsOfUserRequest) (*PagedFriendRelations, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendsOfUser not implemented")
}
func (UnimplementedRelationServiceServer) GetFriendCount(context.Context, *GetFriendCountRequest) (*GetFriendCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendCount not implemented")
}
func (UnimplementedRelationServiceServer) GetManyFriendCount(context.Context, *GetManyFriendCountRequest) (*GetManyFriendCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManyFriendCount not implemented")
}
func (UnimplementedRelationServiceServer) FavorParty(context.Context, *FavorPartyRequest) (*FavoriteParty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavorParty not implemented")
}
func (UnimplementedRelationServiceServer) DefavorParty(context.Context, *FavorPartyRequest) (*common.SuccessIndicator, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DefavorParty not implemented")
}
func (UnimplementedRelationServiceServer) GetFavoritePartiesByUser(context.Context, *GetFavoritePartiesByUserRequest) (*PagedFavoriteParties, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavoritePartiesByUser not implemented")
}
func (UnimplementedRelationServiceServer) GetFavorisingUsersByParty(context.Context, *GetFavorisingUsersByPartyRequest) (*PagedFavoriteParties, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavorisingUsersByParty not implemented")
}
func (UnimplementedRelationServiceServer) mustEmbedUnimplementedRelationServiceServer() {}

// UnsafeRelationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelationServiceServer will
// result in compilation errors.
type UnsafeRelationServiceServer interface {
	mustEmbedUnimplementedRelationServiceServer()
}

func RegisterRelationServiceServer(s grpc.ServiceRegistrar, srv RelationServiceServer) {
	s.RegisterService(&RelationService_ServiceDesc, srv)
}

func _RelationService_FriendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).FriendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.RelationService/FriendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).FriendRequest(ctx, req.(*FriendRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_AcceptFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).AcceptFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.RelationService/AcceptFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).AcceptFriend(ctx, req.(*AcceptFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_RemoveFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).RemoveFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.RelationService/RemoveFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).RemoveFriend(ctx, req.(*RemoveFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetFriendRelation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendRelationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetFriendRelation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.RelationService/GetFriendRelation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetFriendRelation(ctx, req.(*GetFriendRelationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetFriendsOfUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendsOfUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetFriendsOfUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.RelationService/GetFriendsOfUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetFriendsOfUser(ctx, req.(*GetFriendsOfUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetFriendCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetFriendCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.RelationService/GetFriendCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetFriendCount(ctx, req.(*GetFriendCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetManyFriendCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManyFriendCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetManyFriendCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.RelationService/GetManyFriendCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetManyFriendCount(ctx, req.(*GetManyFriendCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_FavorParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavorPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).FavorParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.RelationService/FavorParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).FavorParty(ctx, req.(*FavorPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_DefavorParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavorPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).DefavorParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.RelationService/DefavorParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).DefavorParty(ctx, req.(*FavorPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetFavoritePartiesByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFavoritePartiesByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetFavoritePartiesByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.RelationService/GetFavoritePartiesByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetFavoritePartiesByUser(ctx, req.(*GetFavoritePartiesByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetFavorisingUsersByParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFavorisingUsersByPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetFavorisingUsersByParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relation.RelationService/GetFavorisingUsersByParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetFavorisingUsersByParty(ctx, req.(*GetFavorisingUsersByPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RelationService_ServiceDesc is the grpc.ServiceDesc for RelationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RelationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "relation.RelationService",
	HandlerType: (*RelationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FriendRequest",
			Handler:    _RelationService_FriendRequest_Handler,
		},
		{
			MethodName: "AcceptFriend",
			Handler:    _RelationService_AcceptFriend_Handler,
		},
		{
			MethodName: "RemoveFriend",
			Handler:    _RelationService_RemoveFriend_Handler,
		},
		{
			MethodName: "GetFriendRelation",
			Handler:    _RelationService_GetFriendRelation_Handler,
		},
		{
			MethodName: "GetFriendsOfUser",
			Handler:    _RelationService_GetFriendsOfUser_Handler,
		},
		{
			MethodName: "GetFriendCount",
			Handler:    _RelationService_GetFriendCount_Handler,
		},
		{
			MethodName: "GetManyFriendCount",
			Handler:    _RelationService_GetManyFriendCount_Handler,
		},
		{
			MethodName: "FavorParty",
			Handler:    _RelationService_FavorParty_Handler,
		},
		{
			MethodName: "DefavorParty",
			Handler:    _RelationService_DefavorParty_Handler,
		},
		{
			MethodName: "GetFavoritePartiesByUser",
			Handler:    _RelationService_GetFavoritePartiesByUser_Handler,
		},
		{
			MethodName: "GetFavorisingUsersByParty",
			Handler:    _RelationService_GetFavorisingUsersByParty_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relation/relation.proto",
}
