package main

import (
	"context"
	"fmt"

	"github.com/tagur87/go-netbox/netbox"
)

func NetboxClient() (apiClient *netbox.APIClient) {
	configuration := netbox.NewConfiguration()
	configuration.Host = "demo.netbox.dev"
	// configuration.Debug = true
	configuration.UserAgent = "test-go-netbox-client/0.1"
	configuration.Scheme = "https"
	configuration.AddDefaultHeader(
		"Authorization",
		fmt.Sprintf("Token %s", "07e163ca3d44abec3407f488f7798e4744a7d1c8"))
	apiClient = netbox.NewAPIClient(configuration)
	return
}

func main() {
	ctx := context.Background()
	client := NetboxClient()
	resp, _, err := client.DcimAPI.DcimVirtualChassisList(ctx).Execute()
	// fmt.Println(test)
	// var resp []byte

	// inte, err := resp.Body.Read([]byte{})
	fmt.Println(resp.Results[0].Display)
	fmt.Println(err)
}
