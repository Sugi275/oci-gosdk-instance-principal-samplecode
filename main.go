package main

import (
	"os"
	"context"
	"fmt"

	"github.com/oracle/oci-go-sdk/core"
	_ "github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/common/auth"
)

func main() {
	provider, _ := auth.InstancePrincipalConfigurationProvider()

	compartmentID := os.Getenv("COMPARTMENT_ID")

	request := core.ListInstancesRequest{
		CompartmentId: &compartmentID,
		LifecycleState: core.InstanceLifecycleStateRunning,
	}

	client, _ := core.NewComputeClientWithConfigurationProvider(provider)

	// Override the region, this is an optional step.
	// the InstancePrincipalsConfigurationProvider defaults to the region
	// in which the compute instance is currently running
	// client.SetRegion(string(common.RegionLHR))

	listInstancesResponse, err := client.ListInstances(context.Background(), request)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range listInstancesResponse.Items {
		fmt.Printf("list of Compute Instance: %s \n", *item.DisplayName)
	}
}