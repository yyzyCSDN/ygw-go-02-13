package query

import (
	"fmt"
	"sort"
	"strings"

	"example.com/paymentledger/model"
)

type Filter struct {
	Status      string
	Tag         string
	Active      *bool
	MinPriority *int
	MaxPriority *int
	Name        string
}

type SortField string

const (
	SortByID       SortField = "id"
	SortByName     SortField = "name"
	SortByPriority SortField = "priority"
	SortByAmount   SortField = "amount"
)

func Select(values []model.Payment, filter Filter) []model.Payment {
	result := make([]model.Payment, 0)
	for _, value := range values {
		if filter.Status != "" && value.Status != filter.Status {
			continue
		}
		if filter.Tag != "" && !value.HasTag(filter.Tag) {
			continue
		}
		if filter.Active != nil && value.Active != *filter.Active {
			continue
		}
		if filter.MinPriority != nil && value.Priority < *filter.MinPriority {
			continue
		}
		if filter.MaxPriority != nil && value.Priority > *filter.MaxPriority {
			continue
		}
		if filter.Name != "" && !strings.Contains(strings.ToLower(value.Name), strings.ToLower(filter.Name)) {
			continue
		}
		result = append(result, value.Clone())
	}
	return result
}

func Sort(values []model.Payment, field SortField, descending bool) ([]model.Payment, error) {
	result := make([]model.Payment, len(values))
	for index, value := range values {
		result[index] = value.Clone()
	}
	less := func(i, j int) bool { return false }
	switch field {
	case SortByID:
		less = func(i, j int) bool { return result[i].ID < result[j].ID }
	case SortByName:
		less = func(i, j int) bool { return result[i].Name < result[j].Name }
	case SortByPriority:
		less = func(i, j int) bool { return result[i].Priority < result[j].Priority }
	case SortByAmount:
		less = func(i, j int) bool { return result[i].Amount < result[j].Amount }
	default:
		return nil, fmt.Errorf("unsupported sort field %q", field)
	}
	sort.SliceStable(result, func(i, j int) bool {
		if descending {
			return less(j, i)
		}
		return less(i, j)
	})
	return result, nil
}

func Page(values []model.Payment, offset, limit int) ([]model.Payment, error) {
	if offset < 0 || limit < 0 {
		return nil, fmt.Errorf("offset and limit cannot be negative")
	}
	if offset >= len(values) || limit == 0 {
		return []model.Payment{}, nil
	}
	end := offset + limit
	if end > len(values) {
		end = len(values)
	}
	result := make([]model.Payment, end-offset)
	for index := range result {
		result[index] = values[offset+index].Clone()
	}
	return result, nil
}

func GroupByStatus(values []model.Payment) map[string][]model.Payment {
	result := map[string][]model.Payment{}
	for _, value := range values {
		result[value.Status] = append(result[value.Status], value.Clone())
	}
	return result
}

func IDs(values []model.Payment) []string {
	result := make([]string, len(values))
	for index, value := range values {
		result[index] = value.ID
	}
	return result
}
