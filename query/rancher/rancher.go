package rancher

import (
	"context"
	"fmt"
	"github.com/david-vtuk/prometheus-rancher-exporter/semver"
	"github.com/tidwall/gjson"
	"io"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"net/http"
)

var (
	settingGVR  = schema.GroupVersionResource{Group: "management.cattle.io", Version: "v3", Resource: "settings"}
	clustersGVR = schema.GroupVersionResource{Group: "management.cattle.io", Version: "v3", Resource: "clusters"}
	nodeGVR     = schema.GroupVersionResource{Group: "management.cattle.io", Version: "v3", Resource: "nodes"}
	clusterGVR  = schema.GroupVersionResource{Group: "management.cattle.io", Version: "v3", Resource: "cluster"}
)

type Client struct {
	Client dynamic.Interface
	Config *rest.Config
}

func (r Client) GetRancherVersion() (map[string]int64, error) {

	res, err := r.Client.Resource(settingGVR).Get(context.Background(), "server-version", v1.GetOptions{})
	if err != nil {
		return nil, err
	}

	version, _, _ := unstructured.NestedString(res.UnstructuredContent(), "value")

	result, err := semver.Parse(TrimVersionChar(version))

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r Client) GetNumberOfManagedClusters() (int, error) {

	res, err := r.Client.Resource(clustersGVR).List(context.Background(), v1.ListOptions{})
	if err != nil {
		return 0, err
	}

	return len(res.Items) - 1, nil
}

func (r Client) GetK8sDistributions() (map[string]int, error) {

	distributions := make(map[string]int)

	res, err := r.Client.Resource(clustersGVR).List(context.Background(), v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, v := range res.Items {
		labels := v.GetLabels()
		distribution := labels["provider.cattle.io"]
		distributions[distribution] += 1
	}

	return distributions, nil
}

func (r Client) GetLatestRancherVersion() (map[string]int64, error) {
	resp, err := http.Get("https://api.github.com/repos/rancher/rancher/releases/latest")

	fmt.Println(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	val := gjson.Get(string(bodyBytes), "tag_name")

	result, err := semver.Parse(TrimVersionChar(val.String()))

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r Client) GetNumberOfManagedNodes() (int, error) {

	res, err := r.Client.Resource(nodeGVR).List(context.Background(), v1.ListOptions{})
	if err != nil {
		return 0, err
	}

	return len(res.Items), nil
}

func (r Client) GetClusterConnectedState() (map[string]bool, error) {
	clusterStatus := make(map[string]bool)

	res, err := r.Client.Resource(clustersGVR).List(context.Background(), v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	// Iterate through each cluster management object
	for _, cluster := range res.Items {

		// Grab Cluster name
		clusterName, _, err := unstructured.NestedString(cluster.Object, "spec", "displayName")

		if err != nil {
			return nil, err
		}

		// Ignore if it's the "Local" cluster
		if clusterName != "local" {

			clusterStatus[clusterName] = false

			// Grab status.conditions slide from object
			statusSlice, _, _ := unstructured.NestedSlice(cluster.Object, "status", "conditions")

			// Iterate through each status slice to determine if cluster is connected
			for _, value := range statusSlice {

				// Determine whether we find both conditions in each status message
				// We're looking for both when type == connected and status == true
				// to identify if a cluster is connected to this Rancher instance

				// Reset values when iterating through
				foundStatus := false
				foundType := false

				for k, v := range value.(map[string]interface{}) {

					if k == "type" && v.(string) == "Connected" {
						foundType = true
					}

					if k == "status" && v.(string) == "True" {
						foundStatus = true
					}

					if foundStatus == true && foundType == true {
						clusterStatus[clusterName] = true
					}

				}
			}

		}
	}

	return clusterStatus, nil
}

// Version returned from CRD is in the format of "vN.N.N", trim the leading "v"
func TrimVersionChar(version string) string {
	for i := range version {
		if i > 0 {
			return version[i:]
		}
	}
	return version[:0]
}

/*
func (r Client) GetK8sDistributions() (map[string]int, error) {

	distributions := make(map[string]int)

	res, err := r.Client.Resource(clustersGVR).List(context.Background(), v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, v := range res.Items {
		labels := v.GetLabels()
		distribution := labels["provider.cattle.io"]
		distributions[distribution] += 1
	}

	return distributions, nil
}
*/
