package service

import "server/internal/models"

// AddGoal adds a new user workout goal to the database
func (svc *service) AddGoal(uid int, goal *models.Goal) (*models.Goal, error) {
	goal.UserID = uid
	record, err := svc.db.InsertGoal(goal)
	if err != nil {
		return nil, err
	}

	return record, nil
}

// UpdateGoal updates an existing goal
func (svc *service) UpdateGoal(goal *models.Goal) (*models.Goal, error) {
	record, err := svc.db.UpdateGoal(goal)
	if err != nil {
		return nil, err
	}

	return record, nil
}

// GetGoals returns all the user's goals
func (svc *service) GetGoals(uid int) (*[]models.Goal, error) {
	goals, err := svc.db.GetGoalsByUser(uid)
	if err != nil {
		return nil, err
	}

	return &goals, nil
}

// DeleteGoal deletes a goal by the user ID and goal ID
func (svc *service) DeleteGoal(uid int, goalID int) error {
	return svc.db.DeleteGoalByID(uid, goalID)
}
