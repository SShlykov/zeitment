package entity

import "github.com/jackc/pgx/v5"

type UserProfile struct {
	ID         string
	UserID     string
	FirstName  string
	LastName   string
	SecondName string
	Locale     string
	Gender     int
	Phone      string
	Notes      string
}

func (u UserProfile) TableName() string {
	return "user_profiles"
}

func (u UserProfile) AllFields() []string {
	return []string{"id", "user_id", "first_name", "last_name", "second_name", "locale",
		"gender", "phone", "notes"}
}

func (u UserProfile) InsertOrUpdateFields() []string {
	return []string{"user_id", "first_name", "last_name", "second_name", "locale", "gender", "phone", "notes"}
}

func (u UserProfile) EntityToInsertValues(impl *UserProfile) []interface{} {
	return []interface{}{
		impl.UserID, impl.FirstName, impl.LastName, impl.SecondName, impl.Locale, impl.Gender,
		impl.Phone, impl.Notes}
}

func (u UserProfile) ReadItem(row pgx.Row) (UserProfile, error) {
	var userProfile UserProfile
	err := row.Scan(&userProfile.ID, &userProfile.UserID, &userProfile.FirstName, &userProfile.LastName,
		&userProfile.SecondName, &userProfile.Locale, &userProfile.Gender, &userProfile.Phone, &userProfile.Notes)

	if err != nil {
		return UserProfile{}, err
	}
	return userProfile, nil
}

func (u UserProfile) ReadList(rows pgx.Rows) ([]UserProfile, error) {
	var userProfiles []UserProfile
	for rows.Next() {
		userProfile, err := u.ReadItem(rows)
		if err != nil {
			return nil, err
		}
		userProfiles = append(userProfiles, userProfile)
	}
	return userProfiles, nil
}
