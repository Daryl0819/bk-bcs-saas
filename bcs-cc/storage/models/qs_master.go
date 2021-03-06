// Code generated by go-queryset. DO NOT EDIT.
package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set ManagerMasterQuerySet

// ManagerMasterQuerySet is an queryset type for ManagerMaster
type ManagerMasterQuerySet struct {
	db *gorm.DB
}

// NewManagerMasterQuerySet constructs new ManagerMasterQuerySet
func NewManagerMasterQuerySet(db *gorm.DB) ManagerMasterQuerySet {
	return ManagerMasterQuerySet{
		db: db.Model(&ManagerMaster{}),
	}
}

func (qs ManagerMasterQuerySet) w(db *gorm.DB) ManagerMasterQuerySet {
	return NewManagerMasterQuerySet(db)
}

// Create is an autogenerated method
// nolint: dupl
func (o *ManagerMaster) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *ManagerMaster) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// All is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) All(ret *[]ManagerMaster) error {
	return qs.db.Find(ret).Error
}

// BackupEq is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) BackupEq(backup string) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("backup = ?", backup))
}

// BackupIn is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) BackupIn(backup ...string) ManagerMasterQuerySet {
	if len(backup) == 0 {
		qs.db.AddError(errors.New("must at least pass one backup in BackupIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("backup IN (?)", backup))
}

// BackupNe is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) BackupNe(backup string) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("backup != ?", backup))
}

// BackupNotIn is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) BackupNotIn(backup ...string) ManagerMasterQuerySet {
	if len(backup) == 0 {
		qs.db.AddError(errors.New("must at least pass one backup in BackupNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("backup NOT IN (?)", backup))
}

// ClusterIDEq is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) ClusterIDEq(clusterID string) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("cluster_id = ?", clusterID))
}

// ClusterIDIn is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) ClusterIDIn(clusterID ...string) ManagerMasterQuerySet {
	if len(clusterID) == 0 {
		qs.db.AddError(errors.New("must at least pass one clusterID in ClusterIDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("cluster_id IN (?)", clusterID))
}

// ClusterIDNe is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) ClusterIDNe(clusterID string) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("cluster_id != ?", clusterID))
}

// ClusterIDNotIn is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) ClusterIDNotIn(clusterID ...string) ManagerMasterQuerySet {
	if len(clusterID) == 0 {
		qs.db.AddError(errors.New("must at least pass one clusterID in ClusterIDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("cluster_id NOT IN (?)", clusterID))
}

// Count is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) CreatedAtEq(createdAt time.Time) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) CreatedAtGt(createdAt time.Time) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) CreatedAtGte(createdAt time.Time) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) CreatedAtLt(createdAt time.Time) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) CreatedAtLte(createdAt time.Time) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) CreatedAtNe(createdAt time.Time) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) Delete() error {
	return qs.db.Delete(ManagerMaster{}).Error
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(ManagerMaster{})
	return db.RowsAffected, db.Error
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(ManagerMaster{})
	return db.RowsAffected, db.Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) DeletedAtEq(deletedAt time.Time) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) DeletedAtGt(deletedAt time.Time) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) DeletedAtGte(deletedAt time.Time) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) DeletedAtIsNotNull() ManagerMasterQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) DeletedAtIsNull() ManagerMasterQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) DeletedAtLt(deletedAt time.Time) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) DeletedAtLte(deletedAt time.Time) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) DeletedAtNe(deletedAt time.Time) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// ExtendedInfoEq is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) ExtendedInfoEq(extendedInfo string) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("extended_info = ?", extendedInfo))
}

// ExtendedInfoIn is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) ExtendedInfoIn(extendedInfo ...string) ManagerMasterQuerySet {
	if len(extendedInfo) == 0 {
		qs.db.AddError(errors.New("must at least pass one extendedInfo in ExtendedInfoIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("extended_info IN (?)", extendedInfo))
}

// ExtendedInfoNe is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) ExtendedInfoNe(extendedInfo string) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("extended_info != ?", extendedInfo))
}

// ExtendedInfoNotIn is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) ExtendedInfoNotIn(extendedInfo ...string) ManagerMasterQuerySet {
	if len(extendedInfo) == 0 {
		qs.db.AddError(errors.New("must at least pass one extendedInfo in ExtendedInfoNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("extended_info NOT IN (?)", extendedInfo))
}

// ExtraEq is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) ExtraEq(extra string) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("extra = ?", extra))
}

// ExtraIn is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) ExtraIn(extra ...string) ManagerMasterQuerySet {
	if len(extra) == 0 {
		qs.db.AddError(errors.New("must at least pass one extra in ExtraIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("extra IN (?)", extra))
}

// ExtraNe is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) ExtraNe(extra string) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("extra != ?", extra))
}

// ExtraNotIn is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) ExtraNotIn(extra ...string) ManagerMasterQuerySet {
	if len(extra) == 0 {
		qs.db.AddError(errors.New("must at least pass one extra in ExtraNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("extra NOT IN (?)", extra))
}

// GetDB is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) GetDB() *gorm.DB {
	return qs.db
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) GetUpdater() ManagerMasterUpdater {
	return NewManagerMasterUpdater(qs.db)
}

// HostnameEq is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) HostnameEq(hostname string) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("hostname = ?", hostname))
}

// HostnameIn is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) HostnameIn(hostname ...string) ManagerMasterQuerySet {
	if len(hostname) == 0 {
		qs.db.AddError(errors.New("must at least pass one hostname in HostnameIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("hostname IN (?)", hostname))
}

// HostnameNe is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) HostnameNe(hostname string) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("hostname != ?", hostname))
}

// HostnameNotIn is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) HostnameNotIn(hostname ...string) ManagerMasterQuerySet {
	if len(hostname) == 0 {
		qs.db.AddError(errors.New("must at least pass one hostname in HostnameNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("hostname NOT IN (?)", hostname))
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) IDEq(ID uint) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) IDGt(ID uint) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) IDGte(ID uint) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) IDIn(ID ...uint) ManagerMasterQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id IN (?)", ID))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) IDLt(ID uint) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) IDLte(ID uint) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) IDNe(ID uint) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) IDNotIn(ID ...uint) ManagerMasterQuerySet {
	if len(ID) == 0 {
		qs.db.AddError(errors.New("must at least pass one ID in IDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", ID))
}

// InnerIPEq is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) InnerIPEq(innerIP string) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("inner_ip = ?", innerIP))
}

// InnerIPIn is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) InnerIPIn(innerIP ...string) ManagerMasterQuerySet {
	if len(innerIP) == 0 {
		qs.db.AddError(errors.New("must at least pass one innerIP in InnerIPIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("inner_ip IN (?)", innerIP))
}

// InnerIPNe is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) InnerIPNe(innerIP string) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("inner_ip != ?", innerIP))
}

// InnerIPNotIn is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) InnerIPNotIn(innerIP ...string) ManagerMasterQuerySet {
	if len(innerIP) == 0 {
		qs.db.AddError(errors.New("must at least pass one innerIP in InnerIPNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("inner_ip NOT IN (?)", innerIP))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) Limit(limit int) ManagerMasterQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// NumEq is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) NumEq(num int) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("num = ?", num))
}

// NumGt is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) NumGt(num int) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("num > ?", num))
}

// NumGte is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) NumGte(num int) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("num >= ?", num))
}

// NumIn is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) NumIn(num ...int) ManagerMasterQuerySet {
	if len(num) == 0 {
		qs.db.AddError(errors.New("must at least pass one num in NumIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("num IN (?)", num))
}

// NumLt is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) NumLt(num int) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("num < ?", num))
}

// NumLte is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) NumLte(num int) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("num <= ?", num))
}

// NumNe is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) NumNe(num int) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("num != ?", num))
}

// NumNotIn is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) NumNotIn(num ...int) ManagerMasterQuerySet {
	if len(num) == 0 {
		qs.db.AddError(errors.New("must at least pass one num in NumNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("num NOT IN (?)", num))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) Offset(offset int) ManagerMasterQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs ManagerMasterQuerySet) One(ret *ManagerMaster) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) OrderAscByCreatedAt() ManagerMasterQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) OrderAscByDeletedAt() ManagerMasterQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) OrderAscByID() ManagerMasterQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByNum is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) OrderAscByNum() ManagerMasterQuerySet {
	return qs.w(qs.db.Order("num ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) OrderAscByUpdatedAt() ManagerMasterQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) OrderDescByCreatedAt() ManagerMasterQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) OrderDescByDeletedAt() ManagerMasterQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) OrderDescByID() ManagerMasterQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByNum is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) OrderDescByNum() ManagerMasterQuerySet {
	return qs.w(qs.db.Order("num DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) OrderDescByUpdatedAt() ManagerMasterQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// ProjectIDEq is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) ProjectIDEq(projectID string) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("project_id = ?", projectID))
}

// ProjectIDIn is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) ProjectIDIn(projectID ...string) ManagerMasterQuerySet {
	if len(projectID) == 0 {
		qs.db.AddError(errors.New("must at least pass one projectID in ProjectIDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("project_id IN (?)", projectID))
}

// ProjectIDNe is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) ProjectIDNe(projectID string) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("project_id != ?", projectID))
}

// ProjectIDNotIn is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) ProjectIDNotIn(projectID ...string) ManagerMasterQuerySet {
	if len(projectID) == 0 {
		qs.db.AddError(errors.New("must at least pass one projectID in ProjectIDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("project_id NOT IN (?)", projectID))
}

// StatusEq is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) StatusEq(status string) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("status = ?", status))
}

// StatusIn is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) StatusIn(status ...string) ManagerMasterQuerySet {
	if len(status) == 0 {
		qs.db.AddError(errors.New("must at least pass one status in StatusIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("status IN (?)", status))
}

// StatusNe is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) StatusNe(status string) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("status != ?", status))
}

// StatusNotIn is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) StatusNotIn(status ...string) ManagerMasterQuerySet {
	if len(status) == 0 {
		qs.db.AddError(errors.New("must at least pass one status in StatusNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("status NOT IN (?)", status))
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) UpdatedAtEq(updatedAt time.Time) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) UpdatedAtGt(updatedAt time.Time) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) UpdatedAtGte(updatedAt time.Time) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) UpdatedAtLt(updatedAt time.Time) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) UpdatedAtLte(updatedAt time.Time) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs ManagerMasterQuerySet) UpdatedAtNe(updatedAt time.Time) ManagerMasterQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// SetBackup is an autogenerated method
// nolint: dupl
func (u ManagerMasterUpdater) SetBackup(backup string) ManagerMasterUpdater {
	u.fields[string(ManagerMasterDBSchema.Backup)] = backup
	return u
}

// SetClusterID is an autogenerated method
// nolint: dupl
func (u ManagerMasterUpdater) SetClusterID(clusterID string) ManagerMasterUpdater {
	u.fields[string(ManagerMasterDBSchema.ClusterID)] = clusterID
	return u
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u ManagerMasterUpdater) SetCreatedAt(createdAt time.Time) ManagerMasterUpdater {
	u.fields[string(ManagerMasterDBSchema.CreatedAt)] = createdAt
	return u
}

// SetDeletedAt is an autogenerated method
// nolint: dupl
func (u ManagerMasterUpdater) SetDeletedAt(deletedAt *time.Time) ManagerMasterUpdater {
	u.fields[string(ManagerMasterDBSchema.DeletedAt)] = deletedAt
	return u
}

// SetExtendedInfo is an autogenerated method
// nolint: dupl
func (u ManagerMasterUpdater) SetExtendedInfo(extendedInfo string) ManagerMasterUpdater {
	u.fields[string(ManagerMasterDBSchema.ExtendedInfo)] = extendedInfo
	return u
}

// SetExtra is an autogenerated method
// nolint: dupl
func (u ManagerMasterUpdater) SetExtra(extra string) ManagerMasterUpdater {
	u.fields[string(ManagerMasterDBSchema.Extra)] = extra
	return u
}

// SetHostname is an autogenerated method
// nolint: dupl
func (u ManagerMasterUpdater) SetHostname(hostname string) ManagerMasterUpdater {
	u.fields[string(ManagerMasterDBSchema.Hostname)] = hostname
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u ManagerMasterUpdater) SetID(ID uint) ManagerMasterUpdater {
	u.fields[string(ManagerMasterDBSchema.ID)] = ID
	return u
}

// SetInnerIP is an autogenerated method
// nolint: dupl
func (u ManagerMasterUpdater) SetInnerIP(innerIP string) ManagerMasterUpdater {
	u.fields[string(ManagerMasterDBSchema.InnerIP)] = innerIP
	return u
}

// SetNum is an autogenerated method
// nolint: dupl
func (u ManagerMasterUpdater) SetNum(num int) ManagerMasterUpdater {
	u.fields[string(ManagerMasterDBSchema.Num)] = num
	return u
}

// SetProjectID is an autogenerated method
// nolint: dupl
func (u ManagerMasterUpdater) SetProjectID(projectID string) ManagerMasterUpdater {
	u.fields[string(ManagerMasterDBSchema.ProjectID)] = projectID
	return u
}

// SetStatus is an autogenerated method
// nolint: dupl
func (u ManagerMasterUpdater) SetStatus(status string) ManagerMasterUpdater {
	u.fields[string(ManagerMasterDBSchema.Status)] = status
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u ManagerMasterUpdater) SetUpdatedAt(updatedAt time.Time) ManagerMasterUpdater {
	u.fields[string(ManagerMasterDBSchema.UpdatedAt)] = updatedAt
	return u
}

// SetInstanceID is an autogenerated method
// nolint: dupl
func (u ManagerMasterUpdater) SetInstanceID(instanceID string) ManagerMasterUpdater {
	u.fields[string(ManagerMasterDBSchema.InstanceID)] = instanceID
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u ManagerMasterUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u ManagerMasterUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// ===== END of query set ManagerMasterQuerySet

// ===== BEGIN of ManagerMaster modifiers

// ManagerMasterDBSchemaField describes database schema field. It requires for method 'Update'
type ManagerMasterDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f ManagerMasterDBSchemaField) String() string {
	return string(f)
}

// ManagerMasterDBSchema stores db field names of ManagerMaster
var ManagerMasterDBSchema = struct {
	ID           ManagerMasterDBSchemaField
	CreatedAt    ManagerMasterDBSchemaField
	UpdatedAt    ManagerMasterDBSchemaField
	DeletedAt    ManagerMasterDBSchemaField
	Extra        ManagerMasterDBSchemaField
	InnerIP      ManagerMasterDBSchemaField
	ClusterID    ManagerMasterDBSchemaField
	ProjectID    ManagerMasterDBSchemaField
	ExtendedInfo ManagerMasterDBSchemaField
	Num          ManagerMasterDBSchemaField
	Backup       ManagerMasterDBSchemaField
	Hostname     ManagerMasterDBSchemaField
	Status       ManagerMasterDBSchemaField
	InstanceID   ManagerMasterDBSchemaField
}{

	ID:           ManagerMasterDBSchemaField("id"),
	CreatedAt:    ManagerMasterDBSchemaField("created_at"),
	UpdatedAt:    ManagerMasterDBSchemaField("updated_at"),
	DeletedAt:    ManagerMasterDBSchemaField("deleted_at"),
	Extra:        ManagerMasterDBSchemaField("extra"),
	InnerIP:      ManagerMasterDBSchemaField("inner_ip"),
	ClusterID:    ManagerMasterDBSchemaField("cluster_id"),
	ProjectID:    ManagerMasterDBSchemaField("project_id"),
	ExtendedInfo: ManagerMasterDBSchemaField("extended_info"),
	Num:          ManagerMasterDBSchemaField("num"),
	Backup:       ManagerMasterDBSchemaField("backup"),
	Hostname:     ManagerMasterDBSchemaField("hostname"),
	Status:       ManagerMasterDBSchemaField("status"),
	InstanceID:   ManagerMasterDBSchemaField("instance_id"),
}

// Update updates ManagerMaster fields by primary key
// nolint: dupl
func (o *ManagerMaster) Update(db *gorm.DB, fields ...ManagerMasterDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":            o.ID,
		"created_at":    o.CreatedAt,
		"updated_at":    o.UpdatedAt,
		"deleted_at":    o.DeletedAt,
		"extra":         o.Extra,
		"inner_ip":      o.InnerIP,
		"cluster_id":    o.ClusterID,
		"project_id":    o.ProjectID,
		"extended_info": o.ExtendedInfo,
		"num":           o.Num,
		"backup":        o.Backup,
		"hostname":      o.Hostname,
		"status":        o.Status,
		"instance_id":   o.InstanceID,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update ManagerMaster %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// ManagerMasterUpdater is an ManagerMaster updates manager
type ManagerMasterUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewManagerMasterUpdater creates new ManagerMaster updater
// nolint: dupl
func NewManagerMasterUpdater(db *gorm.DB) ManagerMasterUpdater {
	return ManagerMasterUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&ManagerMaster{}),
	}
}

// ===== END of ManagerMaster modifiers

// ===== END of all query sets
