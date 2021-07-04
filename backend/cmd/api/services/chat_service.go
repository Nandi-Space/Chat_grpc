package services

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/pborman/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"backend/cmd/api/pb"
	"backend/cmd/model"
)

type chatServer struct {
	pb.UnimplementedChatServiceServer
	channel map[string][]chan *pb.Message
}

func NewChatServer() *chatServer {
	return &chatServer{
		channel: make(map[string][]chan *pb.Message),
	}
}
func (s *chatServer) JoinChannel(ch *pb.Channel, msgStream pb.ChatService_JoinChannelServer) error {

	msgChannel := make(chan *pb.Message)
	s.channel[ch.Name] = append(s.channel[ch.Name], msgChannel)

	// doing this never closes the stream
	for {
		select {
		case <-msgStream.Context().Done():
			return nil
		case msg := <-msgChannel:

			fmt.Printf("GO ROUTINE (got message): %v \n", msg)
			msgStream.Send(msg)
		}
	}
}

func (s *chatServer) SendMessage(stream pb.ChatService_SendMessageServer) error {
	msg, err := stream.Recv()

	chat := &model.Chat{
		ID:         uuid.New(),
		SenderID:   msg.SenderID,
		SenderName: msg.Channel.SendersName,
		Msg:        msg.Message,
		Meta:       msg.Meta,
		File:       msg.File,
		InboxHash:  msg.Channel.Name,
		CreatedAt:  time.Now().Unix(),
	}

	if err := chat.Insert(); err != nil {
		return err
	}

	if err == io.EOF {
		return nil
	}

	if err != nil {
		return err
	}

	ack := pb.MessageAck{Status: "SENT"}
	stream.SendAndClose(&ack)

	go func() {
		streams := s.channel[msg.Channel.Name]
		for _, msgChan := range streams {
			msgChan <- msg
		}
	}()
	return nil
}

func (c *chatServer) DeleteChat(ctx context.Context, req *pb.DeleteRequest) (*pb.Empty, error) {
	chat := &model.Chat{
		ID:        req.ChatID,
		DeletedBy: req.DeletedBy,
	}
	err := chat.Delete()
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot create user %s", err)
	}
	return &pb.Empty{}, nil
}

func (c *chatServer) GetAllChats(ctx context.Context, req *pb.Channel) (*pb.ListOfChats, error) {

	chat := model.Chat{
		InboxHash: req.Name,
	}
	cl := model.ChatList{}
	if err := chat.Fetch(&cl); err != nil {
		return nil, err
	}
	messaegs := []*pb.Message{}

	for _, c := range cl {
		messaegs = append(
			messaegs,
			&pb.Message{
				SenderID: c.SenderID,
				Channel: &pb.Channel{
					Name: c.SenderName},
				Message: c.Msg,
				Meta:    c.Meta,
				File: c.File})
	}

	return &pb.ListOfChats{
		Messages: messaegs,
	}, nil

}
