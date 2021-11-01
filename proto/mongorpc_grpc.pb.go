// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MongoRPCClient is the client API for MongoRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MongoRPCClient interface {
	// ListCollections lists the collections in a database.
	ListCollections(ctx context.Context, in *ListCollectionsRequest, opts ...grpc.CallOption) (*ListCollectionsResponse, error)
	// GetDocument gets a document from a collection.
	GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*GetDocumentResponse, error)
	// ListDocuments lists the documents in a collection.
	ListDocuments(ctx context.Context, in *ListDocumentsRequest, opts ...grpc.CallOption) (*ListDocumentsResponse, error)
	// CreateDocument creates a document in a collection.
	CreateDocument(ctx context.Context, in *CreateDocumentRequest, opts ...grpc.CallOption) (*CreateDocumentResponse, error)
	// UpdateDocument updates a document in a collection.
	UpdateDocument(ctx context.Context, in *UpdateDocumentRequest, opts ...grpc.CallOption) (*UpdateDocumentResponse, error)
	// DeleteDocument deletes a document from a collection.
	DeleteDocument(ctx context.Context, in *DeleteDocumentRequest, opts ...grpc.CallOption) (*DeleteDocumentResponse, error)
	// Returns the count of documents that match the query for a collection or view.
	CountDocuments(ctx context.Context, in *ListDocumentsRequest, opts ...grpc.CallOption) (*CountDocumentsResponse, error)
	// Listen listens for changes to a document in a collection.
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (MongoRPC_ListenClient, error)
	// Creates indexes on collections.
	CreateIndex(ctx context.Context, in *CreateIndexRequest, opts ...grpc.CallOption) (*CreateIndexResponse, error)
	// Lists indexes on collections.
	GetIndexes(ctx context.Context, in *GetIndexesRequest, opts ...grpc.CallOption) (*GetIndexesResponse, error)
	// Deletes indexes on collections.
	DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...grpc.CallOption) (*DeleteIndexResponse, error)
	// Reindexes indexes on collections.
	Reindex(ctx context.Context, in *ReindexRequest, opts ...grpc.CallOption) (*ReindexResponse, error)
	// Ping is used to test the connection to the server.
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// CollectionStats returns stats about a collection.
	CollectionStats(ctx context.Context, in *ListCollectionsRequest, opts ...grpc.CallOption) (*ListCollectionsResponse, error)
	// CreateCollection creates a collection.
	CreateCollection(ctx context.Context, in *CreateCollectionRequest, opts ...grpc.CallOption) (*CreateCollectionResponse, error)
	// RenameCollection renames a collection.
	RenameCollection(ctx context.Context, in *RenameCollectionRequest, opts ...grpc.CallOption) (*RenameCollectionResponse, error)
	// DeleteCollection drops a collection.
	DeleteCollection(ctx context.Context, in *DeleteCollectionRequest, opts ...grpc.CallOption) (*DeleteCollectionResponse, error)
}

type mongoRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewMongoRPCClient(cc grpc.ClientConnInterface) MongoRPCClient {
	return &mongoRPCClient{cc}
}

func (c *mongoRPCClient) ListCollections(ctx context.Context, in *ListCollectionsRequest, opts ...grpc.CallOption) (*ListCollectionsResponse, error) {
	out := new(ListCollectionsResponse)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/ListCollections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoRPCClient) GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*GetDocumentResponse, error) {
	out := new(GetDocumentResponse)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/GetDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoRPCClient) ListDocuments(ctx context.Context, in *ListDocumentsRequest, opts ...grpc.CallOption) (*ListDocumentsResponse, error) {
	out := new(ListDocumentsResponse)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/ListDocuments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoRPCClient) CreateDocument(ctx context.Context, in *CreateDocumentRequest, opts ...grpc.CallOption) (*CreateDocumentResponse, error) {
	out := new(CreateDocumentResponse)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/CreateDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoRPCClient) UpdateDocument(ctx context.Context, in *UpdateDocumentRequest, opts ...grpc.CallOption) (*UpdateDocumentResponse, error) {
	out := new(UpdateDocumentResponse)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/UpdateDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoRPCClient) DeleteDocument(ctx context.Context, in *DeleteDocumentRequest, opts ...grpc.CallOption) (*DeleteDocumentResponse, error) {
	out := new(DeleteDocumentResponse)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/DeleteDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoRPCClient) CountDocuments(ctx context.Context, in *ListDocumentsRequest, opts ...grpc.CallOption) (*CountDocumentsResponse, error) {
	out := new(CountDocumentsResponse)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/CountDocuments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoRPCClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (MongoRPC_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &MongoRPC_ServiceDesc.Streams[0], "/mongorpc.MongoRPC/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &mongoRPCListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MongoRPC_ListenClient interface {
	Recv() (*ListenResponse, error)
	grpc.ClientStream
}

type mongoRPCListenClient struct {
	grpc.ClientStream
}

func (x *mongoRPCListenClient) Recv() (*ListenResponse, error) {
	m := new(ListenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mongoRPCClient) CreateIndex(ctx context.Context, in *CreateIndexRequest, opts ...grpc.CallOption) (*CreateIndexResponse, error) {
	out := new(CreateIndexResponse)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/CreateIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoRPCClient) GetIndexes(ctx context.Context, in *GetIndexesRequest, opts ...grpc.CallOption) (*GetIndexesResponse, error) {
	out := new(GetIndexesResponse)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/GetIndexes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoRPCClient) DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...grpc.CallOption) (*DeleteIndexResponse, error) {
	out := new(DeleteIndexResponse)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/DeleteIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoRPCClient) Reindex(ctx context.Context, in *ReindexRequest, opts ...grpc.CallOption) (*ReindexResponse, error) {
	out := new(ReindexResponse)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/Reindex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoRPCClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoRPCClient) CollectionStats(ctx context.Context, in *ListCollectionsRequest, opts ...grpc.CallOption) (*ListCollectionsResponse, error) {
	out := new(ListCollectionsResponse)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/CollectionStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoRPCClient) CreateCollection(ctx context.Context, in *CreateCollectionRequest, opts ...grpc.CallOption) (*CreateCollectionResponse, error) {
	out := new(CreateCollectionResponse)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/CreateCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoRPCClient) RenameCollection(ctx context.Context, in *RenameCollectionRequest, opts ...grpc.CallOption) (*RenameCollectionResponse, error) {
	out := new(RenameCollectionResponse)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/RenameCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoRPCClient) DeleteCollection(ctx context.Context, in *DeleteCollectionRequest, opts ...grpc.CallOption) (*DeleteCollectionResponse, error) {
	out := new(DeleteCollectionResponse)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/DeleteCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MongoRPCServer is the server API for MongoRPC service.
// All implementations must embed UnimplementedMongoRPCServer
// for forward compatibility
type MongoRPCServer interface {
	// ListCollections lists the collections in a database.
	ListCollections(context.Context, *ListCollectionsRequest) (*ListCollectionsResponse, error)
	// GetDocument gets a document from a collection.
	GetDocument(context.Context, *GetDocumentRequest) (*GetDocumentResponse, error)
	// ListDocuments lists the documents in a collection.
	ListDocuments(context.Context, *ListDocumentsRequest) (*ListDocumentsResponse, error)
	// CreateDocument creates a document in a collection.
	CreateDocument(context.Context, *CreateDocumentRequest) (*CreateDocumentResponse, error)
	// UpdateDocument updates a document in a collection.
	UpdateDocument(context.Context, *UpdateDocumentRequest) (*UpdateDocumentResponse, error)
	// DeleteDocument deletes a document from a collection.
	DeleteDocument(context.Context, *DeleteDocumentRequest) (*DeleteDocumentResponse, error)
	// Returns the count of documents that match the query for a collection or view.
	CountDocuments(context.Context, *ListDocumentsRequest) (*CountDocumentsResponse, error)
	// Listen listens for changes to a document in a collection.
	Listen(*ListenRequest, MongoRPC_ListenServer) error
	// Creates indexes on collections.
	CreateIndex(context.Context, *CreateIndexRequest) (*CreateIndexResponse, error)
	// Lists indexes on collections.
	GetIndexes(context.Context, *GetIndexesRequest) (*GetIndexesResponse, error)
	// Deletes indexes on collections.
	DeleteIndex(context.Context, *DeleteIndexRequest) (*DeleteIndexResponse, error)
	// Reindexes indexes on collections.
	Reindex(context.Context, *ReindexRequest) (*ReindexResponse, error)
	// Ping is used to test the connection to the server.
	Ping(context.Context, *Empty) (*Empty, error)
	// CollectionStats returns stats about a collection.
	CollectionStats(context.Context, *ListCollectionsRequest) (*ListCollectionsResponse, error)
	// CreateCollection creates a collection.
	CreateCollection(context.Context, *CreateCollectionRequest) (*CreateCollectionResponse, error)
	// RenameCollection renames a collection.
	RenameCollection(context.Context, *RenameCollectionRequest) (*RenameCollectionResponse, error)
	// DeleteCollection drops a collection.
	DeleteCollection(context.Context, *DeleteCollectionRequest) (*DeleteCollectionResponse, error)
	mustEmbedUnimplementedMongoRPCServer()
}

// UnimplementedMongoRPCServer must be embedded to have forward compatible implementations.
type UnimplementedMongoRPCServer struct {
}

func (UnimplementedMongoRPCServer) ListCollections(context.Context, *ListCollectionsRequest) (*ListCollectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCollections not implemented")
}
func (UnimplementedMongoRPCServer) GetDocument(context.Context, *GetDocumentRequest) (*GetDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDocument not implemented")
}
func (UnimplementedMongoRPCServer) ListDocuments(context.Context, *ListDocumentsRequest) (*ListDocumentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDocuments not implemented")
}
func (UnimplementedMongoRPCServer) CreateDocument(context.Context, *CreateDocumentRequest) (*CreateDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDocument not implemented")
}
func (UnimplementedMongoRPCServer) UpdateDocument(context.Context, *UpdateDocumentRequest) (*UpdateDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDocument not implemented")
}
func (UnimplementedMongoRPCServer) DeleteDocument(context.Context, *DeleteDocumentRequest) (*DeleteDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDocument not implemented")
}
func (UnimplementedMongoRPCServer) CountDocuments(context.Context, *ListDocumentsRequest) (*CountDocumentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountDocuments not implemented")
}
func (UnimplementedMongoRPCServer) Listen(*ListenRequest, MongoRPC_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedMongoRPCServer) CreateIndex(context.Context, *CreateIndexRequest) (*CreateIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIndex not implemented")
}
func (UnimplementedMongoRPCServer) GetIndexes(context.Context, *GetIndexesRequest) (*GetIndexesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndexes not implemented")
}
func (UnimplementedMongoRPCServer) DeleteIndex(context.Context, *DeleteIndexRequest) (*DeleteIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIndex not implemented")
}
func (UnimplementedMongoRPCServer) Reindex(context.Context, *ReindexRequest) (*ReindexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reindex not implemented")
}
func (UnimplementedMongoRPCServer) Ping(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedMongoRPCServer) CollectionStats(context.Context, *ListCollectionsRequest) (*ListCollectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectionStats not implemented")
}
func (UnimplementedMongoRPCServer) CreateCollection(context.Context, *CreateCollectionRequest) (*CreateCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCollection not implemented")
}
func (UnimplementedMongoRPCServer) RenameCollection(context.Context, *RenameCollectionRequest) (*RenameCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameCollection not implemented")
}
func (UnimplementedMongoRPCServer) DeleteCollection(context.Context, *DeleteCollectionRequest) (*DeleteCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCollection not implemented")
}
func (UnimplementedMongoRPCServer) mustEmbedUnimplementedMongoRPCServer() {}

// UnsafeMongoRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MongoRPCServer will
// result in compilation errors.
type UnsafeMongoRPCServer interface {
	mustEmbedUnimplementedMongoRPCServer()
}

func RegisterMongoRPCServer(s grpc.ServiceRegistrar, srv MongoRPCServer) {
	s.RegisterService(&MongoRPC_ServiceDesc, srv)
}

func _MongoRPC_ListCollections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCollectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).ListCollections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/ListCollections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).ListCollections(ctx, req.(*ListCollectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoRPC_GetDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).GetDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/GetDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).GetDocument(ctx, req.(*GetDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoRPC_ListDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDocumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).ListDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/ListDocuments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).ListDocuments(ctx, req.(*ListDocumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoRPC_CreateDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).CreateDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/CreateDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).CreateDocument(ctx, req.(*CreateDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoRPC_UpdateDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).UpdateDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/UpdateDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).UpdateDocument(ctx, req.(*UpdateDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoRPC_DeleteDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).DeleteDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/DeleteDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).DeleteDocument(ctx, req.(*DeleteDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoRPC_CountDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDocumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).CountDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/CountDocuments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).CountDocuments(ctx, req.(*ListDocumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoRPC_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MongoRPCServer).Listen(m, &mongoRPCListenServer{stream})
}

type MongoRPC_ListenServer interface {
	Send(*ListenResponse) error
	grpc.ServerStream
}

type mongoRPCListenServer struct {
	grpc.ServerStream
}

func (x *mongoRPCListenServer) Send(m *ListenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MongoRPC_CreateIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).CreateIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/CreateIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).CreateIndex(ctx, req.(*CreateIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoRPC_GetIndexes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIndexesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).GetIndexes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/GetIndexes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).GetIndexes(ctx, req.(*GetIndexesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoRPC_DeleteIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).DeleteIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/DeleteIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).DeleteIndex(ctx, req.(*DeleteIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoRPC_Reindex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReindexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).Reindex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/Reindex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).Reindex(ctx, req.(*ReindexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoRPC_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoRPC_CollectionStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCollectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).CollectionStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/CollectionStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).CollectionStats(ctx, req.(*ListCollectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoRPC_CreateCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).CreateCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/CreateCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).CreateCollection(ctx, req.(*CreateCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoRPC_RenameCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).RenameCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/RenameCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).RenameCollection(ctx, req.(*RenameCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoRPC_DeleteCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).DeleteCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/DeleteCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).DeleteCollection(ctx, req.(*DeleteCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MongoRPC_ServiceDesc is the grpc.ServiceDesc for MongoRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MongoRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mongorpc.MongoRPC",
	HandlerType: (*MongoRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCollections",
			Handler:    _MongoRPC_ListCollections_Handler,
		},
		{
			MethodName: "GetDocument",
			Handler:    _MongoRPC_GetDocument_Handler,
		},
		{
			MethodName: "ListDocuments",
			Handler:    _MongoRPC_ListDocuments_Handler,
		},
		{
			MethodName: "CreateDocument",
			Handler:    _MongoRPC_CreateDocument_Handler,
		},
		{
			MethodName: "UpdateDocument",
			Handler:    _MongoRPC_UpdateDocument_Handler,
		},
		{
			MethodName: "DeleteDocument",
			Handler:    _MongoRPC_DeleteDocument_Handler,
		},
		{
			MethodName: "CountDocuments",
			Handler:    _MongoRPC_CountDocuments_Handler,
		},
		{
			MethodName: "CreateIndex",
			Handler:    _MongoRPC_CreateIndex_Handler,
		},
		{
			MethodName: "GetIndexes",
			Handler:    _MongoRPC_GetIndexes_Handler,
		},
		{
			MethodName: "DeleteIndex",
			Handler:    _MongoRPC_DeleteIndex_Handler,
		},
		{
			MethodName: "Reindex",
			Handler:    _MongoRPC_Reindex_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _MongoRPC_Ping_Handler,
		},
		{
			MethodName: "CollectionStats",
			Handler:    _MongoRPC_CollectionStats_Handler,
		},
		{
			MethodName: "CreateCollection",
			Handler:    _MongoRPC_CreateCollection_Handler,
		},
		{
			MethodName: "RenameCollection",
			Handler:    _MongoRPC_RenameCollection_Handler,
		},
		{
			MethodName: "DeleteCollection",
			Handler:    _MongoRPC_DeleteCollection_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _MongoRPC_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mongorpc.proto",
}
