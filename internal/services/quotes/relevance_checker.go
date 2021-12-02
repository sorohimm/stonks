package quotes_service

import "time"

const freshHours = 36

func (s *QuotesService) CheckRelevance(lastUpd string) bool {
	last, _ := time.Parse("2006-01-02", lastUpd)
	diff := s.Config.CurrentTime.Sub(last)

	return diff.Hours() >= freshHours
}
