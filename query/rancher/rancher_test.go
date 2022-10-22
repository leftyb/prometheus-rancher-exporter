package rancher

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"reflect"
	"testing"
)

func TestClient_GetClusterConnectedState(t *testing.T) {
	type fields struct {
		Client dynamic.Interface
		Config *rest.Config
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Client{
				Client: tt.fields.Client,
				Config: tt.fields.Config,
			}
			got, err := r.GetClusterConnectedState()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClusterConnectedState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetClusterConnectedState() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetDownstreamClusterVersions(t *testing.T) {
	type fields struct {
		Client dynamic.Interface
		Config *rest.Config
	}
	tests := []struct {
		name    string
		fields  fields
		want    []clusterVersion
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Client{
				Client: tt.fields.Client,
				Config: tt.fields.Config,
			}
			got, err := r.GetDownstreamClusterVersions()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDownstreamClusterVersions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDownstreamClusterVersions() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetInstalledRancherVersion(t *testing.T) {
	type fields struct {
		Client dynamic.Interface
		Config *rest.Config
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Client{
				Client: tt.fields.Client,
				Config: tt.fields.Config,
			}
			got, err := r.GetInstalledRancherVersion()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInstalledRancherVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetInstalledRancherVersion() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetK8sDistributions(t *testing.T) {
	type fields struct {
		Client dynamic.Interface
		Config *rest.Config
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Client{
				Client: tt.fields.Client,
				Config: tt.fields.Config,
			}
			got, err := r.GetK8sDistributions()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetK8sDistributions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetK8sDistributions() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetLatestRancherVersion(t *testing.T) {
	type fields struct {
		Client dynamic.Interface
		Config *rest.Config
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Client{
				Client: tt.fields.Client,
				Config: tt.fields.Config,
			}
			got, err := r.GetLatestRancherVersion()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLatestRancherVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetLatestRancherVersion() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetNumberOfManagedClusters(t *testing.T) {
	type fields struct {
		Client dynamic.Interface
		Config *rest.Config
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Client{
				Client: tt.fields.Client,
				Config: tt.fields.Config,
			}
			got, err := r.GetNumberOfManagedClusters()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNumberOfManagedClusters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetNumberOfManagedClusters() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetNumberOfManagedNodes(t *testing.T) {
	type fields struct {
		Client dynamic.Interface
		Config *rest.Config
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Client{
				Client: tt.fields.Client,
				Config: tt.fields.Config,
			}
			got, err := r.GetNumberOfManagedNodes()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNumberOfManagedNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetNumberOfManagedNodes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetNumberOfTokens(t *testing.T) {
	type fields struct {
		Client dynamic.Interface
		Config *rest.Config
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Client{
				Client: tt.fields.Client,
				Config: tt.fields.Config,
			}
			got, err := r.GetNumberOfTokens()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNumberOfTokens() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetNumberOfTokens() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetNumberOfUsers(t *testing.T) {
	type fields struct {
		Client dynamic.Interface
		Config *rest.Config
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Client{
				Client: tt.fields.Client,
				Config: tt.fields.Config,
			}
			got, err := r.GetNumberOfUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNumberOfUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetNumberOfUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}