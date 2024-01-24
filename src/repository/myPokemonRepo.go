package repository

import (
	"context"
	"gorm.io/gorm"
	"miniProject/src/model"
	"miniProject/src/utils"
)

type (
	MyPokemonRepo interface {
		MyPokemonList(ctx context.Context, paramQuery utils.ParamQuery) ([]*model.MyPokemon, error)
		AddMyPokemon(ctx context.Context, data model.AddMyPokemonRequest) error
		DeletePokemon(ctx context.Context, id uint) error
	}

	MyPokemonRepoImpl struct {
		db *gorm.DB
	}
)

func NewMyPokemonRepo(db *gorm.DB) MyPokemonRepo {
	return &MyPokemonRepoImpl{db: db}
}

func (m *MyPokemonRepoImpl) DeletePokemon(ctx context.Context, id uint) error {
	resp := m.db.WithContext(ctx).Delete(&model.MyPokemon{}, id)
	return resp.Error
}

func (m *MyPokemonRepoImpl) AddMyPokemon(ctx context.Context, data model.AddMyPokemonRequest) error {
	myPokemon := model.MyPokemon{
		Name:    data.Name,
		Picture: data.Picture,
	}
	resp := m.db.WithContext(ctx).Create(&myPokemon)
	return resp.Error
}

func (m *MyPokemonRepoImpl) MyPokemonList(ctx context.Context, paramQuery utils.ParamQuery) ([]*model.MyPokemon, error) {
	var myPokeList []*model.MyPokemon
	resp := m.db.WithContext(ctx).Offset(paramQuery.Offset).Limit(paramQuery.Limit).Find(&myPokeList)
	return myPokeList, resp.Error
}
