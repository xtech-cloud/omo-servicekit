package model

type Account struct {
	ID       int    `gorm:"column:id;primary_key" json:"id"`
	UUID     string `gorm:"column:uuid;not null;unique" json:"uuid"`
	Username string `gorm:"column:username;not null" json:"username"`
	Password string `gorm:"column:password;not null" json:"password"`
	Name     string `gorm:"column:name;not null" json:"name"`
	Avatar   string `gorm:"column:avatar;not null" json:"avatar"`
	Profile  string `gorm:"column:profile" json:"profile"`
}

func (Account) TableName() string {
	return "Accounts"
}

type AccountDAO struct {
}

func NewAccountDAO() *AccountDAO {
	return &AccountDAO{}
}

func MigrateAccount() error {
	err := db.AutoMigrate(&Account{}).Error
	return err
}

func (AccountDAO) Upsert(_account Account) error {
	var account Account
	err := db.Where(Account{UUID: _account.UUID}).FirstOrCreate(&account).Error
	if nil != err {
		return err
	}
	err = db.Model(&account).Updates(_account).Error
	return err
}

func (AccountDAO) List() ([]Account, error) {
	var accounts []Account
	err := db.Find(&accounts).Error
	return accounts, err
}

func (AccountDAO) Find(_uuid string) (Account, error) {
	var account Account
	err := db.Where("uuid = ?", _uuid).First(&account).Error
	return account, err
}

func (AccountDAO) WhereUsername(_username string) (Account, error) {
	var account Account
	res := db.Where("username= ?", _username).First(&account)
	if res.RecordNotFound() {
		return Account{}, nil
	}
	return account, res.Error
}
