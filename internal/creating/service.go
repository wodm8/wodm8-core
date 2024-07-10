package creating

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	crossfit "github.com/wodm8/wodm8-core/internal"
	"github.com/wodm8/wodm8-core/internal/platform/server/handler"
)

type ExerciseService struct {
	exerciseRepository crossfit.ExerciseRepository
}

func NewExerciseService(exerciseRepository crossfit.ExerciseRepository) ExerciseService {
	return ExerciseService{
		exerciseRepository: exerciseRepository,
	}
}

func (w ExerciseService) CreateExercise(ctx *gin.Context, exerciseId, exerciseName string) error {
	exercise, err := crossfit.NewExercise(exerciseId, exerciseName)
	if err != nil {
		return err
	}

	return w.exerciseRepository.Save(ctx, exercise)
}

type WodService struct {
	wodRepository         crossfit.WodRepository
	exerciseWodRepository crossfit.ExerciseWodRepository
}

func NewWodService(wodRepository crossfit.WodRepository, exerciseWodRepository crossfit.ExerciseWodRepository) WodService {
	return WodService{
		wodRepository:         wodRepository,
		exerciseWodRepository: exerciseWodRepository,
	}
}

func (w WodService) CreateWod(ctx *gin.Context, wodDto handler.CreateWodRequest) error {
	wod, err := crossfit.NewWod(wodDto.ID, wodDto.Name, wodDto.Rounds, wodDto.NumberSections, wodDto.TimerType)
	if err != nil {
		return err
	}

	if err := w.wodRepository.Save(ctx, wod); err != nil {
		return err
	}

	for _, exercise := range wodDto.Exercises {
		fmt.Println(exercise)

		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		exerciseWod, err := crossfit.NewExerciseWod(id.String(), wod.ID().String(), exercise.ID, exercise.Reps, exercise.Weight, exercise.WeightUnit, exercise.Section, exercise.SectionTimerType, exercise.SectionCap)
		if err != nil {
			return err
		}

		if err := w.exerciseWodRepository.Save(ctx, exerciseWod); err != nil {
			return err
		}
	}
	return nil
}
