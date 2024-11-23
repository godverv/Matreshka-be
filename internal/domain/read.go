package domain

import (
	"github.com/Red-Sock/evon"

	api "github.com/godverv/matreshka-be/pkg/matreshka_be_api"
)

type ListConfigsRequest struct {
	Paging Paging
	Sort   Sort

	SearchPattern string
}

type ListConfigsResponse struct {
	List         []ConfigListItem
	TotalRecords uint32
}

type ConfigListItem struct {
	Name    string
	Version string
}

type ConfigDescription struct {
	Id          int64
	ServiceName string
}

type ConfigEnvVals struct {
	ConfigDescription
	Nodes *evon.Node
}

type Paging struct {
	Limit  uint32
	Offset uint32
}

type Sort struct {
	SortType api.Sort_Type
	Desc     bool
}
