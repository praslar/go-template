package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"math"
	postgres2 "go-template/internal/pkg/db/postgres"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// ======================================================
// Time
// ======================================================
type Time struct {
	time.Time
}

// Json support
// ---------------------------------------------
func (t *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "" {
		t.Time = time.Time{}
		return
	}
	t.Time, err = time.Parse(time.RFC3339, s)
	return
}

func (t *Time) MarshalJSON() ([]byte, error) {
	if t.Time.UnixNano() == (time.Time{}).UnixNano() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", t.Time.Format(time.RFC3339))), nil
}

// Gorm support
// ---------------------------------------------
func (t *Time) Scan(src interface{}) (err error) {
	var tt Time
	switch src.(type) {
	case time.Time:
		tt.Time = src.(time.Time)
	default:
		return errors.New("incompatible type for Skills")
	}
	*t = tt
	return
}

func (t Time) Value() (driver.Value, error) {
	if !t.IsSet() {
		return "null", nil
	}
	return t.Time, nil
}

func (t *Time) IsSet() bool {
	return t.UnixNano() != (time.Time{}).UnixNano()
}

func jsonb(v string) postgres.Jsonb {
	return postgres.Jsonb{RawMessage: json.RawMessage(v)}
}

func floorOf(a, b int) int {
	return int(math.Floor(float64(a) / float64(b)))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mustBeInt(v string) int {
	res, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}
	return res
}

func CurrentUser(c beego.Controller) (uuid.UUID, error) {
	userIdStr := c.Ctx.Request.Header.Get("x-user-id")
	return uuid.Parse(userIdStr)
}

func paginationQuery(c beego.Controller) (int, int) {
	currentPage := 1
	perPage := 30
	if c.Ctx.Request.Method == http.MethodGet || c.Ctx.Request.Method == http.MethodDelete {
		currentPage = mustBeInt(strings.Trim(c.GetString("page"), " "))
		perPage = mustBeInt(strings.Trim(c.GetString("size"), " "))
	} else if c.Ctx.Request.Method == http.MethodPost ||
		c.Ctx.Request.Method == http.MethodPatch ||
		c.Ctx.Request.Method == http.MethodPut {
		currentPage = mustBeInt(strings.Trim(c.Ctx.Request.FormValue("page"), " "))
		perPage = mustBeInt(strings.Trim(c.Ctx.Request.FormValue("size"), " "))
	}
	// validate
	if currentPage < 1 {
		currentPage = 1
	}
	if perPage < 1 {
		perPage = 30
	}
	return currentPage, perPage
}

//=========================================================================
// Minimal Chain Data Getter For Beego Framework
// @by Namkazt
//=========================================================================
type chain struct {
	ctx     beego.Controller
	q       *gorm.DB
	err     error
	checkOn string
	updater map[string]interface{}
}

func With(c beego.Controller) *chain {
	return &chain{ctx: c}
}

func (c *chain) On(checkOn string) *chain {
	c.checkOn = checkOn
	return c
}

func (c *chain) Get(name string) string {
	if c.checkOn == "path" {
		return strings.Trim(c.ctx.Ctx.Input.Param(":"+name), " ")
	} else if c.checkOn == "query" {
		return strings.Trim(c.ctx.Ctx.Input.Query(name), " ")
	} else if c.checkOn == "form" {
		return strings.Trim(c.ctx.Ctx.Request.FormValue(name), " ")
	} else {
		if c.ctx.Ctx.Request.Method == http.MethodGet || c.ctx.Ctx.Request.Method == http.MethodDelete {
			return strings.Trim(c.ctx.Ctx.Input.Query(name), " ")
		} else if c.ctx.Ctx.Request.Method == http.MethodPost ||
			c.ctx.Ctx.Request.Method == http.MethodPatch ||
			c.ctx.Ctx.Request.Method == http.MethodPut {
			return strings.Trim(c.ctx.Ctx.Request.FormValue(name), " ")
		}
	}
	return ""
}

// ===========================
func Sync(from interface{}, to interface{}) interface{} {
	_from := reflect.ValueOf(from)
	_fromType := _from.Type()
	_to := reflect.ValueOf(to)

	for i := 0; i < _from.NumField(); i++ {
		fromName := _fromType.Field(i).Name
		field := _to.Elem().FieldByName(fromName)
		if !_from.Field(i).IsNil() && field.IsValid() && field.CanSet() {
			fromValue := _from.Field(i).Elem()
			fromType := reflect.TypeOf(fromValue.Interface())
			if fromType.String() == "uuid.UUID" {
				if fromValue.Interface() != uuid.Nil {
					field.Set(fromValue)
				}
			} else if fromType.String() == "string" {
				if field.Kind() == reflect.Ptr {
					tmp := fromValue.String()
					field.Set(reflect.ValueOf(&tmp))
				} else {
					field.Set(fromValue)
				}
			} else if fromType.String() == "service.Time" {
				tmp := fromValue.Interface().(Time)
				if tmp.IsSet() {
					if field.Kind() == reflect.Ptr {
						field.Set(reflect.ValueOf(&tmp))
					} else {
						field.Set(fromValue)
					}
				}
			} else {
				field.Set(fromValue)
			}
		}
	}
	return to
}

func (c *chain) Uuid(name string) (uuid.UUID, error) {
	v := c.Get(name)
	return uuid.Parse(v)
}
func (c *chain) int(name string) (int, error) {
	v := c.Get(name)
	return strconv.Atoi(v)
}
func (c *chain) date(name string) (time.Time, error) {
	v := c.Get(name)
	return time.Parse(time.RFC3339, v)
}
func (c *chain) time(name string) (time.Time, error) {
	v := c.Get(name)
	return time.Parse("15:04", v)
}
func (c *chain) bool(name string) bool {
	v := c.Get(name)
	return strings.ToLower(v) == "true" || v == "1"
}

//---------------------------------------------------------------
// Special methods apply directly to query
//---------------------------------------------------------------
func (c *chain) db(query *gorm.DB) *chain {
	c.q = query
	return c
}

// ------------------------------------------------

func (c *chain) query() *gorm.DB {
	return c.q
}

// return DB error and Chain error if have
func (c *chain) Preload(table string) *chain {
	c.q = c.q.Preload(table)
	return c
}

// return DB error and Chain error if have
func (c *chain) Find(output interface{}) (error, error) {
	return c.q.Find(output).Error, c.err
}

// return DB error and Chain error if have
func (c *chain) update(m interface{}) (error, error) {
	return c.q.Model(m).Update(c.updater).Error, c.err
}

// ------------------------------------------------

func (c *chain) Order(name string) *chain {
	if c.q == nil {
		c.q, _ = postgres2.GetDB("default")
	}
	v := c.Get(name)
	if v != "" {
		c.q = c.q.Order(v, true)
	}
	return c
}

func (c *chain) Pagination(model interface{}) *chain {
	if c.err != nil {
		return c
	}
	if c.q == nil {
		c.q, _ = postgres2.GetDB("default")
	}
	currentPage, perPage := paginationQuery(c.ctx)

	limit := perPage
	page := currentPage

	totalRecords := 0
	c.q.Model(model).Count(&totalRecords)
	lastPage := floorOf(totalRecords, perPage)
	if lastPage*perPage < totalRecords {
		lastPage++
	}
	lastPage = max(1, lastPage)
	currentPage = min(currentPage, lastPage)
	nextPage := currentPage + 1
	if nextPage > lastPage {
		nextPage = 0
	}
	// add to query
	c.q = c.q.Limit(limit).Offset((page - 1) * limit)
	// set header
	c.ctx.Ctx.ResponseWriter.Header().Add("X-Page", fmt.Sprint(currentPage))
	c.ctx.Ctx.ResponseWriter.Header().Add("X-Per-Page", fmt.Sprint(perPage))
	c.ctx.Ctx.ResponseWriter.Header().Add("X-Next-Page", fmt.Sprint(nextPage))
	c.ctx.Ctx.ResponseWriter.Header().Add("X-Last-Page", fmt.Sprint(lastPage))
	c.ctx.Ctx.ResponseWriter.Header().Add("X-Total-Items", fmt.Sprint(totalRecords))
	return c
}

//---------------------------------------------------------------
// Almighty combine methods
//---------------------------------------------------------------
func (c *chain) validateChain() *chain {
	if c.err != nil {
		return c
	}
	if c.q == nil {
		c.q, _ = postgres2.GetDB("default")
	}
	return nil
}

// Example: searchUUID("product_id", []string{"product_id", "=", "?"})
func (c *chain) WhereUUID(name string, domain []string) *chain {
	if c := c.validateChain(); c != nil {
		return c
	}
	if c.Get(name) != "" {
		if v, err := c.Uuid(name); err != nil {
			c.err = err
		} else {
			c.q = c.q.Where(strings.Join(domain, " "), v)
		}
	}
	return c
}

func (c *chain) WhereString(name string, domain []string) *chain {
	if c := c.validateChain(); c != nil {
		return c
	}
	if v := c.Get(name); v != "" {
		c.q = c.q.Where(strings.Join(domain, " "), v)
	}
	return c
}

func (c *chain) WhereStringAdv(name string, domain []string, pre string, sub string) *chain {
	if c := c.validateChain(); c != nil {
		return c
	}
	if v := c.Get(name); v != "" {
		c.q = c.q.Where(strings.Join(domain, " "), pre+v+sub)
	}
	return c
}

func (c *chain) WhereDate(name string, domain []string) *chain {
	if c := c.validateChain(); c != nil {
		return c
	}
	if c.Get(name) != "" {
		if v, err := c.date(name); err != nil {
			c.err = err
		} else {
			c.q = c.q.Where(strings.Join(domain, " "), v)
		}
	}
	return c
}

func (c *chain) WhereTime(name string, domain []string) *chain {
	if c := c.validateChain(); c != nil {
		return c
	}
	if c.Get(name) != "" {
		if _, err := c.time(name); err != nil {
			c.err = err
		} else {
			c.q = c.q.Where(strings.Join(domain, " "), c.Get(name))
		}
	}
	return c
}

func (c *chain) WhereInt(name string, domain []string) *chain {
	if c := c.validateChain(); c != nil {
		return c
	}
	if c.Get(name) != "" {
		if v, err := c.int(name); err != nil {
			c.err = err
		} else {
			c.q = c.q.Where(strings.Join(domain, " "), v)
		}
	}
	return c
}

func (c *chain) WhereBool(name string, domain []string) *chain {
	if c := c.validateChain(); c != nil {
		return c
	}
	if c.Get(name) != "" {
		c.q = c.q.Where(strings.Join(domain, " "), c.bool(name))
	}
	return c
}

func (c *chain) WhereSearch(name string) *chain {
	if c := c.validateChain(); c != nil {
		return c
	}
	if search := c.Get(name); search != "" {
		c.q = c.q.Where("admin_name ilike ? OR input_type ilike ? OR display_names->>'en' ilike ? OR display_names->>'vi' ilike ? OR display_labels->>'en' ilike ? OR display_labels->>'vi' ilike ?", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}
	return c
}

//---------------------------------------------------------------
// Almighty combine methods
//---------------------------------------------------------------

func (c *chain) updateUUIDDirect(name string, v uuid.UUID) *chain {
	if c := c.validateChain(); c != nil {
		return c
	}
	c.updater[name] = v
	return c
}

func (c *chain) updateUUID(name string) *chain {
	if c := c.validateChain(); c != nil {
		return c
	}
	if c.Get(name) != "" {
		if v, err := c.Uuid(name); err != nil {
			c.err = err
		} else {
			c.updater[name] = v
		}
	}
	return c
}

func (c *chain) updateString(name string) *chain {
	if c := c.validateChain(); c != nil {
		return c
	}
	if c.updater == nil {
		c.updater = make(map[string]interface{})
	}
	if c.Get(name) != "" {
		if v := c.Get(name); v != "" {
			c.updater[name] = v
		}
	}
	return c
}

func (c *chain) updateDate(name string) *chain {
	if c := c.validateChain(); c != nil {
		return c
	}
	if c.updater == nil {
		c.updater = make(map[string]interface{})
	}
	if c.Get(name) != "" {
		if v, err := c.date(name); err != nil {
			c.err = err
		} else {
			c.updater[name] = v
		}
	}
	return c
}

func (c *chain) updateTime(name string) *chain {
	if c := c.validateChain(); c != nil {
		return c
	}
	if c.updater == nil {
		c.updater = make(map[string]interface{})
	}
	if c.Get(name) != "" {
		if _, err := c.time(name); err == nil {
			c.updater[name] = c.Get(name)
		} else {
			c.err = err
		}
	}
	return c
}

func (c *chain) updateInt(name string) *chain {
	if c := c.validateChain(); c != nil {
		return c
	}
	if c.updater == nil {
		c.updater = make(map[string]interface{})
	}
	if c.Get(name) != "" {
		if v, err := c.int(name); err != nil {
			c.err = err
		} else {
			c.updater[name] = v
		}
	}
	return c
}

func (c *chain) updateBool(name string) *chain {
	if c := c.validateChain(); c != nil {
		return c
	}
	if c.updater == nil {
		c.updater = make(map[string]interface{})
	}
	if c.Get(name) != "" {
		c.updater[name] = c.bool(name)
	}
	return c
}

func (c *chain) updateJsonB(name string) *chain {
	if c := c.validateChain(); c != nil {
		return c
	}
	if c.updater == nil {
		c.updater = make(map[string]interface{})
	}
	if v := c.Get(name); v != "" {
		c.updater[name] = jsonb(v)
	}
	return c
}
