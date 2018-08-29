package operator

import (
	"context"
	"time"

	prom "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"github.com/operator-framework/operator-metering/pkg/operator/prestostore"
)

const (
	// Keep a cap on the number of time ranges we query per reconciliation.
	// If we get to defaultMaxPromTimeRanges, it means we're very backlogged,
	// or we have a small chunkSize and making tons of small queries all one
	// after another will cause undesired resource spikes, or both.  This will
	// make it take longer to catch up, but should help prevent memory from
	// exploding when we end up with a ton of time ranges.

	// defaultMaxPromTimeRanges is the number of time ranges for 24 hours if we
	// query in 5 minute chunks (the default).
	defaultMaxPromTimeRanges = (24 * 60) / 5 // 24 hours, 60 minutes per hour, default chunkSize is 5 minutes

	defaultMaxTimeDuration = 24 * time.Hour
)

var (
	prometheusReportDatasourceLabels = []string{
		"reportdatasource",
		"reportprometheusquery",
		"table_name",
	}

	prometheusReportDatasourceMetricsScrapedCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "metering",
			Name:      "prometheus_reportdatasource_metrics_scraped_total",
			Help:      "Number of Prometheus metrics returned by a PrometheusQuery for a ReportDataSource.",
		},
		prometheusReportDatasourceLabels,
	)

	prometheusReportDatasourceMetricsImportedCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "metering",
			Name:      "prometheus_reportdatasource_metrics_imported_total",
			Help:      "Number of Prometheus ReportDatasource metrics imported.",
		},
		prometheusReportDatasourceLabels,
	)

	prometheusReportDatasourceTotalImportsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "metering",
			Name:      "prometheus_reportdatasource_imports_total",
			Help:      "Number of Prometheus ReportDatasource metrics imports.",
		},
		prometheusReportDatasourceLabels,
	)

	prometheusReportDatasourceFailedImportsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "metering",
			Name:      "prometheus_reportdatasource_failed_imports_total",
			Help:      "Number of failed Prometheus ReportDatasource metrics imports.",
		},
		prometheusReportDatasourceLabels,
	)

	prometheusReportDatasourceTotalPrometheusQueriesCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "metering",
			Name:      "prometheus_reportdatasource_prometheus_queries_total",
			Help:      "Number of Prometheus ReportDatasource Prometheus queries made for the ReportDataSource since start up.",
		},
		prometheusReportDatasourceLabels,
	)

	prometheusReportDatasourceFailedPrometheusQueriesCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "metering",
			Name:      "prometheus_reportdatasource_failed_prometheus_queries_total",
			Help:      "Number of failed Prometheus ReportDatasource Prometheus queries made for the ReportDataSource since start up.",
		},
		prometheusReportDatasourceLabels,
	)

	prometheusReportDatasourceTotalPrestoStoresCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "metering",
			Name:      "prometheus_reportdatasource_presto_stores_total",
			Help:      "Number of Prometheus ReportDatasource calls to store all metrics collected into Presto.",
		},
		prometheusReportDatasourceLabels,
	)

	prometheusReportDatasourceFailedPrestoStoresCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "metering",
			Name:      "prometheus_reportdatasource_failed_presto_stores_total",
			Help:      "Number of failed Prometheus ReportDatasource calls to store all metrics collected into Presto.",
		},
		prometheusReportDatasourceLabels,
	)

	prometheusReportDatasourceImportDurationHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "metering",
			Name:      "prometheus_reportdatasource_import_duration_seconds",
			Help:      "Duration to import Prometheus metrics into Presto.",
			Buckets:   []float64{30.0, 60.0, 300.0},
		},
		prometheusReportDatasourceLabels,
	)

	prometheusReportDatasourcePrometheusQueryDurationHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "metering",
			Name:      "prometheus_reportdatasource_prometheus_query_duration_seconds",
			Help:      "Duration for a Prometheus query to return metrics to reporting-operator.",
			Buckets:   []float64{2.0, 10.0, 30.0, 60.0},
		},
		prometheusReportDatasourceLabels,
	)

	prometheusReportDatasourcePrestoreStoreDurationHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "metering",
			Name:      "prometheus_reportdatasource_presto_store_duration_seconds",
			Help:      "Duration to store all metrics fetched into Presto.",
			Buckets:   []float64{2.0, 10.0, 30.0, 60.0, 300.0},
		},
		prometheusReportDatasourceLabels,
	)

	prometheusReportDatasourceRunningImportsGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "metering",
			Name:      "prometheus_reportdatasource_running_imports",
			Help:      "Number of Prometheus ReportDatasource imports currently running.",
		},
	)
)

func init() {
	prometheus.MustRegister(prometheusReportDatasourceMetricsScrapedCounter)
	prometheus.MustRegister(prometheusReportDatasourceMetricsImportedCounter)
	prometheus.MustRegister(prometheusReportDatasourceTotalImportsCounter)
	prometheus.MustRegister(prometheusReportDatasourceFailedImportsCounter)
	prometheus.MustRegister(prometheusReportDatasourceTotalPrometheusQueriesCounter)
	prometheus.MustRegister(prometheusReportDatasourceFailedPrometheusQueriesCounter)
	prometheus.MustRegister(prometheusReportDatasourceTotalPrestoStoresCounter)
	prometheus.MustRegister(prometheusReportDatasourceFailedPrestoStoresCounter)
	prometheus.MustRegister(prometheusReportDatasourceImportDurationHistogram)
	prometheus.MustRegister(prometheusReportDatasourcePrometheusQueryDurationHistogram)
	prometheus.MustRegister(prometheusReportDatasourcePrestoreStoreDurationHistogram)
	prometheus.MustRegister(prometheusReportDatasourceRunningImportsGauge)
}

func (op *Reporting) runPrometheusImporterWorker(stopCh <-chan struct{}) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// run a go routine that waits for the stopCh to be closed and propagates
	// the shutdown to the collectors by calling cancel()
	go func() {
		<-stopCh
		// if the stopCh is closed while we're waiting, cancel and wait for
		// everything to return
		cancel()
	}()
	op.startPrometheusImporter(ctx)
}

type prometheusImporterFunc func(ctx context.Context, start, end time.Time) error

type prometheusImporterTimeRangeTrigger struct {
	start, end time.Time
	errCh      chan error
}

func (op *Reporting) triggerPrometheusImporterForTimeRange(ctx context.Context, start, end time.Time) error {
	errCh := make(chan error)
	select {
	case op.prometheusImporterTriggerForTimeRangeCh <- prometheusImporterTimeRangeTrigger{start, end, errCh}:
		return <-errCh
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (op *Reporting) startPrometheusImporter(ctx context.Context) {
	logger := op.logger.WithField("component", "PrometheusImporter")
	logger.Infof("PrometheusImporter worker started")
	workers := make(map[string]*prometheusImporterWorker)
	importers := make(map[string]*prestostore.PrometheusImporter)

	const concurrency = 4
	// create a channel to act as a semaphore to limit the number of
	// imports happening in parallel
	semaphore := make(chan struct{}, concurrency)

	defer logger.Infof("PrometheusImporterWorker shutdown")

	if op.cfg.DisablePromsum {
		logger.Infof("Periodic Prometheus ReportDataSource importing disabled")
	}

	for {
		select {
		case <-ctx.Done():
			logger.Infof("got shutdown signal, shutting down PrometheusImporters")
			return
		case trigger := <-op.prometheusImporterTriggerForTimeRangeCh:
			// manually triggered import for a specific time range, usually from HTTP API

			g, ctx := errgroup.WithContext(ctx)
			for dataSourceName, importer := range importers {
				importer := importer
				dataSourceName := dataSourceName
				// collect each dataSource concurrently
				g.Go(func() error {
					return importPrometheusDataSourceData(ctx, logger, semaphore, dataSourceName, importer, func(ctx context.Context, importer *prestostore.PrometheusImporter) ([]prom.Range, error) {
						return importer.ImportMetrics(ctx, trigger.start, trigger.end, true)
					})
				})
			}
			err := g.Wait()
			if err != nil {
				logger.WithError(err).Errorf("PrometheusImporter worker encountered errors while importing data")
			}
			trigger.errCh <- err

		case dataSourceName := <-op.prometheusImporterDeletedDataSourceQueue:
			// if we have a worker for this ReportDataSource then we need to
			// stop it and remove it from our map
			if worker, exists := workers[dataSourceName]; exists {
				worker.stop()
				delete(workers, dataSourceName)
			}
			if _, exists := importers[dataSourceName]; exists {
				delete(importers, dataSourceName)
			}
		case reportDataSource := <-op.prometheusImporterNewDataSourceQueue:
			if reportDataSource.Spec.Promsum == nil {
				logger.Error("expected only Promsum ReportDataSources")
				continue
			}

			dataSourceName := reportDataSource.Name
			queryName := reportDataSource.Spec.Promsum.Query
			tableName := dataSourceTableName(dataSourceName)

			dataSourceLogger := logger.WithFields(logrus.Fields{
				"queryName":        queryName,
				"reportDataSource": dataSourceName,
				"tableName":        tableName,
			})

			reportPromQuery, err := op.informers.Metering().V1alpha1().ReportPrometheusQueries().Lister().ReportPrometheusQueries(reportDataSource.Namespace).Get(queryName)
			if err != nil {
				op.logger.WithError(err).Errorf("unable to ReportPrometheusQuery %s for ReportDataSource %s", queryName, dataSourceName)
				continue
			}

			promQuery := reportPromQuery.Spec.Query

			chunkSize := op.cfg.PrometheusQueryConfig.ChunkSize.Duration
			stepSize := op.cfg.PrometheusQueryConfig.StepSize.Duration
			queryInterval := op.cfg.PrometheusQueryConfig.QueryInterval.Duration

			queryConf := reportDataSource.Spec.Promsum.QueryConfig
			if queryConf != nil {
				if queryConf.ChunkSize != nil {
					chunkSize = queryConf.ChunkSize.Duration
				}
				if queryConf.StepSize != nil {
					stepSize = queryConf.StepSize.Duration
				}
				if queryConf.QueryInterval != nil {
					queryInterval = queryConf.QueryInterval.Duration
				}
			}

			// round to the nearest second for chunk/step sizes
			chunkSize = chunkSize.Truncate(time.Second)
			stepSize = stepSize.Truncate(time.Second)

			cfg := prestostore.Config{
				PrometheusQuery:       promQuery,
				PrestoTableName:       tableName,
				ChunkSize:             chunkSize,
				StepSize:              stepSize,
				MaxTimeRanges:         defaultMaxPromTimeRanges,
				MaxQueryRangeDuration: defaultMaxTimeDuration,
			}

			importer, exists := importers[dataSourceName]
			if exists {
				dataSourceLogger.Debugf("ReportDataSource %s already has an importer, updating configuration", dataSourceName)
				importer.UpdateConfig(cfg)
			} else {
				promLabels := prometheus.Labels{
					"reportdatasource":      dataSourceName,
					"reportprometheusquery": reportPromQuery.Name,
					"table_name":            tableName,
				}

				totalImportsCounter := prometheusReportDatasourceTotalImportsCounter.With(promLabels)
				failedImportsCounter := prometheusReportDatasourceFailedImportsCounter.With(promLabels)

				totalPrometheusQueriesCounter := prometheusReportDatasourceTotalPrometheusQueriesCounter.With(promLabels)
				failedPrometheusQueriesCounter := prometheusReportDatasourceFailedPrometheusQueriesCounter.With(promLabels)

				totalPrestoStoresCounter := prometheusReportDatasourceTotalPrestoStoresCounter.With(promLabels)
				failedPrestoStoresCounter := prometheusReportDatasourceFailedPrestoStoresCounter.With(promLabels)

				promQueryMetricsScrapedCounter := prometheusReportDatasourceMetricsScrapedCounter.With(promLabels)
				promQueryDurationHistogram := prometheusReportDatasourcePrometheusQueryDurationHistogram.With(promLabels)

				metricsImportedCounter := prometheusReportDatasourceMetricsImportedCounter.With(promLabels)
				importDurationHistogram := prometheusReportDatasourceImportDurationHistogram.With(promLabels)

				prestoStoreDurationHistogram := prometheusReportDatasourcePrestoreStoreDurationHistogram.With(promLabels)

				metricsCollectors := prestostore.ImporterMetricsCollectors{
					TotalImportsCounter:     totalImportsCounter,
					FailedImportsCounter:    failedImportsCounter,
					ImportDurationHistogram: importDurationHistogram,

					TotalPrometheusQueriesCounter:    totalPrometheusQueriesCounter,
					FailedPrometheusQueriesCounter:   failedPrometheusQueriesCounter,
					PrometheusQueryDurationHistogram: promQueryDurationHistogram,

					TotalPrestoStoresCounter:     totalPrestoStoresCounter,
					FailedPrestoStoresCounter:    failedPrestoStoresCounter,
					PrestoStoreDurationHistogram: prestoStoreDurationHistogram,

					MetricsScrapedCounter:  promQueryMetricsScrapedCounter,
					MetricsImportedCounter: metricsImportedCounter,
				}

				importer = prestostore.NewPrometheusImporter(dataSourceLogger, op.promConn, op.prestoQueryer, op.clock, cfg, metricsCollectors)
				importers[dataSourceName] = importer
			}

			if !op.cfg.DisablePromsum {
				worker, workerExists := workers[dataSourceName]
				if workerExists && worker.queryInterval != queryInterval {
					// queryInterval changed stop the existing worker from
					// collecting data, and create it with updated config
					worker.stop()
				} else if workerExists {
					// config hasn't changed skip the update
					continue
				}

				worker = newPromImportWorker(queryInterval)
				workers[dataSourceName] = worker

				// launch a go routine that periodically triggers a collection
				go worker.start(ctx, dataSourceLogger, semaphore, dataSourceName, importer)
			}
		}
	}
}

type prometheusImporterWorker struct {
	stopCh        chan struct{}
	doneCh        chan struct{}
	queryInterval time.Duration
}

func newPromImportWorker(queryInterval time.Duration) *prometheusImporterWorker {
	return &prometheusImporterWorker{
		queryInterval: queryInterval,
		stopCh:        make(chan struct{}),
		doneCh:        make(chan struct{}),
	}
}

// start begins periodic importing with the configured importer.
func (w *prometheusImporterWorker) start(ctx context.Context, logger logrus.FieldLogger, semaphore chan struct{}, dataSourceName string, importer *prestostore.PrometheusImporter) {
	ticker := time.NewTicker(w.queryInterval)
	defer close(w.doneCh)
	defer ticker.Stop()

	logger.Infof("Importing data for ReportDataSource %s every %s", dataSourceName, w.queryInterval)
	for {
		select {
		case <-w.stopCh:
			return
		case _, ok := <-ticker.C:
			if !ok {
				return
			}
			err := importPrometheusDataSourceData(ctx, logger, semaphore, dataSourceName, importer, func(ctx context.Context, importer *prestostore.PrometheusImporter) ([]prom.Range, error) {
				return importer.ImportFromLastTimestamp(ctx, false)
			})
			if err != nil {
				logger.WithError(err).Errorf("error collecting Prometheus DataSource data")
			}
		case <-ctx.Done():
			return
		}
	}
}

func (w *prometheusImporterWorker) stop() {
	close(w.stopCh)
	<-w.doneCh
}

type importFunc func(context.Context, *prestostore.PrometheusImporter) ([]prom.Range, error)

func importPrometheusDataSourceData(ctx context.Context, logger logrus.FieldLogger, semaphore chan struct{}, dataSourceName string, prometheusImporter *prestostore.PrometheusImporter, runImport importFunc) error {
	// blocks trying to increment the semaphore (sending on the
	// channel) or until the context is cancelled
	select {
	case semaphore <- struct{}{}:
	case <-ctx.Done():
		return ctx.Err()
	}
	dataSourceLogger := logger.WithField("reportDataSource", dataSourceName)
	// decrement the semaphore at the end
	defer func() {
		dataSourceLogger.Infof("finished import for Prometheus ReportDataSource %s", dataSourceName)
		prometheusReportDatasourceRunningImportsGauge.Dec()
		<-semaphore
	}()
	dataSourceLogger.Infof("starting import for Prometheus ReportDataSource %s", dataSourceName)
	prometheusReportDatasourceRunningImportsGauge.Inc()
	_, err := runImport(ctx, prometheusImporter)
	return err
}
