package model

import (
	storageclient "github.com/Azure/azure-sdk-for-go/storage"
)

type ServiceInstance struct {
	Id                  string                            `json:"id"`
	DashboardUrl        string                            `json:"dashboard_url"`
	ResourceGroupName   string                            `json:"resource_group_name, omitempty"`
	StorageAccountName  string                            `json:"storage_account_name, omitempty"`
	ContainerAccessType storageclient.ContainerAccessType `json:"container_access_type, omitempty"`
	AcceptsIncomplete   string                            `json:"accepts_incomplete"`
	ServiceId           string                            `json:"service_id"`
	PlanId              string                            `json:"plan_id"`
	OrganizationGuid    string                            `json:"organization_guid"`
	SpaceGuid           string                            `json:"space_guid"`

	LastOperation *LastOperation `json:"last_operation, omitempty"`

	Parameters interface{} `json:"parameters, omitempty"`
}

type LastOperation struct {
	State                    string `json:"state"`
	Description              string `json:"description"`
	AsyncPollIntervalSeconds int    `json:"async_poll_interval_seconds, omitempty"`
}

type CreateServiceInstanceResponse struct {
	DashboardUrl  string         `json:"dashboard_url"`
	LastOperation *LastOperation `json:"last_operation, omitempty"`
}
