package services


import (
	"errors"


	"github.com/google/uuid"
	"gorm.io/gorm"


	"campus-connect-backend/internal/db"
	"campus-connect-backend/internal/models"
)


var (
	ErrAlreadyMember = errors.New("already a member or pending approval")
	ErrNotMember = errors.New("not a club member")
	ErrNotClubAdmin = errors.New("not club admin")
)

type MembershipService struct{
	db *gorm.DB
}

func NewMembershipService() *MembershipService{
	return &MembershipService{
		db: db.DB,
	}
}

func (s *MembershipService) RequestJoin(userID, clubID uuid.UUID) error{
	var count int64

	s.db.Model(&models.ClubMembership{}).
		Where("user_id=? AND club_id=?",userID,clubID).
		Count(&count)

	if count>0{
		return ErrAlreadyMember
	}

	m:=models.ClubMembership{
		UserID: userID,
		ClubID: clubID,
		Role: "member",
		Status: "pending",
	}

	return s.db.Create(&m).Error
}

func (s *MembershipService) ApproveMember(adminID, clubID, userID uuid.UUID) error{
	ok, err :=s.isClubAdmin(adminID,clubID)

	if err!=nil || !ok{
		return ErrNotClubAdmin
	}

	return s.db.Model(&models.ClubMembership{}).
		Where("user_id=? And club_id=?",userID,clubID).Error
}


func (s *MembershipService) isClubAdmin(userID, clubID uuid.UUID)(bool, error){
	var count int64

	err:=s.db.Model(&models.ClubMembership{}).
		Where("user_id=? AND club_id=? AND role=? AND status=?",
			userID,clubID,"admin","approved").
		Count(&count).Error

	return count>0,err
}

func (s *MembershipService) LeaveClub(userID, clubID uuid.UUID) error{
	res:=s.db.Where("user_id=? And club_id=?",userID,clubID).
			Delete(&models.ClubMembership{})

	if res.RowsAffected==0{
		return ErrNotMember
	}
	return res.Error
}


	
