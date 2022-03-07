package model

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Tags struct {
	labels map[string]bool
}

func NewTags(tags ...string) *Tags {
	labels := make(map[string]bool, len(tags))
	for _, tag := range tags {
		labels[tag] = true
	}
	return &Tags{labels: labels}
}

func (t *Tags) Includes(other *Tags) bool {
	for tag := range other.labels {
		_, ok := t.labels[tag]
		if !ok {
			return false
		}
	}
	return true
}

func (t *Tags) Excludes(other *Tags) bool {
	for tag := range t.labels {
		_, ok := other.labels[tag]
		if !ok {
			return false
		}
	}
	return true
}

func (t *Tags) UnmarshalJSON(data []byte) error {
	var arr []string
	err := json.Unmarshal(data, &arr)
	if err != nil {
		return errors.Wrap(err, "tags.unmarshal")
	}

	t.labels = make(map[string]bool, len(arr))
	for _, tag := range arr {
		t.labels[tag] = true
	}

	return nil
}

func (t *Tags) MarshalJSON() ([]byte, error) {
	arr := []string{}
	for tag := range t.labels {
		arr = append(arr, tag)
	}

	data, err := json.Marshal(arr)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return data, nil
}

func (t *Tags) Len() int {
	return len(t.labels)
}

func (t *Tags) Include(other *Tags) {
	for tag := range other.labels {
		t.labels[tag] = true
	}
}

func (t *Tags) Exclude(other *Tags) {
	for tag := range other.labels {
		delete(t.labels, tag)
	}
}
