<?php

include 'simple_html_dom.php';

define('HOST', 'https://botanwang.com');

$terms = [
    'https://botanwang.com/taxonomy/term/11792', # 微历史
#   'https://botanwang.com/taxonomy/term/11834'  # 微视野
];

foreach ($terms as $term) {
    $list = downloadList($term, 5);

    foreach ($list as $page) {
        $html = downloadPage($page);
        $imgs = downloadImages($html);

        foreach ($imgs as $img) {
            #resizeImage($img);
        }
    }
}

function downloadList($base, $maxpgs)
{
    for ($p = 0; $p < $maxpgs; $p++) {
        $url = ($p == 0) ? $base : "$base?page=$p";
        $html = downloadPage($url);
        {
            $dom = str_get_html($html);
            $links = $dom->find('#block-system-main .view-content .item-list a');
            foreach ($links as $a) {
                $link = HOST . $a->href;
                yield $link;
            }
        }
    }
}

function downloadImages($html)
{
    $files = [];

    $dom = str_get_html($html);
    $imgs = $dom->find('p.rtecenter img');

    foreach ($imgs as $img) {
        $flag = ' cached';
        $imgurl = HOST . $img->src;

        $filename = 'image/'.basename($imgurl);
        if (!file_exists($filename)) {
            $flag = '';
            $imgdat = file_get_contents($imgurl);
            file_put_contents($filename, $imgdat);
            $files[] = $filename;
            resizeImage($filename);
        }

        echo $imgurl, $flag, PHP_EOL;
    }

    return $files;
}

function downloadPage($pageurl)
{
    echo $pageurl;

    $filename = 'html/'.md5($pageurl).'.html';
    if (file_exists($filename)) {
        echo ' cached', PHP_EOL;
        return file_get_contents($filename);
    }

    echo PHP_EOL;

    $html = file_get_contents($pageurl);
    file_put_contents($filename, $html);
    return $html;
}

function resizeImage($filename)
{
    $newname = str_replace('.', '-new.', $filename);
    if (file_exists($newname)) {
        return;
    }

    // Get new sizes
    list($width, $height) = getimagesize($filename);
    $newwidth = $width;
    $newheight = $height - 20;

    // Load
    $newimg = imagecreatetruecolor($newwidth, $newheight);
    $source = imagecreatefromjpeg($filename);

    // Resize
    imagecopyresized($newimg, $source, 0, 0, 0, 0, $newwidth, $newheight, $width, $height-20);
    imagecopyresized($newimg, $source, 0, $newheight-3, 0, $height-3, $newwidth, 3, $width, 3);

    // Output
    imagejpeg($newimg, str_replace('.', '-new.', $filename), 100);
}
