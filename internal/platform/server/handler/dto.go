package handler

type ExerciseInWod struct {
	ID               string  `json:"id" binding:"required"`
	Reps             int32   `json:"reps" binding:"required"`
	Weight           float32 `json:"weight"`
	WeightUnit       string  `json:"weight_unit"`
	Section          int32   `json:"section" binding:"required"`
	SectionTimerType int32   `json:"section_timer_type"`
	SectionCap       int32   `json:"section_cap"`
}

type CreateWodRequest struct {
	ID             string          `json:"id" binding:"required"`
	Name           string          `json:"name" binding:"required"`
	Rounds         int32           `json:"rounds" binding:"required"`
	NumberSections int32           `json:"number_sections" binding:"required"`
	TimerType      int32           `json:"timer_type_id" binding:"required"`
	Exercises      []ExerciseInWod `json:"exercises" binding:"required"`
}
