package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-route-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToRouteCollection(raw []byte, l *logger.Logger) ([]RouteCollection, error) {
	pm := &responses.RouteCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to RouteCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	routeCollection := make([]RouteCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		routeCollection = append(routeCollection, RouteCollection{
			ObjectID:                    data.ObjectID,
			ETag:                        data.ETag,
			ID:                          data.ID,
			Name:                        data.Name,
			RouteTypeCode:               data.RouteTypeCode,
			RouteTypeCodeText:           data.RouteTypeCodeText,
			StartDate:                   data.StartDate,
			ExcludeMondayIndicator:      data.ExcludeMondayIndicator,
			ExcludeTuesdayIndicator:     data.ExcludeTuesdayIndicator,
			ExcludeWednesdayIndicator:   data.ExcludeWednesdayIndicator,
			ExcludeThursdayIndicator:    data.ExcludeThursdayIndicator,
			ExcludeFridayIndicator:      data.ExcludeFridayIndicator,
			ExcludeSaturdayIndicator:    data.ExcludeSaturdayIndicator,
			ExcludeSundayIndicator:      data.ExcludeSundayIndicator,
			SalesTerritoryID:            data.SalesTerritoryID,
			SalesOrganisationID:         data.SalesOrganisationID,
			OrganizerPartyID:            data.OrganizerPartyID,
			DefaultStartTime:            data.DefaultStartTime,
			DefaultPreparationTime:      data.DefaultPreparationTime,
			DefaultDuration:             data.DefaultDuration,
			VisitTypeCode:               data.VisitTypeCode,
			VisitTypeCodeText:           data.VisitTypeCodeText,
			Status:                      data.Status,
			StatusText:                  data.StatusText,
			ProcessingStatus:            data.ProcessingStatus,
			ProcessingStatusText:        data.ProcessingStatusText,
			OwnerPartyID:                data.OwnerPartyID,
			OwnerPartyUUID:              data.OwnerPartyUUID,
			PerfectStoreIndicator:       data.PerfectStoreIndicator,
			DistributionChannelCode:     data.DistributionChannelCode,
			DistributionChannelCodeText: data.DistributionChannelCodeText,
			DivisionCode:                data.DivisionCode,
			DivisionCodeText:            data.DivisionCodeText,
			AutomaticResequencing:       data.AutomaticResequencing,
			EndDate:                     data.EndDate,
			WorkingDayCalendarCode:      data.WorkingDayCalendarCode,
			WorkingDayCalendarCodeText:  data.WorkingDayCalendarCodeText,
		})
	}

	return routeCollection, nil
}

func ConvertToRouteAccountCollection(raw []byte, l *logger.Logger) ([]RouteAccountCollection, error) {
	pm := &responses.RouteAccountCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to RouteAccountCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	routeAccountCollection := make([]RouteAccountCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		routeAccountCollection = append(routeAccountCollection, RouteAccountCollection{
			ObjectID:          data.ObjectID,
			ParentObjectID:    data.ParentObjectID,
			ETag:              data.ETag,
			RouteID:           data.RouteID,
			AccountID:         data.AccountID,
			AccountUUID:       data.AccountUUID,
			AllDayEvent:       data.AllDayEvent,
			DayNumber:         data.DayNumber,
			DriveTime:         data.DriveTime,
			Duration:          data.Duration,
			EndTime:           data.EndTime,
			StartTime:         data.StartTime,
			PreparationTime:   data.PreparationTime,
			VisitDate:         data.VisitDate,
			VisitTypeCode:     data.VisitTypeCode,
			VisitTypeCodeText: data.VisitTypeCodeText,
			VisitUUID:         data.VisitUUID,
			VisitID:           data.VisitID,
		})
	}

	return routeAccountCollection, nil
}

func ConvertToRouteInvolvedPartiesCollection(raw []byte, l *logger.Logger) ([]RouteInvolvedPartiesCollection, error) {
	pm := &responses.RouteInvolvedPartiesCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to RouteInvolvedPartiesCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	routeInvolvedPartiesCollection := make([]RouteInvolvedPartiesCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		routeInvolvedPartiesCollection = append(routeInvolvedPartiesCollection, RouteInvolvedPartiesCollection{
			ObjectID:                        data.ObjectID,
			ParentObjectID:                  data.ParentObjectID,
			ETag:                            data.ETag,
			RouteID:                         data.RouteID,
			PartyName:                       data.PartyName,
			RoleCode:                        data.RoleCode,
			RoleCodeText:                    data.RoleCodeText,
			RoleCategoryCode:                data.RoleCategoryCode,
			RoleCategoryCodeText:            data.RoleCategoryCodeText,
			PartyTypeCode:                   data.PartyTypeCode,
			PartyTypeCodeText:               data.PartyTypeCodeText,
			Address:                         data.Address,
			Email:                           data.Email,
			FormattedPhoneNumberDescription: data.FormattedPhoneNumberDescription,
		})
	}

	return routeInvolvedPartiesCollection, nil
}