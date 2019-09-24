<?php

function get_holidays($year)
{
    $holiday_formats = array(
        'New Years Day' => 'january 1 %d',
        'Family Day'      => 'third monday of february %d',

        'Good Friday' => function($year) {
            return date("F j, Y", easter_date($year) - 1*24*3600);
        },

        'Easter Monday' => function($year) {
            return date("F j, Y", easter_date($year) + 2*24*3600);
        },

        'Victoria Day'         => 'last monday may 25 %d',
        'Canada Day'           => 'july 1 %d',
        'August Civic Holiday' => 'first monday of august %d',
        'Labour Day'           => 'first monday of september %d',
        'Thanksgiving Day'     => 'second monday of october %d',
        'Christmas Day'        => 'december 25 %d',
        'Boxing Day'           => 'december 26 %d',
    );

    $holidays = array();
    foreach ($holiday_formats as $day => $timestring) {
        if (is_callable($timestring)) {
            $str = $timestring($year);
        } else {
            $str = sprintf($timestring, $year);
        }
        $d = strftime('%Y-%m-%d', strtotime($str));
        $holidays[$d] = $day;
    }

    return $holidays;
}

function is_holiday($date)
{
    list($y, $m, $d) = explode('-', $date);
    $holidays = get_holidays($y);
    
    return isset($holidays[$date]);
}

print_r(get_holidays(2019));
var_dump(is_holiday('2019-01-01'));
