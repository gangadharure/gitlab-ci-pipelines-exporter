package controller

import (
	"context"

	"github.com/mvisonneau/gitlab-ci-pipelines-exporter/pkg/schemas"
	"github.com/mvisonneau/gitlab-ci-pipelines-exporter/pkg/store"
	log "github.com/sirupsen/logrus"
)

func metricLogFields(m schemas.Metric) log.Fields {
	return log.Fields{
		"metric-kind":   m.Kind,
		"metric-labels": m.Labels,
	}
}

func storeGetMetric(ctx context.Context, s store.Store, m *schemas.Metric) {
	if err := s.GetMetric(ctx, m); err != nil {
		log.WithFields(
			metricLogFields(*m),
		).WithField(
			"error", err.Error(),
		).Errorf("reading metric from the store")
	}
}

func storeSetMetric(ctx context.Context, s store.Store, m schemas.Metric) {
	if err := s.SetMetric(ctx, m); err != nil {
		log.WithFields(
			metricLogFields(m),
		).WithField(
			"error", err.Error(),
		).Errorf("writing metric in the store")
	}
}

func storeDelMetric(ctx context.Context, s store.Store, m schemas.Metric) {
	if err := s.DelMetric(ctx, m.Key()); err != nil {
		log.WithFields(
			metricLogFields(m),
		).WithField(
			"error", err.Error(),
		).Errorf("deleting metric from the store")
	}
}
