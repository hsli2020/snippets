<?php

$images = [
    'kenshin.jpg',
    'sunset.jpg',
    'fruit.jpg',
];

foreach ($images as $image) {
    watermarkImage($image, "Watermark", "w-$image");
}

function watermarkImage($sourceFile, $watermarkText, $outputFile)
{
    $image = imagecreatefromjpeg($sourceFile);

    list($width, $height) = getimagesize($sourceFile);
    $newimg = imagecreatetruecolor($width, $height);
    imagecopyresampled($newimg, $image, 0, 0, 0, 0, $width, $height, $width, $height); 

    $black = imagecolorallocate($newimg,   0,   0,   0);
    $white = imagecolorallocate($newimg, 255, 255, 255);

    $font = 'Garuda.ttf';
    $fontsize = 10; 

    // Set the offset x and y for the text position
    $offsetX = 10;
    $offsetY = 0;

    // Get the size of the text area
    $dims = imagettfbbox($fontsize, 0, $font, $watermarkText);

    $textWidth  = $dims[4] - $dims[6] + $offsetX;
    $textHeight = $dims[3] - $dims[5] + $offsetY;

    $x = $width  - $textWidth;
    $y = $height - $textHeight;

    imagettftext($newimg, $fontsize, 0, $x,   $y,   $black, $font, $watermarkText);
    imagettftext($newimg, $fontsize, 0, $x-1, $y-1, $white, $font, $watermarkText);

    if ($outputFile <> '') {
        imagejpeg($newimg, $outputFile, 100); 
    } else {
        header('Content-Type: image/jpeg');
        imagejpeg($newimg, null, 100);
    };

    imagedestroy($image); 
    imagedestroy($newimg); 
}
