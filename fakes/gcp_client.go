package fakes

import compute "google.golang.org/api/compute/v1"

type GCPClient struct {
	GetProjectCall struct {
		CallCount int
		Receives  struct {
			ProjectID string
		}
		Returns struct {
			Project *compute.Project
			Error   error
		}
	}
	SetCommonInstanceMetadataCall struct {
		CallCount int
		Receives  struct {
			ProjectID string
			Metadata  *compute.Metadata
		}
		Returns struct {
			Operation *compute.Operation
			Error     error
		}
	}
}

func (g *GCPClient) GetProject(projectID string) (*compute.Project, error) {
	g.GetProjectCall.CallCount++
	g.GetProjectCall.Receives.ProjectID = projectID
	return g.GetProjectCall.Returns.Project, g.GetProjectCall.Returns.Error
}

func (g *GCPClient) SetCommonInstanceMetadata(projectID string, metadata *compute.Metadata) (*compute.Operation, error) {
	g.SetCommonInstanceMetadataCall.CallCount++
	g.SetCommonInstanceMetadataCall.Receives.ProjectID = projectID
	g.SetCommonInstanceMetadataCall.Receives.Metadata = metadata
	return g.SetCommonInstanceMetadataCall.Returns.Operation, g.SetCommonInstanceMetadataCall.Returns.Error
}