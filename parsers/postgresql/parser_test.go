package postgresql

import (
	"testing"
)

var createTableMap map[string]string = map[string]string{
	"1":  "id:false:integer",
	"2":  "cascade_id:false:text",
	"3":  "alarm_sent_at:false:timestamp with time zone",
	"4":  "clear_alarm:true:boolean",
	"5":  "alarm_cleared_at:false:timestamp with time zone",
	"6":  "api_recovery_event_id:true:integer",
	"7":  "api_status_code_problem:true:integer",
	"8":  "api_status_code_recovery:true:integer",
	"9":  "api_response_problem:true:text",
	"10": "api_response_recovery:true:text",
	"11": "created_at:false:timestamp with time zone",
	"12": "updated_at:false:timestamp with time zone",
}

var expected string = "package alarm_state_tracking " +
	"\n\n" +
	"type alarm_state_tracking struct {\n" +
	"Id int64 `gorm:\"column:id\"`\n" +
	"CascadeId string `gorm:\"column:cascade_id\"`\n" +
	"AlarmSentAt time.Time `gorm:\"column:alarm_sent_at\"`\n" +
	"ClearAlarm sql.NullBool `gorm:\"column:clear_alarm\"`\n" +
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
	p := PostgresDDL{
		Table: "alarm_state_tracking",
	}
	finishedStruct := p.Parse(createTableMap)

	if finishedStruct != expected {
		t.Logf("expected and actual are not equal. expected: \n%s\n \nactual: \n%s", expected, finishedStruct)
		t.Fail()
	}
}
