package chargeback

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	api "github.com/coreos-inc/kube-chargeback/pkg/apis/chargeback/v1alpha1"
	"github.com/coreos-inc/kube-chargeback/pkg/presto"
)

var ErrReportIsRunning = errors.New("the report is still running")

type server struct {
	chargeback *Chargeback
	logger     log.FieldLogger
	httpServer *http.Server
}

func newServer(c *Chargeback, logger log.FieldLogger) *server {
	logger = logger.WithField("component", "api")
	mux := http.NewServeMux()
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	srv := &server{
		chargeback: c,
		logger:     logger,
		httpServer: httpServer,
	}
	mux.HandleFunc("/api/v1/reports/get", srv.getReportHandler)
	mux.HandleFunc("/api/v1/reports/run", srv.runReportHandler)
	mux.HandleFunc("/api/v1/collect/prometheus", srv.collectPromsumDataHandler)

	// TODO(cgag): use this format and gorilla/mux
	// /api/v2/reports/$REPORT_NAME/default
	// /api/v2/reports/$REPORT_NAME/table
	mux.HandleFunc("/api/v2/reports/default", srv.getReportDefaultHandler)
	mux.HandleFunc("/api/v2/reports/table", srv.getReportTableHandler)

	mux.HandleFunc("/ready", srv.readinessHandler)
	return srv
}

func (srv *server) start() {
	srv.logger.Infof("HTTP server started")
	srv.logger.WithError(srv.httpServer.ListenAndServe()).Info("HTTP server exited")
}

func (srv *server) stop() error {
	return srv.httpServer.Shutdown(context.TODO())
}

func (srv *server) newLogger(r *http.Request) log.FieldLogger {
	return srv.logger.WithFields(log.Fields{
		"method": r.Method,
		"url":    r.URL.String(),
	}).WithFields(newLogIdentifier())
}

func (srv *server) logRequest(logger log.FieldLogger, r *http.Request) {
	logger.Infof("%s %s", r.Method, r.URL.String())
}

type errorResponse struct {
	Error string `json:"error"`
}

func (srv *server) writeErrorResponse(
	logger log.FieldLogger,
	w http.ResponseWriter,
	r *http.Request,
	status int,
	message string,
	args ...interface{},
) {
	msg := fmt.Sprintf(message, args...)
	srv.writeResponseWithBody(logger, w, status, errorResponse{Error: msg})
}

// writeResponseWithBody attempts to marshal an arbitrary thing to JSON then write
// it to the http.ResponseWriter
func (srv *server) writeResponseWithBody(logger log.FieldLogger, w http.ResponseWriter, code int, resp interface{}) {
	enc, err := json.Marshal(resp)
	if err != nil {
		logger.WithError(err).Error("failed JSON-encoding HTTP response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err = w.Write(enc); err != nil {
		logger.WithError(err).Error("failed writing HTTP response")
	}
}

func (srv *server) getReportHandler(w http.ResponseWriter, r *http.Request) {
	logger := srv.newLogger(r)
	srv.logRequest(logger, r)
	if r.Method != "GET" {
		srv.writeErrorResponse(logger, w, r, http.StatusNotFound, "Not found")
		return
	}
	err := r.ParseForm()
	if err != nil {
		srv.writeErrorResponse(logger, w, r, http.StatusBadRequest, "couldn't parse URL query params: %v", err)
		return
	}
	vals := r.Form
	err = checkForFields([]string{"name", "format"}, vals)
	if err != nil {
		srv.writeErrorResponse(logger, w, r, http.StatusBadRequest, "%v", err)
		return
	}
	switch vals["format"][0] {
	case "json", "csv":
		break
	default:
		srv.writeErrorResponse(logger, w, r, http.StatusBadRequest, "format must be one of: csv, json")
		return
	}
	srv.getReport(logger, vals["name"][0], vals["format"][0], w, r)
}

func (srv *server) getReportDefaultHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(cgag): copied and pasted from getReportHandler.  Factor out?
	// not worht it until we switch to gorilla mux
	logger := srv.newLogger(r)

	logger.Debugf("curtis was here, in getReportDefaultHandler")

	srv.logRequest(logger, r)
	if r.Method != "GET" {
		srv.writeErrorResponse(logger, w, r, http.StatusNotFound, "Not found")
		return
	}
	err := r.ParseForm()
	if err != nil {
		srv.writeErrorResponse(logger, w, r, http.StatusBadRequest, "couldn't parse URL query params: %v", err)
		return
	}
	vals := r.Form
	err = checkForFields([]string{"name", "format"}, vals)
	if err != nil {
		srv.writeErrorResponse(logger, w, r, http.StatusBadRequest, "%v", err)
		return
	}
	switch vals["format"][0] {
	case "json", "csv":
		break
	default:
		srv.writeErrorResponse(logger, w, r, http.StatusBadRequest, "format must be one of: csv, json")
		return
	}
	srv.getReportDefault(logger, vals["name"][0], vals["format"][0], w, r)
}

func (srv *server) getReportTableHandler(w http.ResponseWriter, r *http.Request) {
	// TODO(cgag): copied and pasted from getReportHandler.  Factor out?
	// not worht it until we switch to gorilla mux
	logger := srv.newLogger(r)
	srv.logRequest(logger, r)
	if r.Method != "GET" {
		srv.writeErrorResponse(logger, w, r, http.StatusNotFound, "Not found")
		return
	}
	err := r.ParseForm()
	if err != nil {
		srv.writeErrorResponse(logger, w, r, http.StatusBadRequest, "couldn't parse URL query params: %v", err)
		return
	}
	vals := r.Form
	err = checkForFields([]string{"name", "format"}, vals)
	if err != nil {
		srv.writeErrorResponse(logger, w, r, http.StatusBadRequest, "%v", err)
		return
	}
	switch vals["format"][0] {
	case "json", "csv":
		break
	default:
		srv.writeErrorResponse(logger, w, r, http.StatusBadRequest, "format must be one of: csv, json")
		return
	}
	srv.getReportTable(logger, vals["name"][0], vals["format"][0], w, r)
}

func (srv *server) runReportHandler(w http.ResponseWriter, r *http.Request) {
	logger := srv.newLogger(r)
	srv.logRequest(logger, r)
	if r.Method != "GET" {
		srv.writeErrorResponse(logger, w, r, http.StatusNotFound, "Not found")
		return
	}
	err := r.ParseForm()
	if err != nil {
		srv.writeErrorResponse(logger, w, r, http.StatusBadRequest, "couldn't parse URL query params: %v", err)
		return
	}
	vals := r.Form
	err = checkForFields([]string{"query", "start", "end"}, vals)
	if err != nil {
		srv.writeErrorResponse(logger, w, r, http.StatusBadRequest, "%v", err)
		return
	}
	srv.runReport(logger, vals["query"][0], vals["start"][0], vals["end"][0], w)
}

func checkForFields(fields []string, vals url.Values) error {
	var missingFields []string
	for _, f := range fields {
		if len(vals[f]) == 0 || vals[f][0] == "" {
			missingFields = append(missingFields, f)
		}
	}
	if len(missingFields) != 0 {
		return fmt.Errorf("the following fields are missing or empty: %s", strings.Join(missingFields, ","))
	}
	return nil
}

func (srv *server) getReport(logger log.FieldLogger, name, format string, w http.ResponseWriter, r *http.Request) {
	// Get the current report to make sure it's in a finished state
	report, err := srv.chargeback.informers.reportLister.Reports(srv.chargeback.namespace).Get(name)
	if err != nil {
		logger.WithError(err).Errorf("error getting report: %v", err)
		srv.writeErrorResponse(logger, w, r, http.StatusInternalServerError, "error getting report: %v", err)
		return
	}
	switch report.Status.Phase {
	case api.ReportPhaseError:
		err := fmt.Errorf(report.Status.Output)
		logger.WithError(err).Errorf("the report encountered an error")
		srv.writeErrorResponse(logger, w, r, http.StatusInternalServerError, "the report encountered an error: %v", err)
		return
	case api.ReportPhaseFinished:
		// continue with returning the report if the report is finished
	case api.ReportPhaseWaiting, api.ReportPhaseStarted:
		fallthrough
	default:
		logger.Errorf(ErrReportIsRunning.Error())
		srv.writeErrorResponse(logger, w, r, http.StatusAccepted, ErrReportIsRunning.Error())
		return
	}

	reportTable := reportTableName(name)
	getReportQuery := fmt.Sprintf("SELECT * FROM %s", reportTable)
	results, err := presto.ExecuteSelect(srv.chargeback.prestoConn, getReportQuery)
	if err != nil {
		logger.WithError(err).Errorf("failed to perform presto query")
		srv.writeErrorResponse(logger, w, r, http.StatusInternalServerError, "failed to perform presto query (see chargeback logs for more details): %v", err)
		return
	}

	switch format {
	case "json":
		srv.writeResponseWithBody(logger, w, http.StatusOK, results)
		return
	case "csv":
		// Get generation query to get the list of columns
		genQuery, err := srv.chargeback.informers.reportGenerationQueryLister.ReportGenerationQueries(report.Namespace).Get(report.Spec.GenerationQueryName)
		if err != nil {
			logger.WithError(err).Errorf("error getting report generation query: %v", err)
			srv.writeErrorResponse(logger, w, r, http.StatusInternalServerError, "error getting report generation query: %v", err)
			return
		}

		if len(results) > 0 && len(genQuery.Spec.Columns) != len(results[0]) {
			logger.WithError(err).Errorf("report results schema doesn't match expected schema")
			srv.writeErrorResponse(logger, w, r, http.StatusInternalServerError, "report results schema doesn't match expected schema")
			return
		}

		buf := &bytes.Buffer{}
		csvWriter := csv.NewWriter(buf)

		// Write headers
		var keys []string
		if len(results) >= 1 {
			for _, column := range genQuery.Spec.Columns {
				keys = append(keys, column.Name)
			}
			err := csvWriter.Write(keys)
			if err != nil {
				logger.WithError(err).Errorf("failed to write headers")
				return
			}
		}

		// Write the rest
		for _, row := range results {
			vals := make([]string, len(keys))
			for i, key := range keys {
				val, ok := row[key]
				if !ok {
					logger.WithError(err).Errorf("report results schema doesn't match expected schema, unexpected key: %q", key)
					srv.writeErrorResponse(logger, w, r, http.StatusInternalServerError, "report results schema doesn't match expected schema, unexpected key: %q", key)
					return
				}
				switch v := val.(type) {
				case string:
					vals[i] = v
				case []byte:
					vals[i] = string(v)
				case uint, uint8, uint16, uint32, uint64,
					int, int8, int16, int32, int64,
					float32, float64,
					complex64, complex128,
					bool:
					vals[i] = fmt.Sprintf("%v", v)
				case time.Time:
					vals[i] = v.String()
				case nil:
					vals[i] = ""
				default:
					logger.Errorf("error marshalling csv: unknown type %#T for value %v", val, val)
					srv.writeErrorResponse(logger, w, r, http.StatusInternalServerError, "error marshalling csv (see chargeback logs for more details)", err)
					return
				}
			}
			err := csvWriter.Write(vals)
			if err != nil {
				logger.Errorf("failed to write csv row: %v", err)
				return
			}
		}

		csvWriter.Flush()
		w.Header().Set("Content-Type", "text/csv")
		w.Write(buf.Bytes())
	}
}

func (srv *server) getReportTable(logger log.FieldLogger, name, format string, w http.ResponseWriter, r *http.Request) {
	// TODO(cgag): largely copied from getReport, factor

	// get the current report to make sure it's finished
	report, err := srv.chargeback.informers.reportLister.Reports(srv.chargeback.namespace).Get(name)
	if err != nil {
		logger.WithError(err).Errorf("error getting report: %v", err)
		srv.writeErrorResponse(logger, w, r, http.StatusInternalServerError, "error getting report: %v", err)
		return
	}

	// TODO(cgag): this is just testing for the filtering in getReportTable/getReportGraph, don't need these values here(?)
	genQuery, err := srv.
		chargeback.
		informers.
		reportGenerationQueryLister.
		ReportGenerationQueries(srv.chargeback.namespace).
		Get(report.Spec.GenerationQueryName)
	if err != nil {
		logger.WithError(err).Errorf("error getting reportGenerationQuery")
		srv.writeErrorResponse(logger, w, r, http.StatusInternalServerError, "error getting ReportGenerationQuery: %v", err)
		return
	}

	var hiddenColumns []string
	for _, col := range genQuery.Spec.Columns {
		if col.TableHidden {
			logger.Debugf("FOUND TABLE HIDDEN COLUMN: %s", col.Name)
			hiddenColumns = append(hiddenColumns, col.Name)
		}
	}
	logger.Debugf("FOUND TABLE HIDDEN COLUMNS: %s", hiddenColumns)

	switch report.Status.Phase {
	case api.ReportPhaseError:
		err := fmt.Errorf(report.Status.Output)
		logger.WithError(err).Errorf("the report encountered an error")
		srv.writeErrorResponse(logger, w, r, http.StatusInternalServerError, "the report encountered an error %v", err)
		return
	case api.ReportPhaseFinished:
	// continue if it's finished
	case api.ReportPhaseWaiting, api.ReportPhaseStarted:
		fallthrough
	default:
		logger.Errorf(ErrReportIsRunning.Error())
		srv.writeErrorResponse(logger, w, r, http.StatusAccepted, ErrReportIsRunning.Error())
		return
	}

	getReportQuery := fmt.Sprintf("SELECT * FROM %s", reportTableName(name))
	results, err := presto.ExecuteSelect(srv.chargeback.prestoConn, getReportQuery)
	if err != nil {
		logger.WithError(err).Errorf("failed to perform presto query")
		srv.writeErrorResponse(logger, w, r, http.StatusInternalServerError, "failed to perform presto query (see chargeback logs for more details): %v", err)
		return
	}

	switch format {
	case "csv":
		// TODO(cgag): what should really happen here
		srv.writeErrorResponse(logger, w, r, http.StatusBadRequest, "CSV not supported", fmt.Errorf("CSV not supported"))
		return
	// TODO(cgag): could we even get an empty format?
	case "json", "":
		formatted := format(results, "table")
		srv.writeResponseWithBody(logger, w, http.StatusOK, formatted)
		return
	}
}

func (srv *server) getReportDefault(logger log.FieldLogger, name, format string, w http.ResponseWriter, r *http.Request) {
	// TODO(cgag): largely copied from getReport, factor

	// get the current report to make sure it's finished
	/// TODO(cgag): delete
	reports, _ := srv.chargeback.informers.reportLister.Reports(srv.chargeback.namespace).List(nil)
	for _, report := range reports {
		logger.Debugf("report: %s", report.Name)
	}

	report, err := srv.chargeback.informers.reportLister.Reports(srv.chargeback.namespace).Get(name)
	if err != nil {
		logger.WithError(err).Errorf("error getting report: %v", err)
		srv.writeErrorResponse(logger, w, r, http.StatusInternalServerError, "error getting report: %v", err)
		return
	}

	// TODO(cgag): this is just testing for the filtering in getReportTable/getReportGraph, don't need these values here(?)
	genQuery, err := srv.
		chargeback.
		informers.
		reportGenerationQueryLister.
		ReportGenerationQueries(srv.chargeback.namespace).
		Get(report.Spec.GenerationQueryName)
	if err != nil {
		logger.WithError(err).Errorf("error getting reportGenerationQuery")
		srv.writeErrorResponse(logger, w, r, http.StatusInternalServerError, "error getting ReportGenerationQuery: %v", err)
		return
	}
	// TODO(cgag): just leave hiddenColumns blank in default handler
	var hiddenColumns []string
	for _, col := range genQuery.Spec.Columns {
		if col.TableHidden {
			logger.Debugf("FOUND TABLE HIDDEN COLUMN: %s", col.Name)
			hiddenColumns = append(hiddenColumns, col.Name)
		}
	}
	logger.Debugf("FOUND TABLE HIDDEN COLUMNS: %s", hiddenColumns)
	// TODO(cgag): then go delete those keys from the map from before writing response

	// report.Spec.GenerationQueryName
	switch report.Status.Phase {
	case api.ReportPhaseError:
		err := fmt.Errorf(report.Status.Output)
		logger.WithError(err).Errorf("the report encountered an error")
		srv.writeErrorResponse(logger, w, r, http.StatusInternalServerError, "the report encountered an error %v", err)
		return
	case api.ReportPhaseFinished:
	// continue if it's finished
	case api.ReportPhaseWaiting, api.ReportPhaseStarted:
		fallthrough
	default:
		logger.Errorf(ErrReportIsRunning.Error())
		srv.writeErrorResponse(logger, w, r, http.StatusAccepted, ErrReportIsRunning.Error())
		return
	}

	getReportQuery := fmt.Sprintf("SELECT * FROM %s", reportTableName(name))
	results, err := presto.ExecuteSelect(srv.chargeback.prestoConn, getReportQuery)
	if err != nil {
		logger.WithError(err).Errorf("failed to perform presto query")
		srv.writeErrorResponse(logger, w, r, http.StatusInternalServerError, "failed to perform presto query (see chargeback logs for more details): %v", err)
		return
	}

	switch format {
	case "csv":
		// TODO(cgag): what should really happen here
		srv.writeErrorResponse(logger, w, r, http.StatusBadRequest, "CSV not supported", fmt.Errorf("CSV not supported"))
		return
	// TODO(cgag): could we even get an empty format?
	case "json", "":
		formatted := format(results, hiddenColumns)
		srv.writeResponseWithBody(logger, w, http.StatusOK, formatted)
		return
	}
}

// TODO(cgag): Result is a garbage name
// TODO(cgag): json encoding tags
type Result struct {
	Name  string
	Value interface{}
	Unit  string
}

// data is an array of maps of column names to values
func format(data []map[string]interface{}, format string) map[string][]map[string][]Result {
	//	{
	//    "results": [
	//      {
	//        "values": [
	//        {
	//          "name": "pod",
	//          "value": "chargeback-8f4cfcf8b-n2g7j",
	//          "unit": "k8s_pod_name"
	//        },
	//				...
	//			}
	//	}
	results := map[string][]map[string][]Result{
		"results": []map[string][]Result{},
	}

	for _, m := range data {
		var tmp []Result
		for colName, val := range m {
			tableHidden := colName == "table_hidden" && val == true
			graphHidden := colName == "graph_hidden" && val == true

			if !(format == "table" && tableHidden) && !(format == "graph" && graphHidden) {
				// TODO(cgag): unit ??
				r := Result{
					Name:  colName,
					Value: val,
				}
				tmp = append(tmp, r)
			}
		}
		results["results"] = append(results["results"], map[string][]Result{
			"values": tmp,
		})
	}

	return results
}

func (srv *server) runReport(logger log.FieldLogger, query, start, end string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("method not yet implemented"))
}

type CollectPromsumDataRequest struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

func (srv *server) collectPromsumDataHandler(w http.ResponseWriter, r *http.Request) {
	logger := srv.newLogger(r)
	srv.logRequest(logger, r)

	decoder := json.NewDecoder(r.Body)
	var req CollectPromsumDataRequest
	err := decoder.Decode(&req)
	if err != nil {
		srv.writeErrorResponse(logger, w, r, http.StatusInternalServerError, "unable to decode response as JSON: %v", err)
		return
	}

	timeBoundsGetter := promsumDataSourceTimeBoundsGetter(func(dataSource *api.ReportDataSource) (startTime, endTime time.Time, err error) {
		return req.StartTime, req.EndTime, nil
	})

	srv.chargeback.collectPromsumData(context.Background(), logger, timeBoundsGetter)

	srv.writeResponseWithBody(logger, w, http.StatusOK, struct{}{})
}
