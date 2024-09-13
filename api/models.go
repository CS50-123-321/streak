package api

type Habit struct {
	Name             string `json:"name" redis:"name" binding:"required"`
	HabitName        string `json:"habitName" redis:"habit_name" binding:"required"`
	CommitmentPeriod string `json:"commitmentPeriod"  redis:"commitment_Period" binding:"required"`
	RedisId          int    `redis:"redis_id" binding:"required"`
	TeleID           int
}
