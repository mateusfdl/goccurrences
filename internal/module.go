package internal

import (
	"github.com/mateusfdl/go-poc/internal/achievements"
	"github.com/mateusfdl/go-poc/internal/grpc"
	"github.com/mateusfdl/go-poc/internal/logger"
	"github.com/mateusfdl/go-poc/internal/mongo"
	"github.com/mateusfdl/go-poc/internal/occurrences"
	"go.uber.org/fx"
)

var CoreModules = fx.Options(
	occurrences.Module,
	achievements.Module,
)

var AdapterModules = fx.Options(
	mongo.Module,
	logger.Module,
	grpc.Module,
)
