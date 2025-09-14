package link

import (
	"fmt"
	"main-service/internal/models"
)

type LinkService struct {
	repo *LinkRepository
}

func NewLinkService(repo *LinkRepository) *LinkService {
	return &LinkService{repo: repo}
}

func (s *LinkService) Create(originLink string) (string, error) {
	if originLink == "" {
		return "", fmt.Errorf("error: Empty link")
	}
	exists, err := s.repo.Exist(originLink)
	if err != nil {
		return "", fmt.Errorf("error checking link existence: %v", err)
	}
	if exists {
		return "", fmt.Errorf("error: Link already exists")
	}
	pseudoLink := generatePseudoLink()
	return pseudoLink, s.repo.Create(&models.Link{
		OriginLink: originLink,
		PseudoLink: pseudoLink,
	})
}

func (s *LinkService) GetLink(link string) (*models.Link, error) {
	if link == "" {
		return nil, fmt.Errorf("error: Empty link")
	}
	linkObject, err := s.repo.Find(link)
	return linkObject, err
}

func (s *LinkService) DeleteLink(originLink string) error {
	if originLink == "" {
		return fmt.Errorf("error: Empty link")
	}
	return s.repo.Delete(originLink)

}
