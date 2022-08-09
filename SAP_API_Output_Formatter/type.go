package sap_api_output_formatter

type Route struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	RouteCode     string `json:"route_code"`
	Deleted       bool   `json:"deleted"`
}

type RouteCollection struct {
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
}

type RouteAccountCollection struct {
	ObjectID          string `json:"ObjectID"`
	ParentObjectID    string `json:"ParentObjectID"`
	ETag              string `json:"ETag"`
	RouteID           string `json:"RouteID"`
	AccountID         string `json:"AccountID"`
	AccountUUID       string `json:"AccountUUID"`
	AllDayEvent       bool   `json:"AllDayEvent"`
	DayNumber         int    `json:"DayNumber"`
	DriveTime         string `json:"DriveTime"`
	Duration          string `json:"Duration"`
	EndTime           string `json:"EndTime"`
	StartTime         string `json:"StartTime"`
	PreparationTime   string `json:"PreparationTime"`
	VisitDate         string `json:"VisitDate"`
	VisitTypeCode     string `json:"VisitTypeCode"`
	VisitTypeCodeText string `json:"VisitTypeCodeText"`
	VisitUUID         string `json:"VisitUUID"`
	VisitID           string `json:"VisitID"`
}

type RouteInvolvedPartiesCollection struct {
	ObjectID                        string `json:"ObjectID"`
	ParentObjectID                  string `json:"ParentObjectID"`
	ETag                            string `json:"ETag"`
	RouteID                         string `json:"RouteID"`
	PartyID                         string `json:"PartyID"`
	PartyName                       string `json:"PartyName"`
	RoleCode                        string `json:"RoleCode"`
	RoleCodeText                    string `json:"RoleCodeText"`
	RoleCategoryCode                string `json:"RoleCategoryCode"`
	RoleCategoryCodeText            string `json:"RoleCategoryCodeText"`
	PartyTypeCode                   string `json:"PartyTypeCode"`
	PartyTypeCodeText               string `json:"PartyTypeCodeText"`
	Address                         string `json:"Address"`
	Email                           string `json:"Email"`
	FormattedPhoneNumberDescription string `json:"FormattedPhoneNumberDescription"`
}
