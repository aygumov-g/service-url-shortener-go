package create_link

type CreateLinkUsecase interface {
	Execute(original string) (string, error)
}
