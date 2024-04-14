package data

import (
	"database/sql"
	"fmt"
)

type DBModel struct {
	db *sql.DB
}

func NewDBModel(db *sql.DB) *DBModel {
	return &DBModel{db: db}
}

func (m *DBModel) Insert(info *ModuleInfo) error {
	_, err := m.db.Exec("INSERT INTO your_table_name (id, created_at, updated_at, module_name, module_duration, exam_type, version) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		info.ID, info.CreatedAt, info.UpdatedAt, info.ModuleName, info.ModuleDuration, info.ExamType, info.Version)
	if err != nil {
		return fmt.Errorf("Insert failed: %v", err)
	}
	return nil
}

func (m *DBModel) Retrieve(id uint) (*ModuleInfo, error) {
	row := m.db.QueryRow("SELECT id, created_at, updated_at, module_name, module_duration, exam_type, version FROM your_table_name WHERE id = $1", id)
	var moduleInfo ModuleInfo
	err := row.Scan(&moduleInfo.ID, &moduleInfo.CreatedAt, &moduleInfo.UpdatedAt, &moduleInfo.ModuleName, &moduleInfo.ModuleDuration, &moduleInfo.ExamType, &moduleInfo.Version)
	if err != nil {
		return nil, fmt.Errorf("Retrieve failed: %v", err)
	}
	return &moduleInfo, nil
}

func (m *DBModel) Update(info *ModuleInfo) error {
	_, err := m.db.Exec("UPDATE your_table_name SET created_at = $2, updated_at = $3, module_name = $4, module_duration = $5, exam_type = $6, version = $7 WHERE id = $1",
		info.ID, info.CreatedAt, info.UpdatedAt, info.ModuleName, info.ModuleDuration, info.ExamType, info.Version)
	if err != nil {
		return fmt.Errorf("Update failed: %v", err)
	}
	return nil
}

func (m *DBModel) Delete(id uint) error {
	_, err := m.db.Exec("DELETE FROM your_table_name WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("Delete failed: %v", err)
	}
	return nil
}
