<?php

$create = <<<EOS
CREATE TABLE IF NOT EXISTS `docs` ( 
    `id` int(6) unsigned NOT NULL,
    `rev` int(3) unsigned NOT NULL,
    `content` varchar(200) NOT NULL,
    PRIMARY KEY (`id`,`rev`) 
) DEFAULT CHARSET=utf8; 

EOS;

$insert = "INSERT INTO `docs` (`id`, `rev`, `content`) VALUES\n";

//echo $create;

$i = 1;
$new = 1;

while ($i <=1000000) {
    if ($new) {
        echo $insert;
        $new = 0;
    }

    echo "\t('$i', '1', 'a')"; 

    if ($i%1000 == 0) {
        echo ";\n";
        $new = 1;
    } else {
        echo ",\n";
    }

    $i += 1;
}
