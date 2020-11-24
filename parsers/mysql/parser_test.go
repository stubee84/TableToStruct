package mysql

import (
	"testing"
)

var createTableString string = "CREATE TABLE `alarm_state_tracking` (\n" +
	"`id` int(11) NOT NULL,\n" +
	"`cascade_id` varchar(32) NOT NULL,\n" +
	"`alarm_sent_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,\n" +
	"`clear_alarm` TINYINT(1) DEFAULT NULL,\n" +
	"`alarm_cleared_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,\n" +
	"`api_recovery_event_id` int(11) DEFAULT NULL,\n" +
	"`api_status_code_problem` int(11) DEFAULT NULL,\n" +
	"`api_status_code_recovery` int(11) DEFAULT NULL,\n" +
	"`api_response_problem` varchar(500) DEFAULT NULL,\n" +
	"`api_response_recovery` varchar(500) DEFAULT NULL,\n" +
	"`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,\n" +
	"`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,\n" +
	"PRIMARY KEY (`id`,`cascade_id`)\n" +
	") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;"

var expected string = "package alarm_state_tracking " +
	"\n\n" +
	"type alarm_state_tracking struct {\n" +
	"Id int64 `gorm:\"column:id\"`\n" +
	"CascadeId string `gorm:\"column:cascade_id\"`\n" +
	"AlarmSentAt time.Time `gorm:\"column:alarm_sent_at\"`\n" +
	"ClearAlarm sql.NullInt32 `gorm:\"column:clear_alarm\"`\n" +
	"AlarmClearedAt time.Time `gorm:\"column:alarm_cleared_at\"`\n" +
	"ApiRecoveryEventId sql.NullInt64 `gorm:\"column:api_recovery_event_id\"`\n" +
	"ApiStatusCodeProblem sql.NullInt64 `gorm:\"column:api_status_code_problem\"`\n" +
	"ApiStatusCodeRecovery sql.NullInt64 `gorm:\"column:api_status_code_recovery\"`\n" +
	"ApiResponseProblem sql.NullString `gorm:\"column:api_response_problem\"`\n" +
	"ApiResponseRecovery sql.NullString `gorm:\"column:api_response_recovery\"`\n" +
	"CreatedAt time.Time `gorm:\"column:created_at\"`\n" +
	"UpdatedAt time.Time `gorm:\"column:updated_at\"`\n" +
	"}" +
	"\n\n" +
	"func (alarm_state_tracking *alarm_state_tracking) TableName() string {\n" +
	"return \"alarm_state_tracking\"\n" +
	"}"

func TestParser(t *testing.T) {
	m := MysqlDDL{}
	finishedStruct := m.Parse(map[string]string{
		"alarm_state_tracking": createTableString})

	if finishedStruct != expected {
		t.Logf("expected and actual are not equal. expected: \n%s\n \nactual: \n%s", expected, finishedStruct)
		t.Fail()
	}
}
