<?php

session_start();

header("Cache-Control: no-store,no-cache, must-revalidate");  
header('Content-type: image/png'); 

$str = '012345ABCDEFGHIJ765439KLMNOPQRS0984351TUVWXYZ876542';
$rand = rand(0,20); 
$captcha_code = substr(str_shuffle($str),$rand,6);

$_SESSION["code"] = $captcha_code;

if (!function_exists('gd_info') ) {
   throw new Exception('Required GD library is missing');
}    
   
$layer = imagecreatetruecolor(150, 50);
$bg_color = imagecolorallocate($layer, 255, 255, 255);  ; 
$fg_color = imagecolorallocate($layer, 0,0,0);
$line_color =  imagecolorallocate($layer, 64,64,64);   
$pixel_color = imagecolorallocate($layer, 255,125,325);

for($i=0;$i<1500;$i++) {
   imagesetpixel($layer,rand()%150,rand()%50,$pixel_color);
} 

imagefill($layer, 0, 0, $bg_color);  
imagestring($layer, 5, 7, 9, $captcha_code, $fg_color); 
imagepng($layer);  
imagedestroy($layer); 
