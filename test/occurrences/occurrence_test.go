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
	mongoClient *mongo.Mongo
	TestApp     = fx.New(
		occurrences.Module,
		grpcServer.Module,
		logger.Module,
		mongo.Module,
		fx.Populate(&mongoClient),
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
	Expect(mongoClient.DB.Client().Database("rewards-poc").Collection("occurrences").Drop(context.Background())).To(Succeed())
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

	Describe("ListUserOccurences", func() {
		It("returns all user occurrences", func() {
			userId := "65cbcd82f5cec8b2f2b1b29f"
			otherUserId := "65cbcd82f5cec8b2f2b1b28f"

			var occurreceIDs = make([]string, 10)

			for i := 0; i < 10; i++ {
				occurrenceID, err := OccurrenceClient.CreateOccurrence(
					context.Background(),
					&occurrenceGrpc.CreateOccurrenceRequest{
						OccurrenceCode: 0,
						OccurrenceTime: timestamppb.Now(),
						UserId:         userId,
					},
				)
				Expect(err).To(BeNil())
				occurreceIDs = append(occurreceIDs, occurrenceID.GetOccurrenceId())
			}

			otherOccurrenceID, err := OccurrenceClient.CreateOccurrence(
				context.Background(),
				&occurrenceGrpc.CreateOccurrenceRequest{
					OccurrenceCode: 0,
					OccurrenceTime: timestamppb.Now(),
					UserId:         otherUserId,
				},
			)
			Expect(err).To(BeNil())

			resp, err := OccurrenceClient.ListUserOccurrences(context.Background(), &occurrenceGrpc.ListUserOccurrencesRequest{
				UserId: userId,
				Limit:  10,
				Skip:   0,
			})
			Expect(err).To(BeNil())

			Expect(resp.GetOccurrences()).To(HaveLen(10))

			for _, occurrence := range resp.GetOccurrences() {
				Expect(occurreceIDs).To(ContainElement(occurrence.GetOccurrenceId()))
			}
			Expect(occurreceIDs).To(Not(ContainElement(otherOccurrenceID.GetOccurrenceId())))
		})
	})
})
