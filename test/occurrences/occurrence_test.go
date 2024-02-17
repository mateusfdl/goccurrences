package occurrences_test

import (
	"context"
	"testing"

	occurrencesgrpc "buf.build/gen/go/matheusslima/go-poc/grpc/go/occurrences/v1/occurrencesv1grpc"
	occurrenceGrpc "buf.build/gen/go/matheusslima/go-poc/protocolbuffers/go/occurrences/v1"
	grpcServer "github.com/mateusfdl/go-poc/internal/grpc"
	"github.com/mateusfdl/go-poc/internal/logger"
	"github.com/mateusfdl/go-poc/internal/mongo"
	"github.com/mateusfdl/go-poc/internal/occurrences"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	TestApp = fx.New(
		occurrences.Module,
		grpcServer.Module,
		logger.Module,
		mongo.Module,
	)

	clientConn       *grpc.ClientConn
	OccurrenceClient occurrencesgrpc.OccurrenceServiceClient
)

func TestOccurrenceRPCHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Occurrence rpc call")
}

var _ = BeforeSuite(func() {
	Expect(TestApp.Start(context.Background())).To(Succeed())

	var dialErr error
	clientConn, dialErr = grpc.Dial(
		"localhost:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	Expect(dialErr).NotTo(HaveOccurred())

	OccurrenceClient = occurrencesgrpc.NewOccurrenceServiceClient(clientConn)
})

var _ = AfterSuite(func() {
	Expect(clientConn.Close()).To(Succeed())
	Expect(TestApp.Stop(context.Background())).To(Succeed())
})

var _ = Describe("grpc occurrence handler", func() {
	Describe("CreateOccurrence", func() {
		It("creates a occurrence", func() {
			userId := "65cbcd82f5cec8b2f2b1b29f"
			resp, err := OccurrenceClient.CreateOccurrence(
				context.Background(),
				&occurrenceGrpc.CreateOccurrenceRequest{
					OccurrenceCode: 0,
					OccurrenceTime: timestamppb.Now(),
					UserId:         userId,
				},
			)
			Expect(err).To(BeNil())

			Expect(resp.GetOccurrenceId()).To(Not(BeEmpty()))
		})
	})
})
