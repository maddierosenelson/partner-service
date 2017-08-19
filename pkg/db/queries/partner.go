package queries

import (
	"fmt"

	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	"jaxf-github.fanatics.corp/apparel/partner-service/pkg/db/models"
)

func GetPartnerDataFromKeyValue(key, value string, conn *pgx.Conn) (int32, string, error) {

	partnerModel := new(models.Partner)
	//Statement to find the id and code that correspond to the given key and value.
	statement := "SELECT partners.Id, partners.Code FROM partner_mappings INNER JOIN partners on partners.Id = partner_mappings.partner_id WHERE key_id = (select id from keys where name = $1) and value =$2;"

	rows, err := conn.Query(statement, key, value)
	//hasRows is needed because err will return nil even when there are no rows to return from the above statement.
	hasRows := false
	//This for loop checks that rows contains at least one corresponding partner row. If it does, hasRows is set to true.
	for rows.Next() {
		hasRows = true
		err = rows.Scan(&partnerModel.Id, &partnerModel.Code)
		if rows.Next() == true {
			rows.Close()
			err = errors.Wrap(errors.New(""), fmt.Sprintf("Multiple partners matched for given key: %s and value: %s", key, value))
			return 0, "", err
		}
	}

	if err != nil {
		err = errors.Wrap(errors.New(""), fmt.Sprintf("Failed to query PartnerId from key: %s and value: %s in the database", key, value))
		return 0, "", err
	}
	//If hasRows was not reset to true, we want to return an error as this means there was no corresponding row for the entered key value pair.
	if !hasRows {
		err = errors.Wrap(errors.New(""), fmt.Sprintf("No rows returned from key: %s and value: %s in the database", key, value))
		return 0, "", err
	}

	partner := partnerModel.Gen(nil)
	return partner.Id, partner.Code, err
}

func GetPartnerDataByIDOrCode(id int32, code string, conn *pgx.Conn) (int32, string, error) {

	partnerModel := new(models.Partner)
	//For GetDataById in service.go, the partner data can be found using either the partnerId or partnerCode, as long as one entry is a valid entry (eg, non-negative, non-bad)
	//Thus the data can be selected using id or code.
	statement := "SELECT Id, Code FROM partners WHERE id = $1 or code = $2"

	err := conn.QueryRow(statement, id, code).Scan(&partnerModel.Id, &partnerModel.Code)

	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("failed to query Partnerdata from id: %d or partnerCode: %s ", id, code))
		return 0, "", err
	}

	partner := partnerModel.Gen(nil)

	if partner.Id == 0 {
		err = errors.Wrap(errors.New(""), fmt.Sprintf("No partner with id: %d and/pr partnerCode: %s ", id, code))
		return 0, "", err
	}

	return partner.Id, partner.Code, nil
}

func GetAllAttributesForPartner(id int32, conn *pgx.Conn) (map[string]string, error) {

	statement := "SELECT keys.name, partner_mappings.value FROM partner_mappings INNER JOIN keys on keys.id = partner_mappings.key_id WHERE partner_id = $1"
	rows, err := conn.Query(statement, id)

	attrMap := make(map[string]string)

	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("failed to query Attributes for partnerId: %d", id))
		return attrMap, err
	}
	hasRows := false
	for rows.Next() {
		hasRows = true

		attr := &models.Attribute{}
		err = rows.Scan(&attr.Name, &attr.Value)
		if err != nil {
			err = errors.Wrap(errors.New(""), fmt.Sprintf("Failed to scan Name and Value into Attributes"))
			return attrMap, err
		}
		attrMap[attr.Name.String] = attr.Value.String
	}
	if !hasRows {
		err = errors.Wrap(errors.New(""), fmt.Sprintf("No rows returned from id: %d", id))
		return make(map[string]string), err
	}

	return attrMap, err
}

func GetGroupAttributesForPartner(id int32, group string, conn *pgx.Conn) (map[string]string, error) {

	statement := "SELECT keys.name, partner_mappings.value FROM partner_mappings INNER JOIN keys ON keys.id = partner_mappings.key_id WHERE partner_id = $1 AND key_id = ANY(SELECT key_id FROM groups_to_keys WHERE group_id = (SELECT id FROM groups WHERE name = $2 LIMIT 1));"
	rows, err := conn.Query(statement, id, group)

	if err != nil {

		err = errors.Wrap(errors.New(""), fmt.Sprintf("fFiled to query attributes from id: %d and group: %s in db", id, group))
		return make(map[string]string), err
	}
	attrMap := make(map[string]string)
	hasRows := false
	for rows.Next() {
		hasRows = true
		attr := &models.Attribute{}
		err = rows.Scan(&attr.Name, &attr.Value)
		if err != nil {
			err = errors.Wrap(errors.New(""), fmt.Sprintf("Failed to scan Name and Value into Attributes"))
			return nil, err
		}
		attrMap[attr.Name.String] = attr.Value.String
	}
	if !hasRows {
		err = errors.Wrap(errors.New(""), fmt.Sprintf("No rows retured from id: %d and group: %s in db", id, group))
		return make(map[string]string), err
	}

	return attrMap, nil
}

func GetCheckPartnerIDEqualsPartnerCode(id int32, code string, conn *pgx.Conn) (bool, error) {

	partnerModel := new(models.Partner)
	//Before entering GetCheckPartnerIDEqualsPartnerCode, id and code are found to be non-nil. Check that the non-nil inputs correspond
	//to the same row in the DB.
	statement := "SELECT id, code FROM partners WHERE id = $1 and code = $2"

	rows, err := conn.Query(statement, id, code)
	hasRows := false
	for rows.Next() {
		hasRows = true
		err = rows.Scan(&partnerModel.Id, &partnerModel.Code)
	}

	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("failed to query Partnerdata from id: %d or partnerCode: %s when checking if the two correspond to the same row", id, code))
		return true, err
	}
	if !hasRows {
		err = errors.Wrap(errors.New(""), fmt.Sprintf("failed to query Partnerdata from id: %d or partnerCode: %s when checking if the two correspond to the same row because bad id or code", id, code))
	}
	return hasRows, err
}
