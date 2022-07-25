package responses

type RouteCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                    string `json:"ObjectID"`
			ETag                        string `json:"ETag"`
			ID                          string `json:"ID"`
			Name                        string `json:"Name"`
			RouteTypeCode               string `json:"RouteTypeCode"`
			RouteTypeCodeText           string `json:"RouteTypeCodeText"`
			StartDate                   string `json:"StartDate"`
			ExcludeMondayIndicator      bool   `json:"ExcludeMondayIndicator"`
			ExcludeTuesdayIndicator     bool   `json:"ExcludeTuesdayIndicator"`
			ExcludeWednesdayIndicator   bool   `json:"ExcludeWednesdayIndicator"`
			ExcludeThursdayIndicator    bool   `json:"ExcludeThursdayIndicator"`
			ExcludeFridayIndicator      bool   `json:"ExcludeFridayIndicator"`
			ExcludeSaturdayIndicator    bool   `json:"ExcludeSaturdayIndicator"`
			ExcludeSundayIndicator      bool   `json:"ExcludeSundayIndicator"`
			SalesTerritoryID            string `json:"SalesTerritoryID"`
			SalesOrganisationID         string `json:"SalesOrganisationID"`
			OrganizerPartyID            string `json:"OrganizerPartyID"`
			DefaultStartTime            string `json:"DefaultStartTime"`
			DefaultPreparationTime      string `json:"DefaultPreparationTime"`
			DefaultDuration             string `json:"DefaultDuration"`
			VisitTypeCode               string `json:"VisitTypeCode"`
			VisitTypeCodeText           string `json:"VisitTypeCodeText"`
			Status                      string `json:"Status"`
			StatusText                  string `json:"StatusText"`
			ProcessingStatus            string `json:"ProcessingStatus"`
			ProcessingStatusText        string `json:"ProcessingStatusText"`
			OwnerPartyID                string `json:"OwnerPartyID"`
			OwnerPartyUUID              string `json:"OwnerPartyUUID"`
			PerfectStoreIndicator       bool   `json:"PerfectStoreIndicator"`
			DistributionChannelCode     string `json:"DistributionChannelCode"`
			DistributionChannelCodeText string `json:"DistributionChannelCodeText"`
			DivisionCode                string `json:"DivisionCode"`
			DivisionCodeText            string `json:"DivisionCodeText"`
			AutomaticResequencing       bool   `json:"AutomaticResequencing"`
			EndDate                     string `json:"EndDate"`
			WorkingDayCalendarCode      string `json:"WorkingDayCalendarCode"`
			WorkingDayCalendarCodeText  string `json:"WorkingDayCalendarCodeText"`
		} `json:"results"`
	} `json:"d"`
}
