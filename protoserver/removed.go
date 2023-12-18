package protoserver

/*func RPCServer(listen []string, maxc int, mc *domain.Metchain, bc *blockchain.Blockchain) (*rpcServer, error) {
	//listen on the port
	lis, err := net.Listen("tcp", listen[0])

	if err != nil {
		logrpc.Criticalf("Failed to start server %v", err)
	}
	// create a new gRPC server
	newGRPCServer := newGRPCServer()
	// register the greet service
	rpcServers := &rpcServer{gRPCServer: *newGRPCServer, Metchain: mc, Blockchain: bc}
	pb.RegisterRPCServer(newGRPCServer.server, rpcServers)

	logrpc.Infof("Server started ats %v", lis.Addr())
	//list is the port, the grpc server needs to start there
	/*spawn("rpc.Serve", func() {
		if err := newGRPCServer.server.Serve(lis); err != nil {
			logrpc.Criticalf("Failed to start: %v", err)
		}
	})

	return rpcServers, nil
}*/
/*
func P2PServer(listen []string, mc *domain.Metchain, bc *blockchain.Blockchain) (*p2pServer, error) {
	//listen on the port

	lis, err := net.Listen("tcp", "localhost"+listen[0])
	if err != nil {
		logp2p.Criticalf("Failed to start server %v", err)
	}
	// create a new gRPC server
	newGRPCServer := newGRPCServer()
	// register the greet service
	p2pServers := &p2pServer{gRPCServer: *newGRPCServer, Metchain: mc, Blockchain: bc}
	pb.RegisterP2PServer(newGRPCServer.server, p2pServers)

	logp2p.Info("Server started at", lis.Addr())
	//list is the port, the grpc server needs to start there
	/*spawn("grpcServer.Serve", func() {
		if err := newGRPCServer.server.Serve(lis); err != nil {

			log.Errorf("Error starting p2p server:", err)
		}
	})

	if err != nil {
		logp2p.Criticalf("Failed to start: %v", err)
	}
	return p2pServers, nil
}*/
/*
// Connect connects to the RPC server with the given address
func Connect(address string) (*GRPCClient, error) {
	const dialTimeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), dialTimeout)
	defer cancel()

	gRPCConnection, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, errors.Wrapf(err, "error connecting to %s", address)
	}

	grpcClient := NewRPCClient(gRPCConnection)
	stream, err := grpcClient.MessageStream(context.Background(), grpc.UseCompressor(gzip.Name),
		grpc.MaxCallRecvMsgSize(RPCMaxMessageSize), grpc.MaxCallSendMsgSize(RPCMaxMessageSize))
	if err != nil {
		return nil, errors.Wrapf(err, "error getting client stream for %s", address)
	}
	return &GRPCClient{stream: stream, connection: gRPCConnection}, nil
}*/

/*func (s *rpcServer) MessageStream(stream pb.RPC_MessageStreamServer) error {

	//da := stream.Context()
	d := stream.Context()
	d.Done()
	req, err := stream.Recv()
	//req.GetPayload()
	if err != nil {
		return err
	}

	if req.GetGetBlockTemplateRequest() != nil {
		template := &pb.MetchainMessage{}
		template.Payload = req.Payload

		reqs := new(pb.MetchainMessage)
		payload := new(pb.MetchainMessage_GetBlockTemplateResponse)
		payload.GetBlockTemplateResponse = s.NewGetBlockTemplateResponse(template.GetGetBlockTemplateRequest())
		reqs.Payload = payload
		err := stream.Send(reqs)
		if err != nil {
			logrpc.Infof("Error Sending Block Info")
		}

	}
	if req.GetSubmitBlockRequest() != nil {
		template := &pb.MetchainMessage{}
		template.Payload = req.Payload

		reqs := new(pb.MetchainMessage)
		payloads := new(pb.MetchainMessage_SubmitBlockResponse)
		payloads.SubmitBlockResponse = s.GetBlockSubmitResponses(template.GetSubmitBlockRequest())
		reqs.Payload = payloads
		err := stream.Send(reqs)
		if err != nil {
			logrpc.Infof("Error Sending Block Submit Response")
		}

	}
	if req.GetNotifyNewBlockTemplateRequest() != nil {
		Server_GetNotifyNewBlockTemplateRequest(stream)
	}
	if req.GetGetBlockDagInfoRequest() != nil {
		s.Server_GetGetBlockDagInfoRequest(stream)
	}
	if req.GetSubmitSignedTXRequestMessage() != nil {

		template := &pb.MetchainMessage{}
		template.Payload = req.Payload

		reqs := new(pb.MetchainMessage)
		payloads := new(pb.MetchainMessage_SubmitSignedTXResponseMessage)
		payloads.SubmitSignedTXResponseMessage = s.GetSubmitSignedTX(template.GetSubmitSignedTXRequestMessage())
		reqs.Payload = payloads
		err := stream.Send(reqs)
		if err != nil {
			logrpc.Infof("Error Sending Block Submit Response")
		}
	}

	return nil
}

func (s *rpcServer) NewGetBlockTemplateResponse(response *pb.GetBlockTemplateRequestMessage) *pb.GetBlockTemplateResponseMessage {
	addr := response.PayAddress
	lb := blockchain.GetBlockTemplateBC(s.Metchain, addr)

	return &pb.GetBlockTemplateResponseMessage{
		Block: NewRPCBlock(lb),
	}
}
func NewRPCBlock(lb *blockchain.DomainBlock) *pb.RpcBlock {
	return &pb.RpcBlock{
		Header: NewRPCBlockHeader(lb),
	}
}

func NewRPCBlockHeader(lb *blockchain.DomainBlock) *pb.RpcBlockHeader {
	return &pb.RpcBlockHeader{
		Version:              1,
		HashMerkleRoot:       fmt.Sprintf("%x", lb.PreviousHash),
		Bits:                 uint32(lb.Bits),
		Parents:              NewRPCBlockLevelParents(lb),
		AcceptedIdMerkleRoot: fmt.Sprintf("%x", lb.PreviousHash),
		UtxoCommitment:       lb.UtxoCommitment,
		Timestamp:            lb.Timestamp,
		Nonce:                lb.Nonce,
	}
}
func NewRPCBlockLevelParents(lb *blockchain.DomainBlock) []*pb.RpcBlockLevelParents {
	parents := make([]*pb.RpcBlockLevelParents, 2)
	parents[0] = &pb.RpcBlockLevelParents{
		ParentHashes: []string{fmt.Sprintf("%x", lb.Megablock), fmt.Sprintf("%x", lb.Metblock)},
	}

	return parents
}

func (s *rpcServer) GetBlockSubmitResponses(block *pb.SubmitBlockRequestMessage) *pb.SubmitBlockResponseMessage {

	hash, err, reward := blockchain.CreateMiniBlock(block.Block, s.Db, s.Blockchain)
	ConesusHash := fmt.Sprintf("%x", hash)
	ConesusReward := fmt.Sprintf("%.2f", reward)
	if err != nil {
		return &pb.SubmitBlockResponseMessage{
			ConsensusBlockhash: ConesusHash,

			RejectReason: 1,
		}
	}
	return &pb.SubmitBlockResponseMessage{
		ConsensusBlockhash: ConesusHash,
		Consensusreward:    ConesusReward,
		RejectReason:       0,
	}
}

func (s *rpcServer) GetSubmitSignedTX(sender *pb.SubmitSignedTXRequestMessage) *pb.SubmitSignedTXResponseMessage {

	tx := s.AddTransactionRemoveNew(bytestostring(sender.SenderWallet), bytestostring(sender.ReciverWallet), sender.SendersAmount)
	status := "confirmed"
	if tx == [32]byte{} {
		status = "rejected"
	}
	return &pb.SubmitSignedTXResponseMessage{
		SenderWallet:  sender.SenderWallet,
		ReciverWallet: sender.ReciverWallet,
		TxHash:        stringtobyte(bytes32tostring(tx)),
		Status:        status,
	}
}
func stringtobyte(s string) []byte {
	return []byte(s)
}
func bytestostring(b []byte) string {
	return fmt.Sprintf("%s", b)
}
func bytes32tostring(b [32]byte) string {
	return fmt.Sprintf("%s", b)
}
*/
