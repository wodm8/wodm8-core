package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wodm8/wodm8-core/internal/crossfit"
	"github.com/wodm8/wodm8-core/internal/domain"
)

type ExerciseService struct {
	exerciseRepository crossfit.ExerciseRepository
}

func NewExerciseService(exerciseRepository crossfit.ExerciseRepository) ExerciseService {
	return ExerciseService{
		exerciseRepository: exerciseRepository,
	}
}

func (w ExerciseService) CreateExercise(ctx *gin.Context, exerciseName string) error {
	exercise, err := crossfit.NewExercise(exerciseName)
	if err != nil {
		return err
	}

	return w.exerciseRepository.Save(ctx, exercise)
}

type WodService struct {
	wodRepository         crossfit.WodRepository
	exerciseWodRepository crossfit.ExerciseWodRepository
	setRepository         crossfit.WodSetRepository
	roundsRepository      crossfit.WodRoundRepository
}

func NewWodService(wodRepository crossfit.WodRepository, setRepository crossfit.WodSetRepository, roundsRepository crossfit.WodRoundRepository, exerciseWodRepository crossfit.ExerciseWodRepository) WodService {
	return WodService{
		wodRepository:         wodRepository,
		setRepository:         setRepository,
		roundsRepository:      roundsRepository,
		exerciseWodRepository: exerciseWodRepository,
	}
}

func (w WodService) CreateWod(ctx *gin.Context, wodDto domain.CreateWodRequest) error {
	wod, err := crossfit.NewWod(wodDto.ID, wodDto.Name, wodDto.WodDate, wodDto.WodTypeId)
	fmt.Printf("Error: %v\n", err)
	if err != nil {
		return err
	}

	if err := w.wodRepository.Save(ctx, wod); err != nil {
		return err
	}

	for _, set := range wodDto.Sets {
		idSet, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		wodSet, err := crossfit.NewWodSet(idSet.String(), wod.ID().String(), set.SetNumber, set.BuyIn, set.BuyOut, set.EveryMinutes, set.RepsToAddByRound, set.RestTime, set.IsThen)
		if err != nil {
			return err
		}

		if err := w.setRepository.Save(ctx, wodSet); err != nil {
			return err
		}
	}

	for _, round := range wodDto.Rounds {
		idRound, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		wodRound, err := crossfit.NewWodRound(idRound.String(), wod.ID().String(), round.SetNumber, round.RoundNumber, round.RepetitionsByRound, round.Time, round.RestTime, round.RemainingTime)
		if err != nil {
			return err
		}
		if err := w.roundsRepository.Save(ctx, wodRound); err != nil {
			return err
		}
	}

	for _, exercise := range wodDto.Exercises {
		fmt.Println(exercise)

		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		exerciseWod, err := crossfit.NewExerciseWod(
			id.String(),
			wod.ID().String(),
			exercise.ID,
			exercise.SetNumber,
			exercise.RoundNumber,
			exercise.Repetitions,
			exercise.RepetitionsUnit,
			exercise.Weight,
			exercise.WeightUnit,
			exercise.Distance,
			exercise.DistanceUnit)
		if err != nil {
			return err
		}

		if err := w.exerciseWodRepository.Save(ctx, exerciseWod); err != nil {
			return err
		}
	}
	return nil
}

func (w WodService) GetWod(ctx *gin.Context, id string) ([]domain.WodResult, error) {
	res, err := w.wodRepository.Get(ctx, id)
	if err != nil {
		fmt.Printf("error getting wod: %v", err)
		return make([]domain.WodResult, 0), err
	}

	return res, nil
}
