<?php

/**
 * 方法一：将数据格式化成树形结构
 */
function genTree5(Array $items) {
    foreach ($items as $item) {
        $items[$item['pid']]['son'][$item['id']] = &$items[$item['id']];
    }
    return isset($items[0]['son']) ? $items[0]['son'] : array();
}
  
/**
 * 方法二：将数据格式化成树形结构
 */
function genTree9(Array $items) {
    $tree = array();    //格式化好的树
    foreach ($items as $item) {
        if (isset($items[$item['pid']])) {
            $items[$item['pid']]['son'][] = &$items[$item['id']];
        } else {
            $tree[] = &$items[$item['id']];
        }
    }
    return $tree;
}

$items = [
   1 => [ 'id' => 1, 'name' => 'ELECTRONICS',           'pid' => 0 ],
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

#                       ELECTRONICS 
#   TELEVISIONS             			PORTABLE ELECTRONICS
#   	TUBE                				MP3 PLAYERS
#   	LCD                 					FLASH
#   	PLASMA              				CD PLAYERS
#                           				2 WAY RADIOS    

var_dump(genTree5($items));
#var_dump(genTree9($items));
