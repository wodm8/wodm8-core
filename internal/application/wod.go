package application

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wodm8/wodm8-core/internal/crossfit"
	"github.com/wodm8/wodm8-core/internal/domain"
	"github.com/wodm8/wodm8-core/internal/members"
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

	return w.exerciseRepository.Save(exercise)
}

type WodService struct {
	wodRepository         crossfit.WodRepository
	exerciseWodRepository crossfit.ExerciseWodRepository
	setRepository         crossfit.WodSetRepository
	roundsRepository      crossfit.WodRoundRepository
	memberWodRepository   crossfit.MemberWodsRepository
	membersRepository     members.MemberRepository
}

func NewWodService(wodRepository crossfit.WodRepository, setRepository crossfit.WodSetRepository, roundsRepository crossfit.WodRoundRepository, exerciseWodRepository crossfit.ExerciseWodRepository, memberWodRepository crossfit.MemberWodsRepository, membersRepository members.MemberRepository) WodService {
	return WodService{
		wodRepository:         wodRepository,
		setRepository:         setRepository,
		roundsRepository:      roundsRepository,
		exerciseWodRepository: exerciseWodRepository,
		memberWodRepository:   memberWodRepository,
		membersRepository:     membersRepository,
	}
}

func (w WodService) CreateWod(ctx *gin.Context, wodDto domain.CreateWodRequest, userEmail string) error {
	wod, err := crossfit.NewWod(wodDto.ID, wodDto.Name, wodDto.WodDate, wodDto.WodTypeId)
	if err != nil {
		return err
	}

	if err := w.wodRepository.Save(wod); err != nil {
		return err
	}

	for _, set := range wodDto.Sets {
		idSet, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		wodSet, err := crossfit.NewWodSet(idSet.String(), wod.ID, set.SetNumber, set.BuyIn, set.BuyOut, set.EveryMinutes, set.RepsToAddByRound, set.RestTime, set.IsThen)
		if err != nil {
			return err
		}

		if err := w.setRepository.Save(wodSet); err != nil {
			return err
		}
	}

	for _, round := range wodDto.Rounds {
		idRound, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		wodRound, err := crossfit.NewWodRound(idRound.String(), wod.ID, round.SetNumber, round.RoundNumber, round.RepetitionsByRound, round.Time, round.RestTime, round.RemainingTime)
		if err != nil {
			return err
		}
		if err := w.roundsRepository.Save(wodRound); err != nil {
			return err
		}
	}

	for _, exercise := range wodDto.Exercises {

		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		exerciseWod, err := crossfit.NewExerciseWod(
			id.String(),
			wod.ID,
			exercise.ExerciseId,
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

		if err := w.exerciseWodRepository.Save(exerciseWod); err != nil {
			return err
		}
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	member, err := w.membersRepository.FindByEmail(userEmail)
	if err != nil {
		return err
	}

	memberWod, err := crossfit.NewMemberWod(id.String(), member.ID, wod.ID)
	if err != nil {
		return err
	}

	if err := w.memberWodRepository.Save(memberWod); err != nil {
		return err
	}

	return nil
}

func (w WodService) GetWod(ctx *gin.Context, id string) ([]domain.CreatedWod, error) {
	res, err := w.wodRepository.FindByMember(id)
	if err != nil {
		fmt.Printf("error getting wod: %v", err)
		return make([]domain.CreatedWod, 0), err
	}

	return res, nil
}
