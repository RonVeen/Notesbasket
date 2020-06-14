package main

import (
	"time"
)

type Sort int

const (
	SortDate = iota
	SortTitle
)

type Direction int

const (
	AscendingDirection = iota
	DescendingDirection
)



type CaseType int

const (
	Regular = 1 + iota
	Folder
)

type AuditInfo struct {
	CreatedBy string  `json:"createdBy"`
	CreatedOn time.Time  `json:"createdOn"`
	ModifiedBy string  `json:"modifiedBy"`
	ModifiedOn time.Time  `json:"modifiedOn"`
}

type CaseDefinition struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Sort        Sort      `json:"sort"`
	Direction   Direction `json:"order"`
	AuditInfo   `json:"auditInfo"`
}

