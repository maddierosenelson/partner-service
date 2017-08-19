package service

import (
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"jaxf-github.fanatics.corp/apparel/partner-service/pkg/db"
)

func New(logger log.Logger, q db.PartnerServiceQuerier) PartnerService {
	var svc PartnerService
	{
		svc = NewPartnerService(q)
		svc = LoggingMiddleware(logger)(svc)
	}
	return svc
}

type PartnerService interface {
	GetPartnerDataByKeyValue(ctx context.Context, key, value, group string) (int32, string, map[string]string, error)
	GetDataById(ctx context.Context, partnerId int32, partnerCode string, group string) (int32, string, map[string]string, error)
}

// NewPartnerService returns a struct that fulfills the PartnerService interface.
func NewPartnerService(q db.PartnerServiceQuerier) PartnerService {
	return partnerService{
		querier: q,
	}
}

// partnerService has a querier on it so we can easily access db queries in our service functions.
type partnerService struct {
	querier db.PartnerServiceQuerier
}

func (s partnerService) GetPartnerDataByKeyValue(_ context.Context, key, value, group string) (int32, string, map[string]string, error) { //Attribute should be array?
	attributes := make(map[string]string)
	if key == "" {
		return 0, "", attributes, errors.New("key cannot be empty")
	}
	if value == "" {
		return 0, "", attributes, errors.New("value cannot be empty")
	}
	id, code, err := s.querier.FindPartnerDataFromKeyValue(key, value)
	if err != nil {
		return 0, "", attributes, errors.New(fmt.Sprintf("could not find Id or Code from key: %s and value: %s", key, value))
	}
	//If a group is given to the GetPartnerDataByKeyValue function return only the partner attributes for that group.
	if group == "" {
		attributes, err = s.querier.FindAllAttributesForPartner(id)
	} else {
		attributes, err = s.querier.FindPartnerAttribute(id, group)
	}

	return id, code, attributes, err
}

func (s partnerService) GetDataById(_ context.Context, partnerId int32, partnerCode, group string) (int32, string, map[string]string, error) {
	attributes := make(map[string]string)
	if partnerId <= 0 && partnerCode == "" {
		return 0, "", attributes, errors.New("partnerId must be greater than 0")
	}
	if partnerId == 0 && partnerCode == "" {
		return 0, "", attributes, errors.New("partnerId and partnerCode cannot both be empty")
	}
	//If both partnerId and partnerCode are non-nil, check that the two correspond to the same row in the DB.
	if partnerId != 0 && partnerCode != "" {
		areEqual, _ := s.querier.CheckPartnerIDEqualsPartnerCode(partnerId, partnerCode)
		if !areEqual {
			return 0, "", attributes, errors.New("partnerId and partnerCode correspond to different values.")
		}
	}
	id, code, err := s.querier.FindPartnerDataByID(partnerId, partnerCode)

	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("partnerId %d not found", id))
	} else {
		//If a group is given to the GetDataById function return only the partner attributes for that group.
		if group == "" {
			attributes, err = s.querier.FindAllAttributesForPartner(id)

		} else {
			attributes, err = s.querier.FindPartnerAttribute(id, group)
		}
	}
	return id, code, attributes, err
}
