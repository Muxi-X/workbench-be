package service

import (
	"context"

	"muxi-workbench-user/pkg/auth"
	pb "muxi-workbench-user/proto"
)

func (s *UserService) SetClient(ctx context.Context, req *pb.ClientRequest, res *pb.Response) error {
	auth.OauthManager.SetClient(req.Id, req.Secret)
	return nil
}

// type ClientResponse struct {
// 	util.BasicResponse
// 	Data ClientInfo `json:"data"`
// }

// type ClientInfo struct {
// 	ClientID     string `json:"client_id"`
// 	ClientSecret string `json:"client_secret"`
// }

// 注册 oauth 客户端
// func RegisterClient() error {
// 	// clientID = "4b194ad8-7d97-4dca-b078-6c3c65b31c75"
// 	// clientSecret = "8c066b19-e507-4887-88f3-7e7edd99bfd8"

// 	var bodyData = map[string]string{"domain": domain}
// 	body, err := util.SendHTTPRequest(clientStoreURL, "POST", &util.RequestData{
// 		BodyData:    bodyData,
// 		ContentType: util.JsonData,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	var rp ClientResponse
// 	if err := util.MarshalBodyForCustomData(body, &rp); err != nil {
// 		return err
// 	}
// 	if rp.Code != 0 {
// 		return errors.New(rp.Message)
// 	}

// 	clientID = rp.Data.ClientID
// 	clientSecret = rp.Data.ClientSecret

// 	if clientID == "" || clientSecret == "" {
// 		return errors.New("blank client_id or client_secret")
// 	}
// 	return nil
// }
