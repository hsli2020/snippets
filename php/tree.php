<?php

function vd()
{
    ob_start();
    $var = func_get_args(); 
    call_user_func_array('var_dump', $var);
    $output = ob_get_contents();
    ob_end_clean();

    $patterns = [
        "/\"\]/"      => "\"",
        "/\[\"/"      => "\"",
        "/=>\n(\s+)/" => " => ",
    ];

    echo preg_replace(array_keys($patterns), array_values($patterns), $output);
}

function genTree9(Array $items) {
    $tree = array();
    foreach ($items as $item) {
        if (isset($items[$item['pid']])) {
            $items[$item['pid']]['sub'][] = &$items[$item['id']];
        } else {
            $tree[] = &$items[$item['id']];
        }
    }
    return $tree;
}

$items = [
   1 => [ 'id' => 1, 'name' => 'ELECTRONICS',           'pid' => null ],
   2 => [ 'id' => 2, 'name' => 'TELEVISIONS',           'pid' => 1 ],
   3 => [ 'id' => 3, 'name' => 'TUBE',                  'pid' => 2 ],
   4 => [ 'id' => 4, 'name' => 'LCD',                   'pid' => 2 ],
   5 => [ 'id' => 5, 'name' => 'PLASMA',                'pid' => 2 ],
   6 => [ 'id' => 6, 'name' => 'PORTABLE ELECTRONICS',  'pid' => 1 ],
   7 => [ 'id' => 7, 'name' => 'MP3 PLAYERS',           'pid' => 6 ],
   8 => [ 'id' => 8, 'name' => 'FLASH',                 'pid' => 7 ],
   9 => [ 'id' => 9, 'name' => 'CD PLAYERS',            'pid' => 6 ],
  10 => [ 'id' =>10, 'name' => '2 WAY RADIOS',          'pid' => 6 ],
];

const EOL = PHP_EOL;

$tree = genTree9($items);
#vd($tree);

function printTree($tree, $indent) {
    foreach ($tree as $node) {
        echo str_repeat(' ', $indent*4);
        if (isset($node['sub'])) {
            echo '+ ', $node['name'], EOL;
            printTree($node['sub'], $indent+1);
        } else {
            echo '- ', $node['name'], EOL;
        }
    }
}

printTree($tree, 0);
