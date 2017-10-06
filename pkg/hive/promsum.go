package hive

import (
	"errors"
)

var (
	// PromsumTableName is the Hive identi***REMOVED***er to use for usage data.
	PromsumTableName = "promsum"

	// PromsumSerde speci***REMOVED***es the Hadoop serialization/deserialization implementation to be used.
	PromsumSerde = "org.apache.hive.hcatalog.data.JsonSerDe"

	// PromsumSerdeProps de***REMOVED***ne the behavior of the SerDe used with promsum data.
	PromsumSerdeProps = map[string]string{
		"timestamp.formats": "yyyy-MM-dd'T'HH:mm:ss.SSSZ",
	}

	PromsumColumns = []string{
		"query string",
		"amount float",
		"`timestamp` timestamp",
		"`timePrecision` float",
		"labels map<string, string>",
	}
)

// CreatePromsumTable instantiates a new external Hive table for Prometheus observation data stored in S3.
func CreatePromsumTable(conn *Connection, tableName, bucket, pre***REMOVED***x string) error {
	if conn == nil {
		return errors.New("connection to Hive cannot be nil")
	} ***REMOVED*** if conn.session == nil {
		return errors.New("the Hive session has closed")
	}

	location, err := s3Location(bucket, pre***REMOVED***x)
	if err != nil {
		return err
	}
	query := createTable(tableName, location, PromsumSerde, PromsumSerdeProps, PromsumColumns, true, true)
	return conn.Query(query)
}
