package usecase

import "product/internal/repository"

type FirstUsecase struct {
	repo repository.Repo
}

func New(repo repository.Repo) *FirstUsecase {
	return &FirstUsecase{repo: repo}
}

func (ucase *FirstUsecase) getRow(id int) (int, error) {
	return ucase.repo.FindOne(id)
}

func (ucase *FirstUsecase) getRows() ([]int, error) {
	return ucase.repo.FindAll()
}
