package helpers

import (
	"github.com/astaxie/beego"

	"gitlab.com/jamalex-ltd/r-d-department/go-template/internal/app/types"
)

type metaDatabases interface {
	GetListMetaFromType(ctx beego.Controller) (res []*types.MetaData, err error)
}

type MetaDataHelper struct {
	metadataDB metaDatabases
}

func NewMetadataHelper(db metaDatabases) *MetaDataHelper {
	return &MetaDataHelper{
		db,
	}
}

func (s *MetaDataHelper) GetAll(ctx beego.Controller) (res []*types.MetaData, err error) {
	metadata, err := s.metadataDB.GetListMetaFromType(ctx)
	if err != nil {
		return nil, err
	}
	return metadata, nil
}
