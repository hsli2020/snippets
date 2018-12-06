<?php

$table_data = [
    [
        'id' => 1,
        'parent_id' => null,
        'name' => 'Electronics',
    ],
    [
        'id' => 2,
        'parent_id' => 1,
        'name' => 'Mobiles',
    ],
    [
        'id'  => 3,
        'parent_id' => 2,
        'name' => 'Apple',
    ],
    [
        'id'  => 4,
        'parent_id' => 2,
        'name' => 'Android',
    ],
    [
        'id'  => 5,
        'parent_id' => 1,
        'name' => 'MP3 Player',
    ],
    [
        'id'  => 6,
        'parent_id' => 1,
        'name' => 'iPod',
    ],
];

function tree_render ($data)
{
    foreach ($data as $sub)
    {
        if ($sub['parent_id'] == NULL) {
            echo $sub['id'], ' ', $sub['name'], PHP_EOL;
        } else {
            $i = 0;
            $option = $sub['name'];
            $temp_id = $sub['parent_id'];

            while ($i == 0) {
                foreach ($data as $sub_name_another) {
                    if ($sub_name_another['id'] == $temp_id) {
                        if ($sub_name_another['parent_id'] == NULL) {
                            $option = $sub_name_another['name']."->".$option;
                            $i = 1;
                            echo $sub['id'], ' ', $option, PHP_EOL;
                        } else {
                            $option = $sub_name_another['name']."->".$option;
                            $temp_id = $sub_name_another['parent_id'];
                            $i = 0;
                        }
                    }
                }
            }
        }
    }
}

tree_render($table_data);
