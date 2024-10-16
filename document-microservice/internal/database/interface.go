package database

type UserRepository interface {
	CreateHistory(history *History)
	UpdateHistory(idHistory string, history *History)
	GetListHistory(pacientId string) []History
	GetListHistoryById(historyId string) History
}

func CreateHistoryOfVisit(repo UserRepository, history *History) {
	repo.CreateHistory(history)
}

func UpdateHistoryOfVisit(repo UserRepository, idHistory string, history *History) {
	repo.UpdateHistory(idHistory, history)
}
func GetListHistoryByIdPacient(repo UserRepository, idPacient string) []History {
	return repo.GetListHistory(idPacient)
}

func GetListHistoryByIdHistory(repo UserRepository, idHistory string) History {
	return repo.GetListHistoryById(idHistory)
}
