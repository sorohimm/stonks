package quotes_service

import "time"

func (s *QuotesService) CheckRelevance(lastUpd string) bool {
	last, _ := time.Parse("2006-01-02", lastUpd)
	diff := s.Config.CurrentTime.Sub(last)

	return diff.Hours() >= 36
}
