<?php
 
/**
 * MySQL Apache Log Import
 *
 * This script imports apache combined log format logs into MySQL, so you can use standard SQL commands
 * to query your logs. For usage, please run the script with no arguments.
 * 
 * Based in part on http://snippets.dzone.com/posts/show/3721
 *
 * @author David Ordal (david -at- ordal.com)
 * @requires PHP 5.X
 * @requires MySQL 5.X
 *
 */
 
define('VERSION', '1.0');
define('TMP_FILE', '/tmp/mysql_httpd_log_import.tmp');
 
//
// STEP 1: GET CMD LINE ARGS
//
 
// command line arguments; check below for usage
$cmdArgs = getopt('d:t:h:u:p:cxf');
 
// check args
if (!(isset($cmdArgs['d']) && strlen($cmdArgs['d']) > 0 && isset($cmdArgs['t']) && strlen($cmdArgs['t']) > 0))
	displayUsage();
 
// connect to mysql database
$dbHost = isset($cmdArgs['h']) ? $cmdArgs['h'] : ini_get("mysqli.default_host");
$dbUser = isset($cmdArgs['u']) ? $cmdArgs['u'] : ini_get("mysqli.default_user");
$dbPass = isset($cmdArgs['p']) ? $cmdArgs['p'] : ini_get("mysqli.default_pw");
$dbTable = $cmdArgs['t'];
$dbName = $cmdArgs['d'];
$mysqli = dbConnect($dbHost, $dbUser, $dbPass, $dbName);
 
// check to see if we need to drop and/or create the table
$quotedDbTable = dbQuoteIdentifier($mysqli, $dbTable, true);
if (isset($cmdArgs['x'])) {
	$cmdArgs['c'] = true; 				// -x implies -c
	$queryResult = dbQuery($mysqli, "DROP TABLE IF EXISTS {$quotedDbTable}");
}
$quotedDbTable = dbQuoteValue($mysqli, $dbTable, true);
$queryResult = dbQuery($mysqli, "SHOW TABLES LIKE {$quotedDbTable}");
if ($queryResult->num_rows != 1) {
	if(isset($cmdArgs['c'])) {
		dbCreateTable($mysqli, $dbTable);
	} else {
		die ("Database table '{$dbTable}' does not exist. Please rerun the script with the -c option to create it.\n");
	}
}
 
//
// STEP 2: COPY DATA INTO TAB-DELIMITED FILE
//
 
// open the temp CSV file for copying data
$tmpFile = fopen(TMP_FILE, 'w');
 
// read each line of STDIN and process a log
$checkDb = isset($cmdArgs['f']) ? false : true;
while (!feof(STDIN)) {
	$line = fgets(STDIN);
 
	if (empty($line))
		continue;
 
	$results = processLine($line);
 
	// check the first and last entries; print an error if something went wrong
	if (empty($results['fullString']) || empty($results['userAgent']) || !is_numeric($results['status'])) {
		echo "Error! Could not interpret line: ".$line;
		continue;
	}
 
	// convert entries to database format. NOTE: doing the timestamp conversion this way converts
	// each entry to the local timezone on the local box. Stupid MySQL doesn't support storing a timezone
	// with a timestamp, so we covert everything from the web server's timezone to the local box's timezone,
	// and store that.
	$results['date'] = str_replace('/', ' ', $results['date']);
	$logTimestamp = strtotime("{$results['date']} {$results['time']} {$results['timezone']}");
	$sqlTimestamp = date('Y-m-d H:i:s', $logTimestamp);	
	$results['bytes'] = is_numeric($results['bytes']) ? $results['bytes'] : '0';
 
	// run a $mysqli->escape_string() on all the strings to put into the database. We don't want to use
	// dbQuoteValue(), because that also adds quotes, which the LOAD DATA command interepts litterally
	$remote_host = $mysqli->escape_string($results['remoteHost']);
	$ident_user = $mysqli->escape_string($results['identUser']);
	$auth_user = $mysqli->escape_string($results['authUser']);
	$timestamp = $mysqli->escape_string($sqlTimestamp);
	$method = $mysqli->escape_string($results['method']);
	$url = $mysqli->escape_string($results['url']);
	$protocol = $mysqli->escape_string($results['protocol']);
	$status = $mysqli->escape_string($results['status']);
	$bytes = $mysqli->escape_string($results['bytes']);
	$referrer = $mysqli->escape_string($results['referrer']);
	$user_agent = $mysqli->escape_string($results['userAgent']);
 
	// figure out if we should check the database for the first entry. This helps prevent 
	// duplicates. Use -f to override
	if ($checkDb) {
		$quotedDbTable = dbQuoteIdentifier($mysqli, $dbTable);
		$sql = <<<QQ
			SELECT TRUE FROM {$quotedDbTable}
			WHERE 
		    remote_host = '{$remote_host}' AND
		    ident_user = '{$ident_user}' AND
		    auth_user = '{$auth_user}' AND
		    time_stamp = '{$timestamp}' AND
		    request_method = '{$method}' AND
		    request_uri = '{$url}' AND
		    request_protocol = '{$protocol}' AND
		    status = '{$status}' AND
		    bytes = '{$bytes}' AND
		    referer = '{$referrer}' AND
		    user_agent = '{$user_agent}'
QQ;
 
		$queryResult = dbQuery($mysqli, $sql);
		if ($queryResult->num_rows > 0)
			die("Skipping file; the first entry of this log file already appears to be stored in the database. Use -f to override.\n");
 
		// check only the first row
		$checkDb = false;
	}
 
	$logString = "{$remote_host}\t{$ident_user}\t{$auth_user}\t{$timestamp}\t{$method}\t{$url}\t{$protocol}\t{$status}\t{$bytes}\t{$referrer}\t{$user_agent}\n";
	fwrite($tmpFile, $logString);
 
}
fclose($tmpFile);
 
//
// STEP 3: COPY TAB-DELIMITED FILE INTO DB
//
 
// load data into database
$quotedFile = dbQuoteValue($mysqli, TMP_FILE);
$quotedDbTable = dbQuoteIdentifier($mysqli, $dbTable);
$sql = <<<QQ
LOAD DATA LOCAL INFILE {$quotedFile} INTO TABLE {$quotedDbTable}
  FIELDS TERMINATED BY '\t' 
  LINES TERMINATED BY '\n';
QQ;
 
dbQuery($mysqli, $sql);
 
// delete the tmp file after importing
unlink(TMP_FILE);
 
$mysqli->close();
 
/*******************************************************************************
 *************************      INTERNAL FUNCTIONS     *************************
 *******************************************************************************/
 
/**
 * processLine(): processes a line of a log file, returning an associative array
 * with the component parts
 *
 * @param string $line the line of the log
 * @return array associative array of values from log file
 *
 */
function processLine($line) {
	$matches = array();
 
	// process the string. This regular expression was adapted from http://oreilly.com/catalog/perlwsmng/chapter/ch08.html
	preg_match('/^(\S+) (\S+) (\S+) \[([^:]+):(\d+:\d+:\d+) ([^\]]+)\] "(\S+) (.+?) (\S+)" (\S+) (\S+) "([^"]+)" "([^"]+)"$/', $line, $matches);
 
	if (isset($matches[0])) {
		return array('fullString' => $matches[0],
		         'remoteHost' => $matches[1],
		         'identUser' => $matches[2],
		         'authUser' => $matches[3],
		         'date' => $matches[4],
		         'time' => $matches[5],
		         'timezone' => $matches[6],
		         'method' => $matches[7],
		         'url' => $matches[8],
		         'protocol' => $matches[9],
		         'status' => $matches[10],
		         'bytes' => $matches[11],
		         'referrer' => $matches[12],
		         'userAgent' => $matches[13]
		);
	} else {
		return array();
	}
}
 
/**
 * displayUsage(): display a usage message and exit
 *
 */
function displayUsage() {
	$version = VERSION;
	echo <<<QQ
{$_SERVER['SCRIPT_NAME']} v{$version}: Imports an Apache combined log into a MySQL database.
Usage: mysql_httpd_log_import -d <database name> -t <table name> [options] < log_file_name
 -d <database name> The database to use; required
 -t <table name>    The name of the table in which to insert data; required
 -h <host name>     The host to connect to; default is localhost
 -u <username>      The user to connect as
 -p <password>      The user's password
 -c                 Create table if it doesn't exist
 -x                 Drop the existing table if it exists. Implies -c
 -f                 Force load; skip the duplicate check. By default, the software exits if
                    the first entry in a given file already exists in the database
 
QQ;
 
exit;
}
 
/*******************************************************************************
 *************************      DATABASE FUNCTIONS     *************************
 *******************************************************************************/
 
/**
 * dbConnect(): connect to the database
 *
 * @param string $dbHost
 * @param string $dbUser
 * @param string $dbPass
 * @param string $dbName
 * @return mysqli object
 *
 */
function dbConnect($dbHost, $dbUser, $dbPass, $dbName) {
	$mysqli = new mysqli($dbHost, $dbUser, $dbPass, $dbName);
	if ($mysqli->connect_error)
	    die('DB Connect Error (' . $mysqli->connect_errno . ') '. $mysqli->connect_error);
	return $mysqli;
}
 
/**
 * dbQuery(): queries the DB; exits on error
 *
 * @param mysqli $mysqli
 * @param string $query
 * @return result object
 *
 */
function dbQuery($mysqli, $query) {
	if (!$result = $mysqli->query($query))
		die('DB Error (' . $mysqli->errno . ') '. $mysqli->error);
 
	return $result;
}
 
/**
 * dbQuoteValue(): quotes a value and makes it safe for inserting into
 * the database. Note this function DOES include surrounding quotes, but
 * only when necessary (e.g. O'Reilly will come back as 'O\'Reilly', while
 * 4.5 will come back as 4.5)
 *
 * @param   mixed $value the value to be quoted
 * @param   bool  $alwaysQuote set to force a value to be surrounded by quotes,
 *                             no matter what type it is
 * @return  mixed the $value, ready for insertion into the SQL database
 */
function dbQuoteValue($mysqli, $value, $alwaysQuote=false) {
 
	// check for magic quotes. these should just be off, so we throw an exception
	if (get_magic_quotes_gpc())
		die("magic_quotes_gpc is enabled. Please disable it.\n");
 
	if(!$alwaysQuote) {
		if (is_null($value)) {
			return 'NULL';
		} elseif (is_bool($value)) {
			return $value ? 'TRUE' : 'FALSE';
		} elseif (is_numeric($value)) {
			return $value;
		}
	}
	return "'" . $mysqli->escape_string($value) . "'";
}
 
/**
  * dbQuoteIdentifier(): Quotes a string so it can be safely used as a table or
  * column name
  *
  * @param mysqli $mysqli
  * @param string $value  the identifier name to be quoted
  * @return string  the quoted identifier
  *
  */
function dbQuoteIdentifier($mysqli, $value) {
	return '`' . $mysqli->escape_string($value) . "`";
}
 
/**
 * dbCreateTable(): create the database table used to store log entries
 *
 * @param result $mysqli MySQLi result resource
 * @param string $tableName
 *
 */
function dbCreateTable($mysqli, $tableName) {
	$quotedTableName = dbQuoteIdentifier($mysqli, $tableName, true);
 
	$sql = <<<QQ
		CREATE TABLE {$quotedTableName} (
		    `remote_host` VARCHAR(50),
		    `ident_user` VARCHAR(50),
		    `auth_user` VARCHAR(50),
		    `time_stamp` TIMESTAMP,
		    `request_method` VARCHAR(10),
		    `request_uri` VARCHAR(1024),
		    `request_protocol` VARCHAR(10),
		    `status` INT,
		    `bytes` INT UNSIGNED,
		    `referer` VARCHAR(2048),
		    `user_agent` VARCHAR(2048),
		    `id` BIGINT auto_increment,
		    PRIMARY KEY (`id`)
		    );
QQ;
 
		dbQuery($mysqli, $sql);
 
}
 
?>