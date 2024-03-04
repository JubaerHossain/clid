package utilQuery

import (
	"encoding/json"
	"math"
	"net/http"
	"net/url"
	"strconv"

	"github.com/JubaerHossain/restaurant-golang/internal/core/pkg/utils"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Pagination(query *gorm.DB, queryValues map[string][]string) *gorm.DB {
	q := url.Values(queryValues)
	page, _ := strconv.Atoi(q.Get("page"))
	if page <= 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(q.Get("pageSize"))
	switch {
	case pageSize <= 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	query = query.Offset(offset).Limit(pageSize) // Pagination

	return query

}

func RawPagination(sqlQuery string, queryValues map[string][]string) string {
	q := url.Values(queryValues)
	page, _ := strconv.Atoi(q.Get("page"))
	if page <= 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(q.Get("pageSize"))
	switch {
	case pageSize <= 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	sqlQuery += " LIMIT " + strconv.Itoa(pageSize) + " OFFSET " + strconv.Itoa(offset)

	return sqlQuery
}

func HashPassword(password string) (string, error) {
	bp := []byte(password)
	hp, err := bcrypt.GenerateFromPassword(bp, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hp), nil
}

func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func BodyParse(s interface{}, w http.ResponseWriter, r *http.Request, isValidation bool) error {
	err := json.NewDecoder(r.Body).Decode(s)
	if err != nil {
		utils.WriteJSONError(w, http.StatusBadRequest, "Invalid JSON")
		return err
	}

	if isValidation {
		validate := validator.New()
		validateErr := validate.Struct(s)
		if validateErr != nil {
			utils.WriteJSONEValidation(w, http.StatusBadRequest, validateErr.(validator.ValidationErrors))
			return validateErr
		}
	}
	return nil
}

func Round(num float64, places int) float64 {
	if places < 0 {
		panic("places cannot be negative")
	}
	pow := math.Pow(10, float64(places))
	rounded := math.Round(num*pow) / pow
	return rounded
}

func OrderBy(queryValues map[string][]string) (string) {
	q := url.Values(queryValues)
	orderBy := "created_at"
	if conditions, ok := q["orderBy"]; ok && len(conditions) > 0 {
		orderBy = conditions[0]
	}

	sortOrder := "asc"
	if conditions, ok := q["sortBy"]; ok && len(conditions) > 0 {
		sortOrder = conditions[0]
	}

	orderBy = orderBy + " " + sortOrder

	return orderBy
}
