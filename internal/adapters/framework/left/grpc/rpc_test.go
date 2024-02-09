package grpc

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	googlegrpc "google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"hexarch/internal/adapters/app/api"
	"hexarch/internal/adapters/core/arithmetic"
	"hexarch/internal/adapters/framework/left/grpc/pb"
	"hexarch/internal/adapters/framework/right/db"
	"hexarch/internal/ports"
	"hexarch/utils"
	"net"
	"os"
	"testing"
)

const bufSize = 1024 * 1024

var (
	lis         *bufconn.Listener
	dbCloseFunc func()
)

func setup() {
	log.SetLevel(log.DebugLevel)
	var err error
	lis = bufconn.Listen(bufSize)
	grpcServer := googlegrpc.NewServer()

	// ports
	var dbAdapter ports.DBPort
	var gRPCAdapter ports.GRPCPort
	var appAdapter ports.APIPort
	var coreAdapter ports.ArithmeticPort

	// adapters
	dbDriver := os.Getenv("TEST_DB_DRIVER")
	dbSourceName := os.Getenv("TEST_DB_NAME")

	dbAdapter, err = db.NewAdapter(dbDriver, dbSourceName)
	if err != nil {
		log.WithField("func", utils.GetCallerInfo()).Fatalf("db ping failure: %v", err)
	}
	dbCloseFunc = dbAdapter.CloseDBConnection

	coreAdapter = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbAdapter, coreAdapter)
	gRPCAdapter = NewAdapter(appAdapter)

	pb.RegisterArithmeticServiceServer(grpcServer, gRPCAdapter)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.WithField("func", utils.GetCallerInfo()).Fatalf("test gRPC server failure: %v", err)
		}
	}()
}

func shutdown() {
	dbCloseFunc()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func bufDialer(ctx context.Context, _ string) (net.Conn, error) {
	return lis.DialContext(ctx)
}

func GetGRPCConnection(ctx context.Context, t *testing.T) *googlegrpc.ClientConn {
	conn, err := googlegrpc.DialContext(ctx, "bufnet", googlegrpc.WithContextDialer(bufDialer), googlegrpc.WithInsecure())
	if err != nil {
		t.Fatalf("dail bufnet failure: %v", err)
	}
	return conn
}

func TestGetAddition(t *testing.T) {
	ctx := context.Background()
	conn := GetGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	params := &pb.OperationParameters{A: 1, B: 1}

	answer, err := client.GetAddition(ctx, params)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer.Value, int32(2))
}

func TestGetSubtraction(t *testing.T) {
	ctx := context.Background()
	conn := GetGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	params := &pb.OperationParameters{A: 2, B: 1}

	answer, err := client.GetSubtraction(ctx, params)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer.Value, int32(1))
}

func TestGetMultiplication(t *testing.T) {
	ctx := context.Background()
	conn := GetGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	params := &pb.OperationParameters{A: 2, B: 2}

	answer, err := client.GetMultiplication(ctx, params)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer.Value, int32(4))
}

func TestGetDivision(t *testing.T) {
	ctx := context.Background()
	conn := GetGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{A: 9, B: 3}

	answer, err := client.GetDivision(ctx, params)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer.Value, int32(3))
}

func TestGetDivision_DivideByZero(t *testing.T) {
	ctx := context.Background()
	conn := GetGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{A: 9, B: 0}

	answer, err := client.GetDivision(ctx, params)
	if err == nil {
		t.Fatalf("expected: %v, got: %v", err, nil)
	}

	require.NotNil(t, err)
	require.Empty(t, answer)
	require.Equal(t, err.Error(), "rpc error: code = InvalidArgument desc = missing required")
}
