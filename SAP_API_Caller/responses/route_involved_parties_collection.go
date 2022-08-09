package responses

type RouteInvolvedPartiesCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
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
			Route                           struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"Route"`
		} `json:"results"`
	} `json:"d"`
}