package server

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/onepanelio/core/api"
	v1 "github.com/onepanelio/core/pkg"
	"github.com/onepanelio/core/util"
	"google.golang.org/grpc/codes"
)

type SecretServer struct {
	kubeConfig *v1.Config
}

func NewSecretServer(kubeConfig *v1.Config) *NamespaceServer {
	return &NamespaceServer{kubeConfig: kubeConfig}
}

func apiSecret(s *v1.Secret) *api.Secret {
	return &api.Secret{
		Name: s.Name,
		Data: s.Data,
	}
}

func (s *SecretServer) CreateSecret(ctx context.Context, req *api.CreateSecretRequest) (*empty.Empty, error) {
	client, err := v1.NewClient(s.kubeConfig, "")
	if err != nil {
		return nil, util.NewUserError(codes.PermissionDenied, "Permission denied.")
	}

	err := client.CreateSecret(req.Namespace, &v1.Secret{
		Name: req.Secret.Name,
		Data: req.Secret.Data,
	})
	if errors.As(err, &userError) {
		return nil, userError.GRPCError()
	}
	return &empty.Empty{}, nil
}

func (s *SecretServer) SecretExists(ctx context.Context, req *api.SecretExistsRequest) (secretExists *api.SecretExistsResponse, err error) {
	client, err := v1.NewClient(s.kubeConfig, "")
	if err != nil {
		return nil, util.NewUserError(codes.PermissionDenied, "Permission denied.")
	}

	secretExistsBool, err := client.SecretExists(req.Namespace, req.Name)
	if errors.As(err, &userError) {
		return &api.SecretExistsResponse{
			Exists: false,
		}, userError.GRPCError()
	}
	return &api.SecretExistsResponse{
		Exists: secretExistsBool,
	}, nil
}

func (s *SecretServer) GetSecret(ctx context.Context, req *api.GetSecretRequest) (*api.Secret, error) {
	client, err := v1.NewClient(s.kubeConfig, "")
	if err != nil {
		return nil, util.NewUserError(codes.PermissionDenied, "Permission denied.")
	}

	secret, err := client.GetSecret(req.Namespace, req.Name)
	if errors.As(err, &userError) {
		return nil, userError.GRPCError()
	}
	return apiSecret(secret), nil
}

func (s *SecretServer) ListSecrets(ctx context.Context, req *api.ListSecretsRequest) (*api.ListSecretsResponse, error) {
	client, err := v1.NewClient(s.kubeConfig, "")
	if err != nil {
		return nil, util.NewUserError(codes.PermissionDenied, "Permission denied.")
	}

	secrets, err := client.ListSecrets(req.Namespace)
	if errors.As(err, &userError) {
		return nil, userError.GRPCError()
	}

	var apiSecrets []*api.Secret
	for _, secret := range secrets {
		apiSecrets = append(apiSecrets, apiSecret(secret))
	}

	return &api.ListSecretsResponse{
		Count:   int32(len(apiSecrets)),
		Secrets: apiSecrets,
	}, nil
}

func (s *SecretServer) DeleteSecret(ctx context.Context, req *api.DeleteSecretRequest) (deleted *api.DeleteSecretResponse, err error) {
	client, err := v1.NewClient(s.kubeConfig, "")
	if err != nil {
		return nil, util.NewUserError(codes.PermissionDenied, "Permission denied.")
	}

	isDeleted, err := client.DeleteSecret(req.Namespace, req.Name)
	if errors.As(err, &userError) {
		return &api.DeleteSecretResponse{
			Deleted: false,
		}, userError.GRPCError()
	}
	return &api.DeleteSecretResponse{
		Deleted: isDeleted,
	}, nil
}

func (s *SecretServer) DeleteSecretKey(ctx context.Context, req *api.DeleteSecretKeyRequest) (deleted *api.DeleteSecretKeyResponse, err error) {
	client, err := v1.NewClient(s.kubeConfig, "")
	if err != nil {
		return nil, util.NewUserError(codes.PermissionDenied, "Permission denied.")
	}

	secret := v1.Secret{
		Name: req.Secret.Name,
		Data: map[string]string{
			req.Key: "",
		},
	}
	isDeleted, err := client.DeleteSecretKey(req.Namespace, &secret)
	if err != nil {
		if errors.As(err, &userError) {
			return &api.DeleteSecretKeyResponse{
				Deleted: false,
			}, userError.GRPCError()
		}
	}
	return &api.DeleteSecretKeyResponse{
		Deleted: isDeleted,
	}, nil
}

func (s *SecretServer) AddSecretKeyValue(ctx context.Context, req *api.AddSecretKeyValueRequest) (updated *api.AddSecretKeyValueResponse, err error) {
	client, err := v1.NewClient(s.kubeConfig, "")
	if err != nil {
		return nil, util.NewUserError(codes.PermissionDenied, "Permission denied.")
	}

	secret := &v1.Secret{
		Name: req.Secret.Name,
		Data: req.Secret.Data,
	}
	isAdded, err := client.AddSecretKeyValue(req.Namespace, secret)
	if err != nil {
		if errors.As(err, &userError) {
			return &api.AddSecretKeyValueResponse{
				Inserted: false,
			}, userError.GRPCError()
		}
	}
	return &api.AddSecretKeyValueResponse{
		Inserted: isAdded,
	}, nil
}

func (s *SecretServer) UpdateSecretKeyValue(ctx context.Context, req *api.UpdateSecretKeyValueRequest) (updated *api.UpdateSecretKeyValueResponse, err error) {
	client, err := v1.NewClient(s.kubeConfig, "")
	if err != nil {
		return nil, util.NewUserError(codes.PermissionDenied, "Permission denied.")
	}

	secret := v1.Secret{
		Name: req.Secret.Name,
		Data: req.Secret.Data,
	}
	isUpdated, err := client.UpdateSecretKeyValue(req.Namespace, &secret)
	if errors.As(err, &userError) {
		return &api.UpdateSecretKeyValueResponse{
			Updated: false,
		}, userError.GRPCError()
	}
	return &api.UpdateSecretKeyValueResponse{
		Updated: isUpdated,
	}, nil
}
