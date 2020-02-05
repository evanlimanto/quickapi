package crawlers

import (
	"net/http/cookiejar"

	"github.com/evanlimanto/quickapi/src/database"
)

type BCACrawler struct{}

const loginUrl = "https://m.klikbca.com/login.jsp"

func (crawler *BCACrawler) LoginAndGetAccounts() ([]database.Account, error) {
	_, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
