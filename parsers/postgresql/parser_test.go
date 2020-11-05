package postgresql

var createTableString string = "CREATE TABLE `alarm_state_tracking` (\n" +
	"`id` int(11) NOT NULL,\n" +
	"`cascade_id` varchar(32) NOT NULL,\n" +
	"`alarm_sent_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,\n" +
	"`clear_alarm` TINYINT(1) NULL,\n" +
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

var expected string = "package signatures " +
	"\n" +
	"type signatures struct {" +
	"ID string \\`gorm:\"column:id\"\\`+" +
	"TITLE string \\`gorm:\"column:title\"\\`+" +
	"SOURCE string \\`gorm:\"column:source\"\\`+" +
	"CATEGORYID \\`gorm:\"column:category_id\"\\`+" +
	"VENDOR \\`gorm:\"column:vendor\"\\`+" +
	"RAW string \\`gorm:\"column:raw\"\\`+" +
	"SEVERITY \\`gorm:\"column:severity\"\\`+" +
	"SCORE \\`gorm:\"column:score\"\\`+" +
	"ACTIVE int64 \\`gorm:\"column:active\"\\`+" +
	"SUPPORTED int64 \\`gorm:\"column:supported\"\\`+" +
	"CREATELSE int64 \\`gorm:\"column:create_lse\"\\`+" +
	"CREATEBACKHAULLSE int64 \\`gorm:\"column:create_backhaul_lse\"\\`+" +
	"SUPPRESSED int64 \\`gorm:\"column:suppressed\"\\`+" +
	"SETTOINFO int64 \\`gorm:\"column:set_to_info\"\\`+" +
	"ISVISIBLE int64 \\`gorm:\"column:is_visible\"\\`+" +
	"ISINFO int64 \\`gorm:\"column:is_info\"\\`+" +
	"CREATECASE int64 \\`gorm:\"column:create_case\"\\`+" +
	"INVITEFIRSTRESPONSE int64 \\`gorm:\"column:invite_first_response\"\\`+" +
	"INVITEBACKHAUL int64 \\`gorm:\"column:invite_backhaul\"\\`+" +
	"UPDATEDBY \\`gorm:\"column:updated_by\"\\`+" +
	"DELETEDBY \\`gorm:\"column:deleted_by\"\\`+" +
	"COUNT int64 \\`gorm:\"column:count\"\\`+" +
	"LASTSEEN \\`gorm:\"column:last_seen\"\\`+" +
	"CREATEDAT string \\`gorm:\"column:created_at\"\\`+" +
	"UPDATEDAT string \\`gorm:\"column:updated_at\"\\`+" +
	"DELETEDAT \\`gorm:\"column:deleted_at\"\\`+" +
	"SOURCE \\`gorm:\"column:source\"\\`+" +
	"SUPPORTED `gorm:\"column:supported\"\\`+" +
	"VENDOR \\`gorm:\"column:vendor\"\\`+" +
	"SEVERITY \\`gorm:\"column:severity\"\\`+" +
	"}" +
	"\n" +
	"func (s *signatures) TableName() string {" +
	"return \"signatures\"" +
	"}"

	// func TestParser(t *testing.T) {
	// 	before := time.Now().UnixNano() / (1 * int64(time.Microsecond))
	// 	Postgresql(createTableString, "alarm_state_tracking")
	// 	after := time.Now().UnixNano() / (1 * int64(time.Microsecond))
	// 	fmt.Println(after - before)
	// }
