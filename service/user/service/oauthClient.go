package service

import (
	"context"

	pb "muxi-workbench-user/proto"
	"muxi-workbench-user/util"
)

var (
	clientID     string
	clientSecret string

	clientStoreURL = ".../auth/api/oauth/store"

	domain string
)

func (s *UserService) GetClientID(ctx context.Context, req *pb.Request, res *pb.ClientResponse) error {
	res.Id = clientID
	return nil
}

// 注册 oauth 客户端
func RegisterClient() error {
	// domain := viper.GetString("domain")
	// if domain == "" {
	// 	// ...
	// }

	var bodyData = map[string]string{"domain": domain}
	body, err := util.SendHTTPRequest(clientStoreURL, "POST", &util.RequestData{
		BodyData:    bodyData,
		ContentType: util.JsonData,
	})
	if err != nil {
		return err
	}

	var client struct{ ClientID, ClientSecret string }
	if err := util.MarshalBodyForCustomData(body, &client); err != nil {
		return err
	}
	clientID = client.ClientID
	clientSecret = client.ClientSecret

	return nil
}
