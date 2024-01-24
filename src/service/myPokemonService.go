package service

import (
	"context"
	"math/rand"
	"miniProject/src/model"
	"miniProject/src/repository"
	"miniProject/src/utils"
)

type (
	MyPokemonService interface {
		MyPokemonList(ctx context.Context, paramQuery utils.ParamQuery) ([]*model.MyPokemon, error)
		AddMyPokemon(ctx context.Context, data model.AddMyPokemonRequest) error
		ReleasePokemon(ctx context.Context, id uint) (bool, error)
	}

	MyPokemonServiceImpl struct {
		myPokemonRepo repository.MyPokemonRepo
	}
)

func NewMyPokemonService(p repository.MyPokemonRepo) MyPokemonService {
	return &MyPokemonServiceImpl{myPokemonRepo: p}
}

func (p *MyPokemonServiceImpl) MyPokemonList(ctx context.Context, paramQuery utils.ParamQuery) ([]*model.MyPokemon, error) {
	res, err := p.myPokemonRepo.MyPokemonList(ctx, paramQuery)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (p *MyPokemonServiceImpl) AddMyPokemon(ctx context.Context, data model.AddMyPokemonRequest) error {
	err := p.myPokemonRepo.AddMyPokemon(ctx, data)
	return err
}

func (p *MyPokemonServiceImpl) ReleasePokemon(ctx context.Context, id uint) (bool, error) {
	randomNumber := rand.Intn(100)

	for i := 2; i*i <= randomNumber; i++ {
		if randomNumber%i == 0 {
			return false, nil
		}
	}
	err := p.myPokemonRepo.DeletePokemon(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
