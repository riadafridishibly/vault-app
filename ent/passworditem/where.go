// Code generated by ent, DO NOT EDIT.

package passworditem

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/riadafridishibly/vault-app/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// Avatar applies equality check predicate on the "avatar" field. It's identical to AvatarEQ.
func Avatar(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvatar), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// SiteName applies equality check predicate on the "site_name" field. It's identical to SiteNameEQ.
func SiteName(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSiteName), v))
	})
}

// SiteURL applies equality check predicate on the "site_url" field. It's identical to SiteURLEQ.
func SiteURL(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSiteURL), v))
	})
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// UsernameType applies equality check predicate on the "username_type" field. It's identical to UsernameTypeEQ.
func UsernameType(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsernameType), v))
	})
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.PasswordItem {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.PasswordItem {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.PasswordItem {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.PasswordItem {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// AvatarEQ applies the EQ predicate on the "avatar" field.
func AvatarEQ(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvatar), v))
	})
}

// AvatarNEQ applies the NEQ predicate on the "avatar" field.
func AvatarNEQ(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAvatar), v))
	})
}

// AvatarIn applies the In predicate on the "avatar" field.
func AvatarIn(vs ...string) predicate.PasswordItem {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAvatar), v...))
	})
}

// AvatarNotIn applies the NotIn predicate on the "avatar" field.
func AvatarNotIn(vs ...string) predicate.PasswordItem {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAvatar), v...))
	})
}

// AvatarGT applies the GT predicate on the "avatar" field.
func AvatarGT(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAvatar), v))
	})
}

// AvatarGTE applies the GTE predicate on the "avatar" field.
func AvatarGTE(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAvatar), v))
	})
}

// AvatarLT applies the LT predicate on the "avatar" field.
func AvatarLT(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAvatar), v))
	})
}

// AvatarLTE applies the LTE predicate on the "avatar" field.
func AvatarLTE(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAvatar), v))
	})
}

// AvatarContains applies the Contains predicate on the "avatar" field.
func AvatarContains(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAvatar), v))
	})
}

// AvatarHasPrefix applies the HasPrefix predicate on the "avatar" field.
func AvatarHasPrefix(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAvatar), v))
	})
}

// AvatarHasSuffix applies the HasSuffix predicate on the "avatar" field.
func AvatarHasSuffix(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAvatar), v))
	})
}

// AvatarIsNil applies the IsNil predicate on the "avatar" field.
func AvatarIsNil() predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAvatar)))
	})
}

// AvatarNotNil applies the NotNil predicate on the "avatar" field.
func AvatarNotNil() predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAvatar)))
	})
}

// AvatarEqualFold applies the EqualFold predicate on the "avatar" field.
func AvatarEqualFold(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAvatar), v))
	})
}

// AvatarContainsFold applies the ContainsFold predicate on the "avatar" field.
func AvatarContainsFold(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAvatar), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.PasswordItem {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.PasswordItem {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDescription)))
	})
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDescription)))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// SiteNameEQ applies the EQ predicate on the "site_name" field.
func SiteNameEQ(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSiteName), v))
	})
}

// SiteNameNEQ applies the NEQ predicate on the "site_name" field.
func SiteNameNEQ(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSiteName), v))
	})
}

// SiteNameIn applies the In predicate on the "site_name" field.
func SiteNameIn(vs ...string) predicate.PasswordItem {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSiteName), v...))
	})
}

// SiteNameNotIn applies the NotIn predicate on the "site_name" field.
func SiteNameNotIn(vs ...string) predicate.PasswordItem {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSiteName), v...))
	})
}

// SiteNameGT applies the GT predicate on the "site_name" field.
func SiteNameGT(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSiteName), v))
	})
}

// SiteNameGTE applies the GTE predicate on the "site_name" field.
func SiteNameGTE(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSiteName), v))
	})
}

// SiteNameLT applies the LT predicate on the "site_name" field.
func SiteNameLT(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSiteName), v))
	})
}

// SiteNameLTE applies the LTE predicate on the "site_name" field.
func SiteNameLTE(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSiteName), v))
	})
}

// SiteNameContains applies the Contains predicate on the "site_name" field.
func SiteNameContains(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSiteName), v))
	})
}

// SiteNameHasPrefix applies the HasPrefix predicate on the "site_name" field.
func SiteNameHasPrefix(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSiteName), v))
	})
}

// SiteNameHasSuffix applies the HasSuffix predicate on the "site_name" field.
func SiteNameHasSuffix(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSiteName), v))
	})
}

// SiteNameIsNil applies the IsNil predicate on the "site_name" field.
func SiteNameIsNil() predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSiteName)))
	})
}

// SiteNameNotNil applies the NotNil predicate on the "site_name" field.
func SiteNameNotNil() predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSiteName)))
	})
}

// SiteNameEqualFold applies the EqualFold predicate on the "site_name" field.
func SiteNameEqualFold(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSiteName), v))
	})
}

// SiteNameContainsFold applies the ContainsFold predicate on the "site_name" field.
func SiteNameContainsFold(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSiteName), v))
	})
}

// SiteURLEQ applies the EQ predicate on the "site_url" field.
func SiteURLEQ(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSiteURL), v))
	})
}

// SiteURLNEQ applies the NEQ predicate on the "site_url" field.
func SiteURLNEQ(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSiteURL), v))
	})
}

// SiteURLIn applies the In predicate on the "site_url" field.
func SiteURLIn(vs ...string) predicate.PasswordItem {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSiteURL), v...))
	})
}

// SiteURLNotIn applies the NotIn predicate on the "site_url" field.
func SiteURLNotIn(vs ...string) predicate.PasswordItem {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSiteURL), v...))
	})
}

// SiteURLGT applies the GT predicate on the "site_url" field.
func SiteURLGT(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSiteURL), v))
	})
}

// SiteURLGTE applies the GTE predicate on the "site_url" field.
func SiteURLGTE(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSiteURL), v))
	})
}

// SiteURLLT applies the LT predicate on the "site_url" field.
func SiteURLLT(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSiteURL), v))
	})
}

// SiteURLLTE applies the LTE predicate on the "site_url" field.
func SiteURLLTE(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSiteURL), v))
	})
}

// SiteURLContains applies the Contains predicate on the "site_url" field.
func SiteURLContains(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSiteURL), v))
	})
}

// SiteURLHasPrefix applies the HasPrefix predicate on the "site_url" field.
func SiteURLHasPrefix(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSiteURL), v))
	})
}

// SiteURLHasSuffix applies the HasSuffix predicate on the "site_url" field.
func SiteURLHasSuffix(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSiteURL), v))
	})
}

// SiteURLIsNil applies the IsNil predicate on the "site_url" field.
func SiteURLIsNil() predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSiteURL)))
	})
}

// SiteURLNotNil applies the NotNil predicate on the "site_url" field.
func SiteURLNotNil() predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSiteURL)))
	})
}

// SiteURLEqualFold applies the EqualFold predicate on the "site_url" field.
func SiteURLEqualFold(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSiteURL), v))
	})
}

// SiteURLContainsFold applies the ContainsFold predicate on the "site_url" field.
func SiteURLContainsFold(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSiteURL), v))
	})
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsername), v))
	})
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUsername), v))
	})
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.PasswordItem {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUsername), v...))
	})
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.PasswordItem {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUsername), v...))
	})
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUsername), v))
	})
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUsername), v))
	})
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUsername), v))
	})
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUsername), v))
	})
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUsername), v))
	})
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUsername), v))
	})
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUsername), v))
	})
}

// UsernameIsNil applies the IsNil predicate on the "username" field.
func UsernameIsNil() predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUsername)))
	})
}

// UsernameNotNil applies the NotNil predicate on the "username" field.
func UsernameNotNil() predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUsername)))
	})
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUsername), v))
	})
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUsername), v))
	})
}

// UsernameTypeEQ applies the EQ predicate on the "username_type" field.
func UsernameTypeEQ(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUsernameType), v))
	})
}

// UsernameTypeNEQ applies the NEQ predicate on the "username_type" field.
func UsernameTypeNEQ(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUsernameType), v))
	})
}

// UsernameTypeIn applies the In predicate on the "username_type" field.
func UsernameTypeIn(vs ...string) predicate.PasswordItem {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUsernameType), v...))
	})
}

// UsernameTypeNotIn applies the NotIn predicate on the "username_type" field.
func UsernameTypeNotIn(vs ...string) predicate.PasswordItem {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUsernameType), v...))
	})
}

// UsernameTypeGT applies the GT predicate on the "username_type" field.
func UsernameTypeGT(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUsernameType), v))
	})
}

// UsernameTypeGTE applies the GTE predicate on the "username_type" field.
func UsernameTypeGTE(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUsernameType), v))
	})
}

// UsernameTypeLT applies the LT predicate on the "username_type" field.
func UsernameTypeLT(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUsernameType), v))
	})
}

// UsernameTypeLTE applies the LTE predicate on the "username_type" field.
func UsernameTypeLTE(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUsernameType), v))
	})
}

// UsernameTypeContains applies the Contains predicate on the "username_type" field.
func UsernameTypeContains(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUsernameType), v))
	})
}

// UsernameTypeHasPrefix applies the HasPrefix predicate on the "username_type" field.
func UsernameTypeHasPrefix(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUsernameType), v))
	})
}

// UsernameTypeHasSuffix applies the HasSuffix predicate on the "username_type" field.
func UsernameTypeHasSuffix(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUsernameType), v))
	})
}

// UsernameTypeIsNil applies the IsNil predicate on the "username_type" field.
func UsernameTypeIsNil() predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUsernameType)))
	})
}

// UsernameTypeNotNil applies the NotNil predicate on the "username_type" field.
func UsernameTypeNotNil() predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUsernameType)))
	})
}

// UsernameTypeEqualFold applies the EqualFold predicate on the "username_type" field.
func UsernameTypeEqualFold(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUsernameType), v))
	})
}

// UsernameTypeContainsFold applies the ContainsFold predicate on the "username_type" field.
func UsernameTypeContainsFold(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUsernameType), v))
	})
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPassword), v))
	})
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.PasswordItem {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPassword), v...))
	})
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.PasswordItem {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPassword), v...))
	})
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPassword), v))
	})
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPassword), v))
	})
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPassword), v))
	})
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPassword), v))
	})
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPassword), v))
	})
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPassword), v))
	})
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPassword), v))
	})
}

// PasswordIsNil applies the IsNil predicate on the "password" field.
func PasswordIsNil() predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPassword)))
	})
}

// PasswordNotNil applies the NotNil predicate on the "password" field.
func PasswordNotNil() predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPassword)))
	})
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPassword), v))
	})
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPassword), v))
	})
}

// TagsIsNil applies the IsNil predicate on the "tags" field.
func TagsIsNil() predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTags)))
	})
}

// TagsNotNil applies the NotNil predicate on the "tags" field.
func TagsNotNil() predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTags)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PasswordItem) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PasswordItem) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PasswordItem) predicate.PasswordItem {
	return predicate.PasswordItem(func(s *sql.Selector) {
		p(s.Not())
	})
}
