<?php

$all = range(1, 10);

$include = [];
$exclude = [];

$allowed = '*,1-5,-4,7-10';
$allowed = '*,-4';
$allowed = '1-10, -4, -7';

$sections = explode(',', str_replace(' ', '', $allowed));
foreach ($sections as $section) {
    if (preg_match('/\d+-\d+/', $section)) {
        $parts = explode('-', $section);
        $include = array_merge($include, range($parts[0], $parts[1]));
    } else if ($section == '*') {
        $include = $all;
    }
    else if ($section < 0) {
        $exclude[] = abs($section);
    } else {
        $include[] = $section;
    }
}

echo 'include= ', implode(', ', $include), PHP_EOL;
echo 'exclude= ', implode(', ', $exclude), PHP_EOL;
echo '   diff= ', implode(', ', array_diff($include, $exclude)), PHP_EOL;
echo ' unique= ', implode(', ', array_unique(array_diff($include, $exclude))), PHP_EOL;
