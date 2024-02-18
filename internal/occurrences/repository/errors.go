package repository

import (
	"errors"
)

var (
	ErrAggregationPipeline = errors.New("error in aggregation pipeline")
	ErrListUserOccurrences = errors.New("error listing user occurrences")
)
