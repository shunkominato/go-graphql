// Code generated by ent, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProvider holds the string denoting the provider field in the database.
	FieldProvider = "provider"
	// FieldUID holds the string denoting the uid field in the database.
	FieldUID = "uid"
	// FieldEncryptedPassword holds the string denoting the encrypted_password field in the database.
	FieldEncryptedPassword = "encrypted_password"
	// FieldResetPasswordToken holds the string denoting the reset_password_token field in the database.
	FieldResetPasswordToken = "reset_password_token"
	// FieldResetPasswordSentAt holds the string denoting the reset_password_sent_at field in the database.
	FieldResetPasswordSentAt = "reset_password_sent_at"
	// FieldAllowPasswordChange holds the string denoting the allow_password_change field in the database.
	FieldAllowPasswordChange = "allow_password_change"
	// FieldRememberCreatedAt holds the string denoting the remember_created_at field in the database.
	FieldRememberCreatedAt = "remember_created_at"
	// FieldConfirmationToken holds the string denoting the confirmation_token field in the database.
	FieldConfirmationToken = "confirmation_token"
	// FieldConfirmedAt holds the string denoting the confirmed_at field in the database.
	FieldConfirmedAt = "confirmed_at"
	// FieldConfirmationSentAt holds the string denoting the confirmation_sent_at field in the database.
	FieldConfirmationSentAt = "confirmation_sent_at"
	// FieldUnconfirmedEmail holds the string denoting the unconfirmed_email field in the database.
	FieldUnconfirmedEmail = "unconfirmed_email"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldTokens holds the string denoting the tokens field in the database.
	FieldTokens = "tokens"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldProvider,
	FieldUID,
	FieldEncryptedPassword,
	FieldResetPasswordToken,
	FieldResetPasswordSentAt,
	FieldAllowPasswordChange,
	FieldRememberCreatedAt,
	FieldConfirmationToken,
	FieldConfirmedAt,
	FieldConfirmationSentAt,
	FieldUnconfirmedEmail,
	FieldName,
	FieldNickname,
	FieldImage,
	FieldEmail,
	FieldTokens,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}