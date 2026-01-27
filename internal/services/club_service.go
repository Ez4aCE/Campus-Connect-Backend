package services


import (
	"errors"


	"github.com/google/uuid"
	"gorm.io/gorm"


	"campus-connect-backend/internal/db"
	"campus-connect-backend/internal/models"
)

var(
	ErrClubNotFound = errors.New("club not found")
	ErrUnauthorizedClub = errors.New("unauthorized club action")
)

type ClubService struct{
	db *gorm.DB
}

func NewClubService() *ClubService{
	return &ClubService{
		db:db.DB,
	}
}

func (s *ClubService) CreateClub(club *models.Club) error{
	club.Status="pending"
	return s.db.Create(club).Error
}

func (s *ClubService) AprroveClub(clubID uuid.UUID) error{
	var club models.Club

	if err:=s.db.First(&club,"id=?",clubID).Error ; err!=nil{
		return ErrClubNotFound
	}

	if club.Status=="active"{
		return nil
	}

	tx:=s.db.Begin()

	club.Status="active"

	if err:=tx.Save(&club).Error;err!=nil{
		tx.Rollback()
		return err
	}

	membership:= models.ClubMembership{
		ClubID: clubID,
		UserID: club.CreatedBy,
		Role: "admin",
		Status: "approved",
	}

	if err:=tx.Create(&membership).Error ; err!=nil{
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}