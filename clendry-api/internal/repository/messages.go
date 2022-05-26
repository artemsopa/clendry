package repository

// import (
// 	"github.com/artomsopun/clendry/clendry-api/internal/domain"
// 	"gorm.io/gorm"
// )

// type MessagesRepo struct {
// 	db *gorm.DB
// }

// func NewMessagesRepo(db *gorm.DB) *MessagesRepo {
// 	return &MessagesRepo{
// 		db: db,
// 	}
// }

// func (r *MessagesRepo) GetLastIncomingMessagesByUserID(userID uint) ([]domain.Message, error) {
// 	var messages []domain.Message

// 	err := r.db.Raw("SELECT * FROM messages AS ms WHERE user_id = ? AND "+
// 		"created_at = (SELECT MAX(created_at) FROM messages AS ms2 WHERE ms.addressee_id = ms2.addressee_id) "+
// 		"AND ORDER BY ab.addressee_id", userID).Scan(&messages).Error
// 	if err != nil {
// 		return []domain.Message{}, err
// 	}

// 	return messages, nil
// }

// func (r *MessagesRepo) Create(message domain.Message) error {
// 	err := r.db.Create(&message).Error
// 	return err
// }

// func (r *MessagesRepo) Delete(userID, messageID uint) error {
// 	var chat []domain.Message
// 	err := r.db.Where("id = ? AND user_id = ?", messageID, userID).Find(&chat).Error
// 	return err
// }
