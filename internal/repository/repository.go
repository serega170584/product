package repository

type Repo interface {
	FindOne(id int) (int, error)
	FindAll() ([]int, error)
}
