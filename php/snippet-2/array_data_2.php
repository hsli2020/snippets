<?php

return array(
    'name' => 'toplevelelement', // "name" required, all else optional
    'attributes' => array(
        'foo' => 'bar',
        'fruit' => 'apple',
    ),
    'value' => 'Some random value.',
    array(
        'name' => 'achildelement',
        'value'=> 'Value',
    ),
    array(
        'name' => 'anotherchildelement',
        'attributes' => array(
            'some' => 'attr',
        ),
        array(
            'name' => 'grandchildelement',
        ),
    ),
);
