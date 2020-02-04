// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package log

import (
	"context"

	"github.com/atomix/api/proto/atomix/headers"
	api "github.com/atomix/api/proto/atomix/log"
	"github.com/atomix/go-client/pkg/client/session"
	"google.golang.org/grpc"
)

type sessionHandler struct{}

func (m *sessionHandler) Create(ctx context.Context, s *session.Session) error {
	return s.DoCreate(ctx, func(ctx context.Context, conn *grpc.ClientConn, header *headers.RequestHeader) (*headers.ResponseHeader, interface{}, error) {
		request := &api.CreateRequest{
			Header:  header,
			Timeout: &s.Timeout,
		}
		client := api.NewLogServiceClient(conn)
		response, err := client.Create(ctx, request)
		if err != nil {
			return nil, nil, err
		}
		return response.Header, response, nil
	})
}

func (m *sessionHandler) KeepAlive(ctx context.Context, s *session.Session) error {
	return s.DoKeepAlive(ctx, func(ctx context.Context, conn *grpc.ClientConn, header *headers.RequestHeader) (*headers.ResponseHeader, interface{}, error) {
		request := &api.KeepAliveRequest{
			Header: header,
		}
		client := api.NewLogServiceClient(conn)
		response, err := client.KeepAlive(ctx, request)
		if err != nil {
			return nil, nil, err
		}
		return response.Header, response, nil
	})
}

func (m *sessionHandler) Close(ctx context.Context, s *session.Session) error {
	return s.DoClose(ctx, func(ctx context.Context, conn *grpc.ClientConn, header *headers.RequestHeader) (*headers.ResponseHeader, interface{}, error) {
		request := &api.CloseRequest{
			Header: header,
		}
		client := api.NewLogServiceClient(conn)
		response, err := client.Close(ctx, request)
		if err != nil {
			return nil, nil, err
		}
		return response.Header, response, nil
	})
}

func (m *sessionHandler) Delete(ctx context.Context, s *session.Session) error {
	return s.DoClose(ctx, func(ctx context.Context, conn *grpc.ClientConn, header *headers.RequestHeader) (*headers.ResponseHeader, interface{}, error) {
		request := &api.CloseRequest{
			Header: header,
			Delete: true,
		}
		client := api.NewLogServiceClient(conn)
		response, err := client.Close(ctx, request)
		if err != nil {
			return nil, nil, err
		}
		return response.Header, response, nil
	})
}
