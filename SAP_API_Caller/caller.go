package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-route-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetRoute(iD, accountID, partyID string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "RouteCollection":
			func() {
				c.RouteCollection(iD)
				wg.Done()
			}()
		case "RouteAccountCollection":
			func() {
				c.RouteAccountCollection(accountID)
				wg.Done()
			}()
		case "RouteInvolvedPartiesCollection":
			func() {
				c.RouteInvolvedPartiesCollection(partyID)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) RouteCollection(iD string) {
	routeCollectionData, err := c.callRouteSrvAPIRequirementRouteCollection("RouteCollection", iD)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(routeCollectionData)

}

func (c *SAPAPICaller) callRouteSrvAPIRequirementRouteCollection(api, iD string) ([]sap_api_output_formatter.RouteCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithRouteCollection(req, iD)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToRouteCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) RouteAccountCollection(accountID string) {
	routeAccountCollectionData, err := c.callRouteSrvAPIRequirementRouteAccountCollection("RouteAccountCollection", accountID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(routeAccountCollectionData)

}

func (c *SAPAPICaller) callRouteSrvAPIRequirementRouteAccountCollection(api, accountID string) ([]sap_api_output_formatter.RouteAccountCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithRouteAccountCollection(req, accountID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToRouteAccountCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) RouteInvolvedPartiesCollection(partyID string) {
	routeInvolvedPartiesCollectionData, err := c.callRouteSrvAPIRequirementRouteInvolvedPartiesCollection("RouteInvolvedPartiesCollection", partyID)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(routeInvolvedPartiesCollectionData)

}

func (c *SAPAPICaller) callRouteSrvAPIRequirementRouteInvolvedPartiesCollection(api, partyID string) ([]sap_api_output_formatter.RouteInvolvedPartiesCollection, error) {
	url := strings.Join([]string{c.baseURL, "c4codataapi", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithRouteInvolvedPartiesCollection(req, partyID)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToRouteInvolvedPartiesCollection(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithRouteCollection(req *http.Request, iD string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ID eq '%s'", iD))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithRouteAccountCollection(req *http.Request, accountID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("AccountID eq '%s'", accountID))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithRouteInvolvedPartiesCollection(req *http.Request, partyID string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("PartyID eq '%s'", partyID))
	req.URL.RawQuery = params.Encode()
}
