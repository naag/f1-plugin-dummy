package config

import (
	"github.com/form3tech-oss/f1/pkg/f1/testing"
	"github.com/form3tech-oss/go-form3/v2/pkg/form3"
	"github.com/go-openapi/strfmt"
)

type OrganisationAssociation struct {
	Bic            string
	BankId         string
	BankIdCode     string
	EtsCertificate string
}

var (
	BackofficeBacs = OrganisationAssociation{
		BankId:         "040013",
		BankIdCode:     "GBDSC",
		Bic:            "XXXXGB21",
		EtsCertificate: "starling",
	}
	BackofficePayport = OrganisationAssociation{
		BankId:     "112233",
		BankIdCode: "GBDSC",
		Bic:        "XXXXGB22",
	}
	BackofficeSepaInstant = OrganisationAssociation{
		Bic:        "FTHRGBL1",
		BankId:     "040013",
		BankIdCode: "GBDSC",
	}
	BackofficeSepaCreditTransfer = OrganisationAssociation{
		Bic:        "FTHRGBL1",
		BankId:     "040013",
		BankIdCode: "GBDSC",
	}
	BackofficeSepaDirectDebit = OrganisationAssociation{
		Bic:        "FTHRGBL1",
		BankId:     "040013",
		BankIdCode: "GBDSC",
	}
	BackofficeFpsGatewayLoadTest = OrganisationAssociation{
		Bic:        "XXXXGB33",
		BankId:     "123123",
		BankIdCode: "GBDSC",
	}
	BackofficeLHV = OrganisationAssociation{
		Bic:        "LHVBEE22",
		BankId:     "400399",
		BankIdCode: "GBDSC",
	}
	LoadtestBacs = OrganisationAssociation{
		BankId:         "377777",
		BankIdCode:     "GBDSC",
		Bic:            "XXXXGB78",
		EtsCertificate: "starling",
	}
	FpsGatewayScalingStress = OrganisationAssociation{
		Bic:        "XXXXGB36",
		BankId:     "112233",
		BankIdCode: "GBDSC",
	}

	CleanOnStart = false

	organisationMap = map[string]map[OrganisationAssociationType]*OrganisationAssociation{
		"backoffice":                        {Bacs: &BackofficeBacs, SepaInstant: &BackofficeSepaInstant},
		"backoffice payport":                {Payport: &BackofficePayport},
		"Backoffice FPS Gateway (loadtest)": {FpsGateway: &BackofficeFpsGatewayLoadTest},
		"Form3 Financial Cloud Fake LHV":    {LHV: &BackofficeLHV},
		"loadtest-bacs":                     {Bacs: &LoadtestBacs},
	}
)

type OrganisationAssociationType int

const (
	SepaInstant OrganisationAssociationType = iota
	Bacs
	Payport
	FpsGateway
	LHV
)

func GetOrganisationAssociation(t *testing.T, f3Client *form3.F3, associationType OrganisationAssociationType, defaultAssociation OrganisationAssociation) OrganisationAssociation {
	organisationId := MustGetString(t, "FORM3_ORGANISATION_ID")
	response := f3Client.Organisations.ListUnits().WithFilterOrganisationIds([]strfmt.UUID{strfmt.UUID(organisationId)}).MustDo()
	t.Require.Len(response.Data, 1)
	organisationName := response.Data[0].Attributes.Name

	if associationMap, ok := organisationMap[organisationName]; ok {
		if association, ok := associationMap[associationType]; ok {
			t.Log.Debugf("Using organisation association for organisation name: %s", organisationName)
			return *association
		}
	}

	t.Log.Debug("Using default organisation association")
	return defaultAssociation
}
