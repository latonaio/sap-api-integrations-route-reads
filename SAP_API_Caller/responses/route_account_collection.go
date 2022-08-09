package responses

type RouteAccountCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
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
		} `json:"results"`
	} `json:"d"`
}
