package usecase

import "testing"
import "github.com/golang/mock/gomock"
import "product/internal/repository/mocks"

func TestFindOne(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_repository.NewMockRepo(ctl)
	repo.EXPECT().FindOne(1).Return(1)
	//ucase := New(repo)
}
