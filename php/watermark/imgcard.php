<?php

$image = 'kenshin.jpg';
$image = 'sunset.jpg';
$fontface = 'msyhl.ttc';
$fontface = 'msjhbd.ttc';
$fontface = 'c:/windows/fonts/msyh.ttc';
$fontsize = 12;
$text = "怎样来阻止皮肤的这种反弹式的油脂分泌呢？就是在洗脸后赶快抹上保湿霜，主要是保湿的，而不是油性的。这样，皮肤觉得舒服湿润，就不会大量分泌油脂了。同样的道理，洗脸时别用去油污太强的香皂或洗脸液，不是把皮肤的油脂去的越彻底越好。温和的香皂一次洗不干净就多洗一两次。";

saveImageWithText($image, $text, $fontface, $fontsize);

function saveImageWithText($imgfile, $text, $fontface, $fontsize)
{
    $border = 1;
    $padding = 10;

    $image  = imagecreatefromjpeg($imgfile);
    list($width, $height) = getimagesize($imgfile);

    // Get the size of the text area
    $wraptext = makeTextBlock($text, $fontface, $fontsize, $width - $padding*2);
    $dims = imagettfbbox($fontsize, 0, $fontface, $wraptext);
    $textwidth = $dims[4] - $dims[6];
    $textheight = $dims[3] - $dims[5];

    // Copy and resample the imag
    $newwidth  = $width  + $border*2;
    $newheight = $height + $border*2 + $padding*2 + $textheight;
    $newimg    = imagecreatetruecolor($newwidth, $newheight);

    // Colors
    $textbkcolor = imagecolorallocate($newimg, 255, 255, 255);
    $textColor   = imagecolorallocate($newimg,   0,   0,   0);
    $borderColor = imagecolorallocate($newimg, 127, 127, 127);

    // Draw Border & Image
    imagefilledrectangle($newimg, 0, 0, $newwidth, $newheight, $borderColor);
    imagecopyresampled($newimg, $image, $border, $border, 0, 0, $width, $height, $width, $height);

    // Draw Text
    imagefilledrectangle($newimg, $border, $height+$border, $newwidth-$border*2, $newheight-$border*2, $textbkcolor);
    imagettftext($newimg, $fontsize, 0, $border+$padding, $newheight-$textheight, $textColor, $fontface, $wraptext);

    imagejpeg($newimg, "o-$imgfile", 80);

    imagedestroy($image);
    imagedestroy($newimg);
}

function makeTextBlock($text, $fontfile, $fontsize, $width)
{
    $words = preg_split("//u", $text, -1, PREG_SPLIT_NO_EMPTY);

    $lines = array($words[0]);
    $currentLine = 0;

    for($i = 1; $i < count($words); $i++) {
        $lineSize = imagettfbbox($fontsize, 0, $fontfile, $lines[$currentLine] . $words[$i]);
        if ($lineSize[2] - $lineSize[0] < $width) {
            $lines[$currentLine] .= $words[$i];
        }
        else {
            $currentLine++;
            $lines[$currentLine] = $words[$i];
        }
    }

    return implode("\n", $lines);
}
