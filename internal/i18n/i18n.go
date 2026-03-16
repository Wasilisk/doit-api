package i18n

type Lang string

const (
	EN Lang = "en"
	UA Lang = "ua"
)

var translations = map[string]map[Lang]string{
	// auth errors
	"EMAIL_ALREADY_EXISTS":      {EN: "Email already exists", UA: "Електронна пошта вже існує"},
	"INVALID_CREDENTIALS":       {EN: "Invalid credentials", UA: "Невірні облікові дані"},
	"PASSWORD_HASHING_FAILED":   {EN: "Password hashing failed", UA: "Помилка хешування пароля"},
	"PROFILE_CREATION_FAILED":   {EN: "Profile creation failed", UA: "Не вдалося створити профіль"},
	"UNAUTHORIZED":              {EN: "Unauthorized", UA: "Не авторизовано"},
	"USER_WITH_EMAIL_NOT_FOUND": {EN: "User with given email not found", UA: "Користувача з такою електронною поштою не знайдено"},

	// general errors
	"NOT_FOUND":        {EN: "Not found", UA: "Не знайдено"},
	"INTERNAL_ERROR":   {EN: "Internal server error", UA: "Внутрішня помилка сервера"},
	"VALIDATION_ERROR": {EN: "Validation error", UA: "Помилка валідації"},
	"BAD_REQUEST":      {EN: "Bad request", UA: "Невірний запит"},

	// resource errors
	"TAG_ALREADY_EXISTS": {EN: "Tag with this name already exists", UA: "Тег з такою назвою вже існує"},
	"TAG_NOT_FOUND":      {EN: "Tag not found", UA: "Тег не знайдено"},
	"TASK_NOT_FOUND":     {EN: "Task not found", UA: "Завдання не знайдено"},
	"PROFILE_NOT_FOUND":  {EN: "Profile not found", UA: "Профіль не знайдено"},

	// id errors
	"INVALID_ID": {EN: "Invalid ID format", UA: "Невірний формат ID"},

	// file errors
	"FILE_TYPE_NOT_ALLOWED": {EN: "File type is not allowed", UA: "Тип файлу не дозволено"},
	"AVATAR_UPLOAD_FAILED":  {EN: "Failed to upload avatar", UA: "Не вдалося завантажити аватар"},

	// form errors
	"FORM_PARSE_FAILED": {EN: "Failed to parse form data", UA: "Не вдалося обробити дані форми"},

	// validation field errors
	"FIELD_REQUIRED":      {EN: "%s is required", UA: "%s є обов'язковим"},
	"FIELD_INVALID_EMAIL": {EN: "Please enter a valid email address", UA: "Будь ласка, введіть дійсну електронну адресу"},
	"FIELD_MIN":           {EN: "%s must be at least %s characters", UA: "%s повинен містити щонайменше %s символів"},
	"FIELD_MAX":           {EN: "%s must be at most %s characters", UA: "%s повинен містити не більше %s символів"},
	"FIELD_INVALID":       {EN: "%s is invalid", UA: "%s є недійсним"},
}

// Translate returns the localized message for a given code and language.
// Falls back to EN if the language is not found.
func Translate(code string, lang Lang) string {
	msgs, ok := translations[code]
	if !ok {
		return code
	}

	msg, ok := msgs[lang]
	if !ok {
		return msgs[EN]
	}
	return msg
}

// ParseLang converts a language string to a Lang type.
// Defaults to EN for unsupported languages.
func ParseLang(s string) Lang {
	switch s {
	case "ua":
		return UA
	default:
		return EN
	}
}
