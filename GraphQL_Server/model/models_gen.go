// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type EnterpriseInfo2 struct {
	ID             string `json:"id"`
	EnterpriseID   string `json:"enterprise_id"`
	EnterpriseName string `json:"enterprise_name"`
}

type NewEnterpriseInfo struct {
	EnterpriseName string `json:"enterprise_name"`
}