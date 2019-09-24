<?php

$from='America/Toronto';
$to='UTC';

$now = date('Y-m-d H:i:s');
echo $now, " NOW\n";

$utc = gmdate('Y-m-d H:i:s');
echo $utc, " UTC\n";

//echo convertTime($now, $from, $to), " UTC\n";
//echo convertTime($now, $from, 'EST'), " EST\n";

echo changeTimezone($now, 'America/Toronto', 'UTC'), " UTC\n";
echo changeTimezone($now, 'America/Toronto', 'EST'), " EST\n";

echo changeTimezone($utc, 'UTC', 'EST'), " EST\n";
echo changeTimezone($utc, 'UTC', 'America/Toronto'), " TOR\n";

function convertTime($time, $from, $to)
{
    // $from='UTC';
    // $to='America/New_York';
    // $date=date($time);

	$default = date_default_timezone_get();

    date_default_timezone_set($from);
    $newDatetime = strtotime($time);

    date_default_timezone_set($to);
    $newDatetime = date($format, $newDatetime);
    $format = 'Y-m-d H:i:s';

    date_default_timezone_set($default);
    return $newDatetime;
}

function changeTimezone($time, $from, $to)
{
    $date = new DateTime($time, new DateTimeZone($from));
    $date->setTimezone(new DateTimeZone($to));
    return $date->format('Y-m-d H:i:s');
}
