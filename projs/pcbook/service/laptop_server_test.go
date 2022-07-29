package service_test

import (
	"context"
	"grpbook/pb"
	"grpbook/sample"
	"grpbook/service"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServerCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopNoID := sample.NewLaptop()
	laptopNoID.Id = ""

	lpInvalidID := sample.NewLaptop()
	lpInvalidID.Id = "INVALID"

	lpDuplicateID := sample.NewLaptop()
	storeDuplicateID := service.InMemoryLaptopStore{}
	require.NoError(t, storeDuplicateID.Save(lpDuplicateID))

	tt := map[string]struct {
		laptop *pb.Laptop
		store  service.LaptopStore
		code   codes.Code
	}{
		"Success with ID": {
			laptop: sample.NewLaptop(),
			store:  &service.InMemoryLaptopStore{},
			code:   codes.OK,
		},
		"Success no ID": {
			laptop: laptopNoID,
			store:  &service.InMemoryLaptopStore{},
			code:   codes.OK,
		},
		"Failure invalid ID": {
			laptop: lpInvalidID,
			store:  &service.InMemoryLaptopStore{},
			code:   codes.InvalidArgument,
		},
		"Failure duplicate ID": {
			laptop: lpDuplicateID,
			store:  &storeDuplicateID,
			code:   codes.AlreadyExists,
		},
	}
	for name, tc := range tt {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			t.Log(tc.laptop.Id)
			res, err := service.LaptopServer{Store: tc.store}.
				CreateLaptop(context.Background(),
					&pb.CreateLaptopRequest{Laptop: tc.laptop})
			t.Log(tc.laptop.Id)
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
				if len(tc.laptop.Id) > 0 {
					require.Equal(t, tc.laptop.Id, res.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, st.Code())
			}
		})
	}
}
