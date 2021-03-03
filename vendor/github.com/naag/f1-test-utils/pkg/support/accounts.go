package support

import (
	"fmt"
	"net/http"
	"time"

	"github.com/avast/retry-go"
	"github.com/form3tech-oss/f1/pkg/f1/testing"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/client/accounts"
	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/naag/f1-test-utils/pkg/config"
	"github.com/naag/f1-test-utils/pkg/support/rand"
)

func LoadAccounts(t *testing.T, organisation config.OrganisationAssociation, numberOfAccounts int) []*models.Account {
	f3 := F3ForTest(t)
	existingAccounts, err := f3.Accounts.ListAccounts().
		WithFilterOrganisationID([]strfmt.UUID{*f3.Defaults.OrganisationId}).
		WithPageSize(int64(numberOfAccounts)).
		Do()
	t.Require.Nil(err, "failed to load accounts")

	var accounts = existingAccounts.Data
	n := 0
	for _, account := range accounts {
		if account.Attributes.AccountNumber != "" {
			accounts[n] = account
			n++
		}
	}
	accounts = accounts[:n]

	if len(accounts) < numberOfAccounts {
		t.Log.Info("Insufficient accounts found - creating new accounts")
		for i := 0; i < numberOfAccounts-len(accounts); i++ {
			accounts = append(accounts, CreateAccount(t, organisation))
		}
	}
	t.Log.Infof("Loaded %d accounts", len(accounts))

	return accounts
}

func LoadNonForm3Accounts(t *testing.T, organisation config.OrganisationAssociation) []*models.Account {
	var accounts []*models.Account

	r := CreateAccount(t, organisation)

	for i := 0; i < 100; i++ {
		newUuid := uuid.New()
		accounts = append(accounts, r.WithID(strfmt.UUID(newUuid.String())))
	}

	t.Log.Infof("Loaded %d fake non-Form3 accounts", len(accounts))

	return accounts
}

func CreateAccount(t *testing.T, organisation config.OrganisationAssociation) *models.Account {
	var createdAccount *models.Account
	err := retry.Do(func() error {
		r := BuildAccountRequest(t, organisation)
		acc, err := r.Do()
		if err != nil {
			return err
		}

		createdAccount = acc.Data

		return nil
	}, retry.Attempts(2), retry.Delay(100*time.Millisecond))
	t.Require.NoError(err, "unable to create account")
	t.Require.NotEmpty(createdAccount.Attributes.AccountNumber)
	return createdAccount
}

func RandomAccountNumber() string {
	return fmt.Sprintf("%08d", rand.Intn(100000000))
}

func BuildAccountRequest(t *testing.T, organisation config.OrganisationAssociation) *accounts.CreateAccountRequest {
	r := F3ForTest(t).Accounts.CreateAccount()
	r.Data.Attributes.
		WithoutOrganisationIdentification().
		WithoutPrivateIdentification().
		WithBankID(organisation.BankId).
		WithBankIDCode(organisation.BankIdCode).
		WithBic(organisation.Bic).
		WithCountry("GB").
		WithBaseCurrency("GBP").
		WithAccountNumber(RandomAccountNumber())

	return r
}

func DeleteAccount(t *testing.T, account *models.Account) {
	_, err := F3ForTest(t).Accounts.DeleteAccount().
		WithID(*account.ID).
		WithVersion(*account.Version).
		Do()
	if err != nil {
		switch err := err.(type) {
		case *runtime.APIError:
			if err.Code != http.StatusNotFound {
				t.Log.WithError(err).Errorf("Unable to delete account %s", *account.ID)
			}
		}
	}
}
